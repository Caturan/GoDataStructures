package heaps

/*
	In normal queues : First in First Out
	Priority Queues : Highest priority served first

	A heap that can be expressed as a complete tree
	Complete Tree : Means that all the levels in the tree are full except for the lowest levels some nodes can be empty.

	Max Heap : In the case for max heaps the largest key will be stored in the root and for all the nodes in the tree every parents
	 will have a higher key than its children.

	Min Heap : The opposite of the max heap. Root is the smalles key so in a min heap each parent node will have the smaller key than its children.

	Heaps can be visualized as a tree but actually behind the scenes, the keys are stored as an array.
	Each node of tree correspons to an index of that array.

	So strictly speaking heapse are actually just an array that operates like a tree.
	That is possible because we can just easily calculate the indices of the children base on the parent index.
	For example if we want to get the left child node index of number 39, (39*2)+1 will give 79 that is the left child's index number.

	MaxHeap Insert : Where we want the insert, we are going to add new key the bottom's last index. In an array bottom right key last index.
	We need to rearrenge the nodes so that we can maintain the heap property which is to keep the parent key larger than its children.
	 We will compare the new node to its parent node and swap it if the new node is higher.
	  We follow up the tree and keep on repeating this process until it gets to its place.

	We call this process of rearranging the indices as HEAPIFY.
	The heapify process is also needed to afet taking out an item of tree.

	HeapExtractMax : To extract the highest key in a heap we can just take out root because we know root is the largest key,
	 after that we need to rearrenge tree. First step right after losing the root we will fill the empty root with the last node. And after heapify time, until it gets to its place.

	Time complexity of extracting and inserting a node in the heap. Would depend on heapifying process.
	 Simply adding an element at the end of the array, or taking out a key from the first index of the array.
	 The number of operations needed to heapify up or down would depend on how high the tree is so that's going to be the one that affects the time complexity.
	 So we can express insert and extract as O(h).
	 If we want to express the time complexity with n which would be size of the array.
	 We can just replace h with O(log n) because the height and the numnber of indices have a logarithmic relation.


*/
