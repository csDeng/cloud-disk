package helper

func GetMailRegKey(email string) string {
	return RedisConfigObject.RedisPrefix + "email_register" + email
}
