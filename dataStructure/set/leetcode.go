package set

//349 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	s1 := NewBSTSet()
	var res []int
	for _, v := range nums1 {
		s1.Add(v)
	}
	for _, v := range nums2 {
		if s1.Contains(v) {
			res = append(res, v)
			s1.Remove(v)
		}
	}
	return res
}
