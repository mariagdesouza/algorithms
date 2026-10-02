package main

import "fmt"

/*

4 slots

red(R)
Yellow(Y)
Green(G)
Blue(B)

RGGB


you guess:
hit  - RGGB
pseudo-hit  - YRGB

Return Number of hits and number of pseudo hits

*/

const (
	R = iota
	Y
	G
	B
)

func code(c rune) int {
	switch c {
	case 'R':
		return R
	case 'Y':
		return Y
	case 'G':
		return G
	case 'B':
		return B
	default:
		return -1
	}
}

func main() {
	fmt.Println(getScore("RGBY", "GGRR"))
}

func getScore(solution, guess string) (hits int, pseudohits int) {

	//TODO : explore variable lengths
	if len(solution) != len(guess) {

	}
	freq := make([]int, len(solution))

	for _, c := range solution {

		idx := code(c)
		if idx > 0 {
			freq[idx]++
		}
	}

	for i, g := range guess {
		if rune(solution[i]) == g {
			hits++
			idx := code(rune(solution[i]))
			if idx > 0 {
				freq[idx]--
			}
		} else {
			guesscode := code(g)
			if guesscode > 0 && freq[guesscode] > 0 {
				pseudohits++
			}
		}
	}

	return hits, pseudohits
}
