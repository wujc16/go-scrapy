package scrapy

type Context struct {
	spider *Spider // 所属的 spider
	index  int32
}

func (ctx *Context) Spider() *Spider {
	return ctx.spider
}
