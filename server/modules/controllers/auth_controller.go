package controllers

import (
	"github/sgo-chat/internals/configs"
	"github/sgo-chat/modules/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) RegisterController(ctx *gin.Context) {
	result, err := ac.authService.Register(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := configs.SuccessResponse(result)
	ctx.JSON(int(res.StatusCode), &res)
}
