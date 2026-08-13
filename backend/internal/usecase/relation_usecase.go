package usecase

import (
	"adoend/internal/domain"
	"adoend/internal/repository"
	"errors"
)

type RabbitRelationUsecase struct {
	RelationRepo *repository.RabbitRelationRepository
}

func NewRabbitRelationUsecase(repo *repository.RabbitRelationRepository) *RabbitRelationUsecase {
	return &RabbitRelationUsecase{RelationRepo: repo}
}

func (u *RabbitRelationUsecase) CreateRelation(relation domain.RabbitRelation) error {
	if relation.ParentID == relation.ChildID {
		return errors.New("rabbit cannot be its own parent")
	}

	return u.RelationRepo.CreateRelation(relation)
}

func (u *RabbitRelationUsecase) GetSellerRelations(userID string) ([]domain.RabbitRelationResponse, error) {
	return u.RelationRepo.GetSellerRelations(userID)
}
