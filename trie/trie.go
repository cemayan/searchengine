package trie

import "strings"

const (
	DesirableDepth = 100
)

var (
	insertCounter = 0
	searchCounter = 0
	deleteCounter = 0
	depthCounter  = 1
	// since go doesn't have set as built-in support we can use map
	searchResults = map[string]struct{}{}
)

// Trie represents prefix tree methods
type Trie interface {
	Insert(word string)
	SetRoot(root *trieNode)
	SearchByPrefix(prefix string) []string
	DeleteByPrefix(prefix string)
	ConvertForIndexing(data string) map[string]map[string]int
}

type trie struct {
	root *trieNode
}

type trieNode struct {
	children map[string]*trieNode
	last     bool
}

// Insert inserts new word to trie
func (t *trie) Insert(word string) {
	lower := strings.ToLower(word)
	last := strings.TrimSpace(lower)

	t.insert(last, t.root)
	// For another insert operation insertCounter must be reset.
	insertCounter = 0
}

// insert inserts new word to trie
func (t *trie) insert(word string, root *trieNode) {

	insertCounter++
	currentWord := string([]rune(word)[0:insertCounter])

	// empty node
	node := &trieNode{make(map[string]*trieNode), false}

	// If currentWord already exist in trie node will be assigned as currentChild(for store the previous nodes)
	if currentChild, ok := root.children[currentWord]; ok {
		node = currentChild
	}

	// Optional
	if len(word) == insertCounter {
		node.last = true
	}

	// add new children to current node
	root.children[currentWord] = node

	if len(word)-1 != insertCounter {
		// recursion until reaching to last word
		// example: tea || t -> te -> tea
		t.insert(word, root.children[currentWord])
	}
}

// map2array converts given map to string array
func (t *trie) map2array(node map[string]struct{}) []string {
	var arr []string
	for key := range node {
		arr = append(arr, key)
	}

	return arr
}

// traversalSearch searches in node
// DesirableDepth is used to show how many depth do we expected
func (t *trie) traversalSearch(node *trieNode) {

	for key, childs := range node.children {

		searchResults[key] = struct{}{}

		// If node has a children depthCounter will be increased
		if len(childs.children) > 0 {
			depthCounter++
		}

		// Do traversalSearch until DesirableDepth count has been reached
		if depthCounter != DesirableDepth {
			t.traversalSearch(childs)
		}
	}

	// reset the counter
	depthCounter = 1
}

func (t *trie) search(prefix string, root *trieNode, arr []string) []string {

	searchCounter++
	currentWord := string([]rune(prefix)[0:searchCounter])

	n := root.children[currentWord]

	//This condition for reach the actual node
	if (searchCounter != len(prefix)-1) && len(prefix) != 1 {
		if n != nil {
			arr = t.search(prefix, n, nil)
		}
	} else {

		if n != nil {
			// given word will be appended to result array
			arr = append(arr, prefix)
			t.traversalSearch(n)
			array := t.map2array(searchResults)
			arr = append(arr, array...)
		}
	}

	return arr
}

// SearchByPrefix searches by prefix
// example: t -> tea && t -> ten  => tea,ten
func (t *trie) SearchByPrefix(prefix string) []string {
	// For another search operation searchCounter and searchResults must be reset.
	searchCounter = 0
	searchResults = map[string]struct{}{}
	return t.search(prefix, t.root, nil)
}

func (t *trie) delete(prefix string, root *trieNode) {
	// For another delete operation deleteCounter must be reset.
	deleteCounter++
	currentWord := prefix[0:deleteCounter]

	// This condition for reach the actual node
	if deleteCounter != len(prefix) {
		t.delete(prefix, root.children[currentWord])
		return
	}

	// delete the key from trie that's given
	delete(root.children, currentWord)
}

// DeleteByPrefix deletes key from trie according to given prefix
func (t *trie) DeleteByPrefix(prefix string) {
	t.delete(prefix, t.root)
}

// SetRoot sets the root
// This method is used to set root from outside
func (t *trie) SetRoot(root *trieNode) {
	t.root = root
}

// New creates an empty trie
func New() Trie {
	return &trie{root: &trieNode{
		children: make(map[string]*trieNode),
		last:     false,
	}}
}
