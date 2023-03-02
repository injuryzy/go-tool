package jwtutil

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	ID       uint
	Username string
	jwt.StandardClaims
}

// 生成token
func GenerateToken(claims MyCustomClaims, key interface{}) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	return token, err
}

// 解析token
func ParseToken(token string, verifKey interface{}) (*MyCustomClaims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifKey, nil
	})
	if tokenClaims != nil {
		claims := tokenClaims.Claims.(*MyCustomClaims)
		return claims, nil
	}
	return nil, err
}
