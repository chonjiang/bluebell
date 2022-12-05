package middlewares

import (
	"bluebell/config"
	"bluebell/tools"
	"github.com/gin-gonic/gin"
	ratelimit1 "go.uber.org/ratelimit"
	"time"
)

// RateLimitMiddleware1 基于漏桶的限速中间件（以相同固定速率的方式完成rps限制）
func RateLimitMiddleware1() func(c *gin.Context){
	rl := ratelimit1.New(50) // 传入：每秒请求的次数（rps）
	return func(c *gin.Context){
		// rl.Take() 取水滴，返回上一滴水滴下来（本次水滴可用）的时间
		// 定时阻塞，返回上一次请求阻塞完成后（本次请求可通过）的时间点
		availableTimestamp := rl.Take()
		if time.Now().Sub(availableTimestamp) < 0 {
			// 若此时还没到可取水滴的时间点，则请求不通过
			tools.ResponseErrorWithMsg(c, config.CodeRateLimit, config.CodeRateLimit.Msg())
			c.Abort()
			return
		}

		// 请求通过
		c.Next()
		return
	}
}