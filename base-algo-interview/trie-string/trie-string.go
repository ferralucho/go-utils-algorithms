/*
Trie for String Search and Manipulation
- Tries store strings in a tree-like structure, enabling fast lookups.
- They’re ideal for autocomplete or spell-check features.
- Example: Implement an autocomplete system.
*/

package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("banana")

	fmt.Println("Search for 'apple':", trie.Search("apple"))   // Output: Search for 'apple': true
	fmt.Println("Search for 'app':", trie.Search("app"))       // Output: Search for 'app': true
	fmt.Println("Search for 'banana':", trie.Search("banana")) // Output: Search for 'banana': true
	fmt.Println("Search for 'ap':", trie.Search("ap"))         // Output: Search for 'ap': false

	fmt.Println("StartsWith 'app':", trie.StartsWith("app"))   // Output: StartsWith 'app': true
	fmt.Println("StartsWith 'ban':", trie.StartsWith("ban"))   // Output: StartsWith 'ban': true
	fmt.Println("StartsWith 'aple':", trie.StartsWith("aple")) // Output: StartsWith 'aple': false
}
