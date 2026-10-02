package main

import "fmt"

func main() {
	a := 2
	b := 3

	//a, b = b, a
	fmt.Println("Before:", a, b)
	a, b = swap(a, b)
	fmt.Println("After", a, b)
}

func swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = b ^ a
	return a, b
}
