package http

import (
	"adoend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Usecase *usecase.CartUsecase
}

func NewCartHandler(u *usecase.CartUsecase) *CartHandler {
	return &CartHandler{Usecase: u}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	var req struct {
		ProductID string `json:"product_id"`
	}
	userID := c.GetString("id")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.Usecase.AddToCart(userID, req.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "added to cart"})
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetString("id")
	data, err := h.Usecase.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve cart"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	itemID := c.Param("id")
	err := h.Usecase.RemoveItem(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed"})
}
