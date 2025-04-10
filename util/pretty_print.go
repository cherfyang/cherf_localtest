package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrintBody(body []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		fmt.Println("不是有效的 JSON：")
		fmt.Println(string(body))
		return
	}
	fmt.Println(prettyJSON.String())
}
