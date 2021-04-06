package link_list

import "fmt"

type LinkedListStack struct {
	LinkList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{*NewLinkList()}
}

func (l *LinkedListStack) GetSize() int {
	return l.Size
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.LinkList.IsEmpty()
}

func (l *LinkedListStack) Push(e int) {
	l.AddFirst(e)
}

func (l *LinkedListStack) Pop() int {
	return l.RemoveFirst()
}

func (l *LinkedListStack) Peek() int {
	return l.GetFirst()
}

func (l *LinkedListStack) ToString() string {
	var str string
	str += "Stack: top "
	str += l.LinkList.ToString()
	return str
}

func TestLinkedListStack() {
	stack := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.ToString())
	fmt.Println(stack.Pop())
	fmt.Println(stack.ToString())
}
