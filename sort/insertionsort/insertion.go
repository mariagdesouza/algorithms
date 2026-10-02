package main

import "fmt"

func main() {

	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	insertionsort(s)

	fmt.Println(s)
}

func insertionsort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}
