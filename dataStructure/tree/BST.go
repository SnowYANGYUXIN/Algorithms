package tree

import (
	"fmt"
	"strconv"
)

//二分搜索树
//二分搜索树的每个节点的值大于其左子树所有的值，小于其所有右子树的值
//BinarySearchTree
type BST struct {
	Node
	Root *Node
	Size int
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewBinarySearchTreeNode(e int) *Node {
	return &Node{
		Value: e,
		Left:  nil,
		Right: nil,
	}
}

func NewBinarySearchTree() *BST {
	return &BST{
		Root: nil,
		Size: 0,
	}
}

func (b *BST) GetSize() int {
	return b.Size
}

func (b *BST) IsEmpty() bool {
	return b.Size == 0
}

func (b *BST) Add(e int) {
	if b.Root == nil {
		b.Root = NewBinarySearchTreeNode(e)
		b.Size++
	} else {
		b.add(b.Root, e)
	}
}

func (b *BST) add(node *Node, e int) {
	if node.Value == e {
		return
	} else if e < node.Value && node.Left == nil {
		node.Left = NewBinarySearchTreeNode(e)
		b.Size++
		return
	} else if e > node.Value && node.Right == nil {
		node.Right = NewBinarySearchTreeNode(e)
		b.Size++
		return
	}
	if e < node.Value { //前面已经进行了一次判断，再进行判断会显得臃肿
		b.add(node.Left, e)
	} else {
		b.add(node.Right, e)
	}
}

func (b *BST) AddNew(e int) {
	b.Root = b.addNew(b.Root, e) //这样就不用判断根结点了
}

func (b *BST) addNew(node *Node, e int) *Node {
	if node == nil { //改变递归结束条件，优化代码
		b.Size++
		return NewBinarySearchTreeNode(e)
	}
	if e < node.Value {
		node.Left = b.addNew(node.Left, e)
	} else if e > node.Value {
		node.Right = b.addNew(node.Right, e)
	}
	return node //这个还是根节点
}

func (b *BST) Contains(e int) bool {
	return b.contains(b.Root, e)
}

func (b *BST) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}

	if e == node.Value {
		return true
	} else if e < node.Value {
		return b.contains(node.Left, e)
	} else {
		return b.contains(node.Right, e)
	}
}

//前序遍历
func (b *BST) PreOrder() {
	b.preOrder(b.Root)
}

func (b *BST) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	b.preOrder(node.Left)
	b.preOrder(node.Right)
}

//中序遍历
//其遍历结果是二分搜索树的排序结果
func (b *BST) InOrder() {
	b.inOrder(b.Root)
}

//后序遍历
func (b *BST) PostOrder() {
	b.postOder(b.Root)
}

func (b *BST) postOder(node *Node) {
	if node == nil {
		return
	}
	b.postOder(b.Left)
	b.postOder(b.Right)
	fmt.Println(node.Value)
}

func (b *BST) inOrder(node *Node) {
	if node == nil {
		return
	}
	b.inOrder(b.Left)
	fmt.Println(node.Value)
	b.inOrder(b.Right)
}

//层序遍历 广度优先遍历
//需要借助队列
//意义
//更快找到问题的解
//常用于算法设计中-最短路径
//func (b *BST) levelOrder() {
//	q := queue.NewArrayQueue(10)
//	q.EnQueue(b.Root)
//	for !q.IsEmpty() {
//		cur:=q.DeQueue()
//		fmt.Println(cur.Value)
//		if cur.Left!=nil{
//			q.EnQueue(cur.Left)
//		}
//		if cur.Rihght!=nil{
//			q.EnQueue(cur.Right)
//		}
//	}
//}

//寻找最小元素
func (b *BST) Minimum() int {
	if b.Size == 0 {
		fmt.Println("BST is empty")
		return -1
	}
	return b.minimum(b.Root).Value
}

func (b *BST) minimum(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return b.minimum(node.Left)
}

func (b *BST) RemoveMin() int {
	ret := b.Minimum()
	b.Root = b.removeMin(b.Root)
	return ret
}

func (b *BST) removeMin(node *Node) *Node {
	if node.Left == nil {
		right := node.Right
		b.Size--
		return right
	}
	node.Left = b.removeMin(node.Left)
	return node
}

func (b *BST) RemoveMax() int {
	ret := b.Maximum()
	b.Root = b.removeMax(b.Root)
	return ret
}

func (b *BST) removeMax(node *Node) *Node {
	if node.Right == nil {
		left := node.Left
		b.Size--
		return left
	}
	node.Right = b.removeMax(node.Right)
	return node
}

//寻找最大元素
func (b *BST) Maximum() int {
	if b.Size == 0 {
		fmt.Println("BST is empty")
		return -1
	}
	return b.maximum(b.Root).Value
}

func (b *BST) maximum(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return b.maximum(node.Right)
}

//删除任意元素
//如果该元素左右都有孩子
//1、找到该结点的后继元素(结点右孩子的最小值)或者前驱(结点左孩子的最大值)
//2、删除后继元素
//3、将该后继结点的左右孩子分别连接原结点的左右孩子
//4、在将该后继结点连接在原结点的上一个结点
func (b *BST) Remove(e int) {
	b.Root = b.remove(b.Root, e)
}

func (b *BST) remove(node *Node, e int) *Node {
	if node == nil {
		return nil
	}
	if e < node.Value {
		node.Left = b.remove(node.Left, e)
		return node
	} else if e > node.Value {
		node.Right = b.remove(node.Right, e)
		return node
	} else {
		if node.Left == nil {
			b.Size--
			return node.Right
		}
		if node.Right == nil {
			b.Size--
			return node.Left
		}
		//该元素左右都有孩子
		//1、找到该结点的后继元素(结点右孩子的最小值)
		//2、删除后继元素
		//3、将该后继结点的左右孩子分别连接原结点的左右孩子
		//4、在将该后继结点连接在原结点的上一个结点
		successor := b.minimum(node.Right)
		successor.Right = b.removeMin(node.Right) //里面已执行size--
		successor.Left = node.Left
		return successor
	}
}

func (b *BST) ToString() string {
	var res string
	b.generateBSTString(b.Root, 0, &res)
	return res
}

func (b *BST) generateBSTString(node *Node, depth int, res *string) {
	if node == nil {
		*res += generateDepthString(depth) + "null\n"
		return
	}
	*res += generateDepthString(depth) + strconv.Itoa(node.Value) + "\n"
	b.generateBSTString(node.Left, depth+1, res)
	b.generateBSTString(node.Right, depth+1, res)
}

func generateDepthString(depth int) string {
	var res string
	for i := 0; i < depth; i++ {
		res += "--"
	}
	return res
}

func TestBinarySearchTree() {
	b := NewBinarySearchTree()
	nums := []int{5, 3, 6, 8, 4, 2, 6, 1}
	for _, v := range nums {
		b.AddNew(v)
	}
	//b.PreOrder()
	fmt.Println(b.ToString() + "\n\n")
	b.RemoveMin()
	fmt.Println(b.ToString() + "\n\n")
	b.RemoveMax()
	fmt.Println(b.ToString() + "\n\n")
}
