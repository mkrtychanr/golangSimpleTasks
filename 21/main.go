package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list list.List
	len  int
}

func (q *Queue) Init() {
	q.list.Init()
}

func (q *Queue) Push(el interface{}) {
	q.list.PushBack(el)
	q.len++
}

func (q *Queue) Top() (interface{}, bool) {
	if q.len == 0 {
		return 0, false
	}
	return q.list.Front().Value, true
}

func (q *Queue) Pop() {
	if q.len == 0 {
		return
	}
	q.list.Remove(q.list.Front())
	q.len--
}

func main() {
	var q Queue
	q.Init()
	q.Push("asd")
	q.Push(1)
	fmt.Println(q.Top())
	q.Pop()
	fmt.Println(q.Top())
	q.Pop()
	fmt.Println(q.Top())
}
