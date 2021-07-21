package app

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

var mail *gomail.Dialer

func SetUpMail() {
	// 定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mail = gomail.NewDialer("smtp.qq.com", 465, Global.Mail.User, Global.Mail.Pass)
	m, err := mail.Dial()
	if err != nil {
		log.Fatal("gomail init fail")
	}
	defer m.Close()
}

func SendMail(mailTo []string, subject string, body string) (err error) {
	m := gomail.NewMessage()
	// 这种方式可以添加别名，即“XD Game”。也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("From", fmt.Sprintf("WeiBo<%s>", Global.Mail.User))
	// 发送给多个用户
	m.SetHeader("To", mailTo...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)

	err = mail.DialAndSend(m)

	return
}
