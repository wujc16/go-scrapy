package ds

import "fmt"

// 队列
type Queue struct {
	queue []interface{}
	size  int32
	Head  interface{}
	End   interface{}
}

type Element struct {
	Value interface{}
}

// 入队
func (q *Queue) Enqueue(e interface{}) {
	if q.size == 0 {
		q.Head = e
	}
	q.size += 1
	q.End = e
	q.queue = append(q.queue, e)
}

// 出队，元素从队首移出，返回新的元素
func (q *Queue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, fmt.Errorf("cannot dequeue empty queue")
	}
	q.size -= 1
	q.queue = q.queue[1:]
	lastHead := q.Head
	if q.size == 0 {
		q.End = nil
		q.Head = nil
		return lastHead, nil
	}
	q.Head = q.queue[0]
	return lastHead, nil
}

func (q *Queue) GetSize() int32 {
	return q.size
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
		size:  0,
		Head:  nil,
		End:   nil,
	}
}
