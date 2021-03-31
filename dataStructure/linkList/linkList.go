package linkList

import (
	"fmt"
	"strconv"
)

type LinkList struct {
	Data      int
	DummyHead *LinkList //虚拟头结点 就不用再插入时判断是否为头结点，因为头结点是没有prev
	Next      *LinkList
	Size      int
}

func NewLinkList() *LinkList {
	return &LinkList{0, &LinkList{}, nil, 0}
}

func (l *LinkList) Add(index, e int) {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return
	}

	prev := l.DummyHead //不动l本身，是因为只有l的头结点才有size和head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	//node := &LinkList{Data: e}
	//node.Next = prev.Next
	//prev.Next = node
	prev.Next = &LinkList{Data: e, Next: prev.Next}
	l.Size++

}


func (l *LinkList) AddFirst(e int) {
	l.Add(0, e)
}

func (l *LinkList) AddLast(e int) {
	l.Add(l.Size, e)
}

func (l *LinkList) Get(index int) int {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return -1
	}
	cur := l.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Data
}

func (l *LinkList) GetFirst() int {
	return l.Get(0)
}

func (l *LinkList) GetLast() int {
	return l.Get(l.Size - 1)
}

func (l *LinkList) Remove(index int) int {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return -1
	}
	prev := l.DummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	retNode := prev.Next
	prev.Next = retNode.Next
	retNode.Next = nil
	l.Size--
	return retNode.Data
}

func (l *LinkList) RemoveFirst() int {
	return l.Remove(0)
}

func (l *LinkList) RemoveLast() int {
	return l.Remove(l.Size - 1)
}

func (l *LinkList) RemoveElement(e int) {
	prev := l.DummyHead
	for prev.Next != nil {
		if prev.Next.Data == e {
			l.Size--
			break
		}
		prev = prev.Next
	}
	if prev.Next != nil {
		delNode := prev.Next
		prev.Next = delNode.Next
	}
}

func (l *LinkList) Set(index, e int) {
	if index < 0 || index > l.Size {
		fmt.Println("index err , add e err")
		return
	}
	cur := l.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Data = e
}

func (l *LinkList) Contains(e int) bool {
	cur := l.DummyHead.Next
	for ; cur != nil; {
		if cur.Data == e {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (l *LinkList) GetSize() int {
	return l.Size
}

func (l *LinkList) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkList) ToString() string {
	var str string
	cur := l.DummyHead.Next
	for ; cur != nil; {
		str += strconv.Itoa(cur.Data) + "->"
		cur = cur.Next
	}
	str += "NULL"
	return str
}

func TestLinkList() {
	link := NewLinkList()
	for i := 0; i < 5; i++ {
		link.AddFirst(i)
	}
	fmt.Println(link.ToString())
	link.Remove(2)
	fmt.Println(link.ToString())
}
