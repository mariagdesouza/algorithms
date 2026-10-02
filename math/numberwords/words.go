package main

/*
Convert number to words - TODO complete this
*/

import "fmt"

var ones = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var teens = []string{
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func main() {
	fmt.Println(numberToWords(15))
}

func numberToWords(n int) string {

	x := 10
	if n >= x {
		a := n % x
		return teens[a]
	}
	str := "'"
	if n >= 20 {
		a := n / x
		str = tens[str]
		b = n%10
	} 
		b := n % 1
		return ones[b]
	}
	return ""
}
