package scrapy

import (
	"fmt"
	"github.com/wujc16/go-scrapy/ds"
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
	waitQueue    *ds.Queue
	processQueue *ds.Queue
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
	m := map[string]string{
		InitProcessorName: s.currentUrl,
	}
	s.waitQueue.Enqueue(m)
	fmt.Println("s.waitQueue", s.waitQueue)
	s.flushQueue()
}

func (s *Spider) flushQueue() {
	for true {
		if s.waitQueue.GetSize() > 0 || s.processQueue.GetSize() > 0 {
			if s.processQueue.GetSize() < 5 {
				e, _ := s.waitQueue.Dequeue()
				s.processQueue.Enqueue(e)
				for true {
					if s.processQueue.GetSize() > 0 {
						e, _ := s.processQueue.Dequeue()
						fmt.Println("e", e)
					} else {
						break
					}
				}
			}
		} else {
			break
		}
	}
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
		processQueue: ds.NewQueue(),
		waitQueue:    ds.NewQueue(),
	}
}
