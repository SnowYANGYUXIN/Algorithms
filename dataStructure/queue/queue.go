package queue

import (
	"fmt"
	"main.go/dataStructure/array"
)

type Queue interface {
	EnQueue(e int)
	DeQueue() int //出队
	GetFront() int
	GetSize() int
	IsEmpty() bool
}

func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{*array.NewArray(cap)}
}

type ArrayQueue struct {
	array.Array
}

func (a *ArrayQueue) EnQueue(e int) {
	a.Append(e)
}

func (a *ArrayQueue) DeQueue() int {
	return a.RemoveFirst()
}

func (a *ArrayQueue) GetFront() int {
	return a.GetFirst()
}

func (a *ArrayQueue) GetSize() int {
	return a.Array.GetSize()
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.Array.IsEmpty()
}

func (a *ArrayQueue) ToString() {
	//fmt.Printf("stack len is %d , cap is %d\n", a.Length, a.Capacity)
	fmt.Printf("queue front[")
	for i := 0; i < a.GetSize(); i++ {
		fmt.Printf("%d", a.Data[i])
		if i != a.Length-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]\n")
}

func TestQueue(){
	queue := NewArrayQueue(5) //队列
	queue.EnQueue(5)
	queue.EnQueue(6)
	queue.ToString()
	queue.DeQueue()
	queue.ToString()
}
