package linkedlists

/*
	Data stored in a linked list, it's stored in a sequence of nodes.
	It is often compared to an array
	In linked lists nodes contain an adress indicating where the next node.

	In array each box in an array is indexed and they are phsically next to each other.
	 If we want the change a value in array for example the fourth element 15 to 16 we can simply do that.
	When it comes to linked lists we first need to go to the head and then follow along the linked list after we come to the fourth node then we'll be able to change that value.
	So it would be much slower. Going to the Kth element : Slower O(n)
	If the linked list consist of like 1000 nodes then we'll have to do that process a thosund times.

	So why we use linked list?
	If we use a linked list adding or deleting an ath the beginning of the list would be quite simple = Faster O(1)

	When we was mentioning arrays, we typed in an array the locations are physically to next to each other and the adresses are fixed.
	 So in that case when we want to add something at the beginning of an array all the values will have to shift and that can be a bit more complicated than a linked list.
	So that's an example of why a linked list would be better than array.

	Doubly linked list : In a doubly linked list each node contain the adress of the next and also the previous.
	In this case adding an element at the end of the linked list would be easier than a singly linked list.


*/
