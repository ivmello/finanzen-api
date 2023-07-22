package exception

import "net/http"

type ErrorResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Causes  []*ErrorCauses `json:"causes"`
}

type ErrorCauses struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func NewException(message string, code int, causes []*ErrorCauses) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ErrorResponse {
	return NewException(message, http.StatusInternalServerError, nil)
}

func NewBadGatewayError(message string) *ErrorResponse {
	return NewException(message, http.StatusBadGateway, nil)
}

func NewBadRequestError(message string, causes []*ErrorCauses) *ErrorResponse {
	return NewException(message, http.StatusBadRequest, causes)
}
