package datastruct

// BTNode => binary tree node
type BTNode struct {
	data   int
	parent *BTNode
	left   *BTNode
	right  *BTNode
}

// BinaryTree => binary tree
type BinaryTree struct {
	root *BTNode
}

// InsertItem => inset item in BinaryTree
func (bt *BinaryTree) InsertItem(i int) {
	if bt.root == nil {
		bt.root = &BTNode{data: i}
		return
	}
	currentNode := bt.root

	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &BTNode{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &BTNode{data: i, parent: currentNode}
			}
			currentNode = currentNode.left
		}
	}
}

// SearchItem => search item in binary tree
func (bt *BinaryTree) SearchItem(target int) (*BTNode, bool) {
	if bt.root == nil {
		return nil, false
	}

	currentNode := bt.root

	for currentNode != nil {
		if target == currentNode.data {
			return currentNode, true
		} else if target > currentNode.data {
			currentNode = currentNode.right
		} else if target < currentNode.data {
			currentNode = currentNode.left
		}
	}

	return nil, false
}

// InorderTraversal => inorder traversal
func (bt *BinaryTree) InorderTraversal(subtree *BTNode, callback func(int)) {
	if subtree.left != nil {
		bt.InorderTraversal(subtree.left, callback)
	}

	callback(subtree.data)

	if subtree.right != nil {
		bt.InorderTraversal(subtree.right, callback)
	}
}

// PreorderTraversal => preorder traversal
func (bt *BinaryTree) PreorderTraversal(subtree *BTNode, callback func(int)) {
	callback(subtree.data)

	if subtree.left != nil {
		bt.PreorderTraversal(subtree.left, callback)
	}

	if subtree.right != nil {
		bt.PreorderTraversal(subtree.right, callback)
	}
}

// PostorderTraversal => postorder traversal
func (bt *BinaryTree) PostorderTraversal(subtree *BTNode, callback func(int)) {
	if subtree.left != nil {
		bt.PostorderTraversal(subtree.left, callback)
	}

	if subtree.right != nil {
		bt.PostorderTraversal(subtree.right, callback)
	}

	callback(subtree.data)
}
