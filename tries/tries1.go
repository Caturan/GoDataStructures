package tries

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Struct for node (represent each node in the trie)
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Struct for trie (represent a trie and has a )
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert method (Insert will take in a word and add it to the trie)
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // Take a help from Ascii
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search method (will take in a word and RETURN true is that word is included in the trie)
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // Take a help from Ascii
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex] // Update
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func Demo1() {
	//testTrie := InitTrie()
	//fmt.Println(testTrie.root)

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregan",
		"oregano",
		"oreo",
	}

	myTrie := InitTrie()

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	//fmt.Println(myTrie.Search("aragorn"))
	fmt.Println(myTrie.root)

}
