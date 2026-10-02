package main

import "fmt"

func main() {
	//a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//fmt.Println(a, "kadaneMaxsum", kadaneMaxsum(a))
	//resSum, resArr := kadaneMaxSubArray(a)
	//	fmt.Println(a, "kadaneMaxSubArray", resSum, resArr)

	//b := []int{1, 3, 2, 3, 4, 8, 7, 9}
	//fmt.Println(kadaneMaxSubArray(b))

	//c := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	//resSum, resArr = kadaneMaxSubArray(c)
	//fmt.Println(c, "kadaneMaxsum", kadaneMaxsum(c))
	//fmt.Println(c, "kadaneMaxSubArray", resSum, resArr)

	fmt.Println(maximizeRatings([]int{5, 9, -1, -3, 4}))
	fmt.Println(maximizeRatings([]int{-1, -2, -3, -4}))
	fmt.Println(kadaneMaxsum([]int{-1, -2, -3, -4}))

}

// Complete the maximizeRatings function below.
func maximizeRatings(ratings []int) int {

	if len(ratings) == 0 {
		return 0

	}

	if len(ratings) == 1 {
		return ratings[0]

	}

	//	min := ratings[0]
	//	minIdx := 0

	sum := make([]int, len(ratings))

	// cannot skip more than 1 in a row
	//5 9 3 -1 -3 4 // 14 -1 = 13
	// 5+9 + -1 +  4 = 17
	//  5+9   min -1 minIdx := 2
	// 14-1     min -3 minIdx := 3
	// 13 + 4
	// -1, -2, -3
	// sum -1
	sum[0] = ratings[0]
	sum[1] = max(ratings[1], ratings[0]+ratings[1]) // 5, 5+9
	//fmt.Println(sum[1], ratings[0], ratings[1])
	for i := 2; i < len(ratings); i++ {
		//	if ratings[i] < 0 && sum[i-1] < 0 {
		//	sum[i] = max(max(ratings[i], sum[i-1]), sum[i-2])
		//} else {
		sum[i] = max(ratings[i]+sum[i-1], ratings[i]+sum[i-2]) //either skip i-1 or i-2
		//	}
	}
	fmt.Println(sum)
	return max(sum[len(ratings)-1], sum[len(ratings)-2])

}
