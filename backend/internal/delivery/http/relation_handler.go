package http

import (
	"net/http"

	"adoend/internal/domain"
	"adoend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RabbitRelationHandler struct {
	RelationUsecase *usecase.RabbitRelationUsecase
}

func NewRabbitRelationHandler(usecase *usecase.RabbitRelationUsecase) *RabbitRelationHandler {
	return &RabbitRelationHandler{RelationUsecase: usecase}
}

func (h *RabbitRelationHandler) CreateRelation(c *gin.Context) {
	var relation domain.RabbitRelation

	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.RelationUsecase.CreateRelation(relation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Relation created"})
}

func (h *RabbitRelationHandler) GetSellerRelation(c *gin.Context) {
	userID := c.GetString("id")
	relation, err := h.RelationUsecase.GetSellerRelations(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relation)
}
