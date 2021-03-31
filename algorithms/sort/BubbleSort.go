package sort

import "fmt"

//冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//优化
func BubbleSort2(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		isSwapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}
		if !isSwapped { //如果这一轮数组没有进行交换，说明整个数组已经有序
			break
		}
	}
}

//再优化
func BubbleSort3(arr []int) {
	for i := 0; i < len(arr)-1; {
		lastSwappedIndex := 0
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				lastSwappedIndex = j + 1 //最后一个交换的下标
			}
		}
		i = len(arr) - lastSwappedIndex //i不仅表达的是多少轮，也表示数组最后i位以排好序
		//有了最后交换的下标，表示之后没有进行交换过，表示下标之后的已有序
	}
}

//从后向前的冒泡排序
func BubbleSortEnd(arr []int) {
	for i := 0; i < len(arr)-1; {
		lastSwappedIndex := len(arr) - 1
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				lastSwappedIndex = j - 1
			}
		}
		i = lastSwappedIndex + 1 //因为i是已经拍好序的下标,lastSwappedIndex+1表示从排好序的后一位进行比较
	}
}

func TestBubbleSort() {
	arr := []int{1, 7, 5, 6, 2, 8, 3, 9, 4}
	BubbleSortEnd(arr)
	fmt.Println(arr)
}
