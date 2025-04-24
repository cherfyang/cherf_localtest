package db

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

// 初始化数据库连接，只执行一次
func InitUserDB() *gorm.DB {
	once.Do(func() {
		var err error
		_ = os.MkdirAll(PathDir, os.ModePerm)

		DB_Users, err = gorm.Open(sqlite.Open(UserPath), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败: %v", err))
		}
		// 自动建表
		if err := DB_Users.AutoMigrate(&Users{}); err != nil {
			panic(fmt.Sprintf("AutoMigrate 失败: %v", err))
		}
	})

	return DB_Users
}

// Users 表结构
// 邮箱和 token 才是确认用户的凭证
type Users struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`                      // 主键
	Name      string    `gorm:"column:name;type:varchar(32);not null" json:"name"`       // 姓名
	Email     string    `gorm:"column:email;type:varchar(64);not null" json:"email"`     // 邮箱
	Password  string    `gorm:"column:password;type:varchar(64);not null" json:"-"`      // 密码（不返回给前端）
	Token     string    `gorm:"column:token;type:varchar(64)" json:"token"`              // 安全代码 / 认证 token
	Nickname  string    `gorm:"column:nickname;type:varchar(32)" json:"nickname"`        // 昵称
	Role      string    `gorm:"column:role;type:varchar(20);default:'user'" json:"role"` // 角色（如 admin/user）
	Status    int       `gorm:"column:status;type:int;default:1" json:"status"`          // 状态（1=正常，0=禁用）
	FullPath  string    `gorm:"column:full_path;type:varchar(255)" json:"full_path"`
	Premisson string    `gorm:"column:premisson;type:varchar(255)" json:"premisson"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

// 自定义表名
func (u *Users) TableName() string {
	return "users"
}

// 创建单个用户
func (u *Users) Create() error {
	db := InitUserDB()
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
	db := InitUserDB()
	return db.Create(&users).Error
}
