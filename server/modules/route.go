package modules

import (
	"github/sgo-chat/internals/configs"
	"github/sgo-chat/middlewares"
	"github/sgo-chat/modules/controllers"
	"github/sgo-chat/modules/repositories"
	"github/sgo-chat/modules/services"
	"github/sgo-chat/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, cfg *configs.Variable) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.GlobalExceptionHandler())

	profileRepository := repositories.NewProfileRepository(db.Collection("profiles"))
	authRepository := repositories.NewAccountRepository(db.Collection("accounts"), &utils.Bcrypt{Cost: cfg.Crypto.Cost}, profileRepository)

	authService := services.NewAuthService(authRepository, cfg.JWTSecret)

	authController := controllers.NewAuthController(authService)

	r.POST("/auth/register", authController.RegisterController)

	return r
}
