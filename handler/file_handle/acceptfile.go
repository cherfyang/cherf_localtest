package file_handle

//
import (
	_const "cherf_localtest/const"
	"cherf_localtest/log"
	"cherf_localtest/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	ipUploadMap = make(map[string]int64)
	ipLock      = sync.Mutex{}
)

func cleanupIPData() {
	for {
		time.Sleep(10 * time.Minute)
		ipLock.Lock()
		ipUploadMap = make(map[string]int64) // 清空IP计数
		ipLock.Unlock()
	}
}

func UploadHandler(c *gin.Context) {
	start := time.Now()
	uploadDir := c.Query("namepath")
	from := c.GetHeader("Sec-Ch-Ua-Platform")
	Agent := c.GetHeader("User-Agent")
	util.DebugRequest(c)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file"})
		return
	}

	fileSize := file.Size
	clientIP := c.ClientIP()
	if clientIP != "::1" {
		password := c.PostForm("password")

		// 记录IP上传大小
		ipLock.Lock()
		totalSize := ipUploadMap[clientIP] + fileSize
		if fileSize > _const.MaxFileSize || totalSize > _const.MaxIPSize {
			if password != _const.RequiredPassword {
				ipLock.Unlock()
				c.JSON(http.StatusForbidden, gin.H{"error": "File too large. Enter password to continue."})
				return
			}
		}
		ipUploadMap[clientIP] = totalSize
		ipLock.Unlock()
	}

	// 确保目录存在
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}
	file.Filename, err = util.GetFielName(file.Filename, uploadDir)
	destination := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}
	useTime := util.TimeUsed(start)
	log.LogUpload(clientIP, file.Filename, fileSize, uploadDir, Agent, from, useTime) // 记录上传日志
	fmt.Println("耗时：", time.Since(start))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File uploaded successfully!"})
}
