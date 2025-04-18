package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// 初始化数据库连接，只执行一次
func InitDB() *gorm.DB {
	once.Do(func() {
		var err error
		dbInstance, err = gorm.Open(sqlite.Open("./tables/users.db"), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败: %v", err))
		}
		// 自动建表
		if err := dbInstance.AutoMigrate(&Users{}); err != nil {
			panic(fmt.Sprintf("AutoMigrate 失败: %v", err))
		}
	})
	return dbInstance
}

// Users 表结构
type Users struct {
	Name     string `gorm:"column:name;type:varchar(32)" json:"姓名"`
	Email    string `gorm:"column:email;type:varchar(64)" json:"邮箱"`
	Password string `gorm:"column:password;type:varchar(64)" json:"密码"`
	Token    string `gorm:"column:token;type:varchar(64)" json:"安全代码"`
	Nickname string `gorm:"column:nickname;type:varchar(32)" json:"昵称"`
}

// 自定义表名
func (u *Users) TableName() string {
	return "users"
}

// 创建单个用户
func (u *Users) Create() error {
	db := InitDB()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	return db.Create(u).Error
}

// 批量创建用户
func BatchCreate(users []Users) error {
	if len(users) == 0 {
		return nil
	}
	db := InitDB()
	return db.Create(&users).Error
}
