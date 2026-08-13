package repository

import (
	"context"

	"adoend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type CheckoutRepository struct {
	DB *pgx.Conn
}

func NewCheckoutRepository(db *pgx.Conn) *CheckoutRepository {
	return &CheckoutRepository{DB: db}
}

func (r *CheckoutRepository) GetCartItemsByUserID(userID string) ([]domain.CheckoutItem, error) {
	ctx := context.Background()
	rows, err := r.DB.Query(
		ctx,
		`
		SELECT
			ci.product_id,
			r.seller_id,
			r.price
		FROM carts c
		JOIN cart_items ci
		    ON ci.cart_id = c.id
		JOIN rabbits r
			ON ci.product_id = r.id
		WHERE c.user_id = $1::uuid
		`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.CheckoutItem
	for rows.Next() {
		var item domain.CheckoutItem
		err := rows.Scan(
			&item.RabbitID,
			&item.SellerID,
			&item.Price,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *CheckoutRepository) CreateOrder(userID string, total int) (string, error) {
	ctx := context.Background()
	var orderID string

	err := r.DB.QueryRow(
		ctx,
		`
		INSERT INTO orders (
			user_id,
			total_price,
			status
		)
		VALUES ($1, $2, 'pending')
		RETURNING id
		`,
		userID,
		total,
	).Scan(&orderID)

	return orderID, err
}

func (r *CheckoutRepository) CreateOrderItem(orderID string, item domain.CheckoutItem) error {
	ctx := context.Background()
	_, err := r.DB.Exec(
		ctx,
		`
		INSERT INTO order_items (
			order_id,
			rabbit_id,
			seller_id,
			price
		)
		VALUES ($1, $2, $3, $4)
		`,
		orderID,
		item.RabbitID,
		item.SellerID,
		item.Price,
	)

	return err
}

func (r *CheckoutRepository) ClearCart(userID string) error {
	ctx := context.Background()
	_, err := r.DB.Exec(
		ctx,
		"DELETE FROM carts c WHERE c.user_id = $1::uuid",
		userID,
	)

	return err
}
