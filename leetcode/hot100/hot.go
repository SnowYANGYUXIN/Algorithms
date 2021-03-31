package hot100

import "fmt"

// 1 两数之和
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = i
	}
	return nil
}

//53 最大自序和时
func maxSubArray(nums []int) int {
	res := nums[0]
	agg := 0
	for _, n := range nums {
		agg = max(agg+n, n)
		res = max(res, agg)
	}
	return res
}

func max(num1, num2 int) int {
	if num1 <= num2 {
		return num2
	} else {
		return num1
	}
}

func TestLeetcodeHot() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1,-1, 5, 1, -5, 4}))
}
