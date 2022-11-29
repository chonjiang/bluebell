package routes

import (
	"bluebell/controllers/auth"
	"bluebell/controllers/demo"
	"bluebell/controllers/user"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/auth", auth.AuthHandler)
	r.POST("/signUp", demo.SignUpHandler)
	r.POST("/login", user.LoginHandler)
	r.POST("/register", user.RegisterHandler)

	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/userinfo", user.GetLoginInfo)

	return r
}
