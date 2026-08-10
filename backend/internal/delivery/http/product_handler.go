package http

import (
	"adoend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Usecase *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{Usecase: u}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.Usecase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.Usecase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
