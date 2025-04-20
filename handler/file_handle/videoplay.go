package file_handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func StreamVideo(c *gin.Context) {
	videoPath := c.Query("file")
	if videoPath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文件路径"})
		return
	}

	file, err := os.Open(videoPath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文件未找到"})
		return
	}

	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取文件信息"})
		return
	}

	// 设置 MIME 类型，根据文件扩展名也可以动态设置
	c.Header("Content-Type", "video/mp4")
	// 设置支持 range 请求（关键）
	c.Header("Accept-Ranges", "bytes")

	// 使用 http.ServeContent 会自动处理 Range 请求和分段传输
	http.ServeContent(c.Writer, c.Request, stat.Name(), stat.ModTime(), file)
}
