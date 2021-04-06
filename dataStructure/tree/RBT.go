package tree

import (
	"fmt"
	"strconv"
)

//2-3树
//满足二分搜索树的基本性质
//节点可以存放一个元素或者两个元素，每个节点有两个孩子或者三个孩子
//(两个元素三个孩子，比第一个元素小在左边，两个元素间的在中间，比第二个元素大的在右边)
//2-3树是一棵绝对平衡的树(对于任意节点左右子树都相等)

//红黑树
//每个节点都是红色的或者是黑色的
//根节点是黑色的(如果看到根节点是红色的，其实根节点是它父亲节点的黑色)
//每一个叶子节点(最后的空节点)是黑色的(空节点都是黑色的)
//如果一个节点是红色的，那么它的孩子结点都是黑色的(即使连的是3-结点，其实连的是父亲黑节点)
//任意一个节点到叶子节点，经历的黑色节点是一样的(对应2-3树一个节点到任意一个根节点走过的节点个数一样，因为绝对平衡树深度相同)

//红黑树是一个保持黑平衡的二叉树(性质第五条)，严格意义上讲，不是平衡二叉树。最大高度为2logn，logn的黑加logn的红
const (
	red   = true
	black = false
)

type RBT struct {
	Root *RBTNode
	Size int
}

type RBTNode struct {
	Value int
	Left  *RBTNode
	Right *RBTNode
	Color bool
}

//初始化是红色代表该节点总要和其他节点进行融合
func NewRBTNode(e int) *RBTNode {
	return &RBTNode{
		Value: e,
		Left:  nil,
		Right: nil,
		Color: red,
	}
}

func NewRedBlackTree() *RBT {
	return &RBT{
		Root: nil,
		Size: 0,
	}
}

func isRed(node *RBTNode) bool {
	if node == nil {
		return black
	}
	return node.Color
}

//插入2-节点的左侧，就直接将新节点插入到左孩子，并将节点颜色置为黑色

//左旋转，插入2-结点，如果添加元素在右侧，则要进行左旋转，将添加元素变为父亲节点，原来节点变为左孩子红色
func leftRotateRBT(node *RBTNode) *RBTNode {
	x := node.Right
	node.Right = x.Left
	x.Left = node

	//维护颜色
	x.Color = node.Color
	node.Color = red
	return x
}

//向3-节点插入元素并插入在最右侧，会从中间分裂成两个2-节点，即中间节点颜色会翻转成红色，左右孩子节点颜色会翻转成黑色
func flipColors(node *RBTNode) {
	node.Color = red
	node.Left.Color = black
	node.Right.Color = black
}

//向3-节点插入元素并插入在最中间，即插入到红色孩子的右节点，此时要先将红色孩子进行左旋转，再对父亲节点进行右旋转，最后进行颜色翻转

//右旋转，向3-节点插入元素并插入在最侧，此时要进行右旋转，将父亲节点变成红色孩子的右节点，最后进行颜色翻转
func rightRotateRBT(node *RBTNode) *RBTNode {
	x := node.Left
	node.Left = x.Right
	x.Right = node

	//维护颜色
	x.Color = node.Color
	node.Color = red
	return x
}

func (b *RBT) AddNew(e int) {
	b.Root = b.addNew(b.Root, e) //这样就不用判断根结点了
	b.Root.Color = black         //保持根节点为黑色
}

func (b *RBT) addNew(node *RBTNode, e int) *RBTNode {
	if node == nil { //改变递归结束条件，优化代码
		b.Size++
		return NewRBTNode(e) //默认插入红色节点
	}
	if e < node.Value {
		node.Left = b.addNew(node.Left, e)
	} else if e > node.Value {
		node.Right = b.addNew(node.Right, e)
	}

	//左旋转
	if isRed(node.Right) && !isRed(node.Left) {
		node = leftRotateRBT(node)
	}

	//右旋转
	if isRed(node.Left) && isRed(node.Left.Left) {
		node = rightRotateRBT(node)
	}

	//颜色翻转
	if isRed(node.Left) && isRed(node.Right) {
		flipColors(node)
	}

	return node //这个还是根节点
}

func (b *RBT) Contains(e int) bool {
	return b.contains(b.Root, e)
}

func (b *RBT) contains(RBTNode *RBTNode, e int) bool {
	if RBTNode == nil {
		return false
	}

	if e == RBTNode.Value {
		return true
	} else if e < RBTNode.Value {
		return b.contains(RBTNode.Left, e)
	} else {
		return b.contains(RBTNode.Right, e)
	}
}

//前序遍历
func (b *RBT) PreOrder() {
	b.preOrder(b.Root)
}

func (b *RBT) preOrder(RBTNode *RBTNode) {
	if RBTNode == nil {
		return
	}
	fmt.Println(RBTNode.Value)
	b.preOrder(RBTNode.Left)
	b.preOrder(RBTNode.Right)
}

//中序遍历
//其遍历结果是二分搜索树的排序结果
func (b *RBT) InOrder() {
	b.inOrder(b.Root)
}

//后序遍历
func (b *RBT) PostOrder() {
	b.postOder(b.Root)
}

func (b *RBT) postOder(RBTNode *RBTNode) {
	if RBTNode == nil {
		return
	}
	b.postOder(RBTNode.Left)
	b.postOder(RBTNode.Right)
	fmt.Println(RBTNode.Value)
}

func (b *RBT) inOrder(RBTNode *RBTNode) {
	if RBTNode == nil {
		return
	}
	b.inOrder(RBTNode.Left)
	fmt.Println(RBTNode.Value)
	b.inOrder(RBTNode.Right)
}

//寻找最小元素
func (b *RBT) Minimum() int {
	if b.Size == 0 {
		fmt.Println("RBT is empty")
		return -1
	}
	return b.minimum(b.Root).Value
}

func (b *RBT) minimum(RBTNode *RBTNode) *RBTNode {
	if RBTNode.Left == nil {
		return RBTNode
	}
	return b.minimum(RBTNode.Left)
}

func (b *RBT) RemoveMin() int {
	ret := b.Minimum()
	b.Root = b.removeMin(b.Root)
	return ret
}

func (b *RBT) removeMin(RBTNode *RBTNode) *RBTNode {
	if RBTNode.Left == nil {
		right := RBTNode.Right
		b.Size--
		return right
	}
	RBTNode.Left = b.removeMin(RBTNode.Left)
	return RBTNode
}

func (b *RBT) RemoveMax() int {
	ret := b.Maximum()
	b.Root = b.removeMax(b.Root)
	return ret
}

func (b *RBT) removeMax(RBTNode *RBTNode) *RBTNode {
	if RBTNode.Right == nil {
		left := RBTNode.Left
		b.Size--
		return left
	}
	RBTNode.Right = b.removeMax(RBTNode.Right)
	return RBTNode
}

//寻找最大元素
func (b *RBT) Maximum() int {
	if b.Size == 0 {
		fmt.Println("RBT is empty")
		return -1
	}
	return b.maximum(b.Root).Value
}

func (b *RBT) maximum(RBTNode *RBTNode) *RBTNode {
	if RBTNode.Right == nil {
		return RBTNode
	}
	return b.maximum(RBTNode.Right)
}

//删除任意元素
//如果该元素左右都有孩子
//1、找到该结点的后继元素(结点右孩子的最小值)或者前驱(结点左孩子的最大值)
//2、删除后继元素
//3、将该后继结点的左右孩子分别连接原结点的左右孩子
//4、在将该后继结点连接在原结点的上一个结点
func (b *RBT) Remove(e int) {
	b.Root = b.remove(b.Root, e)
}

func (b *RBT) remove(RBTNode *RBTNode, e int) *RBTNode {
	if RBTNode == nil {
		return nil
	}
	if e < RBTNode.Value {
		RBTNode.Left = b.remove(RBTNode.Left, e)
		return RBTNode
	} else if e > RBTNode.Value {
		RBTNode.Right = b.remove(RBTNode.Right, e)
		return RBTNode
	} else {
		if RBTNode.Left == nil {
			b.Size--
			return RBTNode.Right
		}
		if RBTNode.Right == nil {
			b.Size--
			return RBTNode.Left
		}
		//该元素左右都有孩子
		//1、找到该结点的后继元素(结点右孩子的最小值)
		//2、删除后继元素
		//3、将该后继结点的左右孩子分别连接原结点的左右孩子
		//4、在将该后继结点连接在原结点的上一个结点
		successor := b.minimum(RBTNode.Right)
		successor.Right = b.removeMin(RBTNode.Right) //里面已执行size--
		successor.Left = RBTNode.Left
		return successor
	}
}

func (b *RBT) ToString() string {
	var res string
	b.generateRBTString(b.Root, 0, &res)
	return res
}

func (b *RBT) generateRBTString(RBTNode *RBTNode, depth int, res *string) {
	if RBTNode == nil {
		*res += generateDepthString(depth) + "null\n"
		return
	}
	*res += generateDepthString(depth) + strconv.Itoa(RBTNode.Value) + "\n"
	b.generateRBTString(RBTNode.Left, depth+1, res)
	b.generateRBTString(RBTNode.Right, depth+1, res)
}

func (b *RBT) GetSize() int {
	return b.Size
}

func (b *RBT) IsEmpty() bool {
	return b.Size == 0
}

func TestRedBlackTree() {

}
