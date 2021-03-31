package Map

import "fmt"

//先创建具有key和value的链表
type Node struct {
	Key       int
	Value     int
	Next      *Node
}

type LinkedListMap struct {
	DummyHead *Node
	Size      int
}

func NewNodeKV(key, value int, next *Node) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Next:  next,
	}
}

func NewNode() *Node {
	return &Node{
		Key:   0,
		Value: 0,
		Next:  nil,
	}
}

func (n *Node) ToString() string {
	return fmt.Sprintf("%d : %d", n.Key, n.Value)
}

func NewLinkedListMap() *LinkedListMap {
	return &LinkedListMap{DummyHead: NewNode(), Size: 0}
}

func (n *LinkedListMap) GetSize() int {
	return n.Size
}

func (n *LinkedListMap) IsEmpty() bool {
	return n.Size == 0
}

func (n *LinkedListMap) getNode(key int) *Node {
	cur := n.DummyHead.Next
	for cur != nil {
		if cur.Key == key {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (n *LinkedListMap) Contains(key int) bool {
	return n.getNode(key) != nil
}

func (n *LinkedListMap) Get(key int) int {
	node := n.getNode(key)
	if node == nil {
		fmt.Println("no key")
		return 0
	} else {
		return node.Value
	}
}

func (n *LinkedListMap) Add(key, value int) {
	node := n.getNode(key)
	if node == nil {
		n.DummyHead.Next = NewNodeKV(key, value, n.DummyHead.Next)
		n.Size++
	} else {
		node.Value = value
	}
}

func (n *LinkedListMap) Set(key, value int) {
	node := n.getNode(key)
	if node == nil {
		fmt.Println("no key")
		return
	} else {
		node.Value = value
	}
}

func (n *LinkedListMap) Remove(key int) int {
	pre := n.DummyHead
	for pre.Next != nil {
		if pre.Next.Key == key {
			break
		}
		pre = pre.Next
	}
	if pre.Next != nil {
		delNode := pre.Next
		pre.Next = delNode.Next
		return delNode.Value
	}
	fmt.Println("no key")
	return -1
}

func TestLinkedListMap() {
	m := NewLinkedListMap()
	m.Add(1, 2)
	m.Add(3, 4)
	fmt.Println(m.Size)
	fmt.Println(m.Get(3))
	m.Remove(1)
	fmt.Println(m.GetSize())
	m.Set(3, 5)
	fmt.Println(m.Get(3))
}
