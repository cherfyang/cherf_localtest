package util

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"testing"
)

func Test(t *testing.T) {

	searcher, err := xdb.NewWithFileOnly("ip2region.xdb")
	if err != nil {
		panic(err)
	}
	defer searcher.Close()

	ip := "8.8.8.8"
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println("IP归属地:", region)
}
