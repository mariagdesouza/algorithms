package main

import "fmt"

/* array-like struct - no len available
   contain sorted positive ints
   elementAt(i)
   element beyond the array will give you -1

*/

func main() {
	list := []int{2, 5, 8, 9, 14, 18, -1, -1, -1, -1, -1, -1, -1, -1} // not in Go if using array the lenght of array shoudl be greater than x
	x := 18
	low := 0
	high := ((low + 1) << 1)
	idx := -1
	for list[high] != -1 {
		fmt.Println(low, high)
		idx = search(list, x, low, high)
		if idx != -1 {
			break
		}
		low = high + 1
		high = (high << 1) + 1
		fmt.Println(low, high)
	}
	if idx != -1 && list[idx] != -1 {
		fmt.Println("found value ", list[idx])
	} else {
		fmt.Println("not found")
	}

}

func search(list []int, x, low, high int) int {

	for low < high {
		mid := (low + high) / 2
		fmt.Println(low, high, mid)
		if x > list[mid] {
			low = mid + 1
		} else if x < list[mid] || list[mid] == -1 {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
