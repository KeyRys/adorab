package usecase

import (
	"adoend/internal/domain"
	"adoend/internal/repository"
)

type ProductUsecase struct {
	Repo *repository.ProductRepository
}

func NewProductUsecase(r *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{Repo: r}
}

func (u *ProductUsecase) GetProducts() ([]domain.Product, error) {

	return u.Repo.GetAll()
}

func (u *ProductUsecase) GetProductByID(id string) (*domain.Product, error) {
	return u.Repo.GetByID(id)
}
