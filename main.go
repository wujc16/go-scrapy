package main

import (
	"fmt"
	"github.com/wujc16/go-scrapy/core"
)

func main() {
	// 开始 1. 首先支持一个入口
	// 发现 2. 洋葱模型
	// 调度
	// 终止
	spider, err := scrapy.GetSpider("https://www.qidian.com")
	if err != nil {
		panic(err)
	}
	spider.Run()
	fmt.Println(spider.GetCurrentUrl())
}
