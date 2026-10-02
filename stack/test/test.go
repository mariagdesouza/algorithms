package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/stack"
)

func main() {
	var s stack.Stack
	s.Push(10)
	s.Push(5)
	s.Push(1)
	s.Push(7)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
