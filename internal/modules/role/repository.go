package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Role struct {
	ID   uuid.UUID
	Name string
}

type Repository interface {
	FindByName(ctx context.Context, name string) (*Role, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Role, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByName(ctx context.Context, name string) (*Role, error) {
	var role Role

	err := r.db.QueryRow(
		ctx,
		`SELECT id, name FROM roles WHERE name = $1`,
		name,
	).Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*Role, error) {
	var role Role

	err := r.db.QueryRow(
		ctx,
		`SELECT id, name FROM roles WHERE id = $1`,
		id,
	).Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
