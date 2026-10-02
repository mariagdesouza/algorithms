package main

/*

Median of 2 sorted arrays

*/

func main() {
	arr1 := []int{1, 12, 15, 26, 38}
	arr2 := []int{2, 13, 17, 30, 45}

	getMedian(arr1, arr2)
}

func getMedian(a1, a2 []int) int {
	if len(a1) == 0 {
		return a2[len(a2)/2]
	}
	if len(a2) == 0 {
		return a1[len(a1)/2]
	}

	m1 := median(a1)
	m2 := median(a2)

	if m1 > m2 {
		return getMedian(a1[:len(a1)/2], a2[len(a2)/2:])
	} else {
		return getMedian(a1[len(a1)/2:], a2[:len(a2)/2:])
	}

}

func median(a []int) int {

	n := len(a)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return a[0]
	}

	return a[(n / 2)]
}
