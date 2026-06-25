package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

const (
	createUserQuery = `
		INSERT INTO users (
			id,
			username,
			email,
			password_hash,
			full_name,
			role_id
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	findUserByIDQuery = `
		SELECT
			u.id,
			u.username,
			u.email,
			u.password_hash,
			u.full_name,
			u.role_id,
			r.name,
			u.is_active,
			u.created_at,
			u.updated_at
		FROM users u
		JOIN roles r ON r.id = u.role_id
		WHERE u.id = $1
	`

	findUserByUsernameQuery = `
		SELECT
			u.id,
			u.username,
			u.email,
			u.password_hash,
			u.full_name,
			u.role_id,
			r.name,
			u.is_active,
			u.created_at,
			u.updated_at
		FROM users u
		JOIN roles r ON r.id = u.role_id
		WHERE u.username = $1
	`
)

func (r *repository) Create(
	ctx context.Context,
	user *User,
) error {

	_, err := r.db.Exec(
		ctx, createUserQuery,
		user.ID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.RoleID,
	)

	return err
}

func (r *repository) FindByID(
	ctx context.Context,
	id string,
) (*User, error) {

	var user User

	err := r.db.QueryRow(
		ctx, findUserByIDQuery,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.RoleID,
		&user.RoleName,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByUsername(
	ctx context.Context,
	username string,
) (*User, error) {

	var user User

	err := r.db.QueryRow(
		ctx,
		findUserByUsernameQuery,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.RoleID,
		&user.RoleName,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
