package array

import "fmt"

type Array struct {
	Data     []int
	Length   int
	Capacity int
}

func NewArray(cap int) *Array {
	return &Array{
		Data:     make([]int, cap),
		Length:   0,
		Capacity: cap,
	}
}

func NewInArray(arr []int, cap int) *Array {
	return &Array{
		Data:     arr,
		Length:   len(arr),
		Capacity: cap,
	}
}

func (a *Array) Append(e int) {
	a.AddIndex(a.Length, e)
}

func (a *Array) AddIndex(index int, e int) {
	if index < 0 || index > a.Length {
		fmt.Println("add index failed index err")
		return
	}
	if a.Length == a.Capacity {
		a.resize(a.Length * 2)
	}
	for i := a.Length - 1; i >= index; i-- {
		a.Data[i+1] = a.Data[i]
	}
	a.Data[index] = e
	a.Length++
}

func (a *Array) AddFirst(e int) {
	a.AddIndex(0, e)
}

func (a *Array) AddLast(e int) {
	a.AddIndex(a.Length, e)
}

func (a *Array) IsEmpty() bool {
	return a.Length == 0
}

func (a *Array) Get(index int) int {
	if index < 0 || index > a.Length {
		fmt.Println("get failed index err")
		return -1
	}
	return a.Data[index]
}

func (a *Array) GetFirst() int {
	return a.Get(0)
}

func (a *Array) GetLast() int {
	return a.Get(a.Length - 1)
}

func (a *Array) Set(index int, e int) {
	if index < 0 || index > a.Length {
		fmt.Println("set failed index err")
	}
	a.Data[index] = e
}

func (a *Array) Contains(e int) bool {
	for i := 0; i < a.Length; i++ {
		if a.Data[i] == e {
			return true
		}
	}
	return false
}

//查找元素e的索引
func (a *Array) Find(e int) int {
	for i := 0; i < a.Length; i++ {
		if a.Data[i] == e {
			return i
		}
	}
	return -1
}

func (a *Array) Remove(index int) int {
	if index < 0 || index > a.Length {
		fmt.Println("remove failed index err")
		return -1
	}
	var ret = a.Data[index]
	for i := index; i < a.Length-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	a.Length--
	if a.Length <= a.Capacity/4 { //防止复杂度震荡
		a.resize(a.Length / 2)
	}
	return ret
}

func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

func (a *Array) RemoveLast() int {
	return a.Remove(a.Length - 1)
}

func (a *Array) RemoveElement(e int) {
	num := a.Find(e)
	a.Remove(num)
}

func (a *Array) resize(cap int) {
	newData := make([]int, cap+1)
	for i := 0; i < a.Length; i++ {
		newData[i] = a.Data[i]
	}
	a.Data = newData
	a.Capacity = cap + 1
}

func (a *Array) GetSize() int {
	return a.Length
}

func (a *Array) Swap(i, j int) {
	if i < 0 || i >= a.Length || j < 0 || j >= a.Length {
		fmt.Println("index err")
		return
	}
	a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
}

func (a *Array) ToString() {
	fmt.Printf("array len is %d , cap is %d\n", a.Length, a.Capacity)
	fmt.Printf("[")
	for i := 0; i < a.Length; i++ {
		fmt.Printf("%d", a.Data[i])
		if i != a.Length-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]\n")
}

func TestArray() {
	arr := NewArray(0) //动态数组
	arr.Append(2)
	arr.Append(3)
	arr.Append(5)
	arr.ToString()
}
