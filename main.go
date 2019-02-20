package main

import (
	"os"
)
/*
TODO:
1. 每天自动运行
*/
// -w 爬取并写入文件；-r 读取文件并打印； -t 定时任务
// 1. c creative; 2. p play;3. a apple;4. j jobs;5. t 或者 "" tech
func main() {
	tag, flag := HandleCommand(os.Args)
	ExecuteCommand(tag, flag)
}
