package hashtables

/*
	In the hash table we have a superhero. It is the hash functions.
	We put in our data to hash functions. And hash functions convert to hash code.
	So what is this converting process. For example we have 100 array size of hash table. And our data is words.
	We take a word and convert it's every letter to ascii code and sum it. And we divide that hash table size. We get the remainder, remainder is our answer. And the solution give us to index of that data.

		// Collision Handling
	What about if we find the same index solution. What will we do? We say that collision handling.
	That we have two way solve that.
	1. Open Adressing
	2. Seperate chaining

		// Open Adressing
	If our index is full we will be look the next one, if it okey we will put the data in there. But if not we will do same process until the find free space.
	But in the search algrotihm, open adressing can be a problem for us. So we have one way to collision handling.

		// Seperate chaining
	Simple and easy way to explain seperata chaining would be storing multiple names in one index. How? By using linked lists.
	Storing the names in the indices of the array each index will hold a pointer that points to the head of a linked list has a list of names.
	Linked list will be able to grow and shrink dynamically so we don't have much limitations towards the size of the data we want to store,
	and using the combination of an array and a linked list. Makes seperate chaining a very fast and effective method for hash tables.

	// Insert / Delete / Search in the (seperate chaining) hash table
	Data be stored as keys and the linked lists in the array index will be called buckets. We call the names keys because hash tables are used for storing key and value pairs.
	How insert the key to hash tables?
	We will first get the hash code by putting the ley in a hash function then go to that certain index in the hash table array and insert it right there.
	But we are doing seperate chaining we will have to store it in the bucket, each key in a hash table will have to be unique ,
	so first we need to search the bucket and check that key already exists in bucket or not.
	And ıf the key doesn't exists we will then add the node represented by the key into the linked list.
	If add a data same hash code we will make a new node and we will let that new node become the head node and make that node point to the node that used to be the head node before.

	When we are searching for a key we will first find the bucket that holds the key by putting the key in the hash function ,
	and then go to that bucket and search the bucket one by one until we find the key.

	If we want to delete a key. We will go to that slot just like searching and inserting.
	Look for the key in the bucket and then we disconnect it from the bucket we do that by letting the previous node skip that node we want to delete and make a point to the next node.

		// Time Complexity of hash tables
	We can think of hash tables with seperate chaining as a combination of the benefits of arrays and linked lists.
	Array + linked list == Hash Table

	Worst Case = If all the keys collide with each other so all the keys that we want to store have the same hash code,
	hash table would become a linked list and it would take more time to insert, search and delete.
	So in that case the time complexity would become O(N)

	Best Case : Hash tables the time complexity of hash tables for inserting, deleting and searching is be constant. O(1)


*/
