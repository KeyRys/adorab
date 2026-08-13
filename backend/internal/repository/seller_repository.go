package repository

import (
	"adoend/internal/domain"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5"
)

type SellerRepository struct {
	DB *pgx.Conn
}

func NewSellerRepository(db *pgx.Conn) *SellerRepository {
	return &SellerRepository{DB: db}
}

func (r *SellerRepository) IsSeller(userID string) (bool, error) {
	var exists bool
	ctx := context.Background()
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM sellers
			WHERE user_id = $1::uuid
		)
	`
	err := r.DB.QueryRow(ctx, query, userID).Scan(&exists)
	return exists, err
}

func (r *SellerRepository) CreateSeller(userID, location string) error {
	ctx := context.Background()
	_, err := r.DB.Exec(ctx,
		"INSERT INTO sellers (user_id, location) VALUES ($1::uuid, $2)",
		userID, location,
	)

	return err
}

func (r *SellerRepository) GetSellerIDByUserID(userID string) (string, error) {
	var sellerID string
	ctx := context.Background()
	query := `
		SELECT id
		FROM sellers
		WHERE user_id = $1::uuid
	`
	err := r.DB.QueryRow(ctx, query, userID).Scan(&sellerID)

	return sellerID, err
}

func (r *SellerRepository) CreateRabbit(rabbit *domain.Rabbit) error {
	rabbit.HealthStatus = "healthy"
	query := `
		INSERT INTO rabbits (
			seller_id,
			name,
			breed,
			gender,
			age,
			weight,
			color,
			price,
			purpose,
			description,
			health_status,
			image_url
		)
		VALUES (
			$1::uuid,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12
		)
	`
	_, err := r.DB.Exec(
		context.Background(),
		query,
		rabbit.SellerID,
		rabbit.Name,
		rabbit.Breed,
		rabbit.Gender,
		rabbit.Age,
		rabbit.Weight,
		rabbit.Color,
		rabbit.Price,
		rabbit.Purpose,
		rabbit.Description,
		rabbit.HealthStatus,
		rabbit.ImageURL,
	)
	return err
}

func (r *SellerRepository) GetSellerRabbits(userID string) ([]domain.Rabbit, error) {
	ctx := context.Background()
	query := `
		SELECT
			r.id,
			r.seller_id,
			r.name,
			r.breed,
			r.age,
			r.gender,
			r.weight,
			r.color,
			r.price,
			r.purpose,
			r.health_status,
			r.description,
			r.image_url
		FROM rabbits r
		JOIN sellers s
			ON r.seller_id = s.id
		WHERE s.user_id = $1::uuid
	`
	rows, err := r.DB.Query(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rabbits []domain.Rabbit
	for rows.Next() {

		var rabbit domain.Rabbit
		err := rows.Scan(
			&rabbit.ID,
			&rabbit.SellerID,
			&rabbit.Name,
			&rabbit.Breed,
			&rabbit.Age,
			&rabbit.Gender,
			&rabbit.Weight,
			&rabbit.Color,
			&rabbit.Price,
			&rabbit.Purpose,
			&rabbit.HealthStatus,
			&rabbit.Description,
			&rabbit.ImageURL,
		)

		if err != nil {
			return nil, err
		}
		rabbits = append(rabbits, rabbit)
	}

	return rabbits, nil
}

func (r *SellerRepository) UpdateRabbit(
	rabbitID string,
	sellerID string,
	name string,
	breed string,
	healthstatus string,
	price int,
	description string,
	imageURL string,
) error {
	ctx := context.Background()

	if imageURL != "" {
		_, err := r.DB.Exec(ctx, `
			UPDATE rabbits
			SET
				name = $1,
				breed = $2,
				price = $3,
				description = $4,
				health_status = $5,
				image_url = $6
			WHERE id = $7
			AND seller_id = $8
		`,
			name,
			breed,
			price,
			description,
			healthstatus,
			imageURL,
			rabbitID,
			sellerID,
		)
		return err
	}

	_, err := r.DB.Exec(ctx, `
		UPDATE rabbits
		SET
			name = $1,
			breed = $2,
			price = $3,
			description = $4,
			health_status = $5
		WHERE id = $6
		AND seller_id = $7
	`,
		name,
		breed,
		price,
		description,
		healthstatus,
		rabbitID,
		sellerID,
	)
	return err
}

func (r *SellerRepository) DeleteRabbit(rabbitID string, sellerID string) error {
	ctx := context.Background()
	_, err := r.DB.Exec(ctx,
		"DELETE FROM rabbits WHERE id = $1 AND seller_id = $2",
		rabbitID, sellerID,
	)

	return err
}

func (r *SellerRepository) UploadRabbitImage(fileHeader *multipart.FileHeader, sellerID string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		ext = ".jpg"
	}

	fileName := fmt.Sprintf("%s_%d%s", sellerID, time.Now().Unix(), ext)
	filePath := "rabbits/" + fileName

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	bucketName := "rabbit_img"

	uploadURL := fmt.Sprintf(
		"%s/storage/v1/object/%s/%s",
		supabaseURL,
		bucketName,
		filePath,
	)

	req, err := http.NewRequest("POST", uploadURL, bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "image/jpeg"
	}

	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-upsert", "true")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("supabase upload failed: %s", string(respBody))
	}

	publicURL := fmt.Sprintf(
		"%s/storage/v1/object/public/%s/%s",
		supabaseURL,
		bucketName,
		filePath,
	)
	return publicURL, nil
}
