package response

/*
@Time : 2020-10-14 16:49
@Author : liyongzhen
@File : response
@Software: GoLand
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_gin/common/constant"
	"study_gin/common/library/utils"
)

// 返回成功json数据
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		constant.Code:      constant.Success,
		constant.Msg:       constant.SuccessMsg,
		constant.Data:      data,
		constant.Timestamp: utils.Time(),
	})
}

// 带异常码返回异常json数据
func ErrorCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		constant.Code:      code,
		constant.Msg:       msg,
		constant.Data:      nil,
		constant.Timestamp: utils.Time(),
	})
}

// 不带异常码返回异常json数据
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		constant.Code:      constant.FrontError,
		constant.Msg:       msg,
		constant.Data:      nil,
		constant.Timestamp: utils.Time(),
	})
}

// token异常
func ErrorToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		constant.Code:      constant.TokenError,
		constant.Msg:       constant.TokenErrorMsg,
		constant.Data:      nil,
		constant.Timestamp: utils.Time(),
	})
}

// 带异常码返回异常json数据
func ErrorTokenNew(c *gin.Context, token string) {
	m := make(map[string]string)
	m["token"] = token
	c.JSON(http.StatusOK, gin.H{
		constant.Code:      constant.TokenErrorNew,
		constant.Msg:       constant.SuccessMsg,
		constant.Data:      m,
		constant.Timestamp: utils.Time(),
	})
}
