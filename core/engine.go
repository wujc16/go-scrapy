package scrapy

import (
	"fmt"
	"github.com/wujc16/go-scrapy/network"
)

// 处理器
type ProcessorInfo struct {
	Name      string
	Processor ProcessorFunc
}

type ProcessorInfos []*ProcessorInfo

type Spider struct {
	siteCrawled    int32 // 已经爬过的网站数
	currentUrl     string
	ProcessorInfos ProcessorInfos
}

type ParserFunc func(html string)
type ProcessorFunc func(url string, f ParserFunc)
type ProcessorChain []ProcessorFunc

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

// 注册 Processor 处理器
func (s *Spider) Register(name string, processor ProcessorFunc) {
	processorInfo := &ProcessorInfo{
		Name:      name,
		Processor: processor,
	}
	// TODO 注册的时候需要进行重名检查，最好使用 Map 数据结构
	newProcessorInfos := append(s.ProcessorInfos, processorInfo)
	s.ProcessorInfos = newProcessorInfos
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
