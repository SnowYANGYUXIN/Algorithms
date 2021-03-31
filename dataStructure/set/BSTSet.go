package set

import (
	"fmt"
	"main.go/dataStructure/tree"
)

type BSTSet struct {
	tree.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{*tree.NewBinarySearchTree()}
}

func (b *BSTSet) GetSize() int {
	return b.BST.GetSize()
}

func (b *BSTSet) IsEmpty() bool {
	return b.BST.IsEmpty()
}

func (b *BSTSet) Add(e int) {
	b.BST.AddNew(e)
}

func (b *BSTSet) Contains(e int) bool {
	return b.BST.Contains(e)
}

func (b *BSTSet) Remove(e int) {
	b.BST.Remove(e)
}

func TestBSTSET() {
	b:=NewBSTSet()
	for i:=0;i<10;i++{
		b.Add(i)
	}
	fmt.Println(b.GetSize())
	b.Add(5)
	fmt.Println(b.GetSize())
	b.Add(12)
	fmt.Println(b.GetSize())
}
