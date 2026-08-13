package http

import (
	"adoend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{OrderUsecase: orderUsecase}
}

func (h *OrderHandler) GetBuyerOrders(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orders, err := h.OrderUsecase.GetBuyerOrders(userID.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

type SellerOrderHandler struct {
	SellerOrderUsecase *usecase.SellerOrderUsecase
}

func NewSellerOrderHandler(sellerOrderUsecase *usecase.SellerOrderUsecase) *SellerOrderHandler {
	return &SellerOrderHandler{SellerOrderUsecase: sellerOrderUsecase}
}

func (h *SellerOrderHandler) GetSellerOrders(c *gin.Context) {
	userID, _ := c.Get("id")

	orders, err := h.SellerOrderUsecase.GetSellerOrders(userID.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *SellerOrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.SellerOrderUsecase.UpdateOrderStatus(orderID, body.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}
