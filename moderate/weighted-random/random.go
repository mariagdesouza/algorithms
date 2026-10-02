package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

Given a weighted set, write a function to generate a weighted random number

*/
func main() {
	weightedSet := map[string]int{"A": 90, "B": 5, "C": 100}

	for i := 0; i < 100; i++ {
		fmt.Println(getRandom(weightedSet))
	}

}

func getRandom(weightedSet map[string]int) string {

	var sum int

	if len(weightedSet) == 0 {
		return ""
	}

	var defaultResult string
	var maxWeight int

	// FInd total weights
	for val, item := range weightedSet {
		sum += item
		if item > maxWeight {
			defaultResult = val
		}
	}

	if sum != 0 {
		s := rand.NewSource(time.Now().UnixNano())
		weightDiff := rand.New(s).Intn(sum)

		for key, item := range weightedSet {
			weightDiff = weightDiff - item
			if weightDiff <= 0 {
				return key
			}
		}

	}

	return defaultResult
}
