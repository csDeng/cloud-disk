package define

import "github.com/golang-jwt/jwt/v4"

// token 里面包含的用户信息
type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

type TokenConfig struct {
	TokenTime        int    `ini:"token_time"`
	RefreshTokenTime int    `ini:"refresh_token_time"`
	Secret           string `ini:"secret"`
}
