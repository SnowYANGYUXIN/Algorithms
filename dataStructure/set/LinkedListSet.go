package set

import (
	"fmt"
	"main.go/dataStructure/link_list"
)

type LinkedListSet struct {
	link_list.LinkList
}

func NewLinkedListSet() *LinkedListSet {
	return &LinkedListSet{*link_list.NewLinkList()}
}

func (l *LinkedListSet) GetSize() int {
	return l.LinkList.GetSize()
}

func (l *LinkedListSet) IsEmpty() bool {
	return l.LinkList.IsEmpty()
}

func (l *LinkedListSet) Contains(e int) bool {
	return l.LinkList.Contains(e)
}

func (l *LinkedListSet) Add(e int) {
	if !l.Contains(e) {
		l.LinkList.AddFirst(e) //没有尾指针，头部添加操作复杂度为1
	}
}

func (l *LinkedListSet) Remove(e int) {
	l.LinkList.RemoveElement(e)
}

func TestLinkedListSet() {
	s := NewLinkedListSet()
	for i := 0; i < 10; i++ {
		s.Add(i)
	}
	fmt.Println(s.GetSize())
	s.Add(5)
	fmt.Println(s.GetSize())
	s.Add(12)
	fmt.Println(s.GetSize())
	s.RemoveElement(9)
	fmt.Println(s.GetSize())
}
