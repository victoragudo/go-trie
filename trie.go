package trie

import "fmt"

// TrieNode represents a node in the Trie tree
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the Trie tree
type Trie struct {
	root *TrieNode
}

// InitTrie initializes the Trie tree
func New() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert inserts a word into the Trie tree
func (t *Trie) Insert(word string) {
	node := t.root
	// Traverse each character of the word
	for _, char := range word {
		// Check if the current node has the character in its children
		if _, ok := node.children[char]; !ok {
			// If it doesn't exist, create a new node for the character
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}

		// Move to the next node
		node = node.children[char]
	}

	// Mark the final node as the end of a word
	node.isEnd = true
}

// Search searches for a word in the Trie tree
func (t *Trie) Search(word string) bool {
	node := t.root
	// Traverse each character of the word
	for _, char := range word {
		// Check if the current node has the character in its children
		if _, ok := node.children[char]; !ok {
			// If it doesn't exist, the word is not present in the tree
			return false
		}

		// Move to the next node
		node = node.children[char]
	}

	// Check if the last marked node is the end of a word
	return node.isEnd
}

// Delete deletes a word from the Trie tree
func (t *Trie) Delete(word string) {
	node := t.root
	parents := []*TrieNode{}
	// Find the node corresponding to the word
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			// The word is not present in the tree
			return
		}
		parents = append(parents, node)
		node = node.children[char]
	}

	// Mark the last node as not the end of a word
	node.isEnd = false

	// Remove intermediate nodes that are no longer used
	for len(parents) > 0 && len(node.children) == 0 {
		delete(parents[len(parents)-1].children, rune(word[len(parents)-1]))
		node = parents[len(parents)-1]
		parents = parents[:len(parents)-1]
	}
}

// AutoComplete returns a list of words that match the given prefix
func (t *Trie) AutoComplete(prefix string) []string {
	node := t.root
	// Find the node corresponding to the prefix
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			// The prefix is not present in the tree
			return []string{}
		}
		node = node.children[char]
	}

	// Traverse all nodes starting from the prefix node and collect complete words
	var words []string
	t.collectWords(node, prefix, &words)
	return words
}

// collectWords collects complete words from the given node and prefix
func (t *Trie) collectWords(node *TrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	for char, child := range node.children {
		t.collectWords(child, prefix+string(char), words)
	}
}

// CountWords counts the number of words in the Trie tree
func (t *Trie) CountWords() int {
	count := 0
	t.countWordsDFS(t.root, &count)
	return count
}

// countWordsDFS performs a DFS traversal on the Trie tree to count the words
func (t *Trie) countWordsDFS(node *TrieNode, count *int) {
	if node.isEnd {
		*count++
	}
	for _, child := range node.children {
		t.countWordsDFS(child, count)
	}
}

// PrintTrie prints the entire Trie tree with branches and leaves
func (t *Trie) PrintTrie() {
	t.printNode(t.root, "", "")
}

// printNode is a helper function to print the Trie node recursively with a given prefix and indentation
func (t *Trie) printNode(node *TrieNode, prefix, indent string) {
	if node.isEnd {
		fmt.Println(prefix)
	}

	numberOfChildren := len(node.children)
	for char, child := range node.children {
		numberOfChildren--
		newIndent := indent
		if numberOfChildren > 0 {
			fmt.Print(indent + "├──" + string(char) + "\n")
			newIndent = indent + "│  "
		} else {
			fmt.Print(indent + "└──" + string(char) + "\n")
			newIndent = indent + "   "
		}
		t.printNode(child, prefix+string(char), newIndent)
	}
}
