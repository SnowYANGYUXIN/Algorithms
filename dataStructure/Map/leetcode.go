package Map

// 350两个数组的交集II
func intersect(nums1 []int, nums2 []int) []int {
	m := NewBSTMap()
	var res []int
	for _, v := range nums1 {
		if m.Contains(v) {
			m.Set(v, m.Get(v)+1)
		} else {
			m.Add(v, 1)
		}
	}
	for _, v := range nums2 {
		if m.Contains(v) && m.Get(v) > 0 {
			res = append(res, v)
			m.Set(v, m.Get(v)-1)
		}
	}
	return res
}
