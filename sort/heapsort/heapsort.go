package main

import "fmt"

/*
Heapsort

Binary heap

Complete Binary Tree: Binary tree where every level is filled except the last. All nodes are far left as possible

Binary Heap is a Complete Binary tree where items are stored such that parent node is greater (maxheap) than the two child nodes
Minheap is a binary heap where parent node is less than the two child nodes.

A Complete Binary tree can be represented as an array and is space efficient.

PArent node is at index i (starting at 0)

left node is (2*i)+1
right node is (2*i)+2


Heap Sort Algorithm
1. Build a max heap form the data
2.  The largest item is at the root. Replace it with the last element, reducing the size of the heap
3. Repeat till size > 1

*/

func main() {
	a := []int{12, 11, 13, 5, 6, 7}
	heapsort(a)
	fmt.Println(a)
}

func heapsort(a []int) {

	//build maxheap - initial
	for i := len(a)/2 - 1; i >= 0; i-- {
		maxheapify(a, i)
	}

	// loop from  len(a) - 1 to 1 and swap with 1st element
	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0] //max will be a t first - swap into last position
		maxheapify(a[:i], 0)    ////remove last element and re-maxheapify
	}
}

func maxheapify(a []int, i int) {
	l := (i << 1) + 1
	r := (i << 1) + 2
	largest := i

	if l < len(a) && a[l] > a[largest] {
		largest = l
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxheapify(a, largest)
	}
}
