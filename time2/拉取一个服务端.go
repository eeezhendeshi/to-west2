package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	fmt.Println("成功请求")
}

func get(c *gin.Context) {
	fmt.Println("get请求")
}
func post(c *gin.Context) {
	fmt.Println("post请求")
}
func put(c *gin.Context) {
	fmt.Println("put请求")
}
func Delete(c *gin.Context) {
	fmt.Println("delete请求")
}

func form(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	fmt.Println(string(byteData), err, c.Request.Header.Get("Content-Type"))

}
func jsonM(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	fmt.Println(string(byteData), err)
}

func query(c *gin.Context) {
	byteData, err := json.Marshal(c.Request.URL.Query())
	fmt.Println(string(byteData), err)
}
func head(c *gin.Context) {
	byteData, err := json.Marshal(c.Request.Header)
	fmt.Println(string(byteData), err)
}
func file(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return
	}
	fmt.Println(fileHeader.Filename)
	c.SaveUploadedFile(fileHeader, "uploads/file/"+fileHeader.Filename)
}

func getFile(c *gin.Context) {
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, "image.jpg"))
	c.File("uploads/image.jpg")
}
func getJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "xxx",
		"data": gin.H{},
	})
}
func getHtml(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func douban(c *gin.Context) {
	c.HTML(200, "douban.html", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/**")
	router.GET("/ping", ping)
	router.GET("/get", get)
	router.POST("/post", post)
	router.POST("/form", form)
	router.POST("/json", jsonM)
	router.PUT("/put", put)
	router.DELETE("/delete", Delete)
	router.GET("/query", query)
	router.GET("/head", head)
	router.POST("/file", file)
	router.GET("/get_file", getFile)
	router.GET("/get_json", getJson)
	router.GET("/get_html", getHtml)
	router.GET("/douban", douban)
	router.Run(":7070")
}
