package file_handle

import (
	"cherf_localtest/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func LoadFile(c *gin.Context) {
	filePath := c.Query("file")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件路径不能为空"})
		return
	}
	fileCopyPath := getFileCopyPath(filePath)
	CopyFile(filePath, fileCopyPath)
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取文件失败"})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", contentBytes)
}

func getFileCopyPath(path string) string {
	dir, name, ext := util.SplitFilePath(path)

	if strings.Contains(name, "副本") {
		name, _ := util.GetFielName(path)
		now := time.Now()
		date := now.Format("01/02")
		return fmt.Sprintf("%s\\%s-%s.%s", dir, name, date, ext)
	} else {
		return fmt.Sprintf("%s\\%s-副本.%s", dir, name, ext)
	}
}

func CopyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件（如果已存在会被覆盖）
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 复制内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 确保内容刷新到磁盘
	err = dstFile.Sync()
	if err != nil {
		return err
	}
	return nil
}
