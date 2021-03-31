package sort

import "fmt"

//希尔排序
//基本思想：让数组越来越有序
//不能只处理相邻逆序对
//对元素间距为n/2的所有数组做插入排序，对元素间距为n/4的所有数组做插入排序，...对元素间距为1进行插入排序
//看上去有四重循环，实际上其因为数组会越来越有序，循环n的次数也会越来越少，其复杂度是大于nlogn而小于n²，在数组小的时候其性能可能还会超越nlogn
//并且没有使用递归，仅依靠循环就完成了排序
func ShellSort(arr []int) {
	h := len(arr) / 2
	for h >= 1 {
		for start := 0; start < h; start++ {
			//对data[start,start+h,start+2h...]进行插入排序
			for i := start + h; i < len(arr); i += h {
				var j int
				t := arr[i]
				for j = i; j-h >= 0 && t < arr[j-h]; j -= h {
					arr[j] = arr[j-h]
				}
				arr[j] = t
			}
		}
		h /= 2
	}
}

//优化
//四重循环压缩成三重循环，但是复杂度并没有改变
//不用分别对分开的数组对分别使用插入排序,可以直接交替对数组对使用插入排序
func ShellSort2(arr []int) {
	h := len(arr) / 2
	for h >= 1 {
		//对data[h...]进行插入排序
		for i := h; i < len(arr); i++ {
			var j int
			t := arr[i]
			for j = i; j-h >= 0 && t < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = t
		}
		h /= 2
	}
}

//步长序列
//步长序列不同，复杂度分析不同，因为步长序列是一个超参数
func ShellSort3(arr []int) {
	h := 0
	for h < len(arr) {
		h = 3*h + 1
	}
	for h >= 1 {
		//对data[h...]进行插入排序
		for i := h; i < len(arr); i++ {
			var j int
			t := arr[i]
			for j = i; j-h >= 0 && t < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = t
		}
		h /= 3
	}
}

func TestShellSort() {
	arr := []int{1, 7, 6, 4, 8, 3, 9, 5, 2}
	ShellSort(arr)
	fmt.Println(arr)
}
