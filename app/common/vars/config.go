package vars

import jwtpkg "github.com/golang-jwt/jwt"

type EmailConfig struct {
	ServerName string
	Port       int
	From       string
	Password   string
	Second     int
}

type RedisConfig struct {
	Server      string
	Port        int
	RedisPrefix string
	Password    string
}

type MysqlConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	DB       string
}

type CosConfig struct {
	Server    string
	SecretID  string
	SecretKey string
}

// 验证码长度
var Code_length = 6

// 验证码有效时间
var Code_expire = 300

// 默认分页大小
var PageSize = 20

// token 里面包含的用户信息
type UserClaim struct {
	Id             int
	Identity       string
	Name           string
	RefreshTokenId string // refresh_token 已用
	jwtpkg.StandardClaims
}

type TokenConfig struct {
	TokenTime        int
	RefreshTokenTime int
	Secret           string
}

type AesCfg struct {
	Secret string // 加密密钥
	IV     string // 加密初始向量
}
