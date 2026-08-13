package usecase

import "adoend/internal/repository"

type ProfileUsecase struct {
	Repo *repository.ProfileRepository
}

func NewProfileUsecase(r *repository.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{Repo: r}
}

func (u *ProfileUsecase) GetProfile(userID string) (*repository.ProfileResponse, error) {
	return u.Repo.GetProfileByUserID(userID)
}

func (u *ProfileUsecase) UpdateProfile(
	userID string,
	phone string,
	address string,
) error {
	return u.Repo.UpdateProfile(
		userID,
		phone,
		address,
	)
}
