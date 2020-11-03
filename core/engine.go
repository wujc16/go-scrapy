package scrapy

import (
	"fmt"
	"github.com/wujc16/go-scrapy/network"
)

type Spider struct {
	siteCrawled int32
	currentUrl  string
}

// 全局只有一个 spider
var spider *Spider

func (s *Spider) Run() {
	fmt.Println("Start spider now!")
	res := network.HttpGet(s.currentUrl)
	fmt.Println(res)
}

func (s *Spider) GetSiteCrawled() int32 {
	return s.siteCrawled
}

func (s *Spider) GetCurrentUrl() string {
	return s.currentUrl
}

func GetSpider(url string) (*Spider, error) {
	if spider.currentUrl != "" {
		return spider, nil
	}
	spider.currentUrl = url
	return spider, nil
}

func init() {
	spider = &Spider{
		siteCrawled: 0,
		currentUrl:  "",
	}
}
