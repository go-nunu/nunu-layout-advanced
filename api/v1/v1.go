package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Response{Code: ErrorCodeMap[ErrSuccess], Message: ErrSuccess.Error(), Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: ErrorCodeMap[err], Message: err.Error(), Data: data}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var ErrorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	ErrorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
