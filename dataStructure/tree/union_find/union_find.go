package union_find

//并查集
//解决连接和路径问题 网络中节点间的连接状态 网络是个抽象概念:用户之间形成的网络
//数学中集合类的实现
type UnionFind interface {
	IsConnected(p, q int) bool
	UnionElement(p, q int)
	GetSize() int
}
