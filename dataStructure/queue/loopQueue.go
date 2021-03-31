package queue

import "fmt"


//循环队列
type LoopQueue struct {
	Data     []int
	Front    int
	Tail     int
	Length   int
	Capacity int
}

func NewLoopQueue(cap int) *LoopQueue {
	return &LoopQueue{
		Data:     make([]int, cap+1), //因为循环队列本身会浪费一个空间大小
		Front:    0,
		Tail:     0,
		Length:   0,
		Capacity: cap,
	}
}

func (l *LoopQueue) GetCapacity() int {
	return l.Capacity
}

func (l *LoopQueue) IsEmpty() bool {
	return l.Front == l.Tail
}

func (l *LoopQueue) GetSize() int {
	return l.Length
}

func (l *LoopQueue) EnQueue(e int) {
	if (l.Tail+1)%l.Capacity == l.Front {
		l.resize(l.GetCapacity() * 2)
	}
	l.Data[l.Tail] = e
	l.Tail = (l.Tail + 1) % l.Capacity
	l.Length++
}

func (l *LoopQueue) DeQueue() int {
	if l.IsEmpty() {
		fmt.Println("loop queue is empty")
		return -1
	}
	num := l.Data[l.Front]
	l.Front = (l.Front + 1) % l.Capacity
	l.Length--
	if l.Length <= l.GetCapacity()/4 && l.GetCapacity()/2 != 0 {
		l.resize(l.GetCapacity() / 2)
	}
	return num
}

func (l *LoopQueue) GetFront() int {
	return l.Data[l.Front]
}

func (l *LoopQueue) resize(cap int) {
	newLoop := make([]int, cap+1)
	for i := 0; i < l.Length; i++ {
		newLoop[i] = l.Data[(i+l.Front)%l.Capacity] //将front重新置为0开始
	}
	l.Front = 0
	l.Tail = l.Length
	l.Data = newLoop
	l.Capacity = cap
}

func (l *LoopQueue) ToString() {
	fmt.Printf("loop queue len is %d , cap is %d\n", l.Length, l.GetCapacity())
	fmt.Printf("front [")
	for i := l.Front; i < l.Length+l.Front; i++ {
		fmt.Printf("%d", l.Data[i%l.Capacity])
		if (i+1)%l.Capacity != l.Tail {
			fmt.Printf(",")
		}
	}
	fmt.Printf("] tail\n")
}

func TestLoopQueue() {
	loop := NewLoopQueue(4)
	loop.EnQueue(3)
	loop.EnQueue(2)
	loop.EnQueue(5)
	loop.ToString()
	loop.DeQueue()
	loop.EnQueue(7)
	loop.ToString()
	loop.EnQueue(9)
	loop.ToString()
}
