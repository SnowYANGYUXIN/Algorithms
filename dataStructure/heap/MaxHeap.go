package heap

import (
	"fmt"
	"main.go/dataStructure/array"
)

//二叉堆是一棵完全二叉树
//完全二插树 把元素顺序排列成树的形状(一层一层的装)
//(最大堆)堆中的某个结点总是不大于其父亲结点，但是结点的大小和它所处的层次是没有关系的
//且可以用数组存储二叉堆，且满足
//索引0置空 parent=i/2 left=i*2 right=i*2+1
//索引0可用 parent=(i-1)/2 left=i*2+1 right=i*2+2

type MaxHeap struct {
	array.Array
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{*array.NewArray(capacity)}
}

func (h *MaxHeap) Size() int {
	return h.Array.Length
}

func (h *MaxHeap) IsEmpty() bool {
	return h.Array.IsEmpty()
}

//返回完全二叉树的数组表示中，一个索引所表示的父亲结点所对应的索引
func parent(index int) int {
	if index == 0 {
		fmt.Println("index 0 no parent")
		return 0
	}
	return (index - 1) / 2
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) Add(e int) {
	//因为是完全二叉树，按照一层一层从左到右排序，则对应插入元素就是向数组的最后面添加元素
	h.Array.AddLast(e)
	h.shiftUp(h.Array.GetSize() - 1)
}

//因为二叉堆满足结点总是不大于父亲结点,则如果插入元素大于父亲结点，则需要和父亲结点交换位置，直至其小于父亲结点
func (h *MaxHeap) shiftUp(k int) {
	for k > 0 && h.Array.Get(parent(k)) < h.Array.Get(k) {
		h.Array.Swap(k, parent(k))
		k = parent(k)
	}
}

func (h *MaxHeap) FindMax() int {
	if h.Array.Length == 0 {
		fmt.Println("array length is 0")
		return 0
	}
	return h.Array.Get(0)
}

//取出元素 取出最大的元素 就是删除头结点
func (h *MaxHeap) ExtractMax() int {
	res := h.FindMax()
	//将尾结点和头结点交换，再删除尾结点，相当于删除了头结点
	h.Array.Swap(0, h.Array.Length-1)
	h.Array.RemoveLast()
	//但是此时不满足完全二叉树的性质，需要进一步整理
	h.shiftDown(0)
	return res
}

//此时的头结点就是原来的尾结点
//从头结点就是，找出自己左右孩子中最大的值，如果自己比它小，则进行交换
func (h *MaxHeap) shiftDown(k int) {
	for leftChild(k) < h.Array.GetSize() {
		j := leftChild(k)
		if j+1 < h.Array.GetSize() && h.Array.Get(j) < h.Array.Get(j+1) {
			j = rightChild(k)
		}
		//此时data[j]就是左右孩子中最大的值
		if h.Array.Get(j) <= h.Array.Get(k) {
			break
		}
		h.Array.Swap(k, j)
		k = j
	}
}

//取出最大元素后，放入一个元素
//将头顶元素替换，然后shiftDown  则是一次nlogn
func (h *MaxHeap) Replace(e int) int {
	ret := h.FindMax()
	h.Array.Set(0, e)
	h.shiftDown(0)
	return ret
}

//将任意数组转化成成堆形式
//只需拿到最后一个非叶子结点，依次向前shiftDown (抛弃掉了所有叶子结点，相当于少了一半)
//最后一个非叶子结点只需要最后一个结点再求它的parent
//其复杂度是n
func Heapify(arr []int) {
	h := MaxHeap{*array.NewInArray(arr, len(arr))}
	for i := parent(len(arr) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
}

func TestMaxHeap() {
	//h := NewMaxHeap(10)
	//nums := []int{1, 7, 3, 9, 6, 5, 4, 8, 2}
	//for _, v := range nums {
	//	h.Add(v)
	//}
	//fmt.Println(h.Array.Data)
	//for h.Array.Length > 0 {
	//	fmt.Println(h.ExtractMax())
	//}

	a := []int{1, 7, 3, 9, 6, 5, 4, 8, 2}
	Heapify(a)
	fmt.Println(a)
}
