package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewTrie tests the initialization of a new Trie
func TestNewTrie(t *testing.T) {
	trie := New[string]()
	assert.NotNil(t, trie.root)
	assert.Empty(t, trie.root.children)
	assert.False(t, trie.root.isEnd)
}

// TestInsertAndSearch tests the insertion and search capability
func TestInsertAndSearch(t *testing.T) {
	trie := New[string]()

	// Insert words
	trie.Insert("cat", "animal")
	trie.Insert("car", "vehicle")

	// Search for existing words
	data, found := trie.Search("cat")
	assert.True(t, found)
	assert.Equal(t, "animal", data)

	data, found = trie.Search("car")
	assert.True(t, found)
	assert.Equal(t, "vehicle", data)

	// Search for non-existing words
	data, found = trie.Search("can")
	assert.False(t, found)
	assert.Equal(t, "", data) // zero value for string
}

// TestDelete tests the delete capability
func TestDelete(t *testing.T) {
	trie := New[string]()

	// Insert words
	trie.Insert("cat", "animal")
	trie.Insert("car", "vehicle")

	// Delete a word
	trie.Delete("cat")

	// Check that "cat" is deleted
	data, found := trie.Search("cat")
	assert.False(t, found)
	assert.Equal(t, "", data) // zero value for string

	// Check that "car" is still present
	data, found = trie.Search("car")
	assert.True(t, found)
	assert.Equal(t, "vehicle", data)
}

// TestAutoComplete tests the auto-complete capability
func TestAutoComplete(t *testing.T) {
	trie := New[string]()

	// Insert words
	trie.Insert("cat", "animal")
	trie.Insert("car", "vehicle")
	trie.Insert("dog", "animal")
	trie.Insert("cart", "shopping")

	// AutoComplete for "ca"
	results := trie.AutoComplete("ca")
	assert.Len(t, results, 3)
	assert.ElementsMatch(t, []struct {
		Word string
		Data string
	}{
		{"cat", "animal"},
		{"car", "vehicle"},
		{"cart", "shopping"},
	}, results)

	// AutoComplete for "dog"
	results = trie.AutoComplete("dog")
	assert.Len(t, results, 1)
	assert.Equal(t, "dog", results[0].Word)
	assert.Equal(t, "animal", results[0].Data)

	// AutoComplete for non-existing prefix
	results = trie.AutoComplete("z")
	assert.Len(t, results, 0)
}

// TestCountWords tests counting the words in the trie
func TestCountWords(t *testing.T) {
	trie := New[string]()

	// No words inserted
	assert.Equal(t, 0, trie.CountWords())

	// Insert words
	trie.Insert("cat", "animal")
	trie.Insert("car", "vehicle")
	trie.Insert("dog", "animal")
	trie.Insert("cart", "shopping")

	// Count words
	assert.Equal(t, 4, trie.CountWords())

	// Delete a word
	trie.Delete("car")
	assert.Equal(t, 3, trie.CountWords())
}

// TestGetAllWords tests retrieving all words in the trie
func TestGetAllWords(t *testing.T) {
	trie := New[string]()

	// Insert words
	trie.Insert("cat", "animal")
	trie.Insert("car", "vehicle")
	trie.Insert("dog", "animal")
	trie.Insert("cart", "shopping")

	// Get all words
	results := trie.GetAllWords()
	assert.Len(t, results, 4)
	assert.ElementsMatch(t, []struct {
		Word string
		Data string
	}{
		{"cat", "animal"},
		{"car", "vehicle"},
		{"dog", "animal"},
		{"cart", "shopping"},
	}, results)
}

// TestClear tests clearing the trie
func TestClear(t *testing.T) {
	trie := New[string]()

	// Insert words
	trie.Insert("cat", "animal")
	trie.Insert("car", "vehicle")

	// Clear the trie
	trie.Clear()

	// Ensure the trie is empty
	assert.Equal(t, 0, trie.CountWords())

	data, found := trie.Search("cat")
	assert.False(t, found)
	assert.Equal(t, "", data)
}

// TestInsertAndSearchInt tests the insertion and search functionality with int type
func TestInsertAndSearchInt(t *testing.T) {
	trie := New[int]()

	// Insert words with integer data
	trie.Insert("one", 1)
	trie.Insert("two", 2)
	trie.Insert("three", 3)

	// Search for existing words
	data, found := trie.Search("one")
	assert.True(t, found)
	assert.Equal(t, 1, data)

	data, found = trie.Search("two")
	assert.True(t, found)
	assert.Equal(t, 2, data)

	data, found = trie.Search("three")
	assert.True(t, found)
	assert.Equal(t, 3, data)

	// Search for non-existing word
	data, found = trie.Search("four")
	assert.False(t, found)
	assert.Equal(t, 0, data) // zero value for int
}

// TestInsertAndSearchFloat tests the insertion and search capability with a float64 type
func TestInsertAndSearchFloat(t *testing.T) {
	trie := New[float64]()

	// Insert words with float data
	trie.Insert("pi", 3.14)
	trie.Insert("golden", 1.618)
	trie.Insert("euler", 2.718)

	// Search for existing words
	data, found := trie.Search("pi")
	assert.True(t, found)
	assert.Equal(t, 3.14, data)

	data, found = trie.Search("golden")
	assert.True(t, found)
	assert.Equal(t, 1.618, data)

	data, found = trie.Search("euler")
	assert.True(t, found)
	assert.Equal(t, 2.718, data)

	// Search for non-existing word
	data, found = trie.Search("unknown")
	assert.False(t, found)
	assert.Equal(t, 0.0, data) // zero value for float64
}

// Define a custom struct to test with the trie
type Person struct {
	Name string
	Age  int
}

// TestInsertAndSearchStruct tests the insertion and search functionality with a custom struct type
func TestInsertAndSearchStruct(t *testing.T) {
	trie := New[Person]()

	// Insert words with struct data
	trie.Insert("john", Person{Name: "John", Age: 30})
	trie.Insert("jane", Person{Name: "Jane", Age: 25})

	// Search for existing words
	data, found := trie.Search("john")
	assert.True(t, found)
	assert.Equal(t, Person{Name: "John", Age: 30}, data)

	data, found = trie.Search("jane")
	assert.True(t, found)
	assert.Equal(t, Person{Name: "Jane", Age: 25}, data)

	// Search for non-existing word
	data, found = trie.Search("unknown")
	assert.False(t, found)
	assert.Equal(t, Person{}, data) // zero value for struct
}

// TestInsertAndSearchSlice tests the insertion and search functionality with a slice type
func TestInsertAndSearchSlice(t *testing.T) {
	trie := New[[]int]()

	// Insert words with slice data
	trie.Insert("nums1", []int{1, 2, 3})
	trie.Insert("nums2", []int{4, 5, 6})

	// Search for existing words
	data, found := trie.Search("nums1")
	assert.True(t, found)
	assert.Equal(t, []int{1, 2, 3}, data)

	data, found = trie.Search("nums2")
	assert.True(t, found)
	assert.Equal(t, []int{4, 5, 6}, data)

	// Search for non-existing word
	data, found = trie.Search("nums3")
	assert.False(t, found)
	assert.Equal(t, []int(nil), data) // zero value for slice is nil
}

// TestAutoCompleteWithStruct tests the auto-complete functionality with struct data
func TestAutoCompleteWithStruct(t *testing.T) {
	trie := New[Person]()

	// Insert words with struct data
	trie.Insert("john", Person{Name: "John", Age: 30})
	trie.Insert("jane", Person{Name: "Jane", Age: 25})
	trie.Insert("jack", Person{Name: "Jack", Age: 35})

	// AutoComplete for "ja"
	results := trie.AutoComplete("ja")
	assert.Len(t, results, 2)
	assert.ElementsMatch(t, []struct {
		Word string
		Data Person
	}{
		{"jane", Person{Name: "Jane", Age: 25}},
		{"jack", Person{Name: "Jack", Age: 35}},
	}, results)

	// AutoComplete for "jo"
	results = trie.AutoComplete("jo")
	assert.Len(t, results, 1)
	assert.Equal(t, "john", results[0].Word)
	assert.Equal(t, Person{Name: "John", Age: 30}, results[0].Data)
}

// TestCountWordsWithDifferentTypes tests the CountWords function with different types
func TestCountWordsWithDifferentTypes(t *testing.T) {
	// Test with int data
	trieInt := New[int]()
	trieInt.Insert("one", 1)
	trieInt.Insert("two", 2)
	assert.Equal(t, 2, trieInt.CountWords())

	// Test with float data
	trieFloat := New[float64]()
	trieFloat.Insert("pi", 3.14)
	trieFloat.Insert("golden", 1.618)
	trieFloat.Insert("euler", 2.718)
	assert.Equal(t, 3, trieFloat.CountWords())

	// Test with struct data
	trieStruct := New[Person]()
	trieStruct.Insert("john", Person{Name: "John", Age: 30})
	trieStruct.Insert("jane", Person{Name: "Jane", Age: 25})
	assert.Equal(t, 2, trieStruct.CountWords())
}
