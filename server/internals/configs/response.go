package configs

import (
	"github/sgo-chat/internals/configs/httpres"
)

type RestResponse struct {
	StatusCode httpres.StatusCode      `json:"statusCode"`
	Message    httpres.MessageResponse `json:"message"`
	Error      any                     `json:"error"`
	Data       any                     `json:"data"`
}

func SuccessResponse(data any) *RestResponse {
	return &RestResponse{
		StatusCode: httpres.StatusOK,
		Message:    httpres.Success,
		Error:      nil,
		Data:       data,
	}
}

func CreatedResponse(data any) *RestResponse {
	return &RestResponse{
		StatusCode: httpres.StatusCreated,
		Message:    httpres.Success,
		Error:      nil,
		Data:       data,
	}
}
