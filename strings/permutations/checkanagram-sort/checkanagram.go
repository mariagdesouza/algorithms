package main

import (
	"fmt"
	"sort"
)

// Assume no spaces looking only for characters - sorting O(nlogn)

func main() {
	fmt.Println("Enter String1:")
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Println("Enter String2:")
	fmt.Scan(&str2)
	if len(str1) == len(str2) && checkPermutation(str1, str2) {
		fmt.Println(str2, "is a permutation", str1)
	} else {
		fmt.Println(str2, "is not a permutation", str1)
	}
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortCharactersinStrings(str string) string {
	r := []rune(str)
	sort.Sort(sortRunes(r))
	return string(r)
}

func checkPermutation(str1 string, str2 string) bool {
	if sortCharactersinStrings(str1) == sortCharactersinStrings(str2) {
		return true
	}
	return false
}
