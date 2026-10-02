package main

import "fmt"

/*

You have two lists.
Each list will contain lists of length 2, which represent a range (ie. [3,5] means a range from 3 to 5, inclusive).

You need to return the intersection of all ranges between the sets.

If I give you [1,5] and [0,2], the result would be [1,2].

Within each list, the ranges will always increase and never overlap (i.e. it will be [[0, 2], [5, 10] ... ] never [[0,2], [2,5] ... ])
In general there are no "gotchas" in terms of the ordering or overlapping of the lists.
a = [[0, 2], [5, 10], [13, 23], [24, 25]]
b = [[1, 5], [8, 12], [15, 18], [20, 24]]
Expected output:  [[1, 2], [5, 5], [8, 10], [15, 18], [20, 24]]

0 2 1 5

1 2

5 10 8 12

max(start) and min(end)

8 10

13 23 15 18

max(start) and min(end)

15 18


24 25 20 24

24 24

if  min(end) < max(start)=> there is no intersection

Clarification:

varying lengths?

*/

func main() {
	a := [][2]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	b := [][2]int{{1, 5}, {8, 12}, {15, 18}, {20, 24}}
	//a := [][2]int{{0, 4}, {7, 12}}
	//b := [][2]int{{1, 3}, {5, 8}, {9, 11}}
	FindIntersectingIntervals(a, b)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func FindIntersectingIntervals(a [][2]int, b [][2]int) {

	current := a[0]
	next := b[0]

	for i, j := 0, 0; i < len(a) && j < len(b); {

		//fmt.Println(current, next)
		left := max(current[0], next[0])  // max (start)
		right := min(current[1], next[1]) // min(ends)

		if left <= right {
			fmt.Print("[", left, ",", right, "], ")

		}
		current = next
		//next is min i+1 or j+1
		if i+1 < len(a) && j+1 < len(b) {
			if a[i+1][0] < b[j+1][0] {
				i++
				next = a[i]
			} else {
				j++
				next = b[j]
			}
		} else {
			if i+1 < len(a) {
				i++
				next = a[i]
			} else if j+1 < len(b) {
				j++
				next = b[j]
			} else {
				break
			}
		}
		//fmt.Println(i, j)
	}
}

/*
a = [[0, 2], [5, 10], [13, 23], [24, 25]]
b = [[1, 5], [8, 12], [15, 18], [20, 24]]
Expected output:  [[1, 2], [5, 5], [8, 10], [15, 18], [24, 24]]
*/
