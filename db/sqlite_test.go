package db

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func Test(t *testing.T) {

	token := uuid.New().String()
	user := Users{
		Name:      "杨超华",
		Email:     "2637206496@qq.com",
		Password:  "Yang2580..",
		Token:     token,
		Nickname:  "梦溪",
		Role:      "systemManage",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := user.Create()
	if err == nil {
		fmt.Println("添加成功")
	}
}
