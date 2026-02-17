package modules

import (
	"github/sgo-chat/filters"
	"github/sgo-chat/internals/configs"
	"github/sgo-chat/modules/controllers"
	"github/sgo-chat/modules/repositories"
	"github/sgo-chat/modules/services"
	"github/sgo-chat/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, cfg *configs.Variable) *gin.Engine {
	r := gin.Default()
	r.Use(filters.GlobalExceptionHandler())

	helper := utils.NewHelperUtils(cfg)

	profileRepository := repositories.NewProfileRepository(db.Collection("profiles"))
	authRepository := repositories.NewAccountRepository(db.Collection("accounts"), profileRepository)

	authService := services.NewAuthService(authRepository, helper)

	authController := controllers.NewAuthController(authService)

	r.POST("/auth/register", authController.RegisterController)
	r.POST("/auth/login", authController.LoginController)

	return r
}
