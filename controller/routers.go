package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@Time : 2020-10-14 11:41
@Author : liyongzhen
@File : routers
@Software: GoLand
路由
*/

// 初始化路由
func InitRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	user := r.Group("/user")
	user.GET("/get_user", GetUserInfo)
	user.GET("/create_token", CreateToken)
}
