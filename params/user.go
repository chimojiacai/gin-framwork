package params

/*
@Time : 2020-10-15 17:44
@Author : liyongzhen
@File : user
@Software: GoLand
*/
type GetUserInfo struct {
	UserId int64  `form:"user_id" binding:"required,intValid"`
	Name   string `form:"name" binding:"required,-"`
}
