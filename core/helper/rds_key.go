package helper

func GetMailRegKey(email string) string {
	return RedisConfigObject.RedisPrefix + "email_register" + email
}

func GetRefreshTokenKey(rt string) string {
	return RedisConfigObject.RedisPrefix + "refresh_token:" + rt
}
