// TRIE: Prefix tree, radix tree, digital tree
// variant of a n-ary tree where characters are stored at each node
// each path down the tree represents a word
// we need to have a terminating indicator

package trie

import (
	"strings"
)

const (
	SIZE = 26
)

type TrieNode struct {
	Links       []*TrieNode
	IsEndofWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Links: make([]*TrieNode, SIZE), IsEndofWord: false}
}

// Insert and Search Word run in O(n)  where n is the length of the word
func (t *TrieNode) Insert(word string) {

	if len(word) <= 0 {
		return
	}
	crawler := t

	word = strings.ToLower(word)

	for _, character := range word {
		index := character - 'a' // - first
		if crawler.Links[index] == nil {
			crawler.Links[index] = NewTrieNode()
		}
		crawler = crawler.Links[index]
	}
	crawler.IsEndofWord = true
}

func (t *TrieNode) SearchWord(word string, k int, dictstr string) bool {

	if len(dictstr) > 0 && word == dictstr && t.IsEndofWord == true {
		return true
	}

	if k == len(word) || t.IsEndofWord == true {
		return false
	}

	ch := word[k] - 'a'

	if t.Links[ch] == nil {
		return false
	} else {
		dictstr += string(word[k])
		return t.Links[ch].SearchWord(word, k+1, dictstr)
	}

	return false
}

type Lexicon struct {
	Trie *TrieNode
}

func (l *Lexicon) IsWord(s string) bool {

	word := strings.ToLower(s)

	return l.Trie.SearchWord(word, 0, "")

	//	return false

}

func NewLexicon() *Lexicon {
	dictionary := []string{"GEEKS", "FOR", "QUIZ", "GO"}
	root := NewTrieNode()
	for _, word := range dictionary {
		root.Insert(word)
	}

	return &Lexicon{Trie: root}

}
