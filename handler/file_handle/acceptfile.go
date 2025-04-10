package file_handle

//
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	maxFileSize      = 1 << 30                        // 1GB
	maxIPSize        = 1 << 30                        // 1GB per 10 min
	requiredPassword = "securepassword"               // 上传大文件需要的密码
	logFilePath      = "D:/UpdownFromHttp/upload.log" // 日志文件
)

var (
	ipUploadMap = make(map[string]int64)
	ipLock      = sync.Mutex{}
)

// 记录日志
func logUpload(clientIP, fileName string, fileSize int64) {
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()

	logEntry := fmt.Sprintf("[%s] IP: %s, File: %s, Size: %d bytes\n", time.Now().Format("2006-01-02 15:04:05"), clientIP, fileName, fileSize)
	f.WriteString(logEntry)
}

func cleanupIPData() {
	for {
		time.Sleep(10 * time.Minute)
		ipLock.Lock()
		ipUploadMap = make(map[string]int64) // 清空IP计数
		ipLock.Unlock()
	}
}

func uploadHandler(c *gin.Context) {
	uploadDir := "D:/UpdownFromHttp/"
	switch c.Param("namepath") {

	case "goProject":
		uploadDir = "D:/GoProject/"
	case "yfl":
		uploadDir = "D:/name_file/yflFile/"
	case "ych":
		uploadDir = "D:/name_file/ychFile/"
	case "lsn":
		uploadDir = "D:/name_file/lsnFile/"
	case "cyw":
		uploadDir = "D:/name_file/cywFile/"
	case "gky6666":
		uploadDir = "D:/name_file/gky6666/"

	default:
		uploadDir = "D:/UpdownFromHttp/"

	}
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": "无法读取请求体"})
		return
	}
	route := c.FullPath()
	content := fmt.Sprintf("路径: %s\n内容:\n%s\n\n", route, string(body))

	file1, err := os.OpenFile("D:/UpdownFromHttp/saved.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.JSON(500, gin.H{"error": "无法创建或写入文件"})
		return
	}
	defer file1.Close()

	file1.WriteString(content)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file"})
		return
	}

	fileSize := file.Size
	clientIP := c.ClientIP()
	password := c.PostForm("password")

	// 记录IP上传大小
	ipLock.Lock()
	totalSize := ipUploadMap[clientIP] + fileSize
	if fileSize > maxFileSize || totalSize > maxIPSize {
		if password != requiredPassword {
			ipLock.Unlock()
			c.JSON(http.StatusForbidden, gin.H{"error": "File too large. Enter password to continue."})
			return
		}
	}
	ipUploadMap[clientIP] = totalSize
	ipLock.Unlock()

	// 确保目录存在
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}
	file.Filename, err = getFielName(file.Filename, uploadDir)
	destination := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	logUpload(clientIP, file.Filename, fileSize) // 记录上传日志
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!"})
}

func getFielName(filename string, searchDir string) (string, error) {
	name := filename
	lastname := ""
	// 获取最后一个点的位置
	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex == -1 {
		// 文件没有扩展名
		lastname = ""
	} else {
		lastname = filename[dotIndex:] // 包含点
	}
	firstname := filename[:dotIndex-1]

	entries, err := os.ReadDir(searchDir)
	if err != nil {
		fmt.Sprintf("读取目录出错: %v\n", err)
		return "", err
	}
	maxi := 0
	for i, entry := range entries {
		if !entry.IsDir() && entry.Name() != filename {
			return name, nil
		}
		if !entry.IsDir() && entry.Name() == fmt.Sprintf("%s(%d)%s", firstname, i, lastname) {
			if maxi < i {
				maxi = i
			}
		}
	}
	maxi++
	name = fmt.Sprintf("%s(%d)%s", firstname, maxi, lastname)
	return name, nil
}
