package file_handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FileCategoryByFileNme(c *gin.Context) {
	sourceDir := c.GetHeader("X-Source-Dir")
	//通过-连接多个文件名赋给 filenames
	//分离 filnames 获取需要创建的目录名字
	//正则过滤需要分类的文件名字
	fileNames := c.GetHeader("file-name")
	dirs := strings.Split(fileNames, "-")
	allCategoryDir := sourceDir
	for _, v := range dirs {
		// 创建以扩展名命名的文件夹
		categoryDir := filepath.Join(allCategoryDir, v)
		err := os.MkdirAll(categoryDir, 0755)
		if err != nil {
			log.Fatal("创建分类文件夹失败:", err)
		}
	}

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
	respmsg := make([]string, 0, len(files))

	// 遍历所有文件
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过文件夹
		}
		//获取是否有,没有返回 false,有返回 categoryDir
		categoryDir, isneed := FilterValidFiles(file.Name(), dirs)
		if !isneed {
			continue
		}
		// 构造源文件和目标文件的完整路径
		sourceFilePath := filepath.Join(sourceDir, file.Name())
		destFilePath := filepath.Join(allCategoryDir+categoryDir, file.Name())

		// 移动文件
		err = os.Rename(sourceFilePath, destFilePath)
		if err != nil {
			log.Println("移动文件失败:", err)
			continue
		}

		sg := fmt.Sprintf("{文件 %s 已移动到 %s}", file.Name(), allCategoryDir+categoryDir)
		respmsg = append(respmsg, sg)
	}
	c.JSON(200, gin.H{
		"msg": respmsg,
	})
}

func FilterValidFiles(name string, dirs []string) (string, bool) {

	for _, v := range dirs {
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(v)) // `(?i)` 不区分大小写
		if re.MatchString(name) {
			return "/" + v, true
		}
	}
	return "", false
}
