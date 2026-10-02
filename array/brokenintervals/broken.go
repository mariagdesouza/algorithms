package main

import "fmt"

/*

Given an array of integers and a start point, a end point , print all the broken intervals

E.g. start =1 end = 100

a := {2,3,4,7,8}

print 5-6, 9-100


Is the array sorted?

*/

func main() {
	a := []int{2, 3, 4, 7, 8}
	printIntervals(a, 1, 100)
}

func printIntervals(a []int, start, end int) {

	if len(a) < 1 {
		return
	}

	i := 1
	//prev := a[0]??
	for j := 0; j < len(a); j++ { //O(n)
		if i < a[j] {
			fmt.Print(i)
			if a[j]-i > 1 {
				fmt.Print("-", a[j]-1)
			}
			i = a[j]
			fmt.Print(" ,")
		}
		i++
	}
	if i < end {
		fmt.Println(i, "-", end)
	}
}
