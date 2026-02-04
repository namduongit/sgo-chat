package controllers

import (
	"github/sgo-chat/internal/config"
	"github/sgo-chat/modules/controllers"
	"github/sgo-chat/modules/repositories"
	service "github/sgo-chat/modules/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	authRepository := repositories.NewAccountRepository(db.Collection("accounts"))
	authService := service.NewAuthService(authRepository, cfg.JWTSecret)

	authController := controllers.NewAuthController(authService)

	r.POST("/auth/register", authController.RegisterController)

	return r
}
