package main

import (
	"fmt"
	"strings"
)

/*

Can there  be puctuation in between

*/

func main() {

	str := "getting good at coding needs a lot of practice"

	//reverse the whole string
	reverseStr := reverseString(str)

	// reverse individual words
	words := strings.Fields(reverseStr)
	var reversedWords string
	for _, word := range words {
		reversedWords += reverseString(word)
		reversedWords += " "
	}
	fmt.Println(reversedWords)
}

func reverseString(str string) string {

	c := []rune(str)
	for i, n := 0, len(c)-1; i < n; i, n = i+1, n-1 {
		c[i], c[n] = c[n], c[i]
	}

	return string(c)
}
