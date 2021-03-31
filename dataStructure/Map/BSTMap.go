package Map

import "fmt"

//先创建具有key、value的BST
type NodeMap struct {
	Key   int
	Value int
	Left  *NodeMap
	Right *NodeMap
}

type BSTMap struct {
	Root *NodeMap
	Size int
}

func NewNodeMap(key, value int) *NodeMap {
	return &NodeMap{
		Key:   key,
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func NewBSTMap() *BSTMap {
	return &BSTMap{
		Root: nil,
		Size: 0,
	}
}

func (m *BSTMap) GetSize() int {
	return m.Size
}

func (m *BSTMap) IsEmpty() bool {
	return m.Size == 0
}

func (m *BSTMap) Add(key, value int) {
	m.Root=m.add(key, value, m.Root)
}

func (m *BSTMap) add(key, value int, node *NodeMap) *NodeMap {
	if node == nil {
		m.Size++
		return NewNodeMap(key, value)
	}
	if key < node.Key {
		node.Left = m.add(key, value, node.Left)
	} else if key > node.Key {
		node.Right = m.add(key, value, node.Right)
	} else {
		node.Value = value
	}
	return node
}

func (m *BSTMap) getNode(key int, node *NodeMap) *NodeMap {
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

func (m *BSTMap) Contains(key int) bool {
	return m.getNode(key, m.Root) != nil
}

func (m *BSTMap) Get(key int) int {
	node := m.getNode(key, m.Root)
	if node == nil {
		fmt.Println("no key")
		return -1
	} else {
		return node.Value
	}
}

func (m *BSTMap) Set(key, value int) {
	node := m.getNode(key, m.Root)
	if node == nil {
		fmt.Println("no key")
		return
	} else {
		node.Value = value
	}
}

func (m *BSTMap) miniMum(node *NodeMap) *NodeMap {
	if node.Left == nil {
		return node
	}
	return m.miniMum(node.Left)
}

func (m *BSTMap) removeMin(node *NodeMap) *NodeMap {
	if node.Left == nil {
		right := node.Right
		m.Size--
		return right
	}
	node.Left = m.removeMin(node.Left)
	return node
}

func (m *BSTMap) Remove(key int) int {
	node := m.getNode(key, m.Root)
	if node != nil {
		m.Root = m.remove(key, m.Root)
		return node.Value
	}
	fmt.Println("no key")
	return -1
}

func (m *BSTMap) remove(key int, node *NodeMap) *NodeMap {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = m.remove(key, node.Left)
		return node
	} else if key > node.Key {
		node.Right = m.remove(key, node.Right)
		return node
	} else {
		if node.Left == nil {
			m.Size--
			return node.Right
		}
		if node.Right == nil {
			m.Size--
			return node.Left
		}
		successor := m.miniMum(node.Right)
		successor.Right = m.removeMin(node.Right)
		successor.Left = node.Left
		return successor
	}
}

func TestBSTMap() {
	m := NewBSTMap()
	m.Add(1, 2)
	m.Add(3, 4)
	fmt.Println(m.Size)
	fmt.Println(m.Get(1))
	m.Remove(1)
	fmt.Println(m.GetSize())
	m.Set(3, 5)
	fmt.Println(m.Get(3))
}

