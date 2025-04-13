package util

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetFullpathByParam(name string) string {
	path := "D:/UpdownFromHttp/"
	switch name {
	case "goProject":
		path = "D:/GoProject/"
	case "yfl":
		path = "D:/name_file/yflFile/"
	case "ych":
		path = "D:/name_file/ychFile/"
	case "lsn":
		path = "D:/name_file/lsnFile/"
	case "cyw":
		path = "D:/name_file/cywFile/"
	case "gky6666":
		path = "D:/name_file/gky6666/"
	default:
		path = "D:/HttpPublic/"
	}
	return path

}
func CheckPath(decodedPath string) bool {
	a, err := url.QueryUnescape(decodedPath)
	println(a, err)
	if err != nil {
		// 解码失败就拒绝
		return false
	}
	if (a == "D:\\HttpPublic/" || strings.Contains(a, "D:\\name_file")) && a != "D:\\name_file" {
		return true
	}
	return false
}

// 记录日志

func FileSizeConvert(size int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.3f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.3f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.3f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func DebugRequest(c *gin.Context) {
	// 打印请求方法和URL
	fmt.Println("Method:", c.Request.Method)
	fmt.Println("URL:", c.Request.URL.String())

	// 打印请求头
	fmt.Println("Headers:")
	for k, v := range c.Request.Header {
		fmt.Printf("  %s: %v\n", k, v)
	}

	// 打印请求路径参数
	fmt.Println("Path Params:")
	for _, param := range c.Params {
		fmt.Printf("  %s = %s\n", param.Key, param.Value)
	}

	// 打印Query参数
	fmt.Println("Query Params:")
	for k, v := range c.Request.URL.Query() {
		fmt.Printf("  %s = %v\n", k, v)
	}

	// 打印Form参数（支持 x-www-form-urlencoded、multipart）
	c.Request.ParseForm()
	fmt.Println("Form Params:")
	for k, v := range c.Request.PostForm {
		fmt.Printf("  %s = %v\n", k, v)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file"})
		return
	}
	if file != nil {
		return
	} // 打印Body内容（适用于 JSON 等）
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	fmt.Println("Body:")
	fmt.Println(string(bodyBytes))
	// 注意：读取后 Body 会被“消费”，后续再用需要重新赋值
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
}

func TimeUsed(start time.Time) float64 {
	time.Since(start)
	return time.Since(start).Seconds()
}
