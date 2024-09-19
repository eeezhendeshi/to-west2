package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var client http.Client
var wg sync.WaitGroup

var tag1 string = "https://api.bilibili.com/x/v2/reply/wbi/main?oid=420981979&type=1&mode=3&pagination_str="
var tag2 string = "&seek_rpid=&web_location=1315875&w_rid=73571eb289423ac967c128301ef3fe3b&wts=1726749763"

type contains struct {
	Code int64 `json:"code"`
	Data struct {
		Replies []struct {
			Content struct {
				Device  string        `json:"device"`
				JumpURL struct{}      `json:"jump_url"`
				MaxLine int64         `json:"max_line"`
				Members []interface{} `json:"members"`
				Message string        `json:"message"`
				Plat    int64         `json:"plat"`
			} `json:"content"`
			Count  int64 `json:"count"`
			Folder struct {
				HasFolded bool   `json:"has_folded"`
				IsFolded  bool   `json:"is_folded"`
				Rule      string `json:"rule"`
			} `json:"folder"`
			Like    int64 `json:"like"`
			Replies []struct {
				Action  int64 `json:"action"`
				Assist  int64 `json:"assist"`
				Attr    int64 `json:"attr"`
				Content struct {
					Device  string   `json:"device"`
					JumpURL struct{} `json:"jump_url"`
					MaxLine int64    `json:"max_line"`
					Message string   `json:"message"`
					Plat    int64    `json:"plat"`
				} `json:"content"`
				Rcount  int64       `json:"rcount"`
				Replies interface{} `json:"replies"`
			} `json:"replies"`
			Type int64 `json:"type"`
		} `json:"replies"`
		Cursor struct {
			PaginationReply struct {
				NextOffset string `json:"next_offset"`
			} `json:"pagination_reply"`
		} `json:"cursor"`
	} `json:"data"`
	Message string `json:"message"`
}

func spider(tag chan string, ch chan string, i int) {
	defer wg.Done()
	jsonStr := <-tag
	jsonBytes := []byte(jsonStr)
	encodedStr := url.QueryEscape(string(jsonBytes))
	url := tag1 + encodedStr + tag2
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.bilibili.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("cookie", "buvid3=40F5DB37-ED76-8301-C8FB-3D30FECEE28727739infoc; b_nut=1721444827; _uuid=16CDC26B-9AC5-F792-595D-41A9E391FA10628044infoc; buvid4=D0AE52C8-2E59-D948-1232-873C1CD0918766939-024072003-wKCkktEUR%2FGIFLg7IYe1CA%3D%3D; enable_web_push=DISABLE; header_theme_version=CLOSE; rpdid=|(um~R)R)lRm0J'u~kuRRl|mm; DedeUserID=512779590; DedeUserID__ckMd5=5be93575780ece10; buvid_fp_plain=undefined; hit-dyn-v2=1; CURRENT_BLACKGAP=0; CURRENT_FNVAL=4048; fingerprint=76b02d924aa80ba2d4382049a27434fe; is-2022-channel=1; CURRENT_QUALITY=80; home_feed_column=5; browser_resolution=1659-954; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY4MjcwNTcsImlhdCI6MTcyNjU2Nzc5NywicGx0IjotMX0.F5ujr5IoCzyO0b5n1ctLlMr4pmD7zX3NYoStV4SWFTc; bili_ticket_expires=1726826997; buvid_fp=76b02d924aa80ba2d4382049a27434fe; SESSDATA=62c90afd%2C1742201482%2Ca424c%2A92CjD0qVpU4ltlsafDrUzATBdqmreBnGilUSfLFeQL1X2KfYeMEFST4T1msejbLasftnUSVng1ZUE1ek11azJUVVlPWjNnV1dKc05IaHlsTlhFTGtCb2Q1REs0SFNqdS1FOU50VkNIVHlPUWNfa3hJb2NoQ25vd2hnTWQxNkp2R1VGM29IVVc0SUNBIIEC; bili_jct=d83c2b43bd686f38253fc3d7f3fb3846; sid=8s03onkx; b_lsid=AD8CDCFB_1920A496C37; bp_t_offset_512779590=978872997228052480")
	req.Header.Set("origin", "https://www.bilibili.com")
	req.Header.Set("referer", "https://www.bilibili.com/video/BV1AE4depEab/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="128", "Google Chrome";v="128"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status: %d", resp.StatusCode)
	}

	if resp.ContentLength > 0 {
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var resultList contains
		err = json.Unmarshal(bodyText, &resultList)
		if err != nil {
			log.Fatal(err)
		}
		for _, result := range resultList.Data.Replies {
			fmt.Println("一级评论：", result.Content.Message)
			if result.Replies != nil {
				for _, reply := range result.Replies {
					fmt.Println("二级评论：", reply.Content.Message)
				}
			} else {
				fmt.Println("没有二级评论")
			}
		}
		ch <- resultList.Data.Cursor.PaginationReply.NextOffset
	} else {
		log.Println("响应体为空")
	}
}

func channel() {
	ch := make(chan string)
	tag := make(chan string)
	wg.Add(10)
	for i := 0; i < 200000; i++ {
		go spider(tag, ch, i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	close(tag)
}

func main() {
	go channel()
	wg.Wait()
}
