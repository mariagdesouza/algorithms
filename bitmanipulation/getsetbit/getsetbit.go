package main

import "fmt"

//  1000 & 1000  = 1
func getNthBit(num int, i uint) int {
	if num&(1<<i) != 0 {
		return 1
	}
	return 0
}

// 01000 | 10000 = 11000
func setNthBit(num *int, i uint) int {
	return *num | (1 << i)
}

// 11000 & 01111 = 01000
func clearNthBit(num int, n uint) int {
	mask := ^(1 << n)
	return num & mask
}

func main() {
	num := 8
	fmt.Printf("Original: %b %d\n", num, num)
	index := 3
	fmt.Printf("Is bit at index %d set: %d \n", index, getNthBit(num, uint(index)))

	setnum := setNthBit(&num, uint(index+1))
	fmt.Printf("After setting bit at index %d:  %b %d\n", index+1, setnum, setnum)

	setnum = clearNthBit(setnum, uint(index+1))
	fmt.Printf("After clearing bit at index %d:  %b %d\n", index+1, setnum, setnum)
}
