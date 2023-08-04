package tries

/*
	Tries is a tree structure. It is mostly use for store to words.
	Each node a part of word(letter). By connecting in a path from to root.
	Each node can have 26 (alphabet) children. So each node in a triess can be look like array.
	Each node connect each toher with the pointers.
	 And a parent is a array and array size is 26. But it is point just some letters in the array.
	 So what about the rest of letters. They have nil point. Dont have an adrees for that parent.

	The nodes need to boolean value. When the end point come, bool value will be true and the path of the word will be end.

	// 			Add
	Well, ıf we want to add a new word in the tries. First we look the root an looking the to the array that holds the adrees of the children.
	 And ıf there an adress represent to our word to add. If is yes. We going to that adress and do the same process for the our word's second index letter.
	 Repeat that process when we get back the false value. We add a new children for that parent (with use pointer) and check the other children do we have the next letter in that tries part.
	 If we have connect the path again if don't we create a new node and repeat that process.


 	//   	  Search
	In the search algotihm we do the basic process, start at the root and check by one by the childrens, until the all word match up.

	//  	Time Complexity
	Insert or Search = Time Complexity will be O(m) (m =  Length of the word)
	In a node can have 26 different letter. So we have Trade-Off between time and space.

*/
