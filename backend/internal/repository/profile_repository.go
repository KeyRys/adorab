package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type ProfileRepository struct {
	DB *pgx.Conn
}

func NewProfileRepository(db *pgx.Conn) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

type ProfileResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Isseller bool   `json:"is_seller"`
}

func (r *ProfileRepository) GetProfileByUserID(userID string) (*ProfileResponse, error) {
	var profile ProfileResponse
	ctx := context.Background()
	query := `
		SELECT
			id,
			user_id,
			name,
			phone,
			address,
			EXISTS (
				SELECT *
				FROM seller
				WHERE sellers.user_id = profiles.user_id
			) AS is_seller
		FROM profiles
		WHERE user_id = $1::uuid
	`
	err := r.DB.QueryRow(ctx, query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.Name,
		&profile.Phone,
		&profile.Address,
		&profile.Isseller,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) UpdateProfile(
	userID string,
	phone string,
	address string,
) error {
	ctx := context.Background()
	_, err := r.DB.Exec(ctx, `
		UPDATE profiles
		SET
			phone = $1,
			address = $2
		WHERE user_id = $3::uuid
		`,
		phone, address, userID)
	return err
}
