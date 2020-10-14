package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type Claims struct {
	DeviceUuid string `json:"dev,omitempty"`
	jwt.StandardClaims
}

const expTime = time.Hour * 24 * 60

//***************************************************
//Description : 生成需要封装到token中的claims
//param :       用户id
//param :		多少秒后过期
//return :		claims
//***************************************************
func newClaims(id int64, devId string) *Claims {
	return &Claims{
		DeviceUuid: devId,
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.FormatInt(id, 10),
			ExpiresAt: time.Now().Add(expTime).Unix(),
			Issuer:    "study gin",
		},
	}
}
