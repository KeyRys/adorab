package usecase

import (
	"adoend/internal/domain"
	"adoend/internal/repository"
	"errors"
	"mime/multipart"
)

type SellerUsecase struct {
	Repo *repository.SellerRepository
}

func NewSellerUsecase(r *repository.SellerRepository) *SellerUsecase {
	return &SellerUsecase{Repo: r}
}

func (u *SellerUsecase) BecomeSeller(userID, location string) error {
	exists, err := u.Repo.IsSeller(userID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already seller")
	}

	return u.Repo.CreateSeller(userID, location)
}

func (u *SellerUsecase) CreateRabbit(
	userID string,
	rabbit *domain.Rabbit,
	fileHeader *multipart.FileHeader,
) error {
	sellerID, err := u.Repo.GetSellerIDByUserID(userID)
	if err != nil {
		return err
	}

	rabbit.SellerID = sellerID

	if fileHeader != nil {
		imageURL, err := u.Repo.UploadRabbitImage(fileHeader, sellerID)
		if err != nil {
			return err
		}
		rabbit.ImageURL = imageURL
	}

	return u.Repo.CreateRabbit(rabbit)
}

func (u *SellerUsecase) GetSellerRabbits(userID string) ([]domain.Rabbit, error) {
	return u.Repo.GetSellerRabbits(userID)
}

func (u *SellerUsecase) UpdateRabbit(
	rabbitID string,
	userID string,
	name string,
	breed string,
	healthstatus string,
	price int,
	description string,
	fileHeader *multipart.FileHeader,
) error {
	sellerID, err := u.Repo.GetSellerIDByUserID(userID)
	if err != nil {
		return err
	}

	imageURL := ""
	if fileHeader != nil {
		imageURL, err = u.Repo.UploadRabbitImage(fileHeader, sellerID)
		if err != nil {
			return err
		}
	}

	return u.Repo.UpdateRabbit(
		rabbitID,
		sellerID,
		name,
		breed,
		healthstatus,
		price,
		description,
		imageURL,
	)
}

func (u *SellerUsecase) DeleteRabbit(
	rabbitID string,
	userID string,
) error {

	sellerID, err := u.Repo.GetSellerIDByUserID(userID)
	if err != nil {
		return err
	}

	return u.Repo.DeleteRabbit(
		rabbitID,
		sellerID,
	)
}
