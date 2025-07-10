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
	name := c.Query("namepath")
	//// 路径
	rootPath, _ := util.GetFullpathByParam(name)
	//// 获取查询参数
	path := c.Query("path")
	if path == "" {
		path = rootPath
	}

	//if !isAdmin && path != util.WinPath[name] {
	//	c.JSON(http.StatusOK, gin.H{
	//		"path": "当前目录禁止访问！！！",
	//	})
	//	return
	//}
	//if path == "" {
	//	path = "/Users"
	//}
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
		edit := false
		if IsEditable(entry.Name()) {
			edit = true
		}
		files = append(files, gin.H{
			"name":     entry.Name(),
			"path":     fullPath,
			"isDir":    isDir,
			"size":     size,
			"editable": edit,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
		"path":  path,
	})
}
func IsEditable(fname string) bool {
	switch filepath.Ext(fname) {
	case ".md":
		return true
	case ".txt":
		return true
	}
	return false
}
