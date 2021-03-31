package stack

import (
	"fmt"
	"main.go/dataStructure/array"
)

type Stack interface {
	Push(e int)
	Pop() int
	Peek() int
	GetSize() int
	IsEmpty() bool
}

type ArrayStack struct {
	array.Array
}

func NewArrayStack(cap int) *ArrayStack {
	return &ArrayStack{*array.NewArray(cap)}
}

func (a *ArrayStack) Push(e int) {
	a.AddLast(e)
}

func (a *ArrayStack) Pop() int {
	return a.RemoveLast()
}

func (a *ArrayStack) Peek() int {
	return a.GetLast()
}

func (a *ArrayStack) GetSize() int {
	return a.Array.GetSize()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.Array.IsEmpty()
}

//实现多态，因为只有动态数组为基础的栈才有此方法
func (a *ArrayStack) GetCapacity() int {
	return a.Capacity
}

func (a *ArrayStack) ToString() {
	//fmt.Printf("stack len is %d , cap is %d\n", a.Length, a.Capacity)
	fmt.Printf("stack [")
	for i := 0; i < a.GetSize(); i++ {
		fmt.Printf("%d", a.Data[i])
		if i != a.Length-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("] top\n")
}

func TestStack() {
	stack := NewArrayStack(2) //栈
	stack.Push(3)
	stack.Push(7)
	stack.Push(9)
	stack.ToString()
	stack.Pop()
	stack.ToString()
	stack.Pop() //触发删容机制
	stack.ToString()
}
