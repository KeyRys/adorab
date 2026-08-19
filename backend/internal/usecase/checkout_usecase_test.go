package usecase

import (
	"errors"
	"testing"

	"adoend/internal/domain"
)

type MockCheckoutRepository struct {
	items []domain.CheckoutItem

	createdOrderUserID string
	createdOrderTotal  int
	createdOrderID     string

	createdOrderItems []domain.CheckoutItem

	cartClearedForUser string

	getCartItemsErr    error
	createOrderErr     error
	createOrderItemErr error
	clearCartErr       error
}

func (m *MockCheckoutRepository) GetCartItemsByUserID(userID string) ([]domain.CheckoutItem, error) {
	if m.getCartItemsErr != nil {
		return nil, m.getCartItemsErr
	}
	return m.items, nil
}

func (m *MockCheckoutRepository) CreateOrder(userID string, total int) (string, error) {
	if m.createOrderErr != nil {
		return "", m.createOrderErr
	}
	m.createdOrderUserID = userID
	m.createdOrderTotal = total
	m.createdOrderID = "test-order-id"
	return m.createdOrderID, nil
}

func (m *MockCheckoutRepository) CreateOrderItem(orderID string, item domain.CheckoutItem) error {
	if m.createOrderItemErr != nil {
		return m.createOrderItemErr
	}
	m.createdOrderItems = append(m.createdOrderItems, item)
	return nil
}

func (m *MockCheckoutRepository) ClearCart(userID string) error {
	if m.clearCartErr != nil {
		return m.clearCartErr
	}
	m.cartClearedForUser = userID
	return nil
}

func TestCheckout_Success(t *testing.T) {
	repo := &MockCheckoutRepository{
		items: []domain.CheckoutItem{
			{
				RabbitID: "rabbit-1",
				SellerID: "seller-1",
				Price:    100,
			},
			{
				RabbitID: "rabbit-2",
				SellerID: "seller-2",
				Price:    250,
			},
		},
	}
	usecase := NewCheckoutUsecase(repo)
	err := usecase.Checkout("user-1")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if repo.createdOrderTotal != 350 {
		t.Errorf("Expected order total 350, but got %d", repo.createdOrderTotal)
	}

	if repo.createdOrderUserID != "user-1" {
		t.Errorf("Expected order user ID 'user-1', but got %s", repo.createdOrderUserID)
	}

	if len(repo.createdOrderItems) != 2 {
		t.Errorf("Expected 2 order items, but got %d", len(repo.createdOrderItems))
	}

	if repo.createdOrderItems[0].RabbitID != "rabbit-1" {
		t.Errorf("Expected first order item RabbitID 'rabbit-1', but got %s", repo.createdOrderItems[0].RabbitID)
	}

	if repo.createdOrderItems[1].RabbitID != "rabbit-2" {
		t.Errorf("Expected second order item RabbitID 'rabbit-2', but got %s", repo.createdOrderItems[1].RabbitID)
	}

	if repo.cartClearedForUser != "user-1" {
		t.Errorf("Expected cart cleared for user 'user-1', but got %s", repo.cartClearedForUser)
	}
}

func TestCheckout_GetCartItemsError(t *testing.T) {
	expectedErr := errors.New("failed to get cart items")

	repo := &MockCheckoutRepository{
		getCartItemsErr: expectedErr,
	}

	usecase := NewCheckoutUsecase(repo)

	err := usecase.Checkout("user-1")

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected error %v, but got %v", expectedErr, err)
	}

	if repo.createdOrderID != "" {
		t.Error("expected order not to be created")
	}

	if len(repo.createdOrderItems) != 0 {
		t.Error("expected no order items to be created")
	}

	if repo.cartClearedForUser != "" {
		t.Error("expected cart not to be cleared")
	}
}
