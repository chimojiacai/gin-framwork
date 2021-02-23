package middleware

/*
@Time : 2020-10-14 14:39
@Author : liyongzhen
@File : viper
@Software: GoLand
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func HandlerLoadConfig(conf string) gin.HandlerFunc {
	// 获取配置文件
	viper.SetConfigName(conf)
	viper.SetConfigType("toml") // 文件类型
	viper.AddConfigPath(".")    // 搜索路径
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("read config failed: %v", err)
	}
	// 热加载配置
	viper.WatchConfig()
	return func(context *gin.Context) {

	}
}
