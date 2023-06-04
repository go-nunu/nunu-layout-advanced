package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: 0, Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: code, Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}
