package http

import (
	"adoend/internal/domain"
	"adoend/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	Usecase *usecase.SellerUsecase
}

func NewSellerHandler(u *usecase.SellerUsecase) *SellerHandler {
	return &SellerHandler{Usecase: u}
}

type BecomeSellerRequest struct {
	Location string `json:"location"`
}

func (h *SellerHandler) BecomeSeller(c *gin.Context) {
	var req BecomeSellerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("id").(string)
	err := h.Usecase.BecomeSeller(userID, req.Location)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success become seller"})
}

func (h *SellerHandler) CreateRabbit(c *gin.Context) {
	userID := c.MustGet("id").(string)

	age, err := strconv.Atoi(c.PostForm("age"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age"})
		return
	}

	weight, err := strconv.ParseFloat(c.PostForm("weight"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid weight"})
		return
	}

	price, err := strconv.ParseInt(c.PostForm("price"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	rabbit := domain.Rabbit{
		Name:        c.PostForm("name"),
		Breed:       c.PostForm("breed"),
		Age:         age,
		Gender:      c.PostForm("gender"),
		Weight:      weight,
		Color:       c.PostForm("color"),
		Price:       price,
		Purpose:     c.PostForm("purpose"),
		Description: c.PostForm("description"),
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	err = h.Usecase.CreateRabbit(userID, &rabbit, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "rabbit created",
	})
}

func (h *SellerHandler) GetSellerRabbits(c *gin.Context) {
	userID := c.MustGet("id").(string)
	rabbits, err := h.Usecase.GetSellerRabbits(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rabbits)
}

func (h *SellerHandler) UpdateRabbit(c *gin.Context) {
	rabbitID := c.Param("id")
	userID := c.MustGet("id").(string)

	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	name := c.PostForm("name")
	breed := c.PostForm("breed")
	healthStatus := c.PostForm("health_status")
	description := c.PostForm("description")

	fileHeader, err := c.FormFile("image")
	if err != nil {
		fileHeader = nil
	}

	err = h.Usecase.UpdateRabbit(
		rabbitID,
		userID,
		name,
		breed,
		healthStatus,
		price,
		description,
		fileHeader,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rabbit updated"})
}

func (h *SellerHandler) DeleteRabbit(c *gin.Context) {
	rabbitID := c.Param("id")
	user_id := c.MustGet("id").(string)

	err := h.Usecase.DeleteRabbit(
		rabbitID,
		user_id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "rabbit deleted"})
}
