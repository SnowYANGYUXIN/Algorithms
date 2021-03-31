package queue

import "fmt"

//双端队列
type DeQue struct {
	Data     []int
	Front    int
	Tail     int
	Length   int
	Capacity int
}

func NewDeQue(cap int) *DeQue {
	return &DeQue{
		Data:     make([]int, cap+1),
		Front:    0,
		Tail:     0,
		Length:   0,
		Capacity: cap,
	}
}

func (d *DeQue) AddFront(e int) {
	if (d.Tail+1)%d.Capacity == d.Front {
		d.resize(d.Capacity * 2)
	}
	d.Front = (d.Front + d.Capacity - 1) % d.Capacity
	d.Data[d.Front] = e
	d.Length++
}

func (d *DeQue) AddLast(e int) {
	if (d.Tail+1)%d.Capacity == d.Front {
		d.resize(d.Capacity * 2)
	}
	d.Data[d.Tail] = e
	d.Tail = (d.Tail + 1) % d.Capacity
	d.Length++
}

func (d *DeQue) RemoveFront() int {
	if d.IsEmpty() {
		fmt.Println("queue is empty , can't remove")
		return -1
	}
	num := d.Data[d.Front]
	d.Front = (d.Front + 1) % d.Capacity
	d.Length--
	if d.Length <= d.GetCapacity()/4 && d.GetCapacity()/2 != 0 {
		d.resize(d.GetCapacity() / 2)
	}
	return num
}

func (d *DeQue) RemoveLast() int {
	if d.IsEmpty() {
		fmt.Println("queue is empty , can't remove")
		return -1
	}
	num := d.Data[d.Tail+d.Capacity-1] % d.Capacity
	d.Tail = (d.Tail + d.Capacity - 1) % d.Capacity
	d.Length--
	if d.Length <= d.GetCapacity()/4 && d.GetCapacity()/2 != 0 {
		d.resize(d.GetCapacity() / 2)
	}
	return num
}

func (d *DeQue) resize(cap int) {
	newQue := make([]int, cap+1)
	for i := 0; i < d.Length; i++ {
		newQue[i] = d.Data[(i+d.Front)%d.Capacity]
	}
	d.Data = newQue
	d.Front = 0
	d.Tail = d.Length
	d.Capacity = cap
}

func (d *DeQue) IsEmpty() bool {
	return d.Front == d.Tail
}

func (d *DeQue) GetSize() int {
	return d.Length
}

func (d *DeQue) GetCapacity() int {
	return d.Capacity
}

func (d *DeQue) ToString() {
	fmt.Printf("doble queue len is %d , cap is %d\n", d.Length, d.GetCapacity())
	fmt.Printf("front [")
	for i := d.Front; i < d.Length+d.Front; i++ {
		fmt.Printf("%d", d.Data[i%d.Capacity])
		if (i+1)%d.Capacity != d.Tail {
			fmt.Printf(",")
		}
	}
	fmt.Printf("] tail\n")
}

func TestDeQue(){
	de:=NewDeQue(4)
	de.AddFront(5)
	de.AddLast(4)
	de.AddFront(3)
	de.ToString()
	de.RemoveLast()
	de.ToString()
	de.RemoveFront()
	de.ToString()
}