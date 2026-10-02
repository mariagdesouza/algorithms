package main

import (
	"fmt"
	"strings"
)

//O(n)

// add one for each char in a
// subtract one for each char in b
/// should all be 0 or false

func main() {
	fmt.Println("Enter String1:")
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Println("Enter String2:")
	fmt.Scan(&str2)

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	if len(str1) == len(str2) && isAnagram(str1, str2) {
		fmt.Println(str2, "is an anagram of", str1)
	} else {
		fmt.Println(str2, "is not an anagram of", str1)
	}

}

func isAnagram(str1, str2 string) bool {
	m := make(map[int]int)

	for _, c := range str1 {
		m[int(c-'a')]++
	}
	for _, c := range str2 {
		m[int(c-'a')]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
