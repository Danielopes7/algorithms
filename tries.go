package main

import (
	"fmt"
)
 
const ALPHABET_SIZE = 26
 
type TrieNode struct {
	children   [ALPHABET_SIZE]*TrieNode
	endOfWords bool
}
 
func getNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false
 
	for i := 0; i < ALPHABET_SIZE; i++ {
		node.children[i] = nil
	}
 
	return node
}
 
func insert(root *TrieNode, key string) {
	temp := root
 
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] == nil {
			temp.children[index] = getNode()
		}
		temp = temp.children[index]
	}
 
	temp.endOfWords = true
}
 
func search(root *TrieNode, key string) bool {
	temp := root
 
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}
	}
 
	return (temp != nil && temp.endOfWords)
}

func main() {
  words := []string{"a", "and", "go", "golang", "mango"}

  root := getNode()

  for i := 0; i < len(words); i++ {
    insert(root, words[i])
  }

  fmt.Println("contains the word and", search(root, "and"))
  fmt.Println("contains the word golang", search(root, "golang"))
  fmt.Println("contains the word man", search(root, "man"))
}