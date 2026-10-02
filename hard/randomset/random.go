package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
write a method to rsndomly generate a set of m ints from an array of size n

Each element of n must have equal probability of being chosen -> have to itrate through all of n

*/

func random(lo, hi int) int {
	if hi <= 0 {
		return hi
	}
	s := rand.NewSource(time.Now().UnixNano())
	return lo + rand.New(s).Intn(hi)

}

func randomM(orig []int, m int) []int {
	if m > len(orig) {
		return nil
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		k := random(0, N-1)
		b[i] = orig[k]
	}

	/*	for j := m; j < len(orig); j++ {
			k := random(0, j) // between 0 and j
			if k < m {
				b[k] = orig[j]
			}
		}
	*/

	return b
}

const (
	N = 30
)

func main() {
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}
	var m int
	fmt.Scan(&m)
	b := randomM(a, m)
	fmt.Println(b)
}
