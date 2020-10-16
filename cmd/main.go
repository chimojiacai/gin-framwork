package main

/*
@Time : 2020-10-14 11:32
@Author : liyongzhen
@File : main
@Software: GoLand
*/

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"study_gin/common/config"
	"study_gin/common/middleware"
	"study_gin/controller"
	"syscall"
	"time"
)

// 全局变量
var r *gin.Engine
var srv *http.Server

// 初始化
func init() {
	r = gin.New() // 不用自带的日志系统
	// 要是使用gin自带的日志系统, 使用 r = gin.Default()

	// 加载日志中间件
	r.Use(
		middleware.HandlerLoggerToFile(), // log
		middleware.HandlerLoadConfig(),   // 配置
		middleware.HandlerException,      // 全局异常,返回json
		//middleware.OtherHeaderInterceptor, // 检测header
		//middleware.HandlerJwt,             // jwt
	)
	gin.SetMode(gin.ReleaseMode) // 不实用gin打印控制台

	// 初始化私钥证书
	config.CertInit()
	// 初始化mysql
	config.MysqlInit() // gorm
	// 初始化redis
	config.RedisInit()

}

func main() {
	// 初始化路由
	controller.InitRoutes(r)
	srv = &http.Server{
		Addr:    viper.GetString("web.address"),
		Handler: r,
	}

	//r.Run(viper.GetString("web.address")) // default listen and serve on 0.0.0.0:8080
	// 启动http请求
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("app启动失败:[%v]", err)
			panic(err)
		}
	}()

	// 退出程序
	shutdown()
}

// 优雅的关闭
func shutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGQUIT)
	<-quit
	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}
