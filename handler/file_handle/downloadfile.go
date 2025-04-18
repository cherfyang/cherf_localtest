package file_handle

import (
	"cherf_localtest/log"
	"cherf_localtest/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

// 文件下载接口
func DownloadFile(c *gin.Context) {
	start := time.Now()
	filePath := c.DefaultQuery("file", "")

	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件路径不能为空"})
		return
	}

	// 直接返回文件
	println("这是下载接口，下载路径: " + filePath)
	c.FileAttachment(filePath, filepath.Base(filePath))
	usetime := util.TimeUsed(start)
	log.LogDownload(c, start, usetime)
	println("下载完成")
}
