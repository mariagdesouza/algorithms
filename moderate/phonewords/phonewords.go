package main

import (
	"fmt"
)

/*

Given a keypad as shown in diagram, and a n digit number,
 list all words which are possible by pressing these numbers.

1 2(ABC) 3(DEF)

4(GHI) 5(JKL) 6(MNO)

7PQRS 8(tuv) 9(xyz)

* 0 #

if n is length of number - 4 ^n combninations

234
=> ADG, AEF, AFG, BDG


Each key has atmost 8 choices
*/
var digitMapping = [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"} //0-9

func main() {
	//var results []string
	a := []int{2, 4, 5, 6}
	output := make([]rune, len(a))
	GetPhoneWords(a, 0, output)
}

func GetPhoneWords(a []int, currDigit int, output []rune) {
	if currDigit == len(a) {
		fmt.Println(string(output))
		return
	}
	if a[currDigit] == 0 || a[currDigit] == 1 {
		return
	}
	letters := digitMapping[a[currDigit]]
	for _, c := range letters { // 3 or 4
		output[currDigit] = c
		GetPhoneWords(a, currDigit+1, output) // n times
	}
}
