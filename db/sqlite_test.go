package db

import (
	_ "github.com/gin-gonic/gin"
	"testing"
	"time"
)

func Test(t *testing.T) {
	user := Users{
		ID:        1,
		Name:      "ych",
		Email:     "",
		Password:  "",
		Token:     "",
		Nickname:  "mx",
		Role:      "manager",
		Status:    0,
		FullPath:  "",
		Permisson: "all",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	user.Create()
}
