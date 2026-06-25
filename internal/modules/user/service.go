package user

import (
	"context"
)

type Service interface {
	GetByID(ctx context.Context, id string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetByID (
	ctx context.Context,
	id string,
) (*User, error) {
	return s.repo.FindByID(ctx, id)
}