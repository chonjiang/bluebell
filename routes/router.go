package routes

import (
	"bluebell/controllers/demo"
	"bluebell/controllers/user"
	"bluebell/logger"
	"bluebell/tools/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, snowflake.GenIDString())
	})
	r.POST("/signUp", demo.SignUpHandler)
	r.POST("/login", user.LoginHandler)
	r.POST("/register", user.RegisterHandler)
	return r
}
