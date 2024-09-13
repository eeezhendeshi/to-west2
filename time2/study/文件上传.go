package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func test04() {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	// 读取文件
	file, err := os.Open("server.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建一个新的file
	fileWrite, err := bodyWrite.CreateFormFile("file", "server_1.go")
	// 将上面的file放入现在的file
	_, err = io.Copy(fileWrite, file)
	if err != nil {
		log.Println("err")
		return
	}
	bodyWrite.Close()
	contentType := bodyWrite.FormDataContentType()
	http.Post("http://127.0.0.1:7070/file", contentType, bodyBuf)
}
