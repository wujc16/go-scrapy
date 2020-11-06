package scrapy

import (
	"fmt"
	"github.com/wujc16/go-scrapy/ds"
	"net/http"
)

const InitProcessorName = "@@INIT_PROCESSOR"

// 全局只有一个 spider
var spider *Spider

// 处理器
type ProcessorInfo struct {
	Name      string
	Processor ProcessorFunc
}

type ProcessorInfos []*ProcessorInfo

type Spider struct {
	siteCrawled    int32  // 已经爬过的网站数
	currentUrl     string // 当前需要爬的 URL
	context        *Context
	ProcessorInfos ProcessorInfos // 注册的处理器函数{ProcessorName, ProcessorFunc}
	processorMap   map[string]ProcessorFunc
	waitQueue      *ds.Queue // 要完成的爬取任务会添加到这里
	//processQueue 	*ds.Queue // 正在做的爬取任务会添加到这里
	itemQueue *ds.Queue // 处理 item 的队列
}

type ParserFunc func(html string)

type ProcessorFunc func(ctx *Context, response *http.Response) *ProcessorResult

type ProcessorResult struct {
	ItemMaps      map[string]interface{}
	UrlProcessors map[string]string
	ShallStop     bool
}

type ProcessorChain []ProcessorFunc

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
	e := s.flushQueue()
	if e != nil {
		panic(e)
	}
}

func (s *Spider) flushQueue() error {
	var resp *http.Response
	var err error
	for true {
		if s.waitQueue.GetSize() > 0 {
			e, _ := s.waitQueue.Dequeue()
			eMap := e.(map[string]string)
			mapKeys := getMapKeys(eMap)
			for idx := range mapKeys {
				key := mapKeys[idx]
				url := eMap[key]
				processorF := s.processorMap[key]
				if processorF == nil {
					return fmt.Errorf("processorFunc not registered")
				}
				resp, err = http.Get(url)
				if err != nil {
					panic(err)
				}
				res := processorF(s.context, resp)
				if res != nil {
					for url := range res.UrlProcessors {
						processorName := res.UrlProcessors[url]
						newNamedProcessor := map[string]string{
							processorName: url,
						}
						s.waitQueue.Enqueue(newNamedProcessor)
					}
				}
			}
		} else {
			break
		}
	}
	defer resp.Body.Close()
	return nil
}

func (s *Spider) GetSiteCrawled() int32 {
	return s.siteCrawled
}

func (s *Spider) GetCurrentUrl() string {
	return s.currentUrl
}

func (s *Spider) Context() *Context {
	return s.context
}

func NewSpider(url string, initProcessorFunc ProcessorFunc) (*Spider, error) {
	if spider.currentUrl != "" {
		return spider, nil
	}
	spider.processorMap[InitProcessorName] = initProcessorFunc
	spider.context = &Context{spider: spider, index: 0}
	spider.currentUrl = url
	return spider, nil
}

func init() {
	spider = &Spider{
		siteCrawled:  0,
		currentUrl:   "",
		processorMap: make(map[string]ProcessorFunc),
		waitQueue:    ds.NewQueue(),
	}
}

func getMapKeys(m map[string]string) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
