# Trie Tree in Golang
This repository contains an implementation of a Trie tree data structure in Golang. The Trie tree is a tree-like data structure that is commonly used for efficient string searching and retrieval operations.

## TrieNode Structure
The TrieNode struct represents a node in the Trie tree. It has the following properties:

- children: a map that stores references to child nodes based on the corresponding character.
isEnd: a boolean flag indicating whether the current node represents the end of a word.
Trie Structure
The Trie struct represents the Trie tree itself. It has the following properties:
- root: a pointer to the root node of the Trie tree.

## Functions

### InitTrie
The InitTrie function initializes a new Trie tree and returns it.

### Insert
The Insert function inserts a word into the Trie tree. It takes a string word as a parameter and traverses each character of the word, creating new nodes if necessary.

### Search
The Search function searches for a word in the Trie tree. It takes a string word as a parameter and traverses each character of the word, checking if the characters exist in the Trie tree. It returns a boolean value indicating whether the word is found in the Trie tree or not.

### Delete
The Delete function deletes a word from the Trie tree. It takes a string word as a parameter and traverses each character of the word to find the corresponding node in the Trie tree. Once found, it marks the node as not the end of a word and removes any intermediate nodes that are no longer used.

### AutoComplete
The AutoComplete function returns a list of words that match the given prefix. It takes a string prefix as a parameter and finds the corresponding node in the Trie tree. It then traverses all nodes starting from the prefix node and collects complete words.

### CountWords
The CountWords function counts the number of words in the Trie tree. It performs a depth-first search (DFS) traversal on the Trie tree and increments a count for each node that represents the end of a word.

### PrintTrie
The PrintTrie function prints the entire Trie tree with branches and leaves. It uses a helper function printNode to recursively print each node, prefixing each line with appropriate indentation.

## Usage
Here's an example of how to use the Trie tree:

```
trie := InitTrie()

// Insert words
trie.Insert("apple")
trie.Insert("banana")
trie.Insert("orange")

// Search for words
fmt.Println(trie.Search("apple"))   // Output: true
fmt.Println(trie.Search("banana"))  // Output: true
fmt.Println(trie.Search("orange"))  // Output: true
fmt.Println(trie.Search("grape"))   // Output: false

// Auto-complete
fmt.Println(trie.AutoComplete("ap"))  // Output: [apple]
fmt.Println(trie.AutoComplete("ban")) // Output: [banana]

// Delete word
trie.Delete("apple")
fmt.Println(trie.Search("apple"))   // Output: false

// Count words
fmt.Println(trie.CountWords()) // Output: 2

// Print Trie
trie.PrintTrie()
```

This will output:

```
true
true
true
false
[apple]
[banana]
false
2
├──b
│  ├──a
│  │  └──n
│  │     ├──a
│  │     │  └──n
│  │     │    
```