package db

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

type IPInfo struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"-"`         // 主键ID
	Status      string  `gorm:"type:varchar(20)" json:"status"`            // 查询状态
	Country     string  `gorm:"type:varchar(100)" json:"country"`          // 国家名称
	CountryCode string  `gorm:"type:varchar(10)" json:"countryCode"`       // 国家代码
	Region      string  `gorm:"type:varchar(100)" json:"region"`           // 区域代码
	RegionName  string  `gorm:"type:varchar(100)" json:"regionName"`       // 区域名称
	City        string  `gorm:"type:varchar(100)" json:"city"`             // 城市名称
	Zip         string  `gorm:"type:varchar(20)" json:"zip"`               // 邮政编码
	Lat         float64 `gorm:"type:decimal(10,6)" json:"lat"`             // 纬度
	Lon         float64 `gorm:"type:decimal(10,6)" json:"lon"`             // 经度
	Timezone    string  `gorm:"type:varchar(50)" json:"timezone"`          // 时区
	ISP         string  `gorm:"type:varchar(100)" json:"isp"`              // 网络服务提供商
	Org         string  `gorm:"type:varchar(100)" json:"org"`              // 所属组织
	AS          string  `gorm:"type:varchar(100)" json:"as"`               // 自治系统信息
	Query       string  `gorm:"type:varchar(50);uniqueIndex" json:"query"` // 查询的IP地址
}

func (i *IPInfo) TableName() string {
	return "ip_infos"
}

func InitIPInfoDB() *gorm.DB {
	once.Do(func() {
		var err error
		err = os.MkdirAll("./tables", 0750)
		if err != nil {
			fmt.Println("zzzzzzzzzzzzzzzzzzz")
			panic(err)
		}
		Path := "./tables/ip_infos.db"
		DB_IPInfo, err = gorm.Open(sqlite.Open(Path), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败: %v", err))
		}
		// 自动建表
		if err = DB_IPInfo.AutoMigrate(&IPInfo{}); err != nil {
			panic(fmt.Sprintf("AutoMigrate 失败: %v", err))
		}
	})

	return DB_IPInfo
}

func (i *IPInfo) Create() error {
	db := InitIPInfoDB()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	return db.Create(i).Error
}

func (i *IPInfo) DB() *gorm.DB {
	return InitIPInfoDB()
}

func (i *IPInfo) Find() *[]IPInfo {
	db := InitIPInfoDB()
	var ips *[]IPInfo
	db.Find(ips)
	return ips
}
