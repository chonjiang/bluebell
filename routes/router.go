package routes

import (
	"bluebell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "ok")
	})

	return r
}
