package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
	"time"
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
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`                     // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`                     // 更新时间
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
