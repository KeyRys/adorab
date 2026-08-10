package repository

import (
	"adoend/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	DB *pgx.Conn
}

func NewAuthRepository(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUserWithProfile(
	email, password, name, phone string,
) error {
	ctx := context.Background()

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var uid string

	err = tx.QueryRow(ctx,
		"INSERT INTO users (email, hashed_password) VALUES ($1, $2) RETURNING id",
		email, password,
	).Scan(&uid)

	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		"INSERT INTO profiles (user_id, name, phone) VALUES ($1, $2, $3)",
		uid, name, phone,
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *AuthRepository) GetUserByEmail(email string) (*domain.User, error) {
	ctx := context.Background()
	query := `
	SELECT id,
		email,
		hashed_password
	FROM users
	WHERE email=$1
	`

	row := r.DB.QueryRow(ctx, query, email)

	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
