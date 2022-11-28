package demo

import (
	v "bluebell/tools/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	params := new(SignUpParam)
	if err := c.ShouldBindJSON(&params); err != nil {
		var msg interface{}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			// validator.ValidationErrors类型错误则进行翻译
			msg = v.RemoveTopStruct(errs.Translate(v.Trans))
		} else {
			// 非validator.ValidationErrors类型错误直接返回
			msg = err.Error()
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	return
}
