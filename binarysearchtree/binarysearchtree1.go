package binarysearchtree

import "fmt"

var count int

// Node represent the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// And return true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		// Move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

func Demo1() {
	tree := &Node{Key: 100}
	tree.Insert(150)
	fmt.Println(tree)

	tree.Insert(45)
	tree.Insert(98)
	tree.Insert(362)
	tree.Insert(17)
	tree.Insert(36)
	tree.Insert(78)

	fmt.Println(tree.Search(17))
	fmt.Println(count)

	fmt.Println(tree.Search(101))

	fmt.Println(tree)

}
