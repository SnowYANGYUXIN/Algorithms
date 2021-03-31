package search

import (
	"fmt"
)

//对于有序数列，才能用二分查找法
//其时间复杂度是logn 没有将排序时间算进去
//但是对于多次查找，排一次序其均摊复杂度还是很低的
func BinarySearchR(arr []int, k int) int {
	return binarySearchR(arr, k, 0, len(arr)-1)
}

func binarySearchR(arr []int, k, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if arr[mid] == k {
		return mid
	} else if arr[mid] > k {
		return binarySearchR(arr, k, l, mid-1)
	} else {
		return binarySearchR(arr, k, mid+1, r)
	}
}

//非递归的二分搜索法
func BinarySearch(arr []int, k int) int {
	l := 0
	r := len(arr) - 1
	for ; l <= r; {
		mid := (l + r) / 2
		if arr[mid] == k {
			return mid
		} else if arr[mid] > k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//在arr[l,r)中寻找k
func BinarySearch2(arr []int, k int) int {
	l := 0
	r := len(arr) - 2
	for ; l <= r; {
		mid := (l + r) / 2
		if arr[mid] == k {
			return mid
		} else if arr[mid] > k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//查找大于target的最小值
func Upper(arr []int, k int) int {
	l := 0        //比arr最小值还要小则返回索引0
	r := len(arr) //加1 因为如果传入的值比arr的最大值都要大，则返回不存在的一个索引

	for ; l < r; { //因为当l=r时一定有解且是我们想要的解  而二分查找法中<=是因为即使l=r时，也不一定就找得到
		mid := (l + r) / 2
		if arr[mid] > k {
			r = mid //此时不加1，因为前面可能没有比k大的值，mid可能是我们的结果，不能丢掉
		} else {
			l = mid + 1
		}
	}
	return r
}

//ceil
//如果存在元素，则返回最大的索引(因为可能有重复元素)
//如果不存在，则返回upper
//逻辑和upper一样，只需要找到upper后查看当前索引-1的数是否是查找的数
func Ceil(arr []int, k int) int {
	u := Upper(arr, k)
	if u-1 >= 0 && arr[u-1] == k {
		return u - 1
	} else {
		return u
	}
}

//查找小于target的最大值
func Lower(arr []int, k int) int {
	l := -1 //因为如果k比数组最小值还要小，则返回错误索引
	r := len(arr) - 1
	for ; l < r; {
		mid := (l + r + 1) / 2 //解决方法是改变mid，使得l和r相邻时l!=mid
		if arr[mid] >= k {
			r = mid - 1
		} else {
			l = mid
			//存在一个大坑，即如果l、r相邻，此时mid=l，如果此时arr[mid]<k,则会陷入死循环
			//之所以Upper不会出现这个问题，是因为计算机默认是向下取整，Upper的l始终是会变化的
		}
	}
	return l
}

//lower_floor
//如果数组中存在元素则返回最小索引
//如果数组中不存在元素，则返回lower
func LowerFloor(arr []int, k int) int {
	f := Lower(arr, k)
	if f+1 < len(arr) && arr[f+1] == k {
		return f + 1
	} else {
		return f
	}
}

//leetcode 875爱吃香蕉的珂珂
//对于单调的数组就可以考虑二分搜索法
func minEatingSpeed(piles []int, h int) int {
	r := piles[0]
	for i, v := range piles {
		if v > r {
			r = piles[i]
		}
	}

	l := 1
	for l < r {
		mid := (l + r) / 2
		time := eatingTime(piles, mid)
		if time > h {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func eatingTime(arr []int, k int) int {
	h := 0
	for _, n := range arr {
		if n%k != 0 {
			h += n/k + 1
		} else {
			h += n / k
		}
	}
	return h
}

//leetcode 1011在D天送达包裹的能力
//对于单调的数组就可以考虑二分搜索法
func shipWithinDays(weights []int, D int) int {
	r := 0
	for _, v := range weights {
		r += v
	}
	l := weights[0]
	for i, v := range weights {
		if v > l {
			l = weights[i]
		}
	}

	for l < r {
		mid := (l + r) / 2
		d := shipTime(weights, mid)
		if d > D {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func shipTime(arr []int, k int) int {
	d := 0
	max := 0
	for _, v := range arr {
		if v+max <= k {
			max += v
		} else {
			d++
			max = v
		}
	}
	d++
	return d
}

func TestBinarySearch() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
}

