package main

//url编码:7B%5C%22type%5C%22:1,%5C%22direction%5C%22:1,%5C%22session_id%5C%22:%5C%221768174446670613%5C%22,%5C%22data%5C%22:%7B%7D%7D%

//url=tag1+...+tag2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
	"golang.org/x/exp/rand"
)

type ReplyContainer struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Replies []struct {
			Member struct {
				Mid   string `json:"mid"`   // 用户id
				Uname string `json:"uname"` // 用户姓名
				Sex   string `json:"sex"`   // 性别
			} `json:"member"`
			Content struct {
				Message string        `json:"message"` // 评论内容
				Members []interface{} `json:"members"`
				MaxLine int           `json:"max_line"`
			} `json:"content"`
			ReplyControl struct {
				MaxLine           int    `json:"max_line"`
				SubReplyEntryText string `json:"sub_reply_entry_text"`
				SubReplyTitleText string `json:"sub_reply_title_text"`
				TimeDesc          string `json:"time_desc"` // 评论发布时间
			} `json:"reply_control"`
		} `json:"replies"`
		// 添加分页信息
		Next struct {
			Offset    string `json:"offset"`
			Total     int    `json:"total"`
			Count     int    `json:"count"`
			Pn        int    `json:"pn"`
			Ps        int    `json:"ps"`
			IsEnd     bool   `json:"is_end"`
			TotalList int    `json:"total_list"`
		} `json:"next"`
	} `json:"data"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString 生成一个随机的user-agent
func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//var string{
//	tag1:"https://api.bilibili.com/x/v2/reply/wbi/main?oid=420981979&type=1&mode=3&pagination_str=%7B%22offset%22:%22%"
//	tag2:"22%7D&plat=1&web_location=1315875&w_rid=f2bab233af18fc144d440126666d9d51&wts=1726732984"
//}

func main() {
	c := colly.NewCollector()

	// 设置请求头
	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("authority", "api.bilibili.com")
		req.Headers.Set("accept", "application/json, text/plain, */*")
		req.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Headers.Set("cookie", "buvid3=40F5DB37-ED76-8301-C8FB-3D30FECEE28727739infoc; b_nut=1721444827; _uuid=16CDC26B-9AC5-F792-595D-41A9E391FA10628044infoc; buvid4=D0AE52C8-2E59-D948-1232-873C1CD0918766939-024072003-wKCkktEUR%2FGIFLg7IYe1CA%3D%3D; enable_web_push=DISABLE; header_theme_version=CLOSE; rpdid=|(um~R)R)lRm0J'u~kuRRl|mm; DedeUserID=512779590; DedeUserID__ckMd5=5be93575780ece10; buvid_fp_plain=undefined; hit-dyn-v2=1; CURRENT_BLACKGAP=0; CURRENT_FNVAL=4048; fingerprint=76b02d924aa80ba2d4382049a27434fe; buvid_fp=d02a873a2c2166a802095ed7680dc780; is-2022-channel=1; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY1NDE4MjMsImlhdCI6MTcyNjI4MjU2MywicGx0IjotMX0.dYHo5O5tf1ej8YoYQGNbfsC35-rgEy6cQ1x7997_DM0; bili_ticket_expires=1726541763; SESSDATA=6ecdfd50%2C1741928090%2C0d5ad%2A92CjCbauDu2dd9cXOgdu8ZIFTnUdqpm9X71jqP7LeiluQ7mnfCw7SikhmZRzZx9YFt1PMSVnR5d3BEMXBYZHpLRkVQZWxreldtNlZtemk1alpnTVNFb01kcFVidUt6SUJfVnpObVFTTFNzYUI1ZEF0VlBsWDVMczFWcHZiS2dxLVRkemhCdDc0Rk9RIIEC; bili_jct=7355686361869bd0864f7787192796ad; CURRENT_QUALITY=80; home_feed_column=5; browser_resolution=1659-954; b_lsid=7999C671_191FD86CBBB; bsource=search_bing; bp_t_offset_512779590=977951456160120832; sid=59lwks4x")
		req.Headers.Set("origin", "https://www.bilibili.com")
		req.Headers.Set("referer", "https://www.bilibili.com/video/BV1AE4depEab/?spm_id_from=333.1007.tianma.1-1-1.click&vd_source=50b9cdbd346a4236b5ce3aadb61a163e")

		req.Headers.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
		req.Headers.Set("sec-ch-ua-mobile", "?0")
		req.Headers.Set("sec-ch-ua-platform", `"Windows"`)
		req.Headers.Set("sec-fetch-dest", "empty")
		req.Headers.Set("sec-fetch-mode", "cors")
		req.Headers.Set("sec-fetch-site", "same-site")
		req.Headers.Set("user-agent", RandomString())
	})

	// 结构体 用来存放评论数据
	var container ReplyContainer

	// 初始 URL
	var nextURL string = "https://api.bilibili.com/x/v2/reply/wbi/main?oid=113146383571717&type=1&mode=3&pagination_str=%7B%22offset%22:%22%22%7D&plat=1&seek_rpid=&web_location=1315875&w_rid=302a21c3ab1531d52efa7f4a7f13ecb3&wts=1726535773"

	var wg sync.WaitGroup
	var mu sync.Mutex
	ch1 := make(chan string)
	go func() { ch1 <- nextURL }()
	wg.Add(1)
	f, err := os.Create("comments.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	tag := 0
	for nextURL != "" && tag < 100 {
		tag++
		wg.Add(1)
		go func() {
			defer wg.Done()
			go c.OnResponse(func(r *colly.Response) {
				if r.StatusCode != 200 {
					println(r.StatusCode)
					os.Exit(1)
				}
				fmt.Println("response received", r.StatusCode)
				err := json.Unmarshal(r.Body, &container)
				if err != nil {
					fmt.Println("error", err)
					log.Fatal(err)
				}

				// 处理响应数据
				mu.Lock()
				defer mu.Unlock()
				for _, reply := range container.Data.Replies {
					fmt.Fprintf(f, "姓名: %s, 内容: %s, 时间: %s\n", reply.Member.Uname, reply.Content.Message, reply.ReplyControl.TimeDesc)
				}

				// 更新nextURL为下一个分页的URL
				if !container.Data.Next.IsEnd {
					offset := container.Data.Next.Offset // 提取offset字段
					ch1 <- "https://api.bilibili.com/x/v2/reply/wbi/main?oid=113146383571717&type=1&mode=3&pagination_str={\"offset\":\"" + offset + "\"}&plat=1&seek_rpid=&web_location=1315875&w_rid=302a21c3ab1531d52efa7f4a7f13ecb3&wts=1726535773"
					wg.Add(1)
				}

			})

			// 访问 URL
			go func() {
				url := <-ch1
				c.Visit(url)
				wg.Done()
			}()
		}()
	}
	wg.Wait()
}
