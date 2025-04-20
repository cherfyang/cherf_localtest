package handlers

import (
	"cherf_localtest/db"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取未删除网站列表
func GetWebs(c *gin.Context) {
	var webs []db.Webs

	result := db.InitWebDB().Find(&webs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, webs)
}

// 添加新网站
func AddWeb(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	web := db.Webs{
		Title:       input.Title,
		Description: input.Description,
		URL:         input.URL,
	}

	err := web.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// 删除网站（逻辑删除）
func DeleteWeb(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("即将删除" + id)
	result := db.InitWebDB().Model(&db.Webs{}).Where("id = ?", id).Update("deleted_at", time.Now())
	if result.Error != nil {
		fmt.Println("删除失败")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Println("删除成功")

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func UndoDeleteWeb(c *gin.Context) {
	var web db.Webs
	oneHourAgo := time.Now().Add(-1 * time.Hour)

	result := db.InitWebDB().
		Unscoped(). // 必须取消默认的 soft delete 筛选
		Where("deleted_at > ?", oneHourAgo).
		Order("deleted_at DESC").
		First(&web)
	fmt.Println(web.ID, web.DeletedAt)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "没有可恢复的数据"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	restore := db.InitWebDB().Unscoped().Model(&web).Update("deleted_at", nil)
	fmt.Println(web.ID, web.DeletedAt)

	if restore.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": restore.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "恢复成功",
		"restored": web.URL,
	})
}
