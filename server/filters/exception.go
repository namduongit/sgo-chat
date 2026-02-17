package filters

import (
	"github/sgo-chat/internals/configs"
	"github/sgo-chat/internals/configs/errors"
	"github/sgo-chat/internals/configs/httpres"
	"log"

	"github.com/gin-gonic/gin"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		log.Println("[MIDDLEWARE] Global exception handler")
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			switch e := err.(type) {
			case *errors.RestError:
				log.Print("[Rest error] : ", err.Error())

				ctx.JSON(int(e.Code), configs.RestResponse{
					StatusCode: e.Code,
					Message:    e.Message,
					Error:      e.MessageError(),
					Data:       nil,
				})

			case *errors.MongoError:
				log.Print("[Mongo error] : ", err.Error())

				ctx.JSON(int(httpres.StatusInternalServerError), configs.RestResponse{
					StatusCode: httpres.StatusInternalServerError,
					Message:    httpres.InternalServerError,
					Error:      nil,
					Data:       nil,
				})

			default:
				log.Print("[Unknown error] : ", err.Error())
				ctx.JSON(int(httpres.StatusBadRequest), configs.RestResponse{
					StatusCode: httpres.StatusBadRequest,
					Message:    httpres.BadRequest,
					Error:      nil,
					Data:       nil,
				})
			}
		}
	}
}
