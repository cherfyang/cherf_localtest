package file_handle

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func SendFile(Url, filepath string) error {
	url := Url
	filePath := filepath
	// 创建一个缓冲区存储 multipart 数据
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// 创建 multipart 文件字段
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return err
	}

	// 复制文件内容到 multipart writer
	if _, err := io.Copy(part, file); err != nil {
		fmt.Println("Error copying file data:", err)
		return err
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 创建不验证 SSL 证书的 HTTP 客户端（类似 curl -k）
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	// 输出服务器响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
	return nil
}
func SendFileHandle(c *gin.Context) {
	url := c.Request.Header.Get("send_to_url")
	filePath := c.Request.Header.Get("file_path")
	//保存发送记录
	sendRecord(filePath)
	SendFile(url, filePath)

}
func sendRecord(filePath string) {
	// 要写入的内容
	txtname := time.Now().Format("20060102150405") + ".txt"
	println(txtname)
	pcName, _ := os.Hostname()
	input := fmt.Sprintf("From:%s\nPcPath:%s", pcName, filePath)
	txtfile := "SendRecord/" + txtname
	// 直接写入（如果文件不存在，则创建；如果存在，则覆盖）
	err := os.WriteFile(txtfile, []byte(input), 0644)
	if err != nil {
		fmt.Println(" 写入文件失败:", err)
		return
	}
}
