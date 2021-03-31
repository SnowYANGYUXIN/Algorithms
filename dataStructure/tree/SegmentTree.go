package tree

import (
	"fmt"
	"strconv"
)

//线段树
//线段树存放的是固定区间，不能增加和删除元素
//虽然不是完全二叉树(一层一层从左到右)，但是是一棵平衡二叉树(最大深度和最小深度相差小于等于1)
//如果区间有n个元素，数组表示需要有4n个结点，因为线段树不考虑添加元素，即区间固定，使用4n静态空间即可
//为什么4n，因为n个元素最终都是叶子结点，如果刚好铺满最后一层，则要2n空间，但是，如果在下面一层还有若干叶子结点，则需要4n空间
type Merger func(a, b int) int

type SegmentTree struct {
	Array  []int
	Tree   []*int
	Merger Merger
}

func NewSegmentTree(arr []int, merger Merger) *SegmentTree {
	tree := SegmentTree{
		Array:  arr,
		Tree:   make([]*int, len(arr)*4),
		Merger: merger,
	}
	tree.buildSegmentTree(0, 0, len(arr)-1)
	return &tree
}

func (t *SegmentTree) buildSegmentTree(index, l, r int) {
	if l == r {
		t.Tree[index] = &t.Array[l]
		//t.Tree[index]=new(int)
		//*t.Tree[index] = t.Array[l]
		return
	}
	leftTreeIndex := leftChild(index)
	rightTreeIndex := rightChild(index)

	mid := (l + r) / 2
	t.buildSegmentTree(leftTreeIndex, l, mid)
	t.buildSegmentTree(rightTreeIndex, mid+1, r)
	t.Tree[index] = new(int)
	*t.Tree[index] = t.Merger(*t.Tree[leftTreeIndex], *t.Tree[rightTreeIndex]) //此时tree存的值看具体业务逻辑决定
}

func (t *SegmentTree) Get(index int) int {
	if index < 0 || index >= len(t.Array) {
		fmt.Println("index err")
		return -1
	}
	return t.Array[index]
}

func (t *SegmentTree) GetSize() int {
	return len(t.Array)
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

//返回区间[queryL,queryR]的值
func (t *SegmentTree) Query(queryL, queryR int) int {
	if queryL < 0 || queryL > len(t.Array) || queryR < 0 || queryR > len(t.Array) || queryL > queryR {
		fmt.Println("index err")
		return -1
	}
	return t.query(0, 0, len(t.Array)-1, queryL, queryR)
}

//在以index为根的线段树中[l,r]的范围里搜索[queryL,queryR]
func (t *SegmentTree) query(index, l, r, queryL, queryR int) int {
	if l == queryL && r == queryR {
		return *t.Tree[index]
	}
	mid := (l + r) / 2
	if queryR <= mid {
		return t.query(leftChild(index), l, mid, queryL, queryR)
	} else if queryL >= mid+1 {
		return t.query(rightChild(index), mid+1, r, queryL, queryR)
	} else {
		return t.Merger(t.query(leftChild(index), l, mid, queryL, mid), t.query(rightChild(index), mid+1, r, mid+1, queryR))
	}
}

func (t *SegmentTree) Set(index, e int) {
	if index < 0 || index >= len(t.Array) {
		fmt.Println("index err")
		return
	}
	t.Array[index] = e
	t.set(0, 0, len(t.Array)-1, index, e)
}

func (t *SegmentTree) set(treeIndex, l, r, index, e int) {
	if l == r {
		*t.Tree[treeIndex] = e
		return
	}
	mid := (l + r) / 2
	left := leftChild(treeIndex)
	right := rightChild(treeIndex)
	if index <= mid {
		t.set(left, l, mid, index, e)
	} else {
		t.set(right, mid+1, r, index, e)
	}
	*t.Tree[treeIndex] = t.Merger(*t.Tree[left], *t.Tree[right])
}

func (t *SegmentTree) ToString() string {
	str := ""
	str += "["
	for i := 0; i < len(t.Tree); i++ {
		if t.Tree[i] != nil {
			str += strconv.Itoa(*t.Tree[i])
		} else {
			str += "null"
		}

		if i != len(t.Tree)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

func TestSegmentTree() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	tree := NewSegmentTree(nums, func(a, b int) int {
		return a + b
	})
	fmt.Println(tree.ToString())
	fmt.Println(tree.Query(0, 2))
}
