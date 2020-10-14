package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const reissueTime = time.Hour * 24 * 5 //还有5天过期直接重新签发

func GenerateToken(uid int64, devid string) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, newClaims(uid, devid))
	return tokenClaims.SignedString(PrivateKey)
}

//***************************************************
//Description : 解析token
//param :       token字符串
//return :      jwt.MapClaims: 有值或者为nil, 如果错误则为nil
//              error: 错误信息
//***************************************************
func PraseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, pubkeyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenIllegal
	}
	return claims, nil

}

//***************************************************
//Description : 根据原有Token Claims判断是否要重新签发, 如有需要则重新生成token
//param :       jwt claims
//return :      token: 如果重新签发,则有值, 如果报错或者时间未到则为nil
//              error: 错误信息
//***************************************************
func CheckReissueAndReissue(claims jwt.MapClaims) (string, error) {
	expTime, ok := claims["exp"].(float64)
	if !ok {
		return "", ErrTokenIllegal
	}
	if time.Now().Add(reissueTime).Unix() < int64(expTime) {
		return "", nil
	}
	id, ok := claims["sub"].(string)
	if !ok {
		return "", ErrTokenIllegal
	}
	dev, ok := claims["dev"].(string)
	if !ok {
		return "", ErrTokenIllegal
	}
	idStr, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}
	return GenerateToken(idStr, dev)
}

//***************************************************
//Description : 校验头信息中的设备id是否跟token中一样
//param :       jwt claims
//param :       设备id
//return :      ture=一样  false=不一样
//***************************************************
func CheckDeviceUuid(claims jwt.MapClaims, devId string) bool {
	dev, ok := claims["dev"].(string)
	if !ok {
		return false
	}
	return dev == devId
}

//***************************************************
//Description : 从token中获取uid
//param :       token claims
//return :      uid, error
//***************************************************
func GetUidByClaims(claims jwt.MapClaims) (int64, error) {
	sub, ok := claims["sub"].(string)
	if !ok {
		return 0, ErrTokenIllegal
	}
	uid, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return 0, ErrTokenIllegal
	}
	return uid, nil
}
