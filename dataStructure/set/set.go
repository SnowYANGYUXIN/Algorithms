package set

//集合
//每个元素只能存一次，用来去重
//典型应用：客户统计、词汇量统计
type Set interface {
	Add(e int)
	Remove(e int)
	Contains(e int) bool
	GetSize() int
	IsEmpty() bool
}
