package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"log"
)

// 爬取页面
func GetPosts(posts []Post, url string) []Post {
	res, err := http.Get(url)
	CheckNormal(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckNormal(err)
	// 分析页面
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		posts[i].Title = s.Find(".item_title a").Text()
		posts[i].Node = s.Find(".node").Text()
		posts[i].CountNum = s.Find(".count_livid").Text()
		posts[i].Url, _ = s.Find(".item_title a").Attr("href")
		if posts[i].CountNum == "" {
			posts[i].CountNum = "0"
		}
	})
	return filterData(posts)
}

// 处理数据
func filterData(posts []Post) []Post {
	f, s := 0, 0 // 1. f=0 时，代表当前在有效数据之前；
	// 2、f=1 时，代表初次进入有效数据, 删除之前无效数据；
	// 3、f=2 时，代表从有效数据进入无效数据，记录下第一个位置，令 f = 1
	for k, v := range posts {
		if v.Title != "" && f == 0{
			f = 1
			s = k
			posts = posts[k:]
		}
		if v.Title == "" && f == 1{
			f = 2
			s = k - s - 1
		}
		if v.Title != "" && f == 2{
			f = 1
			posts = append(posts[:s], posts[:k+1]...)
		}
	}
	posts = posts[:s]
	return posts
}
// 1. -c creative; 2. -p play;3. -a apple;4. -j jobs;5. -t tech
func TimerWordMin() {
	tags := []string {
		"creative",
		"play",
		"apple",
		"jobs",
		"tech",
	}
	for _, tag := range tags {
		url := DOMAIN_URL + tag
		spath :=  GenerateFileName(FILE_PATH, tag)
		GetV2ex(url, spath)
		fmt.Println("定时任务执行...")
	}
}