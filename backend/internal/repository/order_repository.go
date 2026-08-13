package repository

import (
	"adoend/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	DB *pgx.Conn
}

func NewOrderRepository(db *pgx.Conn) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) GetBuyerOrders(buyerID string) ([]domain.BuyerOrder, error) {
	ctx := context.Background()
	rows, err := r.DB.Query(
		ctx, `
		SELECT
			oi.id,
			r.name,
			oi.price,
			o.status
		FROM orders o
		JOIN order_items oi
			ON oi.order_id = o.id
		JOIN rabbits r
			ON oi.rabbit_id = r.id
		WHERE o.user_id = $1::uuid

		ORDER BY o.created_at DESC
		`,
		buyerID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.BuyerOrder
	for rows.Next() {

		var order domain.BuyerOrder
		err := rows.Scan(
			&order.OrderID,
			&order.RabbitName,
			&order.Price,
			&order.Status,
		)

		if err != nil {
			return nil, err
		}

		orders = append(
			orders,
			order,
		)
	}
	return orders, nil
}

func (r *OrderRepository) GetSellerOrders(userID string) ([]domain.SellerOrder, error) {
	ctx := context.Background()
	rows, err := r.DB.Query(
		ctx,
		`
		SELECT
			o.id,
			p.name buyer_name,
			rb.name,
			oi.price,
			o.status
		FROM order_items oi
		JOIN orders o
			ON oi.order_id = o.id
		JOIN rabbits rb
			ON oi.rabbit_id = rb.id
		JOIN profiles p
			ON o.user_id = p.user_id
		WHERE oi.seller_id = (Select id from sellers where user_id = $1::uuid)
		`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.SellerOrder
	for rows.Next() {

		var order domain.SellerOrder
		err := rows.Scan(
			&order.OrderID,
			&order.BuyerName,
			&order.RabbitName,
			&order.Price,
			&order.Status,
		)

		if err != nil {
			return nil, err
		}

		orders = append(
			orders,
			order,
		)
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrderStatus(orderID string, status string) error {
	ctx := context.Background()
	_, err := r.DB.Exec(
		ctx,
		`
		UPDATE orders
		SET status = $1
		WHERE id = $2::uuid
		`,
		status, orderID,
	)

	return err
}
