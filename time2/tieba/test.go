// https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=1&wbtreeid=1460
// https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=2&wbtreeid=1460
// https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=3&wbtreeid=1460
// https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=4&wbtreeid=1460
// PAGENUM 变化
// 横向爬取：页为单位
// 纵向爬取：条目为单位
package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	//defer 是一个关键字，用于确保在函数返回之前执行特定的代码
	//defer 后面的函数调用会被延迟到包围它的函数即将返回之前执行
	defer resp.Body.Close()
	//循环读取网页数据 传出给调用者
	for {
		buf := make([]byte, 1024)
		//body:=make([]byte,1024)
		n, err := resp.Body.Read(buf)
		if n == 0 {
			println("读取结束")
			break
		}
		if err != nil && err != io.EOF {
			println(err.Error())
		}
		result += string(buf[:n])
	}
	//将读取的网页保存为文件
	return
}

// 爬取操作
func work(start int, end int) {
	println("爬取ing")
	//循环爬取每一页
	for i := start; i <= end; i++ {
		url := "https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=" + strconv.Itoa(i) + "&wbtreeid=1460"
		result, err := HttpGet(url)
		if err != nil {
			println("爬取失败", err.Error())
			continue
		}
		f, err := os.Create("第" + strconv.Itoa(i) + "页.html")
		if err != nil {
			println("第"+strconv.Itoa(i)+"页创建文件失败", err)
			continue
		}
		f.WriteString(result)
		f.Close() //保存一个文件 关闭一个文件
		//不用defer的原因 会导致文件被打开关闭多次
		println("第" + strconv.Itoa(i) + "页爬取成功")
	}

}

func main() {
	//爬取页数
	//指定爬取 起始 终止页面
	start := 1
	end := 50
	work(start, end)
}
