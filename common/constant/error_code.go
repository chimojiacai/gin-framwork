package constant

/*
@Time : 2020-10-14 16:44
@Author : liyongzhen
@File : error_code
@Software: GoLand
*/

const (
	Success             = 100200         // 成功
	TokenError          = 100300         // token过期或者不对
	TokenErrorNew       = 100301         // 签发新token
	TokenErrorMsg       = "token过期或不对"   // token过期或者不对
	FrontError          = 100400         // 传来参数异常
	FrontHeaderError    = 100401         // header不对
	FrontHeaderErrorMsg = "header wrong" // header不对
	SuccessMsg          = "成功"           // 成功
	ServiceError        = 100500         // 服务器异常
)

// 公共字段
const (
	Code      = "code"
	Msg       = "msg"
	Data      = "data"
	Timestamp = "timestamp"
)
