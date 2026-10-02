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
	var reversedWords string
	reversedWords = reverseString(str)

	words := strings.Fields(reversedWords)
	var rStr string
	for i := 0; i < len(words); i++ {
		rStr += reverseString(words[i])
		rStr += " "
	}
	fmt.Println(rStr)
}

func reverseString(str string) string {
	if len(str) == 0 || str == "" {
		return str
	}

	return string(str[len(str)-1]) + reverseString(str[:len(str)-1])

}
