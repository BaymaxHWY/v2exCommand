package main

import "fmt"

const (
	DOMAIN_URL = "https://www.v2ex.com/?tab="
	FILE_PATH = `D:\go\src\commandlists\fetchv2ex\v2exPosts\`
)

type Post struct {
	Title string 	`json:"title"`
	Node string		`json:"node"`
	Url string		`json:"url"`
	CountNum string	`json:"countNum"`
}

func (p Post) Show() {
	fmt.Print("{\n")
	fmt.Printf("标题： %s\n", p.Title)
	fmt.Printf("节点： %s\n", p.Node)
	fmt.Printf("url： %s\n", p.Url)
	fmt.Printf("回复量： %s\n", p.CountNum)
	fmt.Print("}\n")
}