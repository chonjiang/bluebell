package redis

import (
	"bluebell/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var rdb *redis.Client

func init() {
	cfg := config.Conf.RedisConfig
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("connect redis err: %v", err))
	}
}

// Close 因为全局变量rdb没有对外暴露，所以要封装一个对外暴露的方法Close供主函数使用,这是一个技巧
func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Info("mysql close failed", zap.String("error", err.Error()))
	}
}
