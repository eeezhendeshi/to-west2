package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func test02() {
	// 第一个 Post 请求
	res, err := http.Post("https://www.bilibili.com/json", "application/json", bytes.NewBuffer([]byte(`{"code":0,"name":"go"}`)))
	if err != nil {
		fmt.Println("Error in Post request:", err)
		return
	}
	defer res.Body.Close() // 关闭响应体

	// 检查响应状态码
	if res.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", res.StatusCode)
		return
	}

	// 第二个 PostForm 请求
	param := url.Values{}
	param.Add("name", "lyu")
	res, err = http.PostForm("https://www.bilibili.com/post", param)
	if err != nil {
		fmt.Println("Error in PostForm request:", err)
		return
	}
	defer res.Body.Close() // 关闭响应体

	// 检查响应状态码
	if res.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", res.StatusCode)
		return
	}

	//自己构造请求对象
	resp, _ := http.NewRequest("GET", "https:/www.bilibili.com/qet", nil)
	resp, _ = http.NewRequest("POST", "https:/www.bilibili.com/post", nil)
	resp, _ = http.NewRequest("PUT", "https:/www.bilibili.com/put", nil)
	resp, _ = http.NewRequest("DELETE", "https:/www.bilibili.com/delete", nil)
	//发送请求
	http.DefaultClient.Do(resp)
}
