package queue

import "fmt"

//225 用队列实现栈
type MyStack struct {
	ArrayQueue
	Length int
}

func NewMyStack(cap int) *MyStack {
	return &MyStack{*NewArrayQueue(cap), 0}
}

func (m *MyStack) Push(e int) {
	m.EnQueue(e)
	m.Length++
}

func (m *MyStack) Pop() int {
	if m.IsEmpty() {
		fmt.Println("no pop , stack is empty")
		return -1
	}
	queue := NewMyStack(m.Capacity)
	for i := 0; i < m.Length-1; i++ {
		queue.Push(m.DeQueue())
	}
	num := m.DeQueue()
	m.ArrayQueue = queue.ArrayQueue
	m.Length=queue.Length
	return num
}

func (m *MyStack) Top() int {
	if m.IsEmpty() {
		fmt.Println("stack is empty , top is null")
		return -1
	}
	num := m.Pop()
	m.Push(num)

	return num
}

func (m *MyStack) Empty() bool {
	return m.IsEmpty()
}

func (m *MyStack) ToString() {
	fmt.Printf(" [")
	for i := 0; i < m.Length; i++ {
		fmt.Printf("%d", m.Data[i])
		if i != m.Length-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("] top\n")
}

func TestMyStack() {
	my := NewMyStack(4)
	my.Push(5)
	my.Push(4)
	my.Push(3)
	my.ToString()
	my.Pop()
	my.ToString()

	fmt.Println(my.Top())
}
