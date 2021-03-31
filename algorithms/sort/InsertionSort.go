package sort

//插入排序法
//原理是将arr[i]插入到数组前面合适的位置
//和选择排序相比特性是，程序会提前终止内循环即对于极端情况(如有序数组)则复杂度为n，而选择排序始终是n²
//对于近乎有序数组(银行处理业务)，插入排序法是更好的选择
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j-1 >= 0; j-- {
			if arr[j-1] > arr[j] { //不用记录i的下标，因为比较的两个数始终是相邻的，用j-1表示即可
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

func InsertionSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		//for j := i; j-1 >= 0; j-- {
		//	if arr[j-1] > arr[j] { //不用记录i的下标，因为比较的两个数始终是相邻的，用j-1表示即可
		//		arr[j], arr[j-1] = arr[j-1], arr[j]
		//	} else {
		//		break
		//	}
		//}

		//内循环优化
		//for j := i; j-1 >= 0 && arr[j-1] > arr[j]; j-- {
		//	arr[j], arr[j-1] = arr[j-1], arr[j]
		//}

		//交换元素优化为赋值元素 交换可看成是三次赋值
		var j int
		num := arr[i]
		for j = i; j-1 >= 0 && arr[j-1] > num; j-- {
			arr[j] = arr[j-1]
		}
		arr[j]=num
	}
}


func TestInsertionSort(){
	InsertionSort([]int{6, 4, 3, 5, 1, 2})
}
