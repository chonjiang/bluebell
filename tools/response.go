package tools

import (
	"bluebell/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    config.MyCode `json:"code"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
}

func ResponseError(ctx *gin.Context, c config.MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code config.MyCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    config.CodeSuccess,
		Message: config.CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
