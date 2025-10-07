package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
}

func InitAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

type LoginReq struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// bind inputs
	var input *LoginReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
		return
	}

	// login
	if err := ah.svc.Login(c, &domain.User{
		Email: input.Email,
		Password: input.Password,
	}); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// return ok response
	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in successfully",
	})
}
