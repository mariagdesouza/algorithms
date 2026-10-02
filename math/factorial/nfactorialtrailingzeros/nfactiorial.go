package main

import "fmt"

/*

Compute trailing zeros in n factorial

n factorial - 1*2*3*4 ... n

contirbute to the zeros

10 => 5 *2

Multiples of 5

15 => 1 (5*3)

25 => 2 (5*5)



*/

func factorOf5(n int) int {

	count := 0
	for n%5 == 0 {
		count++
		n = n / 5
	}
	return count
}

func countFactZeros(n int) int {
	if n < 5 {
		return 0
	}
	count := 0
	for i := 5; i < n; i += 5 {
		count += factorOf5(i)
	}
	return count
}

func main() {
	fmt.Println(countFactZeros(135))
}
