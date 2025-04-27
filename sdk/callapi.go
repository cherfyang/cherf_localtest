package sdk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

// PublicCallApi 发送 HTTP 请求
//meth:"GET","POST"...

func PublicCallApi(method, rawurl, headers, body, param string) ([]byte, *http.Response, error) {
	client := &http.Client{}

	// 处理参数
	rawurl, err := buildParam(rawurl, param)
	if err != nil {
		return nil, nil, fmt.Errorf("参数处理失败: %w", err)
	}

	// 创建请求体
	var reqBody io.Reader
	if method != http.MethodGet {
		reqBody = bytes.NewBuffer([]byte(body))
	}

	// 创建请求
	req, err := http.NewRequest(method, rawurl, reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("请求创建失败: %w", err)
	}

	// 设置请求头
	h := buildHeader(headers)
	for k, v := range h {
		req.Header.Set(k, v)
	}

	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("请求发送失败: %w", err)
	}
	defer res.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 打印日志
	fmt.Println("响应内容:", string(respBody))

	return respBody, res, nil
}

// buildHeader 解析头部字符串
func buildHeader(headstr string) map[string]string {
	header := make(map[string]string)
	if headstr == "" {
		return header
	}

	items := strings.Split(headstr, ",")
	for _, item := range items {
		kv := strings.SplitN(item, "=", 2)
		if len(kv) == 2 {
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			header[k] = v
		}
	}

	// 如果没有自定义 Content-Type，默认加上
	if _, ok := header["Content-Type"]; !ok {
		header["Content-Type"] = "application/x-www-form-urlencoded"
	}

	return header
}

// buildParam 把 param 参数拼接到 URL 上
func buildParam(rawurl, param string) (string, error) {
	if param == "" {
		return rawurl, nil
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	items := strings.Split(param, "&")
	for _, item := range items {
		kv := strings.SplitN(item, "=", 2)
		if len(kv) == 2 {
			q.Add(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
