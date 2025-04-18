package sdk

import (
	"cherf_localtest/models"
	"encoding/json"
	"fmt"
	"log"
)

// http://ip-api.com/json/118.249.72.132?lang=zh-CN
func IpToAddress(ip string) string {
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	body, _ := CallApi("GET", url, nil, nil)
	var info models.IPInfo
	if err := json.Unmarshal(body, &info); err != nil {
		log.Fatal("JSON 解析失败:", err)
	}
	return fmt.Sprintf("IP地址：%s\n%s-%s-%s\n运营商：%s", info.Query, info.Country, info.RegionName, info.City, info.ISP)

}
