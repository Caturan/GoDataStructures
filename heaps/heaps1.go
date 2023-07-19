package heaps

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and remove it from heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0] // We are storing the first index in extracted
	l := len(h.array) - 1   // Get the last key(index)

	if len(h.array) == 0 { // If extract from an empty array that would be problem so checking system
		fmt.Println("Cannot extract because array length is 0")
		return -1 // We use return here because with that we can easily out the whole method
	}

	h.array[0] = h.array[l] // Put it in the root
	h.array = h.array[:l]   // Rearrenge the array size(shorter)

	h.maxHeapifyDown(0)
	return extracted // first step is return
}

//	maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // When left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // When left child is larger
			childToCompare = l
		} else { // When right child is larger
			childToCompare = r
		}
		//	Compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func Demo1() {
	m := MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
