package main

import "fmt"

/*

SPARSE SEARCH


Give a sorted array of strings that is interspersed with empty strings
write a method to find location of a give string

IF not for the empty strings, since this is a SORTED array we could use simple binary search and compare the str

With empty strings interspersed, we implement a simple modification where we move to the closest non-empty if mid is empty

*/

func main() {

	list := []string{"at", "", "", "", "ball", "", "car", "", "", "dad", "", ""}
	//list := []string{"at", "ball", "car", "dad"}

	findword := "dad"

	idx := searchword(list, findword, 0, len(list))

	if idx != -1 {
		fmt.Println("Found at index :", idx)
	} else {
		fmt.Println("Not found")
	}

}

func searchword(list []string, w string, low, high int) int {

	if low > high {
		return -1
	}

	for low < high {
		mid := (low + high) / 2

		if list[mid] == "" && list[mid+1] == "" { // additional logic to move mid pointer
			for mid < high && list[mid] == "" {
				mid = mid + 1
			}
		} else if list[mid] == "" && list[mid-1] == "" {
			for mid > low && list[mid] == "" {
				mid = mid - 1
			}
		} else if list[mid] == "" {
			mid = mid + 1
		}

		if w > list[mid] {
			if mid >= high {
				high--
			} else {
				low = mid + 1
			}
		} else if w < list[mid] {
			high = mid - 1
		} else {
			return mid
		}

	}

	return -1
}
