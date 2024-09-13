/*
目的
<p >
         <a href="list.jsp?urltype=tree.TreeTempUrl&wbtreeid=1324" target=_blank class="lm_a" style="float:left;"> 【科技处】</a>
         <a href="content.jsp?urltype=news.NewsContentUrl&wbtreeid=1324&wbnewsid=37022" target=_blank title="关于转发《福建省科协等 21部门关于举办福建省2024年全国科普日活动的通知》的通知" style="">关于转发《福建省科协等 21部门关于举办福建省2024年全国科普日活动的通知》的通知</a>
<span class="fr">2024-08-30</span>
</p>
*/

// 正则表达式
package main

import (
	"fmt"
	"regexp"
)

func main() {
	sstr := "<p >" +
		"<a href=\"list.jsp?urltype=tree.TreeTempUrl&wbtreeid=1324\" target=_blank class=\"lm_a\" style=\"float:left;\"> 【科技处】</a>" +
		"<a href=\"content.jsp?urltype=news.NewsContentUrl&wbtreeid=1324&wbnewsid=37022\" target=_blank title=\"关于转发《福建省科协等 21部门关于举办福建省2024年全国科普日活动的通知》的通知\" style=\"\">关于转发《福建省科协等 21部门关于举办福建省2024年全国科普日活动的通知》的通知</a>" +
		"<span class=\"fr\">2024-08-30</span>" +
		"</p>"
	ret := regexp.MustCompile("[0-9].[a-zA-Z]")
	alls := ret.FindAllStringSubmatch(sstr, -1)
	fmt.Printf("alls: %v\n", alls)
}
