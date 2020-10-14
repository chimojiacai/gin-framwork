package jwt

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

var (
	PrivateKey *rsa.PrivateKey = nil
	PublickKey *rsa.PublicKey  = nil
	pubkeyFunc jwt.Keyfunc     = func(token *jwt.Token) (interface{}, error) {
		return PublickKey, nil
	}
)
