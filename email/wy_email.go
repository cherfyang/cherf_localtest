package email

import (
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmailFromWangyi(msg EmailMsg) {
	m := gomail.NewMessage()
	m.SetHeader("From", "mengxi_email@163.com")
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Title)
	m.SetBody("text/html", msg.Content)

	d := gomail.NewDialer("smtp.163.com", 465, "mengxi_email@163.com", "KEjuLaFJkz93Jy2a")
	d.SSL = true // 网易的 465 端口是 SSL 加密的

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

}

type EmailMsg struct {
	To      string
	Title   string
	Content string
}
