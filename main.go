package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	scrapy "github.com/wujc16/go-scrapy/core"
	"net/http"
)

// 最初入口 url 的 processor，用来获取到各个小说分类的入口 url
func initProcessor(ctx *scrapy.Context, resp *http.Response) *scrapy.ProcessorResult {
	result := &scrapy.ProcessorResult{
		ShallStop:     true,
		UrlProcessors: make(map[string]string),
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find("div.select-list div.type-filter ul li a").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Attr("href")
		result.UrlProcessors["https:"+val] = "category"
	})
	return result
}

// 处理分类 url 的 processor，可以获取到所有小说的名字并打印
func categoryProcessor(ctx *scrapy.Context, resp *http.Response) *scrapy.ProcessorResult {
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find("div.all-book-list ul.all-img-list li h4 a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	return nil
}

func main() {
	// 开始 1. 首先支持一个入口
	// 发现 2. 通过 spider 的 processorMap 获取 namedProcessor
	// 调度
	// 终止
	spider, err := scrapy.NewSpider("https://www.qidian.com/finish", initProcessor)
	if err != nil {
		panic(err)
	}
	spider.Register("category", categoryProcessor)
	spider.Run()
}
