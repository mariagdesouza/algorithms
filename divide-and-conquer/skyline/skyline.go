package main

import "fmt"

/*

Given n rectangular buildings in a 2-dimensional city, compute the skyline of these buildings,
eliminating hidden lines.
The main task is to view buildings from a side and remove all sections
that are not visible.
All buildings share common bottom and every building is represented by
triplet (left, ht, right)

‘left’: is x coordinated of left side (or wall).
‘right’: is x coordinate of right side
‘ht’: is height of building.

Input: Array of buildings
       { (1,11,5), (2,6,7), (3,13,9), (12,7,16), (14,3,25),
         (19,18,22), (23,13,29), (24,4,28) }
Output: Skyline (an array of rectangular strips)
        A strip has x coordinate of left side and height
        (1, 11), (3, 13), (9, 0), (12, 7), (16, 3), (19, 18),
		(22, 3), (25, 0)

		Consider following as another example when there is only one
building
Input:  {(1, 11, 5)}
Output: (1, 11), (5, 0)


*/

type Building struct {
	Left, Right, Height int
}

type Strip struct {
	Left, Height int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findSkyline(b []Building, lo, hi int) []Strip {
	if lo == hi {
		var skyline []Strip
		skyline = append(skyline, Strip{b[lo].Left, b[lo].Height})
		skyline = append(skyline, Strip{b[lo].Right, 0})
		return skyline
	}
	mid := (lo + hi) / 2
	ls := findSkyline(b, lo, mid) // Recur for left and right halves and merge the two results
	rs := findSkyline(b, mid+1, hi)
	return merge(ls, rs)
}
func merge(l, r []Strip) []Strip {
	var a []Strip

	h1, h2 := 0, 0
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i].Left < r[j].Left { //compare x coordinates
			h1 = l[i].Height
			a = appendStrip(a, l[i].Left, max(h1, h2)) //add less x with max height
			i++
		} else {
			h2 = r[j].Height
			a = appendStrip(a, r[j].Left, max(h1, h2))
			j++
		}
	}
	for ; i < len(l); i++ { // add remaining left
		a = appendStrip(a, l[i].Left, l[i].Height)
	}
	for ; j < len(r); j++ { // add remaining right
		a = appendStrip(a, r[j].Left, r[j].Height)
	}
	return a
}
func appendStrip(a []Strip, x, h int) []Strip {
	if len(a) > 0 && a[len(a)-1].Height == h { //if heights of previous are same dont add
		return a
	}
	// if left is same, check height and put the max height
	if len(a) > 0 && a[len(a)-1].Left == x {
		a[len(a)-1].Height = max(a[len(a)-1].Height, h)
		return a
	}
	// Else append new point
	return append(a, Strip{x, h})
}

func main() {
	b := []Building{
		{2, 9, 10}, {3, 6, 15}, {5, 12, 12},
		{13, 16, 10}, {15, 17, 25}}

	//b := []Building{{2, 9, 10}}
	skyline := findSkyline(b, 0, len(b)-1)
	fmt.Println(skyline)
}
