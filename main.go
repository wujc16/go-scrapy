package main

import (
	"fmt"
	scrapy "github.com/wujc16/go-scrapy/core"
)

func initProcessor(resp string) *scrapy.ProcessorResult {
	return nil
}

func articleProcessor(resp string) *scrapy.ProcessorResult {
	fmt.Println("article resp")
	return nil
}

func chapterProcessor(resp string) *scrapy.ProcessorResult {
	fmt.Println(" chapter resp")
	return nil
}

func main() {
	// 开始 1. 首先支持一个入口
	// 发现 2. 洋葱模型
	// 调度
	// 终止
	spider, err := scrapy.NewSpider("https://www.qidian.com", initProcessor)
	fmt.Println(scrapy.InitProcessorName)
	if err != nil {
		panic(err)
	}
	spider.Register("article", articleProcessor)
	spider.Register("chapter", chapterProcessor)
	fmt.Println(spider)
	spider.Run()
}
