package main

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) search(val int) bool {
	if t.root == nil {
		return false
	} else if t.root.data != val {
		t.root.search(val)
	} else {
		return true
	}
}

func (n *BinaryNode) search(val int) bool {
	if n == nil {
		return false
	} else if n.data == val {
		return true
	} else {
		if n.data > val {
			n.left.search(val)
		} else {
			n.right.search(val)
		}
	}
	return true
}
