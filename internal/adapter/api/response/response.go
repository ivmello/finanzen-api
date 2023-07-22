package response

import (
	"github.com/ivmello/finanzen-api/internal/adapter/api/exception"
)

type ControllerResponse struct {
	Code   int                      `json:"code"`
	Sucess bool                     `json:"success"`
	Data   interface{}              `json:"data,omitempty"`
	Errors *exception.ErrorResponse `json:"errors,omitempty"`
}

func NewResponse(code int, data interface{}, errorsResponse *exception.ErrorResponse) *ControllerResponse {
	return &ControllerResponse{
		Code:   code,
		Sucess: code >= 200 && code <= 299,
		Data:   data,
		Errors: errorsResponse,
	}
}
