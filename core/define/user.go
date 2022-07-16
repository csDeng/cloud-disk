package define

import "github.com/golang-jwt/jwt/v4"

// token 里面包含的用户信息
type UserClaim struct {
	Id       int
	Identify string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk"
