package usecase

import (
	"adoend/internal/repository"
)

type CartUsecase struct {
	Repo *repository.CartRepository
}

func NewCartUsecase(r *repository.CartRepository) *CartUsecase {
	return &CartUsecase{Repo: r}
}

func (u *CartUsecase) AddToCart(userID, productID string) error {
	cartID, err := u.Repo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	return u.Repo.AddItem(cartID, productID)
}

func (u *CartUsecase) GetCart(userID string) (interface{}, error) {
	cartID, err := u.Repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	return u.Repo.GetCartItems(cartID)
}

func (u *CartUsecase) RemoveItem(itemID string) error {
	return u.Repo.RemoveItem(itemID)
}
