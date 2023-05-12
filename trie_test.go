package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	trie := InitTrie()
	words := []string{"apple", "banana", "orange", "grape", "peach", "kiwi"}

	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected word '%s' to be in the trie, but was not found", word)
		}
	}

	if trie.Search("pumpkin") {
		t.Error("Expected word 'pumpkin' not to be in the trie, but was found")
	}
}

func TestDelete(t *testing.T) {
	trie := InitTrie()
	words := []string{"apple", "banana", "orange", "grape", "peach", "kiwi"}

	for _, word := range words {
		trie.Insert(word)
	}

	trie.Delete("banana")
	if trie.Search("banana") {
		t.Error("Expected word 'banana' not to be in the trie after deletion, but was found")
	}
}

func TestAutoComplete(t *testing.T) {
	trie := InitTrie()
	words := []string{"apple", "banana", "orange", "grape", "peach", "kiwi", "grapefruit"}

	for _, word := range words {
		trie.Insert(word)
	}

	expected := []string{"grape", "grapefruit"}
	actual := trie.AutoComplete("grap")
	sort.Strings(actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected autocomplete to return %v, but got %v", expected, actual)
	}
}

func TestCountWords(t *testing.T) {
	trie := InitTrie()
	words := []string{"apple", "banana", "orange", "grape", "peach", "kiwi"}

	for _, word := range words {
		trie.Insert(word)
	}

	expected := 6
	actual := trie.CountWords()

	if expected != actual {
		t.Errorf("Expected trie to have %d words, but got %d", expected, actual)
	}
}
