package config

import (
	jwtBag "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"io/ioutil"
	"study_gin/common/library/jwt"
)

/*
@Time : 2020-10-14 18:46
@Author : liyongzhen
@File : cert
@Software: GoLand
*/
func CertInit() {
	privateKeyBytes, err := ioutil.ReadFile(viper.GetString("web.certpath"))
	priKey, err := jwtBag.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("jwt秘钥文件不正确: " + err.Error())
	}

	jwt.PrivateKey = priKey
	jwt.PublickKey = &priKey.PublicKey
}