package main

import "fmt"

/* search for a number ina  rotated array */
// O(logn)
func main() {
	a := []int{15, 16, 20, 25, 1, 3, 4, 5, 6, 7, 10, 14}
	n := 5

	fmt.Println("Searching for", n, "in", a)
	fmt.Println(search(a, n, 0, len(a)-1))
}

func search(a []int, n int, left int, right int) int {
	if right < left {
		return -1
	}
	mid := (left + right) / 2
	if n == a[mid] {
		return mid
	}

	if a[left] < a[mid] { // left normally ordered
		if n >= a[left] && n < a[mid] {
			return search(a, n, left, mid-1)
		} else {
			return search(a, n, mid+1, right)
		}
	} else if a[mid] < a[left] { //// right normally ordered
		if n > a[mid] && n <= a[right] {
			return search(a, n, mid+1, right)
		} else {
			return search(a, n, left, mid-1)
		}
	} else {
		if a[left] == a[mid] {
			if a[mid] != a[right] { //if rigth is different look there
				return search(a, n, mid+1, right)
			} else { //search both sides
				result := search(a, n, left, mid-1)
				if result == -1 {
					return search(a, n, mid+1, right)
				} else {
					return result
				}
			}
		}
	}

	return -1
}
