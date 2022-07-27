package helper

import (
	"core/core/define"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// 生成token
func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}

	// 使用特定的 签名算法加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	// 创建并返回一个完整的,签署的JWT
	t, err := token.SignedString([]byte(TokenConfigObject.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// 解析token
func ParseToken(token string) (uc *define.UserClaim, err error) {
	uc = new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(TokenConfigObject.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}
