package main

import (
	"fmt"
	"os"
)

/*

Diving Board
You are building a diving board by placing a bunch of planks of wood end-to-end

There are two types of planks, one length shorter and one length longer

2 sizes
shorter =10
longer = 20

Use exactly K planks of wood

Write a method to generate all possible lengths for the diving board

for range k
 1 choose one of two types of plank
 2. add to list

Possible lengths of combinations

K=5

s = 10 m=20

10 + 10 +... 10 (50)
20 + 20 +...20 (100)
10 + 20

2 choices k times

2 raise to k choices

*/

func main() {

	var lengths []int

	shorter, longer := 10, 20

	k := 5
	total := 0

	if k == 0 { // required number of planks is zero
		fmt.Println(0)
		os.Exit(0)
	}

	m := make(map[string]struct{})
	lengths = getAllLengths(lengths, shorter, longer, k, total, m)

	fmt.Println(len(lengths))
	for _, l := range lengths {
		fmt.Println(l)
	}

}

//O(2^k)
func getAllLengths(lengths []int, shorter, longer, k, total int, m map[string]struct{}) []int {

	if k == 0 { // required number of planks is zero
		return append(lengths, total)
	}

	key := fmt.Sprintf("%d:%d", k, total)
	if _, ok := m[key]; ok { // we've seen this before
		return lengths
	}

	lengths = getAllLengths(lengths, shorter, longer, k-1, total+shorter, m)
	lengths = getAllLengths(lengths, shorter, longer, k-1, total+longer, m)

	m[key] = struct{}{} // add the key

	return lengths
}
