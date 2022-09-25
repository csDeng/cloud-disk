package user_helper

import (
	"cloud_disk/app/common/vars"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"log"

	"encoding/base64"
	"errors"
)

var aesc *vars.AesCfg
var hasInjectAesCfg = false

func InitAesCfg(in *vars.AesCfg) {
	aesc = in
	hasInjectAesCfg = true
}
func getAesCfg() *vars.AesCfg {
	if !hasInjectAesCfg {
		return nil
	} else if aesc == nil {
		log.Fatal("AesCfgObj nil")
	}
	return aesc
}

func getByte(s string) []byte {
	sum := md5.Sum([]byte(s))
	return sum[:]
}

var secret []byte
var iv []byte

func getAesSecret() []byte {
	var AesCfgObj = getAesCfg()
	if secret == nil {
		secret = getByte(AesCfgObj.Secret)
	}
	return secret
}

func getAesIV() []byte {
	var AesCfgObj = getAesCfg()
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

	// base 64 加密 避免二进制数据无法利用 token 传输
	base64Res := base64.StdEncoding.EncodeToString(cipherText)
	return base64Res, nil
}

func AesDecrypt(cipherText string) (string, error) {
	s := getAesSecret()
	i := getAesIV()
	// fmt.Println("cipherText=>", cipherText)
	// base 64 解码
	base64Res, err := base64.StdEncoding.DecodeString(cipherText)
	// fmt.Println("base64res=> ", base64Res)
	if err != nil {
		return "", errors.New("base64 decode error")
	}

	c, err := aes.NewCipher(s)
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(c, i)
	plainText := make([]byte, len(base64Res))
	decrypter.XORKeyStream(plainText, []byte(base64Res))
	return string(plainText), nil
}
