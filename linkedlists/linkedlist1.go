package linkedlists

import "fmt"

type node struct { // First we need to describe a node
	data int   // node needs a data we are just use a single number
	next *node // It also needs to hold next which is adress of the next, so that we needs to be a pointer
}

type linkedList struct { // We need to describe the list
	head   *node // Hold a head, it is just has to be an adress of the node pointer. A good thing about the link list, we don't need to contain the whole list, we just need to contain the head.
	length int   // Indicate how long the linked list.
}

// Right now we can't do much because we are not able to put in the nodes inside the linked list yet. So let's make a method that can do that.
func (l *linkedList) prepend(n *node) { // The receiver is going to be a linked list. And it will gonna take a node to be addes at the front.
	second := l.head     // We will make temporary place called second and the put the current head in a second.  			// (l *linkedList) this called a method receiver in Go.
	l.head = n           // Set the new node as the head which is an N														// In go we define the method outside of the struct and this is (l * linkedList) how we indicate it for linked list. This called the receiver.
	l.head.next = second // Let the new head point to the old head which is second 										 	// And we can see here receiver is a pointer this means if we want to actually make modifications to this receiver, if we just put a values the we are just working on a copy of it.
	l.length++           // We need to increase the length 																	// So we want to actually work on the actual receiver like changes values and stuff. So we need to put it as a pointer.
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// And in most cases we'll see a pointer to receive and more often that a value receiver.
}

func Demo1() {
	myList := linkedList{}
	node1 := &node{data: 48} // We need to pass in a poniter (&) that is what this for it needs to be a pointer
	node2 := &node{data: 18}
	node3 := &node{data: 16}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList)
	fmt.Println(myList.length)
}
