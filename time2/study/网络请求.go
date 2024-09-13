package main

import (
	"net/http"
)

func test01() {
	http.Get("https://www.bilibili.com/")
}

//不同的请求方式
//post请求用得最多
//kv形式请求
//key value
//json形式请求
