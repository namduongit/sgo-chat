package errors

import (
	"errors"
	"github/sgo-chat/internals/configs/httpres"

	"github.com/go-playground/validator/v10"
)

/*
* Can use struct in cases:
- Valid request
- Validation form
- Duplicate resource
*/
type RestError struct {
	Code    httpres.StatusCode
	Message httpres.MessageResponse

	Err error
}

func (e *RestError) Error() string {
	return "Message error in MessageError() method"
}

func (e *RestError) MessageError() any {
	// Body is null
	if e.Err.Error() == "EOF" {
		message := "Missing request body"
		return message
	}

	// Validation form
	if errs, ok := e.Err.(validator.ValidationErrors); ok {
		var messages []string
		for _, err := range errs {
			messages = append(messages, err.Field()+" "+err.Tag())
		}
		return messages
	}

	return e.Err.Error()
}

/* Create RestError */
func BadRequestError(err string) *RestError {
	return &RestError{
		Code:    httpres.StatusBadRequest,
		Message: httpres.BadRequest,
		Err:     errors.New(err),
	}
}
func UnauthorizedError(err string) *RestError {
	return &RestError{
		Code:    httpres.StatusUnauthorized,
		Message: httpres.Unauthorized,
		Err:     errors.New(err),
	}
}
func ForbiddenError(err string) *RestError {
	return &RestError{
		Code:    httpres.StatusForbidden,
		Message: httpres.Forbidden,
		Err:     errors.New(err),
	}
}
func NotFoundError(err string) *RestError {
	return &RestError{
		Code:    httpres.StatusNotFound,
		Message: httpres.NotFound,
		Err:     errors.New(err),
	}
}
