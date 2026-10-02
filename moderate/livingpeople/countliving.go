package main

import (
	"fmt"
)

/*

Given a list of people with birth and death years

Compute the year with the most number of alive people

Assume years are between 1900 and 2000

E.g. Person birth 1908 death 1909


*/

const (
	MIN = 1900
	MAX = 2000
)

type Person struct {
	Birth, Death int
}

func NewPerson(b, d int) *Person {
	return &Person{Birth: b, Death: d}
}

// max: 1908
//{1908, 1909}
//{1905, 1908}
//{1910, 1942}

func main() {

	p := []Person{
		{1908, 1909},
		{1905, 1908},
		{1910, 1942},
	}
	fmt.Println(bruteForceCount(p))
	fmt.Println(countLiving(p))
}

// sort births and deaths
// 1905,1908, 1910
// 1908, 1909, 1942

// max alive
// max year = 1
// 1905
// alive = 1

func countLiving(people []Person) int {
	maxyear := MIN
	//m := make(map[int]int) // map[year]count

	a := make([]int, (MAX-MIN)+1)

	for _, p := range people { //(4 *50) + 100
		err := incrementYearsLived(a, p.Birth-MIN, p.Death-MIN)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	maxyear = getMax(a) + 1900

	return maxyear
}

func getMax(a []int) int {
	if len(a) == 0 {
		return -1
	}
	max := a[0]
	maxyear := 0
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
			maxyear = i
		}
	}
	return maxyear
}

func incrementYearsLived(a []int, b, d int) error {
	if d < b {
		return fmt.Errorf("invalid data: death precedes birth")
	}
	for i := b; i <= d; i++ {
		a[i]++
	}
	return nil
}

func bruteForceCount(people []Person) int {

	maxalive := 0
	maxyear := MIN
	for year := MIN; year <= MAX; year++ { //100 *4 = 400
		alive := 0
		for _, p := range people {
			if year >= p.Birth && year <= p.Death {
				alive++
			}
		}
		if alive > maxalive {
			maxalive = alive
			maxyear = year
		}
	}
	return maxyear
}
