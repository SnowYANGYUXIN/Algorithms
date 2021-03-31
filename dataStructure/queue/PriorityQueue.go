package queue

import (
	"fmt"
	"main.go/dataStructure/heap"
)

//优先队列
//出队顺序和入队顺序无关，和优先级相关
//动态选择优先级高的任务执行
type PriorityQueue struct {
	heap.MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{*heap.NewMaxHeap(0)}
}

func (q *PriorityQueue) GetSize() int {
	return q.MaxHeap.Size()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.MaxHeap.IsEmpty()
}

func (q *PriorityQueue) GetFront() int {
	return q.MaxHeap.FindMax()
}

func (q *PriorityQueue) EnQueue(e int) {
	q.MaxHeap.Add(e)
}

func (q *PriorityQueue) DeQueue() int {
	return q.MaxHeap.ExtractMax()
}

//leetcode 最小的k个数
//使用优先队列(队首是最大值)，维护当前看到最小的k个数
//对于每一个新的数据，如果比这k个数最大的数还要小，则进行替换
func getLeastNumbers(arr []int, k int) []int {
	q := NewPriorityQueue()
	for i := 0; i < k; i++ {
		q.EnQueue(arr[i])
	}
	for i := k; i < len(arr); i++ {
		if  arr[i] < q.GetFront() {
			q.DeQueue()
			q.EnQueue(arr[i])
		}
	}
	return q.MaxHeap.Array.Data[:k]
}

func TestPriorityQueue() {
	fmt.Println(getLeastNumbers([]int{3,2,1},2))
}
