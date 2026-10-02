package main

import "fmt"

/*

Mutliply using add

*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func multiply(a, b int) int {

	if a == 0 || b == 0 {
		return 0
	}

	var sum, x, y int

	if a <= b { // add b a times
		x, y = abs(a), abs(b)
	} else { // add a b  times
		x, y = abs(b), abs(a)
	}
	for i := x; i > 0; i-- {
		sum = sum + y
	}

	if (a < 0 && b > 0) || (b < 0 && a > 0) {
		sum = -sum
	}

	return sum
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(multiply(a, b))
}
