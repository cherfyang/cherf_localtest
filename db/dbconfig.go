package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	dbMap = make(map[string]*gorm.DB)
	once2 sync.Once
)

const macBasePath = "/Users/developer/GolandProjects/cherf_localtest/cherf_localtest/db/tables/"
const winBasePath = "/Users/developer/GolandProjects/cherf_localtest/cherf_localtest/db/tables/"

type dbConfig struct {
	Name     string
	FileName string
	Models   []interface{}
}

var dbConfigs = []dbConfig{
	{
		Name:   "users",
		Models: []interface{}{&Users{}},
	},
	{
		Name:   "webs",
		Models: []interface{}{&Webs{}},
	},
	// 可以继续添加 IPInfo 等配置
}

func init() {

	basePath := macBasePath
	//if {
	//	basePath = winBasePath
	//}
	_ = os.MkdirAll(basePath, os.ModePerm)

	for _, cfg := range dbConfigs {
		path := basePath + cfg.Name + ".db"
		db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("打开数据库 %s 失败: %v", cfg.Name, err))
		}

		if err := db.AutoMigrate(cfg.Models...); err != nil {
			panic(fmt.Sprintf("AutoMigrate %s 失败: %v", cfg.Name, err))
		}

		dbMap[cfg.Name] = db
		fmt.Printf("✅ 数据库 [%s] 初始化成功\n", cfg.Name)
	}

}

func GetDB(name string) *gorm.DB {
	//InitAllDBs()
	return dbMap[name]
}
