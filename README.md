## v2exCommand
> 在学习go时编写的一个命令行工具
### 功能
1. 按命令从v2ex不同模块爬取帖子并以JSON文件存储在本地
2. 按命令从本地JSON文件中读取并在控制台打印
3. 定时执行任务(还未完全实现)

### 安装
```
cd fetchv2ex // 进入程序目录
go install // 需要go版本在1.11以上
// 运行
fetchv2ex -w t // 爬取
```
### 使用方式
```
// 爬取
fetchv2ex -w (c|p|a|j|t) //从中任选一个（只能选择一个） 分别代表:创意,好玩,Apple, 酷工作，技术分区
// 显示
fetchv2ex -r (c|p|a|j|t) //将会打印出所有有关模块的信息
// 定时任务
fetchv2ex -t  // 目前还在测试
```
### 收获
1. 当函数参数为slice时，如果对slice进行append则在函数外部的slice无法得到改变，因为append会创建一个新的slice
2. 初步了解页面爬取、文件读写等操作