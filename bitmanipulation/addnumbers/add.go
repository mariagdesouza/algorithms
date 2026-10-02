// Add two numbers without using the arithmetic operators

package main

import (
	"fmt"
)

func main() {
	var a, b int64
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println("Sum of ", a, b, " = ", add(a, b))
}

func add(a, b int64) int64 {
	for b != 0 {
		// common bits of a and b go to carry
		carry := a & b
		// xor - sum bits where one is not set
		a = a ^ b
		// shift carry by 1
		b = carry << 1
	}
	return a
}
