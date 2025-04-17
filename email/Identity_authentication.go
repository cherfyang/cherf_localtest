package email

import (
	_const "cherf_localtest/const"
	"fmt"
)

func OwnerVerify(token string) {
	url := _const.Myurl + "?" + token
	msgcontent := fmt.Sprintf("<!DOCTYPE html>\n<html>\n  <body style=\"font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;\">\n    <div style=\"max-width: 600px; margin: auto; background-color: #ffffff; border-radius: 8px; padding: 30px; box-shadow: 0 2px 8px rgba(0,0,0,0.1);\">\n      <h2 style=\"color: #2E86DE;\">您的 Token 已成功更新</h2>\n      <p>您好，</p>\n      <p>您的认证 Token 已更新，请点击以下链接完成操作：</p>\n\n      <p style=\"text-align: center; margin: 30px 0;\">\n        <a href=\"%s\" target=\"_blank\" style=\"background-color: #2E86DE; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px;\">\n          点击访问\n        </a>\n      </p>\n\n      <p style=\"color: #999;\">如果按钮无法点击，请将以下链接复制到浏览器中打开：</p>\n      <p style=\"word-break: break-all;\"><a href=\"%s\" target=\"_blank\">%s</a></p>\n\n      <p>如非您本人操作，请忽略此邮件。</p>\n\n      <p style=\"margin-top: 40px; color: #888;\">—— 来自 mengxiy</p>\n    </div>\n  </body>\n</html>\n", url, url, url)
	msg := EmailMsg{
		Title:   "安全认证",
		To:      "2637206496@qq.com",
		Content: msgcontent,
	}
	SendEmailFromWangyi(msg)
}
