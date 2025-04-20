package db

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

var DBweb *gorm.DB

func InitWebDB() *gorm.DB {
	once.Do(func() {
		var err error
		_ = os.MkdirAll("D:/code/cherf_localtest/db/tables", os.ModePerm)

		dbInstance, err = gorm.Open(sqlite.Open("D:/code/cherf_localtest/db/tables/web.db"), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败: %v", err))
		}
		// 自动建表
		if err := dbInstance.AutoMigrate(&Webs{}); err != nil {
			panic(fmt.Sprintf("AutoMigrate 失败: %v", err))
		}
	})
	return dbInstance
}

type Webs struct {
	ID          uint           `gorm:"column:id"json:"id" gorm:"primaryKey"`
	Title       string         `gorm:"column:title"json:"title"`
	Description string         `gorm:"column:description"json:"description"`
	URL         string         `gorm:"column:url"json:"url"`
	CreatedAt   time.Time      `gorm:"column:created_at"json:"create_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"json:"update_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"json:"delete_at" gorm:"index"`
}

func (w *Webs) TableName() string {
	return "webs"
}
func (w *Webs) Create() error {
	db := InitWebDB()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	return db.Create(w).Error
}
func (w *Webs) DB() *gorm.DB {
	return InitWebDB()
}
func Find(w *[]Webs) {
	db := InitWebDB()
	db.Find(w)

}
