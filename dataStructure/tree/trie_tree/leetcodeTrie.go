package trie_tree

// 211添加与搜索单词
type WordDictionary struct {
	TrieTree
}

//func Constructor() WordDictionary {
//	return WordDictionary{*NewTrieTree()}
//}

func (this *WordDictionary) AddWord(word string) {
	this.TrieTree.Add(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.match(this.TrieTree.Root, word, 0)
}

func (this *WordDictionary) match(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		return node.IsWord
	}
	b := word[index]
	if b != '.' {
		if _, ok := node.Next[b]; !ok {
			return false
		}
		return this.match(node.Next[b], word, index+1)
	} else {
		for i := range node.Next { //如果是 . 则要对map中的所有结点进行遍历，只要存在一个递归到底则返回true
			if this.match(node.Next[i], word, index+1) {
				return true
			}
		}
		return false
	}
}

// 677键值映射
type MapSum struct {
	Root *MapNode
	Size int
}

type MapNode struct {
	Value int
	Next  map[byte]*MapNode
}

func Constructor() MapSum {
	return MapSum{
		Root: NewMapNode(0),
		Size: 0,
	}
}

func NewMapNode(value int) *MapNode {
	return &MapNode{
		Value: value,
		Next:  make(map[byte]*MapNode),
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.Root
	for i := range key {
		b := key[i]
		if _, ok := cur.Next[b]; !ok {
			cur.Next[b] = NewMapNode(0)
		}
		cur = cur.Next[b]
	}
	cur.Value = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.Root
	for i := range prefix {
		b := prefix[i]
		if _, ok := cur.Next[b]; !ok {
			return 0
		}
		cur = cur.Next[b]
	}
	return this.sum(cur)
}

func (this *MapSum) sum(node *MapNode) int {
	if len(node.Next) == 0 {
		return node.Value
	}

	res := node.Value
	for _, v := range node.Next {
		res += this.sum(v)
	}
	return res
}
