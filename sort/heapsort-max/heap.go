package main

import "fmt"

// MaxHeap sort - max is given as 1st element
//O(n logn)

const (
	MAX = 10
)

func main() {
	a := []int{10, 4, 3, 6, 1, 2, 9, 5, 7, 8}

	size := len(a) - 1
	a[0], a[size] = a[size], a[0]
	size--
	buildMaxHeap(a[:size])
	for i := size; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		maxHeapify(a[:i], 0)
	}

	fmt.Println(a)
}

func buildMaxHeap(a []int) {

	for i := len(a)/2 - 1; i >= 0; i-- {
		maxHeapify(a, i)
	}

}

func maxHeapify(a []int, i int) {

	l := (i << 1) + 1
	r := (i << 1) + 2
	largest := i
	if l < len(a) && i >= 0 && a[l] > a[largest] {
		largest = l
	}

	if r < len(a) && r >= 0 && a[r] > a[largest] {
		largest = r
	}

	if largest != i { // changed - exchange
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest)
	}

}
