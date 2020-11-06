# go-scrapy
类似 Scrapy 的 Go 爬虫框架

## 安装
使用以下脚本安装本包
```shell script
go get -u github.com/wujc16/go-scrapy
```

## TODO

- 添加 ItemPipeline 功能，允许用户在 ItemPipeline 中对 Result 中返回的 item 结果进行处理；
- 添加 Config 功能，支持配置 Proxy
- 添加多线程支持
- 添加暂停、重新开始功能，允许实时获取当前队列中的待完成任务的个数，已完成任务的个数