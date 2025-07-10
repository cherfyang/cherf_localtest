package file_handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DownloadFile(c *gin.Context) {
	start := time.Now()
	filePath := c.Query("file")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件路径不能为空"})
		return
	}

	// 替换路径分隔符（适配Windows/Linux）
	filePath = filepath.Clean(strings.ReplaceAll(filePath, "/", string(os.PathSeparator)))

	// 检查文件是否存在
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文件不存在"})
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取文件信息"})
		return
	}

	// 支持断点续传（Range请求）
	rangeHeader := c.GetHeader("Range")
	if rangeHeader != "" {
		// 解析Range头（格式：bytes=0-999）
		startByte, endByte, err := ParseRangeHeader(rangeHeader, stat.Size())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的Range请求"})
			return
		}

		// 设置部分内容响应头
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", startByte, endByte, stat.Size()))
		c.Header("Content-Length", strconv.FormatInt(endByte-startByte+1, 10))
		c.Status(http.StatusPartialContent) // 206

		// 跳转到指定位置并发送数据
		_, err = file.Seek(startByte, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "文件读取失败"})
			return
		}

		_, err = io.CopyN(c.Writer, file, endByte-startByte+1)
		if err != nil {
			log.Printf("下载中断: %v", err)
		}
		return
	}

	// 普通下载（完整文件）
	c.Header("Accept-Ranges", "bytes") // 声明支持断点续传
	c.FileAttachment(filePath, filepath.Base(filePath))

	log.Printf("下载完成: %s (耗时: %v)", filePath, time.Since(start))
}

// 解析Range请求头
func ParseRangeHeader(rangeHeader string, fileSize int64) (int64, int64, error) {
	const prefix = "bytes="
	if !strings.HasPrefix(rangeHeader, prefix) {
		return 0, 0, fmt.Errorf("invalid range header")
	}

	rangeStr := strings.TrimPrefix(rangeHeader, prefix)
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range format")
	}

	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start offset")
	}

	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		end = fileSize - 1 // 如果只指定start（如 "bytes=1000-"）
	}

	if start < 0 || end >= fileSize || start > end {
		return 0, 0, fmt.Errorf("range out of bounds")
	}

	return start, end, nil
}
