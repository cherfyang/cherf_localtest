package db

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once      sync.Once
	DB_Users  *gorm.DB
	DB_Webs   *gorm.DB
	DB_IPInfo *gorm.DB
	DB_       *gorm.DB
)

const (
	//PathDir  = "D:/code/cherf_localtest/db/tables"
	PathDir  = "/Users/developer/GolandProjects/cherf_localtest/cherf_localtest/db/tables"
	UserPath = PathDir + "/users.db"
	WebPath  = PathDir + "/webs.db"
)

var ()
