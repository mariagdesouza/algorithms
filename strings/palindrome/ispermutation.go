package main

import (
	"fmt"
	"strings"
)

// given a string check if its a permtation of a palindrome
// Print all palindrome permutations of a string

// Input : Tact Coa
// Output: "taco cat", "atco cta"

// ASSUMPTIONS:
// From example we dont care about spaces , we'll trim those
// we are working with only alphabets
// Capitals are not differentiated from lowercase

//APPROACH 1 - BRUTE FORCE

// IF ODD LENGTH - EVERY CHARACTER -1  MUST OCCUR EVEN NUMBER OF TIMES
// IF EVEN LENGTH - EVERY CHARACTER MUST OCCUR EVEN NUMBER OF TIMES
// we'll use fixed array

const (
	M = 26
)

func main() {
	str1 := "Tact Coa"
	str1 = strings.ToLower(str1)
	fmt.Println(str1)

	if getAllPermutations(str1) == 0 {
		fmt.Println("No Palindrome permutations possible")
	}
}

func getAllPermutations(origStr string) int {

	count := 0

	var str []rune
	var freq [M]int
	for _, c := range origStr {
		if c != ' ' {
			i := c - 'a'
			fmt.Println(i)
			if i >= 0 && i < 26 {
				freq[i]++
				str = append(str, c)
			} else { // not a valid char
				return 0
			}
		}
	}

	length := len(str)
	var oddCount int
	var oddC rune
	var half []rune
	for i := 0; i < M; i++ {
		if freq[i] > 0 {
			if freq[i]%2 == 0 {
				half = append(half, rune(i+'a')) //half will contain all letters that occur twice - tac
			} else {
				oddCount++
				oddC = rune(i + 'a')
			}
		}
		if oddCount > 1 {
			return 0
		}
	}

	if length%2 == 0 { // even
		if oddCount == 1 {
			return 0
		}
	} //else { //odd -- oddc will be in the middle
	var memo []string
	memo = getallpermutations(string(half), "", memo)
	for _, str := range memo {
		r := make([]rune, 2*len(str)+oddCount)
		begin := 0
		end := len(r) - 1
		for _, c := range str {
			r[begin] = c
			begin++
			r[end] = c
			end--
		}
		if oddCount == 1 {
			r[len(half)] = oddC
		}
		fmt.Println(string(r))
	}
	//}
	return count
}

func getallpermutations(str, prefix string, memo []string) []string {

	if len(str) == 0 {
		memo = append(memo, prefix)
	} else {
		for i := 0; i < len(str); i++ {
			rem := str[0:i] + str[i+1:]
			memo = getallpermutations(rem, prefix+string(str[i]), memo) //ac, t // c, ta //  ,cat
		}
	}
	return memo
}
