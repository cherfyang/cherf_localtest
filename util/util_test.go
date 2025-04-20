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
func TestStructToMap(t *testing.T) {

	//var webs []db.Webs
	var webs []db.Webs
	webs = append(webs, db.Webs{})
	db.Find(&webs)

	println(webs[0].ID)

}
