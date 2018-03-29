package jwt

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

//解析token
func ParseToken(tokenStr string, secret []byte) (map[string]interface{}, error) {
	//校验签名方法
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("未知的签名方法: %v", token.Header["alg"])
		}
		return secret, nil
	})

	//校验token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims类型无效")
	}
	if !token.Valid {
		return nil, errors.New("token验证失败")
	}
	return claims, err
}
