package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 - Int to Other Base System

 1. Divide by base
 2. Get remainder of division by base  (least significat number)

 3. Divide the quotient from step 1 by  base
 4. Record remainder as next digit to the left of step 2


 15 -> binary

 15/2 => q = 7  r = 1   result =1
 7/2 => q=3 r =1 result 11
 3/2  => q=1 r=1 result 111
 1/2  => q=0 r=1 result = 1111

 15/16



- Base system to decimal
 1011

    (1*2^0) + (1*2^1)+ (0*2^2)+ (1*2^3)
-


*/

func main() {
	n := 15
	fmt.Println(n, "base 2", decimalToBase(n, 2))
	fmt.Println(n, "base 8", decimalToBase(n, 8))

	fmt.Println("10111 to decimal", baseTodecimal("10111", 2))
	fmt.Println("15 to decimal", baseTodecimal("17", 8))
}

func decimalToBase(n, base int) string {

	result := ""
	q := n / base
	r := n % base
	for q > 0 {
		result = strconv.Itoa(r) + result
		r = q % base
		q = q / base
	}
	result = strconv.Itoa(r) + result
	return result
}

func baseTodecimal(s string, base int) int64 {
	var result int64

	for i, c := range s {
		result += (int64(c-'0') * int64(math.Pow(float64(base), float64(i))))
	}
	return result
}
