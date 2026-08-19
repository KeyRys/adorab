package usecase

import "adoend/internal/domain"

type CheckoutRepository interface {
	GetCartItemsByUserID(userID string) ([]domain.CheckoutItem, error)
	CreateOrder(userID string, total int) (string, error)
	CreateOrderItem(orderID string, item domain.CheckoutItem) error
	ClearCart(userID string) error
}
