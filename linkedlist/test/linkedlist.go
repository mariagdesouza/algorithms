package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/linkedlist"
)

func main() {

	n := linkedlist.New(0)
	n.Append(3)
	n.Append(6)
	n.Append(9)

	n.Delete(6)

	for m := n; m != nil; m = m.Next {
		fmt.Println(m.Data)
	}

	n1 := linkedlist.New("a")
	n1.Append("b")
	for m := n1; m != nil; m = m.Next {
		fmt.Println(m.Data)
	}

}
