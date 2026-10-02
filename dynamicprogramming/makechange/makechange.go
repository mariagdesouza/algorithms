package main

import (
	"fmt"
)

func main() {

	//	var n int64
	//	fmt.Scan(&n)

	n := int64(10)

	denominations := []int64{2, 5, 3, 6}

	//denominationsAsInt := []int64(denominations)
	//sort.Ints(denominationsAsInt)

	m := make(map[int64][]*int64)
	//fmt.Println(makeChange(n, denominations, 0, m))

	ways := int64(0)
	coinchange(int64(len(denominations))-1, n, denominations, &ways, m)

	fmt.Println(ways)

}

//todo dynamically store result and check

func coinchange(startIndex int64, totalMoney int64, coins []int64, ways *int64, m map[int64][]*int64) {
	if startIndex < 0 {
		return
	}
	if totalMoney < 0 {
		return
	}

	if a, ok := m[totalMoney]; ok {
		if a[startIndex] != nil {
			*ways += *a[startIndex]
		}
	}
	if totalMoney == 0 {
		*ways++
		return
	}
	for i := startIndex; i >= 0; i-- {
		coinchange(i, totalMoney-coins[i], coins, ways, m)
	}

	if _, ok := m[totalMoney]; !ok {
		m[totalMoney] = make([]*int64, len(coins))
	}
	a := *ways
	m[totalMoney][startIndex] = &a

}
