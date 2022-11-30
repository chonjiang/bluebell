package auth

import (
	"bluebell/config"
	"bluebell/controllers/user"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/redis"
	"bluebell/tools"
	"bluebell/tools/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// AuthHandler 鉴权登录
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
	tools.ResponseSuccess(c, gin.H{"auth": "pass", "token": token})
	return
}


// SinglePointAuthHandler 单点登录，redis存储
func SinglePointAuthHandler(c *gin.Context){
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
	// 使用redis存储uid和token的唯一对应关系
	rdb := redis.GetClient()
	rdb.Set(c, strconv.FormatInt(u.UserId, 10), token, 10*time.Minute)
	tools.ResponseSuccess(c, gin.H{"auth": "pass", "token": token})
	return
}

