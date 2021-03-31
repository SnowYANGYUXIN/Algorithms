package linkList

import (
	"fmt"
	"strconv"
)

//递归实现链表功能
type LinkedListR struct {
	Data int
	Head *LinkedListR
	Size int
	Next *LinkedListR
}

func NewLinedListR() *LinkedListR {
	return &LinkedListR{Head: &LinkedListR{}}
}

func (l *LinkedListR) Add(index, e int) {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return
	}
	l.Head = l.AddRec(l.Head, index, e)
	l.Size++
}

func (l *LinkedListR) AddRec(node *LinkedListR, index, e int) *LinkedListR {
	if index == 0 {
		return &LinkedListR{Data: e, Next: node}
	}
	node.Next = l.AddRec(node.Next, index-1, e)
	return node
}

func (l *LinkedListR) Remove(index int) int {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return -1
	}
	_, res := l.RemoveRec(l.Head, index)
	l.Size--
	return res
}

func (l *LinkedListR) RemoveRec(node *LinkedListR, index int) (*LinkedListR, int) {
	var res int
	if index == 0 {
		return node.Next, node.Data
	}
	node.Next, res = l.RemoveRec(node.Next, index-1)
	return node, res
}

func (l *LinkedListR) Get(index int) int {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return -1
	}
	_, res := l.GetRec(l.Head, index)
	return res
}

func (l *LinkedListR) GetRec(node *LinkedListR, index int) (*LinkedListR, int) {
	if index == 0 {
		return node, node.Data
	}
	return l.GetRec(node.Next, index-1)
}

func (l *LinkedListR) Set(index, e int) {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return
	}
	l.Head = l.SetRec(l.Head, index, e)
}

func (l *LinkedListR) SetRec(node *LinkedListR, index, e int) *LinkedListR {
	if index == 0 {
		return &LinkedListR{Data: e,Next: node.Next}
	}
	node.Next = l.SetRec(node.Next, index-1, e)
	return node
}

func (l *LinkedListR) GetSize() int {
	return l.Size
}

func (l *LinkedListR) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkedListR) ToString() string {
	var str string
	cur := l.Head
	for ; cur.Next != nil; {
		str += strconv.Itoa(cur.Data) + "->"
		cur = cur.Next
	}
	str += "NULL"
	return str
}

func TestLinkedListR() {
	link := NewLinedListR()
	for i := 0; i < 6; i++ {
		link.Add(i, i+5)
	}
	fmt.Println(link.ToString())
	link.Remove(2)
	fmt.Println(link.ToString())
	fmt.Println(link.Get(3))
	fmt.Println(link.Get(4))
	link.Set(3, 50)
	fmt.Println(link.ToString())
}
