package tree

import (
	"fmt"
	"math"
)

//平衡二叉树
//平衡二叉树 任意一个节点，左子树和右子树的高度差不能超过1
//标注节点高度(为左右子树中最高的+1) 计算平衡因子(节点左右子树高度差)

type AVLNode struct {
	Key    int
	Value  int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

type AVLTree struct {
	Root *AVLNode
	Size int
}

func NewAVLNode(key, value int) *AVLNode {
	return &AVLNode{
		Key:    key,
		Value:  value,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		Root: nil,
		Size: 0,
	}
}

func (m *AVLTree) GetSize() int {
	return m.Size
}

func getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

//计算node的平衡因子
func getBalanceFactor(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

//判断该二叉树是否是一棵二分搜索树
func (m *AVLTree) isBST() bool {
	keys := make([]int, m.Size)
	m.inOrder(m.Root, keys)
	for i := 1; i < len(keys); i++ {
		if keys[i-1] > keys[i] {
			return false
		}
	}
	return true
}

//如果是二分搜索树，中序遍历会将树顺序打印
func (m *AVLTree) inOrder(node *AVLNode, arr []int) {
	if node == nil {
		return
	}
	m.inOrder(node.Left, arr)
	arr = append(arr, node.Key)
	m.inOrder(node.Right, arr)
}

//判断该二叉树是否是一棵平衡二叉树
func (m *AVLTree) isBalance() bool {
	return m.isBalanced(m.Root)
}

func (m *AVLTree) isBalanced(node *AVLNode) bool {
	if node == nil {
		return true
	}
	balanceFactor := getBalanceFactor(node)
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}
	return m.isBalanced(node.Left) && m.isBalanced(node.Right)
}

func (m *AVLTree) IsEmpty() bool {
	return m.Size == 0
}

func (m *AVLTree) Add(key, value int) {
	m.Root = m.add(key, value, m.Root)
}

func (m *AVLTree) add(key, value int, node *AVLNode) *AVLNode {
	if node == nil {
		m.Size++
		return NewAVLNode(key, value)
	}
	if key < node.Key {
		node.Left = m.add(key, value, node.Left)
	} else if key > node.Key {
		node.Right = m.add(key, value, node.Right)
	} else {
		node.Value = value
	}

	//更新height
	node.Height = 1 + int(math.Max(float64(getHeight(node.Left)), float64(getHeight(node.Right))))

	//计算平衡因子
	balanceFactor := getBalanceFactor(node)

	//右旋转LL
	//插入的元素在不平衡的节点的左侧的左侧，将不平衡节点向右旋转(顺时针)
	if balanceFactor > 1 && getBalanceFactor(node.Left) >= 0 {
		return rightRotate(node)
	}

	//左旋转RR
	//插入的元素在不平衡的节点的右侧的右侧，将不平衡节点向左旋转(逆时针)
	if balanceFactor < -1 && getBalanceFactor(node.Right) <= 0 {
		return leftRotate(node)
	}

	//LR
	//插入的元素在不平衡的节点的左侧的右侧
	if balanceFactor > 1 && getBalanceFactor(node.Left) < 0 {
		node.Left = leftRotate(node.Left) //转换成LL的情况
		return rightRotate(node)
	}

	//RL
	//插入的元素在不平衡的节点的右侧的左侧
	if balanceFactor < -1 && getBalanceFactor(node.Right) > 0 {
		node.Right = rightRotate(node.Right) //转换成RR的情况
		return leftRotate(node)
	}

	return node
}

//右旋转 使之既保持二分搜索树又保持平衡二叉树
func rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	t3 := x.Right

	//向右旋转过程
	x.Right = y
	y.Left = t3

	//更新height 先更新y再更新x 因为此时x是y的父亲节点
	y.Height = int(math.Max(float64(getHeight(y.Left)), float64(getHeight(y.Right))))
	x.Height = int(math.Max(float64(getHeight(x.Left)), float64(getHeight(x.Right))))
	return x
}

//左旋转 使之既保持二分搜索树又保持平衡二叉树
func leftRotate(y *AVLNode) *AVLNode {
	x := y.Right
	t2 := x.Left

	//向左旋转
	x.Left = y
	y.Right = t2
	//更新height 先更新y再更新x 因为此时x是y的父亲节点
	y.Height = int(math.Max(float64(getHeight(y.Left)), float64(getHeight(y.Right))))
	x.Height = int(math.Max(float64(getHeight(x.Left)), float64(getHeight(x.Right))))
	return x
}

func (m *AVLTree) getNode(key int, node *AVLNode) *AVLNode {
	if node == nil {
		return nil
	}
	if key < node.Key {
		return m.getNode(key, node.Left)
	} else if key > node.Key {
		return m.getNode(key, node.Right)
	} else {
		return node
	}
}

func (m *AVLTree) Contains(key int) bool {
	return m.getNode(key, m.Root) != nil
}

func (m *AVLTree) Get(key int) int {
	node := m.getNode(key, m.Root)
	if node == nil {
		fmt.Println("no key")
		return -1
	} else {
		return node.Value
	}
}

func (m *AVLTree) Set(key, value int) {
	node := m.getNode(key, m.Root)
	if node == nil {
		fmt.Println("no key")
		return
	} else {
		node.Value = value
	}
}

func (m *AVLTree) miniMum(node *AVLNode) *AVLNode {
	if node.Left == nil {
		return node
	}
	return m.miniMum(node.Left)
}

func (m *AVLTree) removeMin(node *AVLNode) *AVLNode {
	if node.Left == nil {
		right := node.Right
		m.Size--
		return right
	}
	node.Left = m.removeMin(node.Left)
	return node
}

func (m *AVLTree) Remove(key int) int {
	node := m.getNode(key, m.Root)
	if node != nil {
		m.Root = m.remove(key, m.Root)
		return node.Value
	}
	fmt.Println("no key")
	return -1
}

func (m *AVLTree) remove(key int, node *AVLNode) *AVLNode {
	if node == nil {
		return nil
	}

	var retNode *AVLNode
	if key < node.Key {
		node.Left = m.remove(key, node.Left)
		retNode = node
	} else if key > node.Key {
		node.Right = m.remove(key, node.Right)
		retNode = node
	} else {
		if node.Left == nil {
			m.Size--
			retNode = node.Right
		} else if node.Right == nil {
			m.Size--
			retNode = node.Left
		} else {
			successor := m.miniMum(node.Right)
			successor.Right = m.remove(successor.Key, node.Right) //因为原先的删除可能会破坏自平衡
			successor.Left = node.Left
			retNode = successor
		}
	}

	if retNode == nil { //防止被删除的是叶子节点导致防卫其左右孩子出现空指针
		return nil
	}

	//更新height
	retNode.Height = 1 + int(math.Max(float64(getHeight(retNode.Left)), float64(getHeight(retNode.Right))))

	//计算平衡因子
	balanceFactor := getBalanceFactor(retNode)

	//右旋转LL
	//插入的元素在不平衡的节点的左侧的左侧，将不平衡节点向右旋转(顺时针)
	if balanceFactor > 1 && getBalanceFactor(retNode.Left) >= 0 {
		return rightRotate(retNode)
	}

	//左旋转RR
	//插入的元素在不平衡的节点的右侧的右侧，将不平衡节点向左旋转(逆时针)
	if balanceFactor < -1 && getBalanceFactor(retNode.Right) <= 0 {
		return leftRotate(retNode)
	}

	//LR
	//插入的元素在不平衡的节点的左侧的右侧
	if balanceFactor > 1 && getBalanceFactor(retNode.Left) < 0 {
		retNode.Left = leftRotate(retNode.Left) //转换成LL的情况
		return rightRotate(retNode)
	}

	//RL
	//插入的元素在不平衡的节点的右侧的左侧
	if balanceFactor < -1 && getBalanceFactor(retNode.Right) > 0 {
		retNode.Right = rightRotate(retNode.Right) //转换成RR的情况
		return leftRotate(retNode)
	}

	return retNode
}

func TestAVLTree() {
	m := NewAVLTree()
	m.Add(1, 2)
	m.Add(3, 4)
	m.Add(2, 5)
	fmt.Println(m.isBST())
}
