package main

import "fmt"

// reverse an array O(n/2) Inplace swap

func main() {
	a := []int{4, 6, 2, 7, 9, 8}

	//O(n / 2)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		//swap
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}
