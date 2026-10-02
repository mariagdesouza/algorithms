package main

import "fmt"

// replace spaces in string with %20
// URLify a given string characters new slice and replace space

// simple solution - copy string

// assume the string has enough space to hold the additional characters
// you are given the true length of the string
// Better solution - in-place shifting of characters

func main() {
	str1 := "Mr John Smith    "
	//lastIndexLength := 13

	r := []rune(str1)
	length := len(str1)
	var truelength int
	for truelength = length - 1; truelength >= 0; truelength-- {
		if r[truelength] != ' ' {
			break
		}
	}
	fmt.Println(truelength)
	if truelength > 0 {
		index := length - 1
		for j := truelength; j >= 0; j-- {
			if r[j] == ' ' && index-3 >= 0 {
				r[index] = '0'
				r[index-1] = '2'
				r[index-2] = '%'
				index = index - 3
			} else {
				r[index] = r[j] //"Mr John Smith     h"
				index--
			}
		}
	}

	fmt.Println(string(r))

}
