package scrapy

import (
	"container/list"
	"fmt"
	"github.com/wujc16/go-scrapy/network"
)

const InitProcessorName = "@@INIT_PROCESSOR"

// 处理器
type ProcessorInfo struct {
	Name      string
	Processor ProcessorFunc
}

type ProcessorInfos []*ProcessorInfo

type Spider struct {
	siteCrawled    int32          // 已经爬过的网站数
	currentUrl     string         // 当前需要爬的 URL
	ProcessorInfos ProcessorInfos // 注册的处理器函数{ProcessorName, ProcessorFunc}
	// TODO：待爬取的队列
	processorMap map[string]ProcessorFunc
	waitQueue    *list.List
	processQueue *list.List
}

type ParserFunc func(html string)

type ProcessorFunc func(response string) *ProcessorResult

type ProcessorResult struct {
	ItemMaps      map[string]interface{}
	UrlProcessors map[string]string
	ShallStop     bool
}

type ProcessorChain []ProcessorFunc

// 全局只有一个 spider
var spider *Spider

// 注册 Processor 处理器
func (s *Spider) Register(name string, processor ProcessorFunc) {
	if name == InitProcessorName {
		return
	}
	s.processorMap[name] = processor
}

func (s *Spider) Run() {
	fmt.Println("Start spider now!")
	res := network.HttpGet(s.currentUrl)
	fmt.Println(res)
	for true {
		fmt.Println("Hello World")
		if s.waitQueue.Len() > 0 {
		}
		if s.waitQueue.Len() == 0 && s.processQueue.Len() == 0 {
			break
		}
	}
}

func (s *Spider) flushQueue() {

}

func (s *Spider) GetSiteCrawled() int32 {
	return s.siteCrawled
}

func (s *Spider) GetCurrentUrl() string {
	return s.currentUrl
}

func NewSpider(url string, initProcessorFunc ProcessorFunc) (*Spider, error) {
	if spider.currentUrl != "" {
		return spider, nil
	}
	spider.processorMap[InitProcessorName] = initProcessorFunc
	spider.currentUrl = url
	return spider, nil
}

func init() {
	spider = &Spider{
		siteCrawled:  0,
		currentUrl:   "",
		processorMap: make(map[string]ProcessorFunc),
		processQueue: list.New(),
		waitQueue:    list.New(),
	}
}
