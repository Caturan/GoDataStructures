package binarysearchtree

import "fmt"

type Node2 struct {
	Value int
	Left  *Node2
	Right *Node2
}

type BST struct {
	Root *Node2
}

// Insert
func (bst *BST) Insert(value int) {
	newNode := &Node2{Value: value}

	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.insertNode(bst.Root, newNode)
	}
}

func (bst *BST) insertNode(root, newNode *Node2) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			bst.insertNode(root.Left, newNode)
		}
	} else if newNode.Value > root.Value {
		if root.Right == nil {
			root.Right = newNode
		} else {
			bst.insertNode(root.Right, newNode)
		}
	} else {
		// If the value is already in the tree, we can choose to ignore or handle it as per requirements.
		// For this example, we'll ignore duplicates.
	}
}

// Search
func (bst *BST) Search(value int) bool {
	return bst.searchNode(bst.Root, value)
}

func (bst *BST) searchNode(root *Node2, value int) bool {
	if root == nil {
		return false
	}

	if value < root.Value {
		return bst.searchNode(root.Left, value)
	} else if value > root.Value {
		return bst.searchNode(root.Right, value)
	}

	return true
}

// In order Traversal
func (bst *BST) InOrderTraversal() {
	bst.inOrderTraversal(bst.Root)
}

func (bst *BST) inOrderTraversal(root *Node2) {
	if root != nil {
		bst.inOrderTraversal(root.Left)
		fmt.Printf("%d ", root.Value)
		bst.inOrderTraversal(root.Right)
	}
}

func Demo2() {
	bst := BST{}
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	fmt.Println("In-order Traversal:")
	bst.InOrderTraversal()
	fmt.Println()

	searchValue := 40
	if bst.Search(searchValue) {
		fmt.Printf("%d is found in the binary search tree.\n", searchValue)
	} else {
		fmt.Printf("%d is not found in the binary search tree.\n", searchValue)
	}
}
