package Map

//映射map
//典型应用 数据库 id->信息  词频统计 单词->频率
//储存(键，值)数据对的数据结构(Key,Value)
//根据键(Key),寻找值(Value)
type Map interface {
	Add(key, value int)
	Remove(key int) int
	Contains(key int) bool
	Get(key int) int
	Set(key, value int)
	GetSize() int
	IsEmpty() bool
}
