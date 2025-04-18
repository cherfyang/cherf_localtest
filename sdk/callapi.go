package sdk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// 用于直接调用api，传入方法和url即可
func CallApi(method, url string, headers map[string]string, bodys []byte) ([]byte, *http.Response) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodys))
	if err != nil {
		fmt.Println("请求创建失败:", err)
		return nil, nil
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil, nil
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		fmt.Printf("请求失败，状态码: %d\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil, nil
	}

	fmt.Println("响应内容：", string(body))
	return body, res
}
