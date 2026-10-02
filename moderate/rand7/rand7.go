package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

implement a rand7 given rand5

1. Get 2 rand5
2. add them
3. Mod by 7

*/

func rand5() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(5)
}

func rand7() {
	//fmt.Println(rand5())
	v := rand5() + rand5()
	fmt.Println(v % 7)
}

func main() {
	rand7()
}
