// Author: yann
// Date: 2020/7/28 10:17 上午
// Desc:

package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	privateKeyBytes, err := ioutil.ReadFile("/Users/yann/Desktop/exchange/src/go-exchange/exchange-admin/resource/key.pem")
	priKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("jwt秘钥文件不正确: " + err.Error())
	}
	PrivateKey = priKey
	PublickKey = &priKey.PublicKey
	token, err := GenerateToken(1, "100")
	if err != nil {
		t.Errorf("token 生成失败 %v", err)
		return
	}
	t.Logf("token : %s", token)

}

func TestPraseToken(t *testing.T) {
	privateKeyBytes, err := ioutil.ReadFile("/Users/yann/Desktop/exchange/src/go-exchange/exchange-admin/resource/key.pem")
	priKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("jwt秘钥文件不正确: " + err.Error())
	}
	PrivateKey = priKey
	PublickKey = &priKey.PublicKey
	token, err := GenerateToken(1, "100")
	if err != nil {
		t.Errorf("token 生成失败 %v", err)
		return
	}
	t.Logf("token : %s", token)
	claims, err := PraseToken(token)
	if err != nil {
		t.Errorf("token 解析失败 %v", err)
		return
	}
	dev, ok := claims["dev"].(string)
	if !ok {
		t.Errorf("token 解析dev失败 %v", err)
		return
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		t.Errorf("token 解析sub失败 %v", err)
		return
	}

	fmt.Printf("uid = %s\ndev = %s\n", sub, dev)
}

func TestCheckReissueAndReissue(t *testing.T) {
	privateKeyBytes, err := ioutil.ReadFile("/Users/yann/Desktop/exchange/src/go-exchange/exchange-admin/resource/key.pem")
	priKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("jwt秘钥文件不正确: " + err.Error())
	}
	PrivateKey = priKey
	PublickKey = &priKey.PublicKey
	token, err := GenerateToken(1, "100")
	if err != nil {
		t.Errorf("token 生成失败 %v", err)
		return
	}
	t.Logf("token : %s", token)
	claims, err := PraseToken(token)
	if err != nil {
		t.Errorf("token 解析失败 %v", err)
		return
	}
	dev, ok := claims["dev"].(string)
	if !ok {
		t.Errorf("token 解析dev失败 %v", err)
		return
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		t.Errorf("token 解析sub失败 %v", err)
		return
	}
	expTime, ok := claims["exp"].(float64)
	if !ok {
		t.Errorf("token 解析exp失败 %v", err)
		return
	}
	fmt.Printf("uid = %s\ndev = %s\nexp=%f\n", sub, dev, expTime)

	fmt.Println(CheckReissueAndReissue(claims))
}

func TestCheckDeviceUuid(t *testing.T) {
	privateKeyBytes, err := ioutil.ReadFile("/Users/yann/Desktop/exchange/src/go-exchange/exchange-admin/resource/key.pem")
	priKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("jwt秘钥文件不正确: " + err.Error())
	}
	PrivateKey = priKey
	PublickKey = &priKey.PublicKey
	token, err := GenerateToken(1, "100")
	if err != nil {
		t.Errorf("token 生成失败 %v", err)
		return
	}
	t.Logf("token : %s", token)
	claims, err := PraseToken(token)
	if err != nil {
		t.Errorf("token 解析失败 %v", err)
		return
	}
	if CheckDeviceUuid(claims, "107") {
		t.Logf("设备id相同")
	} else {
		t.Logf("设备id不同")
	}
}

func TestGetUidInClaims(t *testing.T) {
	privateKeyBytes, err := ioutil.ReadFile("/Users/yann/Desktop/exchange/src/go-exchange/exchange-admin/resource/key.pem")
	priKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("jwt秘钥文件不正确: " + err.Error())
	}
	PrivateKey = priKey
	PublickKey = &priKey.PublicKey
	token, err := GenerateToken(1, "100")
	if err != nil {
		t.Errorf("token 生成失败 %v", err)
		return
	}
	t.Logf("token : %s", token)
	claims, err := PraseToken(token)
	if err != nil {
		t.Errorf("token 解析失败 %v", err)
		return
	}
	uid, err := GetUidByClaims(claims)
	if err != nil {
		t.Errorf("token 解析uid失败, %v", err)
	}
	t.Logf("uid = %d", uid)
}
