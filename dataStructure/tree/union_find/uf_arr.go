package union_find

import "fmt"

//Quick Find 查找快速
type UnionFindArr struct {
	ID []int
}

//初始化给每个id赋上不同的值表示最开始大家彼此都在不同的集合
func NewUnionFindArr(size int) *UnionFindArr {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return &UnionFindArr{ID: arr}
}

func (u *UnionFindArr) GetSize() int {
	return len(u.ID)
}

//查找元素p所对应的编号
func (u *UnionFindArr) find(p int) int {
	if p < 0 || p >= len(u.ID) {
		fmt.Println("index err")
		return -1
	}
	return u.ID[p]
}

//查看元素p和元素q是否属于同一集合 复杂度O(1)
func (u *UnionFindArr) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

//合并元素p和元素q的集合 复杂度O(n)
func (u *UnionFindArr) UnionElement(p, q int) {
	pID := u.find(p)
	qID := u.find(q)
	if pID == qID {
		return
	}
	for i := range u.ID {
		if u.ID[i] == pID {
			u.ID[i] = qID
		}
	}
}
