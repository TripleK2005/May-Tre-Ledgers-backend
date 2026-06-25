package token

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, token *RefreshToken) error
	FindByToken(ctx context.Context, token string) (*RefreshToken, error)
	Revoke(ctx context.Context, token string) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, token *RefreshToken) error {
	_, err := r.db.Exec(
		ctx,
		`
			INSERT INTO refresh_tokens (
				id,
				user_id,
				token_hash,
				expires_at
			)
			VALUES ($1, $2, $3, $4)
		`,
		token.ID,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
	)

	return err
}

func (r *repository) FindByToken(ctx context.Context, tokenHash string) (*RefreshToken, error) {
	var token RefreshToken

	err := r.db.QueryRow(
		ctx,
		`
			SELECT
				id,
				user_id,
				token_hash,
				expires_at,
				revoked,
				created_at
			FROM refresh_tokens
			WHERE token_hash = $1
		`,
		tokenHash,
	).Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.Revoked,
		&token.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *repository) Revoke(ctx context.Context, tokenHash string) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE refresh_tokens SET revoked = TRUE WHERE token_hash = $1`,
		tokenHash,
	)

	return err
}
