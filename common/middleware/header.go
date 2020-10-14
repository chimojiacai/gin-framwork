package middleware

import (
	"github.com/gin-gonic/gin"
	"study_gin/common/constant"
	"study_gin/common/library/response"
)

/*
@Time : 2020-10-14 18:28
@Author : liyongzhen
@File : header
@Software: GoLand
*/
func OtherHeaderInterceptor(c *gin.Context) {
	if c.GetHeader(constant.HEADER_LANG) == "" || c.GetHeader(constant.HERDER_VERSION) == "" || c.
		GetHeader(constant.HEADER_USER_AGENT) == "" || c.GetHeader(constant.HEADER_DEVICE_UUID) == "" {
		response.ErrorCode(c, constant.FrontHeaderError, constant.FrontHeaderErrorMsg)
		c.Abort() // 终止代码流程
	}
	c.Next()
}
