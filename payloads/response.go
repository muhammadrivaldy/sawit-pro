package payloads

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseOK is a function for send result to client
func ResponseOK(ctx echo.Context, code int, obj interface{}) {
	ctx.JSON(code, Response{
		Code:    code,
		Message: "success",
		Data:    obj,
	})
}

// ResponseError is a function for send result error to client
func ResponseError(ctx echo.Context, code int, msg error, obj interface{}) {
	ctx.JSON(code, Response{
		Code:    code,
		Message: msg.Error(),
		Data:    obj,
	})
}
