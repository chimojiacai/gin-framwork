package params

import (
	"github.com/go-playground/validator/v10"
)

/*
@Time : 2020-10-15 18:00
@Author : liyongzhen
@File : params_check
@Software: GoLand
*/

// 校验int泛型是否为0
func IntValid(fl validator.FieldLevel) bool {
	if num, ok := fl.Field().Interface().(int64); ok {
		if num <= 0 {
			return false
		}
	}
	return true
}

// 校验string是否为空
func StringValid(fl validator.FieldLevel) bool {
	if str, ok := fl.Field().Interface().(string); ok {
		if str == "" {
			return false
		}
	}
	return true
}
