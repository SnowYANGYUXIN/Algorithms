package sort

import "fmt"

//选择排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var max = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				max = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
	}
	fmt.Println(arr)
}

func TestSelectionSort(){
	SelectionSort([]int{6, 4, 3, 5, 1, 2})
}