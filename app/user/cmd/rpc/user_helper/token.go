package user_helper

import (
	"cloud_disk/app/common/helper"
	"cloud_disk/app/common/vars"
	"log"
	"time"

	jwtpkg "github.com/golang-jwt/jwt"
)

var tokenCfg *vars.TokenConfig
var hasInjectToken = false

func InitTokenCfg(in *vars.TokenConfig) {
	tokenCfg = in
	hasInjectToken = true
}

func getTokenCfg() *vars.TokenConfig {
	if !hasInjectToken {
		return nil
	} else if tokenCfg == nil {
		log.Fatal("please inject tokenCfg first")
	}
	return tokenCfg
}

// 生成token
func GenerateToken(id int, identity string, name string, isRefreshToken bool) (string, error) {
	TokenConfigObject := getTokenCfg()
	t := 0
	if isRefreshToken {
		t = TokenConfigObject.RefreshTokenTime
	} else {
		t = TokenConfigObject.TokenTime
	}
	ex := time.Now().Add(time.Minute * time.Duration(t)).Unix()
	uc := vars.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		RefreshTokenId: helper.GenerateUuid(),
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
func ParseToken(token string) (*vars.UserClaim, error) {
	TokenConfigObject := getTokenCfg()
	plain, err := AesDecrypt(token)
	if err != nil {
		return nil, err
	}
	uc := new(vars.UserClaim)
	claims, err := jwtpkg.ParseWithClaims(plain, uc, func(t *jwtpkg.Token) (interface{}, error) {
		return []byte(TokenConfigObject.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, vars.ErrTokenInvalid
	}
	return uc, nil
}
