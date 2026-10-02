package main

import "fmt"

/*

Find all pairs in an array that sum to a specified value k
*/

func main() {
	a := []int{3, 5, 1, 6, 7, 2, 5, 4, 3, 9}
	//a := []int{1}
	k := 6

	//(5,2) (4,3) (6,1)

	pairs := findPairsSum(a, k)
	for _, p := range pairs {
		fmt.Println(p)
	}
}

type Pair struct {
	X, Y int
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func findPairsSum(a []int, k int) []*Pair {
	results := []*Pair{}
	if len(a) < 2 {
		return results
	}

	unpairedElements := make(map[int]int)

	for _, element := range a {
		c := k - element // now we are looking for c
		// k = 5  element is 2 => we are looking for 3

		m, ok := unpairedElements[c]
		if ok && m > 0 { //we found a pair
			results = append(results, &Pair{c, element})
			unpairedElements[c]-- //decrement unpairedElements[3]-- // incase there was more than one 3
		} else {
			unpairedElements[element]++
		}

	}

	return results
}
