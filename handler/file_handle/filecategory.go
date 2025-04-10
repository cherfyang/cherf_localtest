package file_handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func FileCategory(c *gin.Context) {
	sourceDir := c.GetHeader("X-Source-Dir")
	fmt.Println(sourceDir)

	files, err := os.ReadDir(sourceDir)
	fmt.Println("files:")
	fmt.Println(files)
	if err != nil {
		log.Println("读取源目录失败:", err)
		c.JSON(200, gin.H{
			"code": 500,
			"err":  err,
		})
	}
	allCategoryDir := sourceDir + "/category"

	// 遍历所有文件
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过文件夹
		}

		ext := filepath.Ext(file.Name())
		if ext == "" {
			continue
		}
		ext = ext[1:]

		// 创建以扩展名命名的文件夹
		categoryDir := filepath.Join(allCategoryDir, ext)
		err := os.MkdirAll(categoryDir, 0755)
		if err != nil {
			log.Fatal("创建分类文件夹失败:", err)
		}

		// 构造源文件和目标文件的完整路径
		sourceFilePath := filepath.Join(sourceDir, file.Name())
		destFilePath := filepath.Join(categoryDir, file.Name())

		// 移动文件
		err = os.Rename(sourceFilePath, destFilePath)
		if err != nil {
			log.Fatal("移动文件失败:", err)
		}

		fmt.Printf("文件 %s 已移动到 %s\n", file.Name(), categoryDir)
	}
}
