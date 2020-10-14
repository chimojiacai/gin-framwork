package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
@Time : 2020-10-14 17:06
@Author : liyongzhen
@File : func
@Software: GoLand
*/

// 时间戳.方便更改毫秒/秒
func Time() int64 {
	return time.Now().Unix()
}

// 自定义检测字段是否传
// 返回,1=是否正确,2=缺少的参数名字
func CheckParams(c *gin.Context, needParams []string) (bool, string) {
	var (
		b bool
		s string
	)

	l := len(needParams)
	if l == 0 {
		b = true
		return b, s
	}

	params := make(map[string][]string)
	method := c.Request.Method
	if method == http.MethodGet {
		// 获取所有参数,k-v
		params = c.Request.URL.Query()
	} else {
		params = c.Request.PostForm
	}

	// 判断是否缺少参数
	for i := 0; i < len(needParams); i++ {
		if _, ok := params[needParams[i]]; !ok {
			return b, needParams[i]
		}
	}
	return true, s
}
