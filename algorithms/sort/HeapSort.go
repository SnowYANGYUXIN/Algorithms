package sort

import (
	"fmt"
	"main.go/dataStructure/heap"
)

//堆排序
func HeapSort(arr []int) {
	h := heap.NewMaxHeap(len(arr))
	//nlogn
	for _, v := range arr {
		h.Add(v)
	}
	//nlogn
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = h.ExtractMax()
	}
}

func HeapSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	heap.Heapify(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDownArr(arr, 0, i)
	}
}

func shiftDownArr(arr []int, k, n int) {
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && arr[j] < arr[j+1] {
			j++
		}
		if arr[j] <= arr[k] {
			break
		}
		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}

func TestHeapSort() {
	arr := []int{7, 2, 5, 3, 8, 9, 6, 4, 1}
	HeapSort2(arr)
	fmt.Println(arr)
}
