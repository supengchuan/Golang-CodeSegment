package container

import "sync"

type (
	Stack struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0, sync.Mutex{}}
}

func (this *Stack) Len() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.length
}

func (this *Stack) Peek() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.length == 0 {
		return nil
	}

	return this.top.value
}

func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(val interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	n := &node{value: val, prev: this.top}
	this.top = n
	this.length++
}
