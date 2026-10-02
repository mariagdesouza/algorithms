package main

import (
	"fmt"
	"strconv"
)

/* given an integer you can flip exactly one bit
 */

func main() {
	var i int32
	i = 1775

	// convert to bit sequence
	a := fmt.Sprintf("%b", i)
	fmt.Println(a)

	fmt.Println("findLongestSequence:", findLongestSequence(a))

	// bit sequence in a string back to string
	b, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}

	fmt.Println("flip 12 (1100)", flipBits(12))

}

func flipBits(n int32) int32 {
	return ^n
}

func findLongestSequenceBit(a string) int {

	max := 0

	return max
}

// O(b)
func findLongestSequence(a string) int {
	var prev rune
	var sequence []int
	var s int
	for i, m := range a {
		if i == 0 {
			prev = m
			sequence = append(sequence, 1)
			if a[i] == '0' {
				s = 1
			}
		} else {
			if prev == m {
				sequence[len(sequence)-1] = sequence[len(sequence)-1] + 1
			} else {
				prev = m
				sequence = append(sequence, 1)
			}
		}
	}

	//fmt.Println(sequence)

	max := 0
	for ; s < len(sequence); s += 2 {
		sum := 0

		sum = sequence[s]

		if s+1 < len(sequence) && sequence[s+1] == 1 {
			sum++
			if s+2 < len(sequence) {
				sum += sequence[s+2]

			}
		}
		if sum > max {
			max = sum
		}

	}

	return max
}
