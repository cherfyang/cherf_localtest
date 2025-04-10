package sdk

import "cherf_localtest/util"

func CategoryFILE(path, keyWords string) {
	url := "http://localhost:8080/api/v1/file/categorybyname"
	//path := "/Users/developer/Desktop"
	//keyWords := ""
	header := map[string]string{
		"Content-Type": "application/json",
		"X-Source-Dir": path,
		"file-name":    keyWords,
	}
	resp := CallApi("POST", url, header, []byte{})
	util.PrettyPrintBody(resp)
}
