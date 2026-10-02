package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(addToSubtract(a, b))

}

func addToSubtract(a, b int) int {

	b = -b
	fmt.Println(a, b)
	return (a + b)
}
