package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

// 开始爬取页面并写入文件
func GetV2ex(url, spath string) {
	posts := make([]Post, 100)
	posts = GetPosts(posts, url)
	path := filepath.FromSlash(spath)
	WriteToFile(posts, path)
	fmt.Println("写入完毕")
}
// 从文件中读取并打印到终端
func ShowFile(spath string, tag string) {
	path := filepath.FromSlash(spath) // 把win下的路径转化为linux的路径
	files, err := ioutil.ReadDir(path)
	CheckFatal(err)
	for _, f := range files {
		if strings.Contains(f.Name(), tag) {
			data := ReadFromFile(path + f.Name())
			posts := UnMarshalJson(data)
			for _, p := range posts {
				p.Show()
			}
			fmt.Printf("读出文件%s", f.Name())
			fmt.Println()
		}
	}
}

// 命令处理
// -w 爬取并写入文件；-r 读取文件并打印； -t 定时任务
// 1. c creative; 2. p play;3. a apple;4. j jobs;5. t  tech
func HandleCommand(args []string) (tag string, flag int) {
	switch args[1] {
	case "-w":
		flag = 1
	case "-r":
		flag = 2
	case "-t":
		return "", 3
	default:
		return "", 0
	}
	switch args[2] {
	case "c":
		fmt.Println("creative")
		tag = "creative"
	case "p":
		fmt.Println("play")
		tag = "play"
	case "a":
		fmt.Println("apple")
		tag = "apple"
	case "j":
		fmt.Println("jobs")
		tag = "jobs"
	case "t":
		fmt.Println("tech")
		tag = "tech"
	default:
		return "", 0
	}
	return tag, flag
}

// 命令执行
func ExecuteCommand(tag string, flag int) {
	switch flag {
	case 1: // 爬取v2ex
		url := DOMAIN_URL + tag
		spath :=  GenerateFileName(FILE_PATH, tag)
		GetV2ex(url, spath)
	case 2: // 从已爬取中读
		ShowFile(FILE_PATH, tag)
	case 3:	// 定时任务
		StartTimer(TimerWordMin)
		time.Sleep(time.Minute*10)
		fmt.Println("定时任务开始")
	default:
		fmt.Println("命令错误")
	}
}