package main

import (
	"fmt"
)

func main() {
	str1 := "abcdef8762"

	str2 := "abcbca2892"

	str3 := ""

	str4 := "nhfccs woqoljh"

	fmt.Println("1", str1, ":", hasUniqueChars(str1))
	fmt.Println("2", str2, ":", hasUniqueChars(str2))
	fmt.Println("3", str3, ":", hasUniqueChars(str3))
	fmt.Println("4", str4, ":", hasUniqueChars(str4))

	fmt.Println("--------")

	fmt.Println("1", str1, ":", hasUniqueCharNoSet(str1))
	fmt.Println("2", str2, ":", hasUniqueCharNoSet(str2))
	fmt.Println("3", str3, ":", hasUniqueCharNoSet(str3))
	fmt.Println("4", str4, ":", hasUniqueCharNoSet(str4))
}

// method 1 using hashset - map[rune]bool
// O(n) where n is length of the string
func hasUniqueChars(test string) bool {
	set := make(map[rune]struct{})
	var empty struct{}
	for _, c := range test {
		if _, ok := set[c]; ok {
			//fmt.Println(c)
			return false
		}
		set[c] = empty
	}
	return true
}

// no additionalData structures - use a 32 -bit integer
func hasUniqueCharNoSet(test string) bool {
	var checkDup uint64
	for _, c := range test {
		i := uint64(c - 'a')
		//find if bit is already set using & .. 1 & 1 = 1
		if checkDup&(1<<i) > 0 {
			return false
		}
		// set the bit
		checkDup |= (1 << i)
	}

	return true
}
