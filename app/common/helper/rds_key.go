package helper

func GetMailRegKey(prefix, email string) string {
	return prefix + "email_register" + email
}

func GetRefreshTokenKey(prefix, rt string) string {
	return prefix + "refresh_token:" + rt
}
