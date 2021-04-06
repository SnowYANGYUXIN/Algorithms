package trie_tree

import "fmt"

//字典树、前缀树
//查询，每个条目的时间复杂度和字典中一共有多少条目无关
//时间复杂度为O(w)w为查询单词长度
//是一颗多叉树，每一个结点有若干个指向下一个结点的指针
type TrieTree struct {
	Root *TrieNode
	Size int
}

//需要isWord来判断当前是否是一个单词，比如区分panda中的单词pan
type TrieNode struct {
	IsWord bool
	Next   map[byte]*TrieNode
}

func NewTrieNode(isWord bool) *TrieNode {
	return &TrieNode{
		IsWord: isWord,
		Next:   make(map[byte]*TrieNode),
	}
}

func NewTrieNodeNull() *TrieNode {
	return &TrieNode{
		IsWord: false,
		Next:   make(map[byte]*TrieNode),
	}
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		Root: NewTrieNodeNull(),
		Size: 0,
	}
}

func (t *TrieTree) GetSize() int {
	return t.Size
}

func (t *TrieTree) Add(word string) {
	cur := t.Root
	for i := range word {
		b := word[i]
		if _, ok := cur.Next[b]; !ok {
			cur.Next[b] = NewTrieNodeNull()
		}
		cur = cur.Next[b]
	}
	if !cur.IsWord { //防止重复添加
		cur.IsWord = true
		t.Size++
	}
}

func (t *TrieTree) Contains(word string) bool {
	cur := t.Root
	for i := range word {
		b := word[i]
		if _, ok := cur.Next[b]; !ok {
			return false
		}
		cur = cur.Next[b]
	}
	return cur.IsWord //返回是否存在此单词而不是是否包含此字符串
}

//是否含有前缀
func (t *TrieTree) IsPrefix(prefix string) bool {
	cur := t.Root
	for i := range prefix {
		b := prefix[i]
		if _, ok := cur.Next[b]; !ok {
			return false
		}
		cur = cur.Next[b]
	}
	return true
}

func TestTrieTree() {
	t := NewTrieTree()
	t.Add("hello")
	t.Add("world")
	t.Add("panda")
	fmt.Println(t.Contains("pan"))
	t.Add("pan")
	fmt.Println(t.Contains("pan"))
	fmt.Println(t.Contains("wor"))
	fmt.Println(t.Contains("world"))
}
