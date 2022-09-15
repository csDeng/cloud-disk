package helper

import (
	"core/app/common/vars"
	"core/core/define"
	"errors"
	"time"

	jwtpkg "github.com/golang-jwt/jwt"
)

var TokenConfigObject = getTokenCfg()
var tokenCfg *vars.TokenConfig

func InjectTokenCfg(in *vars.TokenConfig) {
	tokenCfg = in
}

func getTokenCfg() *vars.TokenConfig {

	return tokenCfg
}

// 生成token
func GenerateToken(id int, identity string, name string, isRefreshToken bool) (string, error) {
	t := 0
	if isRefreshToken {
		t = TokenConfigObject.RefreshTokenTime
	} else {
		t = TokenConfigObject.TokenTime
	}
	ex := time.Now().Add(time.Minute * time.Duration(t)).Unix()
	uc := define.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		RefreshTokenId: GenerateUuid(),
		StandardClaims: jwtpkg.StandardClaims{
			ExpiresAt: ex,
		},
	}

	// 使用特定的 签名算法加密
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, uc)

	// fmt.Println("TokenConfigObject.Secret=", TokenConfigObject.Secret)
	// 创建并返回一个完整的,签署的JWT
	tt, err := token.SignedString([]byte(TokenConfigObject.Secret))
	if err != nil {
		return "", err
	}

	res, err := AesEncrypt(tt)
	if err != nil {
		return "", err
	}
	return res, err
}

// 解析token
func ParseToken(token string) (*define.UserClaim, error) {
	plain, err := AesDecrypt(token)
	if err != nil {
		return nil, err
	}
	uc := new(define.UserClaim)
	claims, err := jwtpkg.ParseWithClaims(plain, uc, func(t *jwtpkg.Token) (interface{}, error) {
		return []byte(TokenConfigObject.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("token is invalid")
	}
	return uc, nil
}
