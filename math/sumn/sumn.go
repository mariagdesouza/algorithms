package main

import "fmt"

/* sum of n numbers 0 to n */

func main() {
	var n int64
	fmt.Scan(&n)

	sumn := (n * (n + 1)) / 2
	fmt.Println("Sum", sumn)
	// sum of 1st n squares
	var sumnsquares int64
	sumnsquares = (n * (n + 1) * ((2 * n) + 1)) / 6
	fmt.Println("Sum of squares", sumnsquares)

}
