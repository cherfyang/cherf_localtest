package db

import (
	_ "github.com/gin-gonic/gin"
	"testing"
)

func Test(t *testing.T) {

	//web := Webs{
	//	Title:       "搜狗",
	//	URL:         "http://www.sougou.com",
	//	Description: "一个搜索引擎",
	//}
	var webs []Webs
	//webs = append(webs, web)
	Find(&webs)
	print(webs[0].Title)
}
