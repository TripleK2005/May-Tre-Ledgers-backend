package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"may-tre-ledger-be/internal/core/config"
	"may-tre-ledger-be/internal/modules/role"
	"may-tre-ledger-be/internal/modules/token"
	"may-tre-ledger-be/internal/modules/user"
	"may-tre-ledger-be/internal/utils"

	"github.com/google/uuid"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrInactiveUser       = errors.New("user is inactive")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("expired token")
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) error
	Login(ctx context.Context, req LoginRequest) (*AuthResponse, error)
	RefreshToken(ctx context.Context, token string) (*AuthResponse, error)
	Logout(ctx context.Context, token string) error
}

type service struct {
	userRepo  user.Repository
	roleRepo  role.Repository
	tokenRepo token.Repository

	cfg *config.Config
}

func NewService(
	userRepo user.Repository,
	roleRepo role.Repository,
	tokenRepo token.Repository,
	cfg *config.Config,
) Service {
	return &service{
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		tokenRepo: tokenRepo,
		cfg:       cfg,
	}
}

func (s *service) Register(ctx context.Context, req RegisterRequest) error {
	roleName := strings.ToUpper(strings.TrimSpace(req.Role))
	if roleName == "" {
		roleName = "STAFF"
	}

	userRole, err := s.roleRepo.FindByName(ctx, roleName)
	if err != nil {
		return err
	}

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := &user.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
		FullName:     req.FullName,
		RoleID:       userRole.ID,
	}

	return s.userRepo.Create(ctx, newUser)
}

func (s *service) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	existingUser, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !existingUser.IsActive {
		return nil, ErrInactiveUser
	}

	if err := utils.ComparePassword(existingUser.PasswordHash, req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return s.issueTokens(ctx, existingUser)
}

func (s *service) RefreshToken(ctx context.Context, refreshToken string) (*AuthResponse, error) {
	claims, err := utils.ParseToken(refreshToken, s.cfg.JWTSecret)
	if err != nil {
		return nil, ErrInvalidToken
	}

	storedToken, err := s.tokenRepo.FindByToken(ctx, hashToken(refreshToken))
	if err != nil {
		return nil, ErrInvalidToken
	}

	if storedToken.Revoked {
		return nil, ErrInvalidToken
	}

	if time.Now().After(storedToken.ExpiresAt) {
		return nil, ErrExpiredToken
	}

	existingUser, err := s.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if !existingUser.IsActive {
		return nil, ErrInactiveUser
	}

	if err := s.tokenRepo.Revoke(ctx, hashToken(refreshToken)); err != nil {
		return nil, err
	}

	return s.issueTokens(ctx, existingUser)
}

func (s *service) Logout(ctx context.Context, refreshToken string) error {
	return s.tokenRepo.Revoke(ctx, hashToken(refreshToken))
}

func (s *service) issueTokens(ctx context.Context, user *user.User) (*AuthResponse, error) {
	accessToken, err := utils.GenerateAccessToken(
		user.ID.String(),
		user.RoleName,
		s.cfg.JWTSecret,
		s.cfg.AccessTokenExpire,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(
		user.ID.String(),
		s.cfg.JWTSecret,
		s.cfg.RefreshTokenExpire,
	)
	if err != nil {
		return nil, err
	}

	err = s.tokenRepo.Create(ctx, &token.RefreshToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		TokenHash: hashToken(refreshToken),
		ExpiresAt: time.Now().Add(s.cfg.RefreshTokenExpire),
	})
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
