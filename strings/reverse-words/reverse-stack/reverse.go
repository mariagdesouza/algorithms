package main

import (
	"fmt"
	"strings"
)

/*

Can there  be puctuation in between

*/

func main() {

	str := "getting, good at coding needs a lot of practice"
	fmt.Println("Reversed:", reversewordsEasy(str))

}

type Stack []string

func NewStack(n int) Stack {

	s := make([]string, 0, n)
	return s
}

func (s *Stack) Push(a string) {
	*s = append(*s, a)
}

func (s *Stack) Pop() string {
	a := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return a
}

// Assuming no punctuation
func reversewordsEasy(str string) string {

	words := strings.Fields(str)
	maxwords := len(words)
	if len(words) == 0 {
		return ""
	}
	s := NewStack(maxwords)
	//var s Stack
	for _, word := range words {
		s.Push(word)
	}
	reverseStr := ""
	n := len(s)
	for i := 0; i < n; i++ {
		reverseStr += s.Pop()
		reverseStr += " "
	}
	return reverseStr
}
