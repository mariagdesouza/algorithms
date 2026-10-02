package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
A shuffle deck of cards

[1] [2] [3] [4] [5]

shuffle n elements (recursively)
 shuffle(n-1)
 swap the nth element with 1 random element

*/

// get a random number between 0 to i
func random(lo, hi int) int {
	//fmt.Println(hi)
	if hi == lo {
		return lo
	}
	s1 := rand.NewSource(time.Now().UnixNano())

	return lo + rand.New(s1).Intn(hi)
}

func shuffleCardsRecursive(cards []int, i int) {
	if i == 0 {
		return
	}
	shuffleCardsRecursive(cards, i-1) //shuffle earlier part
	k := random(0, i)
	//swap k and i elements
	cards[k], cards[i] = cards[i], cards[k]
	//Return shuffled array
}

func shuffleCards(cards []int) {
	for i := N - 1; i > 0; i-- {
		k := random(0, i)
		if k != i {
			cards[k], cards[i] = cards[i], cards[k]
		}
	}
}

const (
	N = 52
)

func main() {

	cards := make([]int, N)
	for i := 0; i < N; i++ {
		cards[i] = i
	}
	//shuffleCardsRecursive(cards, len(cards)-1)
	shuffleCards(cards)
	fmt.Println(cards)
}
