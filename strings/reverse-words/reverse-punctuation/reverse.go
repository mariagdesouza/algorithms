package main

import (
	"fmt"
	"regexp"
)

/*

Can there  be puctuation in between

*/

func main() {

	str := "getting good at coding, needs a lot of practice"
	fmt.Println(reversewordsEasy(str))

}

type Stack []string

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

	var s Stack
	//words := strings.Fields(str)

	words := regexp.MustCompile("[\\:\\,\\.\\s]+").Split(str, -1)

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
