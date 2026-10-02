package main

import "fmt"

/*
Find max of 2 mumbers without comparison

2 ^3 = 5
010
011
====
001

010
101
===
000 => 2^ (5 &

*/

func max(x, y int) int {

	signOfdiff := (x << uint(y)) //x > y signOfDiff  = x-y  is positive
	return x ^ ((x ^ y) & -signOfdiff)
}

func min(x, y int) int {
	signOfdiff := (x << uint(y))
	return y ^ ((x ^ y) & -signOfdiff)
}

func main() {
	flipBits
	x, y := -2, -15
	fmt.Println("Max of:", x, y, "is", max(x, y))
	fmt.Println()
	fmt.Println("Min of:", x, y, "is", min(x, y))
}
