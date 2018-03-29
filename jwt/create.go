package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

//创建token
func CreateToken(claims map[string]interface{}, secret []byte) (string, error) {
	//校验claims的类型
	c := jwt.MapClaims(claims)
	//校验claims的有效性
	err := c.Valid()
	if err != nil {
		return "", errors.New("claims验证失败：" + err.Error())
	}
	//创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//根据秘钥生成token字符串
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("token创建失败：" + err.Error())
	}
	return tokenStr, nil
}
