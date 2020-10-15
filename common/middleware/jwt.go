package middleware

import (
	"github.com/gin-gonic/gin"
	"study_gin/common/constant"
	"study_gin/common/library/jwt"
	"study_gin/common/library/response"
)

/*
@Time : 2020-10-14 18:14
@Author : liyongzhen
@File : jwt
@Software: GoLand
*/
type PassPATH map[string]interface{}

//token白名单
var authPass = PassPATH{
	"/ping":              nil,
	"/user/create_token": nil,
}

//Description : 判断请求路径是否需要放行
//param : *restful.Request
//return : true 放行, false校验token
func (p PassPATH) isPass(c *gin.Context) bool {
	_, ok := p[c.Request.URL.Path]
	return ok
}

// 验证token
func HandlerJwt(c *gin.Context) {
	//白名单放行
	if authPass.isPass(c) {
		c.Next()
		return
	}
	//校验token  如果要校验其他信息, 可以放在外面校验, 需求不一样
	claims, err := jwt.PraseToken(c.GetHeader(constant.HEADER_TOKEN))
	if err != nil {
		response.ErrorToken(c)
		c.Abort()
		return
	}

	//如有需要可以判断是否需要重新签发token
	reissue, err := jwt.CheckReissueAndReissue(claims)
	if err != nil {
		response.ErrorToken(c)
		c.Abort()
		return
	}
	if reissue != "" {
		response.ErrorTokenNew(c, reissue)
		c.Abort()
		return
	}

	//如有需要可以判断设备id是否相同
	if !jwt.CheckDeviceUuid(claims, c.GetHeader(constant.HEADER_DEVICE_UUID)) {
		response.ErrorToken(c)
		c.Abort()
		return
	}

	//获取uid
	uid, err := jwt.GetUidByClaims(claims)
	if err != nil {
		response.ErrorToken(c)
		c.Abort()
		return
	}

	c.Set("uid", uid)
	// 校验成功，继续
	c.Next()
}
