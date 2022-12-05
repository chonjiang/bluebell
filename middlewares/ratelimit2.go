package middlewares

import (
	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit"
	"time"
)

// RateLimitMiddleware2 基于令牌桶的限流（令牌桶按固定的速率往桶里放入令牌，并且只要能从桶里取出令牌就能通过，这种方式支持突发流量的快速处理）
func RateLimitMiddleware2(fillInterval time.Duration, cap int64) func(c *gin.Context){
	bucket := ratelimit2.NewBucket(fillInterval, cap)
	return func(c *gin.Context){
		//if bucket.TakeAvailable(1) < 1 {
		//	tools.ResponseErrorWithMsg(c, config.CodeRateLimit, config.CodeRateLimit.Msg())
		//	c.Abort()
		//	return
		//}
		bucket.Wait(1)
		c.Next()
		return
	}
}