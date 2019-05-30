package stack_test

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Queue struct {
	l *list.List
}

func NewQueue() Queue {
	return Queue{
		l: list.New(),
	}
}

func (q *Queue) Push(v int) {
	q.l.PushBack(v)
}

func (q *Queue) Peek() int {
	if q.l.Len() == 0 {
		panic("peek empty queue")
	}
	return q.l.Front().Value.(int)
}

func (q *Queue) Pop() int {
	v := q.Peek()
	q.l.Remove(q.l.Front())
	return v
}

func (q *Queue) Size() int {
	return q.l.Len()
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

type MyStack struct {
	q1, q2 Queue
}

/** Initialize your data structure here. */
// O(1)
func Constructor() MyStack {
	return MyStack{
		q1: NewQueue(),
		q2: NewQueue(),
	}
}

/** Push element x onto stack. */
// O(1)
func (this *MyStack) Push(x int) {
	this.q1.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
// O(n)
func (this *MyStack) Pop() int {
	for this.q1.Size() > 1 {
		v := this.q1.Pop()
		this.q2.Push(v)
	}
	v := this.q1.Pop()
	this.q1, this.q2 = this.q2, this.q1
	return v
}

/** Get the top element. */
// O(n)
func (this *MyStack) Top() int {
	for this.q1.Size() > 1 {
		v := this.q1.Pop()
		this.q2.Push(v)
	}
	return this.q1.Peek()
}

/** Returns whether the stack is empty. */
// O(1)
func (this *MyStack) Empty() bool {
	return this.q1.Size() == 0 && this.q2.Size() == 0
}

func TestStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, false, stack.Empty())
}
