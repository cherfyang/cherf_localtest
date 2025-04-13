package file_handle

import (
	"cherf_localtest/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	_ "path/filepath"
	_ "strings"
)

func ListHandler(c *gin.Context) {
	// 路径
	rootPath := util.GetFullpathByParam(c.Query("namepath"))
	// 获取查询参数
	path := c.Query("path")
	if path == "" {
		path = rootPath
	}

	if !util.CheckPath(path) {
		c.JSON(http.StatusOK, gin.H{
			"path": "当前目录禁止访问！！！",
		})
		return
	}
	// 读取目录内容
	entries, err := os.ReadDir(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 构造返回数据
	var files []gin.H
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		isDir := entry.IsDir()
		size := int64(0)

		if !isDir {
			info, err := entry.Info()
			if err == nil {
				size = info.Size()
			}
		}

		files = append(files, gin.H{
			"name":  entry.Name(),
			"path":  fullPath,
			"isDir": isDir,
			"size":  size,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
		"path":  path,
	})
}
