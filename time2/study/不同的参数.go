package main

import (
	"bytes"
	"net/http"
	"net/url"
)

func test03() {
	//查询参数
	//1 直接写在url上
	//2 构造url.Values{}
	http.Get("http://www.baidu.com?name=zhangsan&age=20")

	//构造查询参数
	params := url.Values{}
	params.Add("name", "zhangsan")
	params.Add("age", "20")

	//使用 http.NewRequest 函数创建一个新的 HTTP GET 请求对象 resq
	resq, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	//将之前构造的 params 对象通过 Encode 方法转换为 URL 编码的字符串，并赋值给 resq.URL.RawQuery
	//从而将查询参数附加到请求的 URL 上
	resq.URL.RawQuery = params.Encode()
	println(params.Encode())

	//使用默认的 HTTP 客户端发送一个 HTTP 请求。
	//Do 方法会执行这个请求，并返回一个 http.Response 对象和一个错误对象
	http.DefaultClient.Do(resq)

	http.Get("http://www.baidu.com/query?" + params.Encode())

	//body参数
	params = url.Values{}
	params.Set("name", "zhangsan")
	http.Post("http://www.baidu.com/form", "application/x-www-form-unlencoded", bytes.NewBuffer([]byte(params.Encode())))
	http.Post("http://www.baidu.com/form", "multipart/form-data", bytes.NewBuffer([]byte(params.Encode())))
}
