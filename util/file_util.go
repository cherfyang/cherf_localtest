package util

import (
	"bytes"
	"cherf_localtest/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var WinPath = map[string]string{}
var unWinPqth = map[string]string{}
var users []db.Users

func init() {

	//err := db.InitUserDB().Find(&users).Error
	//if err == nil && len(users) > 0 {
	//	for _, v := range users {
	//		WinPath[v.Email] = v.FullPath
	//	}
	//}
	//if err != nil {
	//	fmt.Println("初始化用户路径失败:", err)
	//}

}

// 如果不能全部路径就返回路径,否则返回 true
func GetFullpathByParam(name string) (string, bool) {
	//permission := ""
	//for _, v := range users {
	//	if v.Email == name {
	//		permission = v.Permisson
	//	}
	//}
	//if permission == "" {
	//	return "err", false
	//}
	path := "D:"
	//switch permission {
	//case "user":
	//	path = WinPath[name]
	//case "onlyD":
	//	path = "D:"
	//case "all":
	//	return "", true
	//}
	return path, false

}
func CheckPermission(decodedPath string) bool {
	a, err := url.QueryUnescape(decodedPath)
	println(a, err)
	if err != nil {
		// 解码失败就拒绝
		return false
	}
	if a == "D:/name_file" {
		return false
	}
	cond := map[int]bool{
		2: strings.Contains(a, "D:\\HttpPublic"),
		3: strings.Contains(a, "D:/HttpPublic"),
		4: strings.Contains(a, "D:\\name_file\\"),
		5: strings.Contains(a, "D:/name_file/"),
	}
	for i := 2; i < 6; i++ {
		if cond[i] {
			return true
		}
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

func MaxLastBracketIndex(files []os.DirEntry, firstname string, lastname string) int {
	names := make([]string, 0, len(files))

	for _, v := range files {
		pattern := fmt.Sprintf(`^%s(?:.*)?%s$`, regexp.QuoteMeta(firstname), regexp.QuoteMeta(lastname))
		matched, _ := regexp.MatchString(pattern, v.Name())
		if matched {
			names = append(names, v.Name())
		}
	}
	if len(names) == 0 {
		return 1
	}
	re := regexp.MustCompile(`\(([^()]*)\)`)

	maxIndex := 0

	for _, name := range names {
		matches := re.FindAllStringSubmatch(name, -1) // 获取所有匹配
		if len(matches) == 0 {
			continue
		}
		// 取最后一个括号中的数字
		last := matches[len(matches)-1][1]
		num, err := strconv.Atoi(last)
		if err != nil {
			continue
		}
		if num > maxIndex {
			maxIndex = num
		}
	}
	return maxIndex + 1
}
func GetFielName(path string) (string, error) {

	searchDir, firstname, lastname := SplitFilePath(path)
	name := firstname + "." + lastname

	entries, err := os.ReadDir(searchDir)
	if err != nil {
		fmt.Sprintf("读取目录出错: %v\n", err)
		return "", err
	}
	num := 0
	flag := false
	//判断是否有entrie==filename
	for _, entry := range entries {
		if entry.Name() != name {
			continue
		} else {
			flag = true
		}
	}
	if flag {
		num = MaxLastBracketIndex(entries, firstname, lastname)
		return fmt.Sprintf("%s(%d)%s", firstname, num, lastname), nil
	}
	//
	return name, nil
}

// SplitFilePath 分离路径，返回目录、文件名（无后缀）、扩展名（不带点）
func SplitFilePath(fullPath string) (dir, nameWithoutExt, ext string) {
	dir = filepath.Dir(fullPath)
	base := filepath.Base(fullPath)       // 例如：readme.md
	rawExt := filepath.Ext(base)          // 例如：.md
	ext = strings.TrimPrefix(rawExt, ".") // 去掉前缀点
	nameWithoutExt = strings.TrimSuffix(base, rawExt)
	return
}

func RequestString(c *gin.Context) string {
	str := ""
	// 打印请求方法和URL
	str += fmt.Sprintf("Method:%s\n", c.Request.Method)
	str += fmt.Sprintf("URL:%s\n", c.Request.URL.String())

	// 打印请求头
	str += fmt.Sprintf("Headers:\n")
	for k, v := range c.Request.Header {
		str += fmt.Sprintf("  %s: %v\n", k, v)
	}

	// 打印请求路径参数
	str += fmt.Sprintf("Path Params:\n")
	for _, param := range c.Params {
		str += fmt.Sprintf("  %s = %s\n", param.Key, param.Value)
	}

	// 打印Query参数
	str += fmt.Sprintf("Query Params:\n")
	for k, v := range c.Request.URL.Query() {
		str += fmt.Sprintf("  %s = %v\n", k, v)
	}

	// 打印Form参数（支持 x-www-form-urlencoded、multipart）
	c.Request.ParseForm()
	str += fmt.Sprintf("Form Params:\n")
	for k, v := range c.Request.PostForm {
		str += fmt.Sprintf("  %s = %v\n", k, v)
	}

	bodyBytes, _ := io.ReadAll(c.Request.Body)
	str += fmt.Sprintf("Body:\n")
	str += fmt.Sprintf(string(bodyBytes))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return str
}

// 移动文件srcPath到dstPath
func MoveFile(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 复制内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 删除原文件
	err = os.Remove(srcPath)
	if err != nil {
		return err
	}
	return nil
}
