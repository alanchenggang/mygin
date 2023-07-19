package core

import "net/http"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(data any) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func Fail(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
