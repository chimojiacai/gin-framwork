package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"study_gin/common/constant"
	"study_gin/common/library/jwt"
	"study_gin/common/library/response"
	"study_gin/params"
)

/*
@Time : 2020-10-14 11:43
@Author : liyongzhen
@File : user
@Software: GoLand
*/

func GetUserInfo(c *gin.Context) {
	//id := c.Query("id")
	//name := c.Query("name")
	//b, s := utils.CheckParams(c, []string{"id", "name", "age"})
	//if !b {
	//	logrus.Errorf("缺少参数:[%s]", s)
	//	response.Error(c, s)
	//	return
	//}
	//
	//response.Success(c, id+name)
	info := params.GetUserInfo{}
	if err := c.ShouldBindQuery(&info); err != nil {
		response.Error(c, err.Error())
		return
	}
	m := make(map[string]interface{})
	m["user_id"] = info.UserId
	m["name"] = info.Name
	response.Success(c, m)
}

func CreateToken(c *gin.Context) {
	id := c.Query("id")
	token, err := jwt.GenerateToken(cast.ToInt64(id), c.GetHeader(constant.HEADER_DEVICE_UUID))
	if err != nil {
		response.Error(c, "创建token失败")
		return
	}
	m := make(map[string]string)
	m["token"] = token
	response.Success(c, m)
}

func Post(ctx *gin.Context) {
	//json := make(map[string]interface{}) //注意该结构接受的内容
	//ctx.Bind(&json)
	//fmt.Println(ctx.Request.PostFormValue("source"))
	fmt.Println(ctx.Request.PostForm)
}
