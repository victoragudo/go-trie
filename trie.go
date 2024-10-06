package main

type (
	// Node represents a node in the trie
	Node[T any] struct {
		char     rune
		children map[rune]*Node[T]
		isEnd    bool
		data     T // Additional generic information
	}

	// Trie represents the trie data structure
	Trie[T any] struct {
		root *Node[T]
	}
)

// New initializes the trie
func New[T any]() Trie[T] {
	return Trie[T]{
		root: &Node[T]{
			children: make(map[rune]*Node[T]),
			isEnd:    false,
		},
	}
}

// Insert adds a word to the trie along with associated data
func (trie *Trie[T]) Insert(word string, data T) {
	node := trie.root
	// Traverse each character of the word
	for i, char := range word {
		// Check if the current node has the character in its children
		if _, ok := node.children[char]; !ok {
			// If it doesn't exist, create a new node for the character
			node.children[char] = &Node[T]{
				char:     char,
				children: make(map[rune]*Node[T]),
				isEnd:    false,
				data:     data,
			}
		}

		// Move to the next node
		node = node.children[char]

		// If we're at the last character, set the data
		if i == len([]rune(word))-1 {
			node.isEnd = true
			node.data = data
		}
	}
}

// Search looks for a word in the trie and returns the associated data
func (trie *Trie[T]) Search(word string) (T, bool) {
	node := trie.root
	// Traverse each character of the word
	for _, char := range word {
		// Check if the current node has the character in its children
		if _, ok := node.children[char]; !ok {
			// If it doesn't exist, the word isn't present in the trie
			var zero T
			return zero, false
		}

		// Move to the next node
		node = node.children[char]
	}

	// Check if the last node marks the end of a word and return the data
	if node.isEnd {
		return node.data, true
	} else {
		var zero T
		return zero, false
	}
}

// Delete removes a word from the trie
func (trie *Trie[T]) Delete(word string) {
	trie.deleteHelper(trie.root, word, 0)
}

func (trie *Trie[T]) deleteHelper(node *Node[T], word string, depth int) bool {
	if node == nil {
		return false
	}

	// If we've reached the end of the word
	if depth == len(word) {
		if node.isEnd {
			node.isEnd = false
		}

		// If the node has no children, it can be deleted
		return len(node.children) == 0
	}

	char := rune(word[depth])
	if trie.deleteHelper(node.children[char], word, depth+1) {
		delete(node.children, char)
		// Return true if it's not the end of another word and has no more children
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

// AutoComplete returns a list of words that match the given prefix along with their data
func (trie *Trie[T]) AutoComplete(prefix string) []struct {
	Word string
	Data T
} {
	node := trie.root
	// Find the node corresponding to the prefix
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			// The prefix isn't present in the trie
			return []struct {
				Word string
				Data T
			}{}
		}
		node = node.children[char]
	}

	// Collect all words starting from the prefix node
	var results []struct {
		Word string
		Data T
	}
	trie.collectWords(node, prefix, &results)
	return results
}

// collectWords collects complete words from the given node and prefix
func (trie *Trie[T]) collectWords(node *Node[T], prefix string, results *[]struct {
	Word string
	Data T
}) {
	if node.isEnd {
		*results = append(*results, struct {
			Word string
			Data T
		}{
			Word: prefix,
			Data: node.data,
		})
	}

	for char, child := range node.children {
		trie.collectWords(child, prefix+string(char), results)
	}
}

// CountWords counts the number of words in the trie
func (trie *Trie[T]) CountWords() int {
	count := 0
	trie.countWordsDFS(trie.root, &count)
	return count
}

// countWordsDFS performs a DFS traversal on the trie to count the words
func (trie *Trie[T]) countWordsDFS(node *Node[T], count *int) {
	if node.isEnd {
		*count++
	}
	for _, child := range node.children {
		trie.countWordsDFS(child, count)
	}
}

// GetAllWords returns all words in the trie along with their data
func (trie *Trie[T]) GetAllWords() []struct {
	Word string
	Data T
} {
	var results []struct {
		Word string
		Data T
	}
	trie.collectWords(trie.root, "", &results)
	return results
}

// Clear removes all words from the trie, effectively emptying it
func (trie *Trie[T]) Clear() {
	trie.root = &Node[T]{
		children: make(map[rune]*Node[T]),
		isEnd:    false,
	}
}
