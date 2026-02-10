package utils

import (
	"github/sgo-chat/internals/configs/errors"
	"github/sgo-chat/internals/configs/httpres"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ShouldBindReq(ctx *gin.Context, req interface{}) error {
	var result error
	result = validReq(ctx, req)
	if result != nil {
		return result
	}
	result = validForm(req)
	if result != nil {
		return result
	}
	return nil
}

func validReq(ctx *gin.Context, req interface{}) error {
	contentType := ctx.ContentType()
	switch contentType {
	case "application/json":
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			return &errors.RestError{
				Code:    httpres.StatusBadRequest,
				Message: httpres.BadRequest,
				Err:     err,
			}
		}
		return nil
	default:
		return errors.BadRequestError("Invalid request type")
	}
}

func validForm(req interface{}) error {
	if err := validator.New().Struct(req); err != nil {
		return &errors.RestError{
			Code:    httpres.StatusBadRequest,
			Message: httpres.BadRequest,
			Err:     err,
		}
	}
	return nil
}
