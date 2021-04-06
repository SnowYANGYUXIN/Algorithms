package union_find

import "fmt"

//Quick Union 并集快速
//孩子指向父亲的树
//初始化一个size数组用于记录以i为根的集合的个数
//使之优化,让每次结点数少的根指向结点数多的根，保证树的深度很浅(防止极端情况导致树的深度退化成近似链表的情况)
type UnionFindTree struct {
	Parent []int
	Size   []int
}

//初始时每个结点都指向自己表示互不在同一集合，即初始了一个森林
func NewUnionFindTree(size int) *UnionFindTree {
	arr := make([]int, size)
	sz := make([]int, size)
	for i := range arr {
		arr[i] = i
		sz[i] = 1
	}
	return &UnionFindTree{
		Parent: arr,
		Size:   sz,
	}
}

func (u *UnionFindTree) GetSize() int {
	return len(u.Parent)
}

//查找元素p所对应的根节点结合编号 O(h)复杂度 h为树的深度
func (u *UnionFindTree) find(p int) int {
	if p < 0 || p >= len(u.Parent) {
		fmt.Println("index err")
		return -1
	}
	for p != u.Parent[p] {
		p = u.Parent[p]
	}
	return p
}

//压缩使得深度为2，即所有子节点都指向根结点
func (u *UnionFindTree) find2(p int) int {
	if p < 0 || p >= len(u.Parent) {
		fmt.Println("index err")
		return -1
	}
	if p != u.Parent[p] {
		u.Parent[p] = u.find2(u.Parent[p])
	}
	return p
}

func (u *UnionFindTree) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

//复杂度为O(h) h为树的高度
func (u *UnionFindTree) UnionElement(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}
	if u.Size[pRoot] >= u.Size[qRoot] {
		u.Parent[qRoot] = pRoot
		u.Size[pRoot] += u.Size[qRoot]
	} else {
		u.Parent[pRoot] = qRoot
		u.Size[qRoot] += u.Size[pRoot]
	}
}

//基于rank的优化
//rank[i]表示根节点为i树的高度，合并集合时深度低的树向深度高的树合并比以节点数合并更好
type UnionFindTreeRank struct {
	Parent []int
	Rank   []int
}

func NewUnionFindTreeRank(size int) *UnionFindTreeRank {
	arr := make([]int, size)
	rank := make([]int, size)
	for i := range arr {
		arr[i] = i
		rank[i] = 1
	}
	return &UnionFindTreeRank{
		Parent: arr,
		Rank:   rank,
	}
}

func (u *UnionFindTreeRank) UnionElement(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}
	if u.Rank[pRoot] > u.Rank[qRoot] {
		u.Parent[qRoot] = pRoot
	} else if u.Rank[pRoot] < u.Rank[qRoot] {
		u.Parent[pRoot] = qRoot
	} else {
		u.Parent[qRoot] = pRoot
		u.Rank[pRoot] += 1
	}
}

func (u *UnionFindTreeRank) find(p int) int {
	if p < 0 || p >= len(u.Parent) {
		fmt.Println("index err")
		return -1
	}
	for p != u.Parent[p] {
		u.Parent[p] = u.Parent[u.Parent[p]] //压缩路径，减少树的深度
		p = u.Parent[p]
	}
	return p
}

func TestUnionFind() {

}
