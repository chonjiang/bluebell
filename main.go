package main

// web脚手架
// 1.加载配置
// 2.初始化日志
// 3.初始化db连接
// 4.初始化redis连接
// 5.注册路由
// 6.启动服务（优雅关机）

import (
	_ "bluebell/cmd" // 配置命令行参数
	"bluebell/config"
	_ "bluebell/config" // 加载配置
	"bluebell/database"
	_ "bluebell/logger" // 初始化日志
	_ "bluebell/tools/snowflake"
	//_ "bluebell/database" // 初始化db
	//_ "bluebell/redis" // 初始化redis

	"bluebell/redis"
	"bluebell/routes"
	"bluebell/watch"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := watch.WatchConfigChanged()
	if err != nil {
		panic(fmt.Sprintf("watch config file failed : %v", err))
	}
	defer database.Close()
	defer redis.Close()
	defer zap.L().Sync()
	// 注册路由
	router := routes.Setup()

	// 优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Port),
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过10秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
