package http

import (
	"net/http"

	"adoend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	CheckoutUsecase *usecase.CheckoutUsecase
}

func NewCheckoutHandler(checkoutUsecase *usecase.CheckoutUsecase) *CheckoutHandler {
	return &CheckoutHandler{CheckoutUsecase: checkoutUsecase}
}

func (h *CheckoutHandler) Checkout(c *gin.Context) {
	userID := c.GetString("id")
	err := h.CheckoutUsecase.Checkout(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checkout successful"})
}
