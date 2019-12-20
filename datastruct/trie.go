package datastruct

// TrieNode ...
type TrieNode struct {
	children map[rune]*TrieNode
	value    interface{}
}

// NewNode => new node
func NewNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

// Insert => insert
func (trie *TrieNode) Insert(key string, value int) bool {
	node := trie

	for _, ch := range key {
		child, _ := node.children[ch]

		if child == nil {
			child = NewNode()
			node.children[ch] = child
		}
		node = child
	}

	isNewVal := node.value == nil

	node.value = value

	return isNewVal
}

// Contains => contains
func (trie *TrieNode) Contains(key string) bool {
	node := trie

	for _, ch := range key {
		node = node.children[ch]
		if node == nil {
			return false
		}
	}

	return node.value != nil
}

type nodeStr struct {
	node *TrieNode
	part rune
}

// Delete => delete
func (trie *TrieNode) Delete(key string) bool {
	var path []nodeStr
	node := trie
	for _, ch := range key {
		path = append(path, nodeStr{part: ch, node: node})
		node = node.children[ch]
		if node == nil {
			return false
		}
	}

	node.value = nil

	if len(node.children) == 0 {
		for i := len(path) - 1; i >= 0; i-- {
			parent := path[i].node
			part := path[i].part
			delete(parent.children, part)
			if parent.value != nil || !(len(parent.children) == 0) {
				break
			}
		}
	}

	return true
}
