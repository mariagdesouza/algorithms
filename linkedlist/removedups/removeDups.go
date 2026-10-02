package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/linkedlist"
)

func main() {

	n := linkedlist.New(0)
	n.Append(3)
	n.Append(6)
	n.Append(3)
	n.Append(9)
	n.Append(3)
	n.Append(6)
	n.Append(9)

	// can use map and if exists remove
	for m := n; m != nil; m = m.Next {
		fmt.Printf("%d:", m.Data)
	}
	fmt.Println("removing duplicates")

	//use empty map to verify duplicates
	buff := make(map[int]struct{})

	prev := n
	for m := n; m != nil; m = m.Next {
		_, ok := buff[m.Data.(int)]
		if ok {
			//delete
			m.Delete(m.Data)
			prev.Next = m.Next
		} else {
			//add to map
			var e struct{}
			buff[m.Data.(int)] = e
			prev = m
		}
	}

	for m := n; m != nil; m = m.Next {
		fmt.Printf("%d:", m.Data)
	}
}
