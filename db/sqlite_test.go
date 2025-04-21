package db

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"testing"
)

func Test(t *testing.T) {
	ip := IPInfo{
		Status:      "1",
		Country:     "1",
		CountryCode: "1",
		Region:      "1",
		RegionName:  "1",
		City:        "1",
		Zip:         "1",
		Lat:         0,
		Lon:         0,
		Timezone:    "1",
		ISP:         "1",
		Org:         "1",
		AS:          "1",
		Query:       "1",
	}
	fmt.Println(ip)
	InitIPInfoDB()
}
