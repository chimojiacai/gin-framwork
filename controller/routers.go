package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"study_gin/params"
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
	// 参数校验
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("intValid", params.IntValid) // tag要小写
		//_ = v.RegisterValidation("stringValid", params.StringValid) // tag要小写
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	user := r.Group("/user")
	user.GET("/get_user", GetUserInfo)
	user.GET("/create_token", CreateToken)
	user.POST("/test", Post)
}
