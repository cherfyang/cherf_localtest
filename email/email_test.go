package email

import "testing"

func Test(t *testing.T) {
	msgcontent := `<!DOCTYPE html>
<html>
  <body style="font-family:Arial, sans-serif; color:#333;">
    <h2>您好，</h2>
    <p>您正在进行身份验证操作。</p>
    <p>您的验证码是：</p>
    <h1 style="color:#2E86DE;">{{.Code}}</h1>
    <p>请在 <strong>5 分钟</strong> 内完成验证。</p>
    <p>如果您没有进行此操作，请忽略此邮件。</p>
    <br/>
    <p>—— {{.ProductName}} 团队</p>
  </body>
</html>
`
	msg := EmailMsg{
		Title:   "欢迎使用",
		To:      "cherf.yang@bindo.com",
		Content: msgcontent,
	}
	SendEmailFromWangyi(msg)
}
func Test2(t *testing.T) {
	OwnerVerify("8rhdafiuhsdbakiduH")
}
