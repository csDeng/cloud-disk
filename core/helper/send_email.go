package helper

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// 邮箱验证码发送
func SendEmailCode(mail, code string) error {
	emailConfig := EmailConfigObject
	e := email.NewEmail()
	e.From = fmt.Sprintf("<%s>", emailConfig.From)
	e.To = []string{mail}
	e.Subject = "邮箱注册"
	e.HTML = []byte("你的验证码是: " + code)
	err := e.SendWithTLS(
		fmt.Sprintf("%s:%d", emailConfig.ServerName, emailConfig.Port),
		smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.ServerName),
		&tls.Config{InsecureSkipVerify: true, ServerName: emailConfig.ServerName})
	if err != nil {
		log.Printf("邮箱注册: %s, 验证码发送失败, error=%v\r\n", mail, err)
		return err
	}
	return nil
}
