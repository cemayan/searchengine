package main

import (
	"fmt"
	"github.com/cemayan/searchengine/trie"
)

func main() {

	_trie := trie.New()
	_trie.Insert("tea")
	_trie.Insert("ted")
	_trie.Insert("ten")
	_trie.Insert("tent")

	_trie.Insert("search")

	fmt.Println(_trie.SearchByPrefix("te"))
	//_trie.DeleteByPrefix("te")
	//fmt.Println(_trie.SearchByPrefix("se"))

}
