package repository

import (
	"adoend/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	DB *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	ctx := context.Background()
	query := `
	SELECT id,
		seller_id,
		(Select phone from profiles where user_id = (Select user_id from sellers where id = seller_id::uuid)) as phone,
		name,
		breed,
		gender,
		age,
		color,
		price,
		purpose,
		health_status,
		image_url
	FROM rabbits
	WHERE health_status = 'healthy'
	`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.SellerID, &p.Phone, &p.Name, &p.Breed, &p.Gender, &p.Age, &p.Color, &p.Price, &p.Purpose, &p.HealthStatus, &p.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id string) (*domain.Product, error) {
	row := r.DB.QueryRow(context.Background(), "SELECT id, seller_id, (Select phone from profiles where user_id = (Select user_id from sellers where id = seller_id::uuid))as phone, name, breed, age, weight, color, gender, price, description, purpose, health_status, image_url FROM rabbits WHERE id = $1", id)

	var p domain.Product
	err := row.Scan(&p.ID, &p.SellerID, &p.Phone, &p.Name, &p.Breed, &p.Age, &p.Weight, &p.Color, &p.Gender, &p.Price, &p.Description, &p.Purpose, &p.HealthStatus, &p.ImageURL)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
