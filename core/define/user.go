package define

import jwtpkg "github.com/golang-jwt/jwt"

// token 里面包含的用户信息
type UserClaim struct {
	Id             int
	Identity       string
	Name           string
	RefreshTokenId string // refresh_token 已用
	jwtpkg.StandardClaims
}

type TokenConfig struct {
	TokenTime        int    `ini:"token_time"`
	RefreshTokenTime int    `ini:"refresh_token_time"`
	Secret           string `ini:"secret"`
}
