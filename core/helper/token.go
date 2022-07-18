package helper

import (
	"core/core/define"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int, identify string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identify: identify,
		Name:     name,
	}

	// 使用特定的 签名算法加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	// 创建并返回一个完整的,签署的JWT
	t, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

// func ParseToken(token string) (uc *define.UserClaim) {
// 	// jwt.
// }
