package main

/*
@Time : 2020-10-14 11:32
@Author : liyongzhen
@File : main
@Software: GoLand
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"study_gin/common/config"
	"study_gin/common/middleware"
	"study_gin/controller"
)

// 全局变量
var r *gin.Engine

// 初始化
func init() {
	r = gin.New()

	// 加载日志中间件
	r.Use(
		middleware.HandlerLoggerToFile(),  // log
		middleware.HandlerLoadConfig(),    // 配置
		middleware.HandlerException,       // 全局异常,返回json
		//middleware.OtherHeaderInterceptor, // 检测header
		//middleware.HandlerJwt,             // jwt
	)
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
	r.Run(viper.GetString("web.address")) // default listen and serve on 0.0.0.0:8080
}
