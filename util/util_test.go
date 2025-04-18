package util

import (
	"cherf_localtest/db"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := db.Users{
		Name:     "cherf_localtest",
		Email:    "",
		Password: "",
		Token:    "",
		Nickname: "",
	}
	sMap, _ := StructToMap(s)
	fmt.Println(s)
	PrintMap(sMap)
}
