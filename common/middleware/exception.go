package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"study_gin/common/constant"
)

/*
@Time : 2020-10-14 16:37
@Author : liyongzhen
@File : exception
@Software: GoLand
*/
// 异常处理
func HandlerException(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			logrus.Printf("panic: %v\n", r)
			debug.PrintStack() // 打印出异常
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			c.JSON(http.StatusOK, gin.H{
				constant.Code: constant.ServiceError,
				constant.Msg:  errorToString(r),
				constant.Data: nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
