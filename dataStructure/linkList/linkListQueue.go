package linkList

import (
	"fmt"
	"strconv"
)

type LinkedListQueue struct {
	Data int
	Head *LinkedListQueue
	Tail *LinkedListQueue
	Next *LinkedListQueue
	Size int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (l *LinkedListQueue) GetSize() int {
	return l.Size
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkedListQueue) EnQueue(e int) {
	if l.Tail == nil {
		l.Tail = &LinkedListQueue{Data: e} //因为尾结点为空，则头结点肯定也是空的
		l.Head = l.Tail                    //此时头结点和尾结点指向同一个空间
	} else {
		l.Tail.Next = &LinkedListQueue{Data: e}
		l.Tail = l.Tail.Next
	}
	l.Size++
}

func (l *LinkedListQueue) DeQueue() int {
	if l.IsEmpty() {
		fmt.Println("de queue err , queue is empty")
		return -1
	}
	retNode := l.Head
	l.Head = l.Head.Next
	retNode.Next = nil
	if l.Head == nil { //此时链表为空，则将尾结点也置为空
		l.Tail = nil
	}
	l.Size--
	return retNode.Data
}

func (l *LinkedListQueue) GetFront() int {
	if l.IsEmpty() {
		fmt.Println("queue is empty")
		return -1
	}
	return l.Head.Data
}

func (l *LinkedListQueue) ToString() string {
	var str string
	str+="Queue: front "
	cur := l.Head
	for ; cur != nil; {
		str += strconv.Itoa(cur.Data) + "->"
		cur = cur.Next
	}
	str += "NULL tail"
	return str
}

func test(a ,b,c int){

}

func TestLinkedListQueue() {
	//queue := NewLinkedListQueue()
	//for i := 0; i < 5; i++ {
	//	queue.EnQueue(i)
	//}
	//fmt.Println(queue.ToString())
	//fmt.Println(queue.DeQueue())
	//fmt.Println(queue.ToString())

}
