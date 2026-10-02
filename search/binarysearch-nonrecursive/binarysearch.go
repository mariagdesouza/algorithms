package main

import "fmt"

func main() {

	list := []string{"at", "ball", "car", "dad"}

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
		if w > list[mid] {
			low = mid + 1
		} else if w < list[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
