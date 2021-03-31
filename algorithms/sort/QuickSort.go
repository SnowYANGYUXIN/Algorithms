package sort

import (
	"fmt"
	"math/rand"
)

//快速排序法
func sortQuick(arr []int) {
	QuickSort3way(arr, 0, len(arr)-1)
}

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	QuickSort(arr, l, p-1)
	QuickSort(arr, p+1, r)
}

//存在问题，即如果是完全有序的数组，则每次递归左边都是nil，右边是比原先数组少一个元素的数组，即会递归n次
//此时时间复杂度为n²，而栈的深度也会为n，会导致栈溢出的情况
//就是使[l+1,j]<v,[j+1,i]>v
func partition(arr []int, l, r int) int {
	//生成[l,r]之间的随机索引，防止是有序数组而导致算法性能退化
	p := rand.Intn(r-l+1) + l
	arr[p], arr[l] = arr[l], arr[p] //将其随机的数作为数组的一个元素

	j := l //指向比v小的最后一个元素
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++                             //加了1后指向比v大的第一个元素
			arr[i], arr[j] = arr[j], arr[i] //等于将这个数插入到比v小的数组的到最后一个
		}
	}
	arr[l], arr[j] = arr[j], arr[l] //因为j是指向比v小的最后一个元素，这次交换将v变成了左边都比v小右边都比v大
	return j
}

//双路排序法
//快速排序法还是存在问题，即如果对于完全相同的一组数，还是会将其分成一边为空一边少一个元素
//使用双路快速排序法 目的是将数组元素尽可能的分散在标定元素两侧
//目的是arr[l+1,i-1]<=v arr[j+1,r]>=v
func partition2(arr []int, l, r int) int {
	//生成[l,r]之间的随机索引，防止是有序数组而导致算法性能退化
	p := rand.Intn(r-l+1) + l
	arr[p], arr[l] = arr[l], arr[p] //将其随机的数作为数组的一个元素

	//arr[l+1,i-1]<=v arr[j+1,r]>=v
	i := l + 1
	j := r
	for ; ; {
		for ; i <= j && arr[i] < arr[l]; {
			i++
		}
		for ; i <= j && arr[j] > arr[l]; {
			j--
		}
		if i >= j {
			break
		}
		//当i遇到比l大的元素就停下，j遇上比l小的元素就停下，进行交换，使得arr[l+1,i-1]<=v arr[j+1,r]>=v
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//三路快速排序法
//在二路快排中，虽然相同元素的数组中的元素被均摊到了两边，但是可以不必要再对相同元素再进行排序
func QuickSort3way(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := rand.Intn(r-l+1) + l
	arr[p], arr[l] = arr[l], arr[p]
	//arr[l+1,lt]<v arr[lt+1,i-1]==v arr[gt,r]>v
	//lt指向小于v的最后一个元素，gt指向大于v的第一个元素
	lt := l
	i := l + 1
	gt := r + 1
	for ; i < gt; {
		if arr[i] < arr[l] {
			lt++
			arr[lt], arr[i] = arr[i], arr[lt]
			i++
		} else if arr[i] > arr[l] {
			gt--
			arr[gt], arr[i] = arr[i], arr[gt] //此时不用i++，因为此时来了一个新的元素arr[gt]，这个元素还没有比较
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	//arr[l,lt-1]<v arr[lt,gt-1]==v arr[gt,r]>v
	QuickSort3way(arr, l, lt-1)
	QuickSort3way(arr, gt, r)
}

//leetcode 75 颜色分类
func sortColors(nums []int) {
	//if len(nums) <= 1 {
	//	return
	//} else if len(nums) == 2 {
	//	if nums[0] > nums[1] {
	//		nums[0], nums[1] = nums[1], nums[0]
	//	}
	//	return
	//}

	var p = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			p = i
			break
		}
	}

	nums[p], nums[0] = nums[0], nums[p]
	lt := 0
	i := lt + 1
	gt := len(nums)
	for ; i < gt; {
		if nums[i] < nums[0] {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else if nums[i] > nums[0] {
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		} else {
			i++
		}
	}
	nums[0], nums[lt] = nums[lt], nums[0]
}

//leetcode 215 数组中的第k个最大的元素
func findKthLargest(nums []int, k int) int {
	findKthLargestSort(nums, 0, len(nums)-1, k)
	return nums[(k - 1)]
}

func findKthLargestSort(nums []int, l, r, k int) {
	p := findKthLargestPartition(nums, l, r)
	if p+1 == k {
		return
	} else if p+1 < k {
		findKthLargestSort(nums, p+1, r, k)
	} else {
		findKthLargestSort(nums, l, p-1, k)
	}
}

func findKthLargestPartition(arr []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
	arr[p], arr[l] = arr[l], arr[p]

	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] > arr[l] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//leetcode 最小k个数
func smallestK(arr []int, k int) []int {
	smallestKSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func smallestKSort(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	if p == k {
		return
	} else if p < k {
		smallestKSort(nums, p+1, r, k)
	} else {
		smallestKSort(nums, l, p-1, k)
	}
}

func TestQuickSort() {
	//sort := []int{4, 7, 2, 6, 8, 5, 9, 1}
	//sortQuick(sort)
	//fmt.Println(sort)

	arr := smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4)
	fmt.Println(arr)
}
