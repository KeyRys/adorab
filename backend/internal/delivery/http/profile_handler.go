package http

import (
	"adoend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Usecase *usecase.ProfileUsecase
}

func NewProfileHandler(u *usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{Usecase: u}
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	userID := c.GetString("id")

	profile, err := h.Usecase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("id")

	var body struct {
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Usecase.UpdateProfile(
		userID.(string),
		body.Phone,
		body.Address,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}
