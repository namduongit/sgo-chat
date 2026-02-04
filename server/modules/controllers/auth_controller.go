package controllers

import (
	service "github/sgo-chat/modules/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) RegisterController(c *gin.Context) {
	// Route registration logic goes here
	c.JSON(201, gin.H{"statusCode": 201})
}
