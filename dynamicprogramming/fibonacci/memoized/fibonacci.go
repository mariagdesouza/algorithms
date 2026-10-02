package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var n int64
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("0")
	} else {
		memo := make([]int64, n+1)
		fmt.Println(fibonacci(n, memo))
		//fmt.Println(memo)
	}
}

//O(n) - each node computed only once
func fibonacci(n int64, memo []int64) int64 {

	// already computed  - return
	if memo[n] != 0 {
		return memo[n]
	}

	if n == 1 || n == 2 {
		memo[n] = 1
	} else {
		memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	}

	return memo[n]
}
