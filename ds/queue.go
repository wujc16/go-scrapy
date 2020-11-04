package ds

// 队列
type Queue struct {
	Data interface{}
}

// 入队
func (q *Queue) Enqueue(data interface{}) {
}

// 出队，元素从队首移出，返回新的元素
func (q *Queue) Dequeue() interface{} {
	return q
}
