package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func test05() {
	res, _ := http.Get("https://www.fengfengzhidao.com/article/0dlaH4sBEG4v2tWkjmrp#go%E7%88%AC%E8%99%AB")
	//使用io.ReadAll函数读取HTTP响应体的所有内容，并将其存储在byteData变量中。
	//res.Body是响应体的io.ReadCloser接口。
	byteData, _ := io.ReadAll(res.Body)
	//fmt.Println(string(byteData))
	//定义一个map[string]any类型的变量data，
	//用于存储解析后的JSON数据。
	//然后使用json.Unmarshal函数将byteData中的JSON数据解析到data中。
	var data map[string]any
	json.Unmarshal(byteData, &data)
	if !json.Valid(byteData) {
		fmt.Println("Invalid JSON data")
		return
	}

	fmt.Println(data)
}
