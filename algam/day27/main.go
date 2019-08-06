package main

import (
	"fmt"
)

func main() {
	//stack := Constructor()
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Top())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Top())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())

	q := Constructor()
	q.Push(3)
	q.Push(2)
	q.Push(1)
	fmt.Println(q.stack)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
	//fmt.Println(q.Pop())
	//fmt.Println(q.Peek())
}

type MyQueue struct {
	stack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stack = append([]int{x}, this.stack...)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return v
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack[len(this.stack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
