package appresponse

import (
	"net/http"
)

type successResponseModel struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type errorResponseModel struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Success(code int, data any) (int, any) {
	return code, successResponseModel{
		Message: http.StatusText(code),
		Data:    data,
	}
}

func Error(code int, err error) (int, any) {
	return code, errorResponseModel{
		Message: http.StatusText(code),
		Error:   err.Error(),
	}
}
