package file_handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type SaveFileRequest struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func SaveFileContent(c *gin.Context) {
	var req SaveFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求格式错误"})
		return
	}

	err := os.WriteFile(req.Path, []byte(req.Content), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "保存失败：" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
