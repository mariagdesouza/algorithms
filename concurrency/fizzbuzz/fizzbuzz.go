package main

import "fmt"

/*
Multithreaded FizzBuzz

thread1 -  divisible by 3 print Fizz
thread2 - divisisble by 5 print Buzz

thread 3 - divisible by 3 and 5 print FizzBuzz

thread 4 - prints the numbers if not divisible by 3 or 5

each thread
*/

func main() {

	for i := 1; i <= 10; i++ {
		//fizzBuzz(i)
		////fmt.Println()
		fizzBuzzConcurrent(i)

	}
	fmt.Println()
}

func fizzBuzz(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Print("FizzBuzz ")
	} else if n%3 == 0 {
		fmt.Print("Fizz ")
	} else if n%5 == 0 {
		fmt.Print("Buzz ")
	} else {
		fmt.Printf(" %d ", n)
	}
}

func fizzBuzzConcurrent(n int) {

	c := make(chan struct{})
	go func() {
		if n%3 == 0 && n%5 == 0 {
			fmt.Print("FizzBuzz ")
			var n struct{}
			c <- n
		}
	}()
	go func() {
		if n%3 == 0 {
			fmt.Print("Fizz ")
			var n struct{}
			c <- n
		}
	}()

	go func() {
		if n%5 == 0 {
			fmt.Print("Buzz ")
			var n struct{}
			c <- n
		}
	}()
	go func() {
		if n%3 != 0 && n%5 != 0 {
			fmt.Printf(" %d ", n)
			var n struct{}
			c <- n
		}
	}()

	<-c
}
