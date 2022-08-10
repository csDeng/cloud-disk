package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
)

func getByte(s string) []byte {
	sum := md5.Sum([]byte(s))
	return sum[:]
}

var secret []byte
var iv []byte

func getAesSecret() []byte {
	if secret == nil {
		secret = getByte(AesCfgObj.Secret)
	}
	return secret
}

func getAesIV() []byte {
	if iv == nil {
		iv = getByte(AesCfgObj.IV)
	}
	return iv
}

func AesEncrypt(plainText string) (string, error) {
	s := getAesSecret()
	i := getAesIV()

	c, err := aes.NewCipher(s)
	if err != nil {
		return "", err
	}

	// 获取加密块
	cfb := cipher.NewCFBEncrypter(c, i)

	// 加密结果接收的长度需要与明文一样
	cipherText := make([]byte, len(plainText))

	cfb.XORKeyStream(cipherText, []byte(plainText))

	return string(cipherText), nil
}

func AesDecrypt(cipherText string) (string, error) {
	s := getAesSecret()
	i := getAesIV()

	c, err := aes.NewCipher(s)
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(c, i)
	plainText := make([]byte, len(cipherText))
	decrypter.XORKeyStream(plainText, []byte(cipherText))
	return string(plainText), nil
}
