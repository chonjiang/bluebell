package auth

import (
	"bluebell/config"
	"bluebell/controllers/user"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/tools"
	"bluebell/tools/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthHandler(c *gin.Context){
	params := new(user.LoginParam)
	if err := c.ShouldBindJSON(&params); err != nil {
		tools.ResponseErrorWithMsg(c, config.CodeInvalidParams, config.CodeInvalidParams.Msg())
		return
	}

	u := new(models.User)
	u.Username = params.Name
	u.Password = params.Password
	loginUser, err := logic.Login(u)
	if err != nil {
		zap.L().Error("auth failed, err:", zap.Error(err))
		tools.ResponseErrorWithMsg(c, config.CodeInvalidPassword, "鉴权未通过")
		return
	}
	u = loginUser
	token, err := jwt.GenToken(u.Username, u.UserId)
	if err != nil {
		tools.ResponseErrorWithMsg(c, config.CodeInvalidAuthFormat, err.Error())
		return
	}
	fmt.Println("test---", token)
	tools.ResponseSuccess(c, gin.H{"auth": "pass", "token": token})
	return
}