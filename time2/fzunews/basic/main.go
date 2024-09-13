package main

import(
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)
//包含发布时间，作者，标题以及正文。
func main(){
	myurl:="https://info22.fzu.edu.cn/lm_list.jsp?wbtreeid=1460"
	resp,err:=http.Get(myurl)
	data,_:=io.ReadAll(resp.body)
	time:=regexp.
	rwiter:=regexp.
	title:=regexp.
	text:=regexp.
}