package helper

func GetMailRegKey(email string) string {
	return RedisConfigObject.RedisPrefix + "email_register" + email
}

func GetRefreshTokenKey(userIdentity string) string {
	return RedisConfigObject.RedisPrefix + "refresh_token:" + userIdentity
}

func GetTokenKey(token string) string {
	return RedisConfigObject.RedisPrefix + "token:" + token
}
