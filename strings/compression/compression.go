package main

import (
	"fmt"
	"strconv"
)

// aabcccccaaa => a2b1c5a3

func main() {

	in := "aabcccccaaa"

	fmt.Println(compressString(in))
}

func compressString(in string) string {
	out := string(in[0])
	cc := 1
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			cc++
		} else {
			out += strconv.Itoa(cc)
			out += string(in[i+1])
			cc = 1
		}
	}
	out += strconv.Itoa(cc)
	return out
}
