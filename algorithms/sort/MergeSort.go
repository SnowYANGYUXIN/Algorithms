package sort

import (
	"fmt"
	"math"
)

//自顶向下的并归排序
func MergeSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2 //因为
	// l+r如果超出了32位则会有bug
	sort(arr, l, mid)
	sort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func MergeSortBU(arr []int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	n := len(arr)

	//可以先用插入排序对每十五个元素进行排序
	for i := 0; i < n; i += 16 {
		insertionSort(arr, i, int(math.Min(float64(i+15), float64(n-1))))
	}

	//合并区间的长度
	for sz := 1; sz < n; sz *= 2 {
		//合并[i,i+sz-1]和[i+sz,math.Min(i+sz+sz-1, n-1)]
		for i := 0; i+sz < n; i += sz + sz {
			//因为第i+sz+sz-1可能会越界，则需要和数组的最后一个比大小
			//比如有十个元素，分到后面分别为8个和2个，则剩下的2个不能成为i+sz+sz-1(15)，而是取n-1(9)
			merge2(arr, i, i+sz-1, int(math.Min(float64(i+sz+sz-1), float64(n-1))), temp)
		}
	}
}

//自底向上的并归排序
func sortBU(arr []int, l, r int) {
	if r >= len(arr) {
		return
	}
	mid := (l + r) / 2
	sort(arr, l+2, r+2)
	merge(arr, l, mid, r)
}

//归并排序法
func merge(arr []int, l, mid, r int) {
	temp := make([]int, r-l+1)
	copy(temp, arr[l:r+1])
	i := l
	j := mid + 1

	//每轮循环为arr[k]赋值
	for k := l; k <= r; k++ {
		if i > mid { //如果i越界了则说明给定界限的数组前一半以前排完了，则直接将后一半剩下的数依次放入
			arr[k] = temp[j-l] //因为temp是给定界限的数组，有l的偏移量
			j++
			continue
		} else if j > r {
			arr[k] = temp[i-l]
			i++
			continue
		}
		if temp[i-l] <= temp[j-l] {
			arr[k] = temp[i-l]
			i++
		} else {
			arr[k] = temp[j-l]
			j++
		}
	}
}

//归并排序法的优化
func MergeSort2(arr []int) {
	temp := make([]int, len(arr)) //进行内存优化，只开一次空间，防止每次merge时的开辟空间
	copy(temp, arr)
	sort2(arr, 0, len(arr)-1, temp)
}

func sort2(arr []int, l, r int, temp []int) {
	if r-l < 15 {
		insertionSort(arr, l, r) //选择使用插入排序法，因为对于小规模的排序，插入排序法的常数小，反而耗时更少
		return
	}
	mid := l + (r-l)/2 //因为
	// l+r如果超出了32位则会有bug
	sort2(arr, l, mid, temp)
	sort2(arr, mid+1, r, temp)

	if arr[mid] > arr[mid+1] { //如果左数组的最大都比有数组最小还要小，则不需要进行并归
		merge2(arr, l, mid, r, temp)
	}
}

//归并排序法
func merge2(arr []int, l, mid, r int, temp []int) {
	copy(temp[l:r+1], arr[l:r+1])
	i := l
	j := mid + 1

	//每轮循环为arr[k]赋值
	for k := l; k <= r; k++ {
		if i > mid { //如果i越界了则说明给定界限的数组前一半以前排完了，则直接将后一半剩下的数依次放入
			arr[k] = temp[j] //此时，temp也是从l开始取，则不需要偏移量
			j++
			continue
		} else if j > r {
			arr[k] = temp[i]
			i++
			continue
		}
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
	}
}

//如果是有序数组，则merge不会执行则在递归树中，每一个叶子结点复杂度都是1，则总的复杂度为n+2/n+4/n+.... 其复杂度为n
func insertionSort(arr []int, l, r int) {
	for i := l; i <= r; i++ {
		var j int
		num := arr[i]
		for j = i; j-1 >= l && arr[j-1] > num; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = num
	}
}

func TestMergeSort() {
	arr := []int{1, 7, 9, 5, 6, 4, 2, 0}
	MergeSortBU(arr)
	fmt.Println(arr)
}

//leetcode 剑指offer51 求逆序数对个数
var Res int
func reversePairs(nums []int) int {
	Res = 0
	temp := make([]int, len(nums))
	sortLeetcode(nums, 0, len(nums)-1, temp)
	return Res
}

func sortLeetcode(arr []int, l, r int, temp []int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	sortLeetcode(arr, l, mid, temp)
	sortLeetcode(arr, mid+1, r, temp)
	mergeLeetcode(arr, l, mid, r, temp)
}

func mergeLeetcode(arr []int, l, mid, r int, temp []int) {
	copy(temp[l:r+1], arr[l:r+1])
	i := l
	j := mid + 1

	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = temp[j]
			j++
			continue
		} else if j > r {
			arr[k] = temp[i]
			i++
			continue
		}
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			Res+=mid-i+1
			j++
		}
	}
}
