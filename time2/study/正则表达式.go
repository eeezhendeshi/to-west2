package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile(`[a-z]{3}-(\d+)`)
	// 找到 regexp 匹配的第一个字符串  abc-123
	fmt.Println(r.FindString("abc-123|hrd-134"))
	//  返回第一个匹配的原始字符串和括号里面的  [abc-123 123]
	fmt.Println(r.FindStringSubmatch("abc-123|hrd-134"))
	// 返回所有匹配的字符串 [abc-123 hrd-134]
	fmt.Println(r.FindAllString("abc-123|hrd-134", -1))
	//  返回所有匹配的原始字符串和括号里面的 [[abc-123 123] [hrd-134 134]]
	fmt.Println(r.FindAllStringSubmatch("abc-123|hrd-134", -1))
	// 是否匹配成功 true
	fmt.Println(r.MatchString("abc-123"))
}
