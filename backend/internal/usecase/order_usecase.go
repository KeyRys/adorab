package usecase

import (
	"adoend/internal/domain"
	"adoend/internal/repository"
)

type OrderUsecase struct {
	OrderRepo *repository.OrderRepository
}

func NewOrderUsecase(orderRepo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{OrderRepo: orderRepo}
}

func (u *OrderUsecase) GetBuyerOrders(buyerID string) ([]domain.BuyerOrder, error) {
	return u.OrderRepo.GetBuyerOrders(buyerID)
}

type SellerOrderUsecase struct {
	SellerOrderRepo *repository.OrderRepository
}

func NewSellerOrderUsecase(sellerOrderRepo *repository.OrderRepository) *SellerOrderUsecase {
	return &SellerOrderUsecase{SellerOrderRepo: sellerOrderRepo}
}

func (u *SellerOrderUsecase) GetSellerOrders(userID string) ([]domain.SellerOrder, error) {
	return u.SellerOrderRepo.GetSellerOrders(userID)
}

func (u *SellerOrderUsecase) UpdateOrderStatus(orderID string, status string) error {
	return u.SellerOrderRepo.UpdateOrderStatus(orderID, status)
}
