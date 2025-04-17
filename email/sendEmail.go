package email

import (
	"gopkg.in/gomail.v2"
	"log"
)

func sendEmail(msg EmailMsg) {
	// 构建邮件
	m := gomail.NewMessage()

	// 发件人（必须和下面 SMTP 登录账号一致）
	m.SetHeader("From", "mengxiy081@gmail.com")

	// 收件人
	m.SetHeader("To", msg.To)

	// 邮件标题
	m.SetHeader("Subject", msg.Title)

	// 邮件正文（HTML 或纯文本）
	m.SetBody("text/html", msg.Body)
	//m.SetBody("text/html", "<h1>Hello</h1><p>这是一封来自 Go 的测试邮件。</p>")

	// 设置 Gmail SMTP
	d := gomail.NewDialer("smtp.gmail.com", 587, "mengxiy081@gmail.com", "cgaq kmcn octn dll v")

	// 发送
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("发送失败: %v", err)
	}

	log.Println("发送成功")
}

type EmailMsg struct {
	Title string
	Body  string
	To    string
}
