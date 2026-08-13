package repository

import (
	"adoend/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type CartRepository struct {
	DB *pgx.Conn
}

func NewCartRepository(db *pgx.Conn) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) GetOrCreateCart(userID string) (string, error) {
	ctx := context.Background()

	var cartID string

	err := r.DB.QueryRow(ctx,
		"SELECT id FROM carts WHERE user_id=$1",
		userID,
	).Scan(&cartID)

	if err == nil {
		return cartID, nil
	}

	err = r.DB.QueryRow(ctx,
		"INSERT INTO carts (user_id) VALUES ($1) RETURNING id",
		userID,
	).Scan(&cartID)

	if err != nil {
		return "", err
	}

	return cartID, nil
}

func (r *CartRepository) AddItem(cartID, productID string) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx,
		"INSERT INTO cart_items (cart_id, product_id) VALUES ($1, $2)",
		cartID, productID,
	)

	return err
}

func (r *CartRepository) GetCartItems(userID string) ([]domain.CartItemWithProduct, error) {
	ctx := context.Background()
	query := `
		SELECT ci.id, p.id, p.name, p.breed, p.price
		FROM cart_items ci
		JOIN rabbits p 
			ON ci.product_id = p.id
		WHERE ci.cart_id = $1 :: uuid
	`
	rows, err := r.DB.Query(ctx, query, userID)

	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var items []domain.CartItemWithProduct

	for rows.Next() {
		var item domain.CartItemWithProduct

		err := rows.Scan(
			&item.ID,
			&item.Product.ID,
			&item.Product.Name,
			&item.Product.Breed,
			&item.Product.Price,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *CartRepository) RemoveItem(itemID string) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx,
		"DELETE FROM cart_items WHERE id=$1::uuid",
		itemID,
	)

	return err
}

func (r *CartRepository) ClearCart(cartID string) error {
	ctx := context.Background()

	_, err := r.DB.Exec(ctx,
		"DELETE FROM cart_items WHERE cart_id=$1::uuid",
		cartID,
	)

	return err
}
