package user

import (
	"bluebell/config"
	"bluebell/logic"
	"bluebell/middlewares"
	"bluebell/models"
	"bluebell/tools"
	v "bluebell/tools/validator"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// RegisterHandler 注册用户
func RegisterHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	params := new(RegisterParam)
	if err := c.ShouldBindJSON(&params); err != nil {
		// 若非validator.ValidationErrors类型错误直接返回
		var msg = err.Error()
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			// 若为validator.ValidationErrors类型错误则进行翻译
			m := v.RemoveTopStruct(errs.Translate(v.Trans))
			dataType, _ := json.Marshal(m)
			msg = string(dataType)
		}
		tools.ResponseErrorWithMsg(c, config.CodeInvalidParams, msg)
		return
	}

	user := new(models.User)
	user.Username = params.Name
	user.Password = params.Password
	user.Gender = params.Gender
	user.Email = params.Email
	err := logic.Register(user)

	if err != nil {
		zap.L().Error("register failed", zap.Error(err))
		code := config.CodeServerBusy
		if errors.Is(err, config.ErrorUserExist) {
			code = config.CodeUserExist
		}
		tools.ResponseErrorWithMsg(c, code, err.Error())
		return
	}

	tools.ResponseSuccess(c, user)

	return
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	params := new(LoginParam)
	if err := c.ShouldBindJSON(&params); err != nil {
		// 若非validator.ValidationErrors类型错误直接返回
		var msg = err.Error()
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			// 若为validator.ValidationErrors类型错误则进行翻译
			m := v.RemoveTopStruct(errs.Translate(v.Trans))
			dataType, _ := json.Marshal(m)
			msg = string(dataType)
		}
		tools.ResponseErrorWithMsg(c, config.CodeInvalidParams, msg)
		return
	}

	user := new(models.User)
	user.Username = params.Name
	user.Password = params.Password
	loginUser, err := logic.Login(user)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		tools.ResponseErrorWithMsg(c, config.CodeInvalidPassword, err.Error())
		return
	}
	user = loginUser
	data, _ := json.Marshal(user)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)

	tools.ResponseSuccess(c, m)
	return
}

// GetLoginInfo 获取登录用户信息
func GetLoginInfo(c *gin.Context) {
	_userName, ok := c.Get(middlewares.ContextUserNameKey)
	if !ok {
		tools.ResponseErrorWithMsg(c, config.CodeNotLogin, config.CodeNotLogin.Msg())
		return
	}
	userName, ok := _userName.(string)
	if !ok{
		tools.ResponseErrorWithMsg(c, config.CodeNotLogin, config.CodeNotLogin.Msg())
		return
	}

	_userID, ok := c.Get(middlewares.ContextUserIDKey)
	if !ok {
		tools.ResponseErrorWithMsg(c, config.CodeNotLogin, config.CodeNotLogin.Msg())
		return
	}
	userID, ok := _userID.(int64)
	if !ok{
		tools.ResponseErrorWithMsg(c, config.CodeNotLogin, config.CodeNotLogin.Msg())
		return
	}

	user := &models.User{
		UserId: userID,
		Username: userName,
	}

	userInfo, err := logic.GetUserInfo(user)
	if err != nil {
		tools.ResponseErrorWithMsg(c, config.CodeUserExist, err.Error())
		return
	}

	tools.ResponseSuccess(c, userInfo)
	return
}