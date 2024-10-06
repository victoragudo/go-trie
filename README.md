
# Trie Implementation in Go

This project implements a generic Trie (prefix tree) data structure in Go, supporting generic types. It allows inserting, searching, deleting, and autocompleting words with associated data. The implementation is type-safe and flexible, making it possible to store any kind of data in the trie.

## Features

- **Insert**: Add words to the trie with associated data of any type.
- **Search**: Look up words in the trie and retrieve the associated data.
- **Delete**: Remove words from the trie.
- **Autocomplete**: Retrieve all words starting with a given prefix.
- **CountWords**: Count the total number of words in the trie.
- **GetAllWords**: Retrieve all the words in the trie along with their data.
- **Clear**: Reset the trie, removing all words and data.

## Table of Contents

- [Trie Implementation in Go](#trie-implementation-in-go)
    - [Features](#features)
    - [Installation](#installation)
    - [Usage](#usage)
        - [Inserting Words](#inserting-words)
        - [Searching for Words](#searching-for-words)
        - [Deleting Words](#deleting-words)
        - [Autocompleting Words](#autocompleting-words)
        - [Counting Words](#counting-words)
        - [Clearing the Trie](#clearing-the-trie)
    - [Example](#example)
    - [Running Tests](#running-tests)
    - [License](#license)

## Installation

To use this Trie implementation, you need to have Go installed. You can download and install Go from [here](https://golang.org/dl/).

### Clone the Repository

```bash
git clone https://github.com/yourusername/trie-go
cd trie-go
```

### Install Dependencies

This project uses `testify` for testing. You can install it by running:

```bash
go get github.com/stretchr/testify
```

## Usage

The trie can store any type of data associated with words. Below are examples showing how to use the `Trie` data structure with different types.

### Inserting Words

You can insert words into the trie, associating each word with any type of data. For example:

```go
trie := New[int]()
trie.Insert("apple", 5)
trie.Insert("banana", 10)
```

### Searching for Words

You can search for words in the trie and retrieve their associated data:

```go
data, found := trie.Search("apple")
if found {
    fmt.Println("Found 'apple' with data:", data)
} else {
    fmt.Println("'apple' not found")
}
```

### Deleting Words

Words can be deleted from the trie:

```go
trie.Delete("banana")
```

### Autocompleting Words

You can find all words that match a given prefix:

```go
results := trie.AutoComplete("app")
for _, result := range results {
    fmt.Printf("Word: %s, Data: %v\n", result.Word, result.Data)
}
```

### Counting Words

To count the total number of words stored in the trie:

```go
count := trie.CountWords()
fmt.Println("Total words in trie:", count)
```

### Clearing the Trie

To clear all data from the trie:

```go
trie.Clear()
```

## Example

Here is a full example that demonstrates how to create a trie, insert words with different types of data, search, and autocomplete:

```go
package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    // Initialize a Trie with int data
    trieInt := New[int]()
    trieInt.Insert("one", 1)
    trieInt.Insert("two", 2)
    
    // Search for a word
    data, found := trieInt.Search("one")
    if found {
        fmt.Println("Found 'one' with data:", data)
    } else {
        fmt.Println("'one' not found")
    }

    // Autocomplete example
    trieInt.Insert("three", 3)
    results := trieInt.AutoComplete("t")
    for _, result := range results {
        fmt.Printf("Word: %s, Data: %d\n", result.Word, result.Data)
    }

    // Trie with custom struct data
    trieStruct := New[Person]()
    trieStruct.Insert("john", Person{Name: "John", Age: 30})
    trieStruct.Insert("jane", Person{Name: "Jane", Age: 25})

    person, found := trieStruct.Search("john")
    if found {
        fmt.Printf("Found 'john' with data: Name=%s, Age=%d\n", person.Name, person.Age)
    } else {
        fmt.Println("'john' not found")
    }
}
```

## Running Tests

This project includes a suite of tests to ensure that all functionality works as expected. The tests cover insertion, searching, deletion, autocomplete, and counting words for various data types (`int`, `float64`, structs, slices, etc.).

### Run Tests

To run the tests, you can use the following command:

```bash
go test ./...
```

The tests use the `testify` library for assertions. If you haven't installed it yet, you can do so by running:

```bash
go get github.com/stretchr/testify
```

### Test Examples

Here are a few test cases included in the project:

- **TestInsertAndSearchInt**: Verifies the insertion and search functionality for `int` data.
- **TestInsertAndSearchFloat**: Tests insertion and search for `float64` data.
- **TestInsertAndSearchStruct**: Ensures that the trie works with custom structs.
- **TestAutoCompleteWithStruct**: Validates that autocomplete works with struct data.

You can find the full list of tests in the `trie_test.go` file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
