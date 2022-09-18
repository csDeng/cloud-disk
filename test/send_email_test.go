package test

// import (
// 	"core/core/helper"
// 	"crypto/tls"
// 	"fmt"
// 	"net/smtp"
// 	"testing"

// 	"github.com/jordan-wright/email"
// )

// func TestSendEmail(t *testing.T) {
// 	emailConfig := helper.EmailConfigObject
// 	e := email.NewEmail()
// 	e.From = fmt.Sprintf("<%s>", emailConfig.From)

// 	e.To = []string{"553537528@qq.com"}
// 	// 秘密抄送
// 	// e.Bcc = []string{"test_bcc@example.com"}

// 	// 抄送
// 	// e.Cc = []string{"test_cc@example.com"}
// 	e.Subject = "Awesome Subject"
// 	e.Text = []byte("Text Body is, of course, supported!")
// 	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")

// 	err := e.SendWithTLS(
// 		fmt.Sprintf("%s:%d", emailConfig.ServerName, emailConfig.Port),
// 		smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.ServerName),
// 		&tls.Config{InsecureSkipVerify: true, ServerName: emailConfig.ServerName})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// }
