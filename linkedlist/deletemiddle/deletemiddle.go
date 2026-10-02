package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/linkedlist"
)

// Delete middle node of a linked list

// not first or last but anyone int he middle

func main() {
	n := linkedlist.New(0)
	n.Append(3)
	n.Append(6)
	n.Append(5)
	n.Append(11)
	n.Append(7)
	n.Append(8)
	n.Append(2)

	i, m := 0, n

	n.Print()

	for ; m != nil; m, i = m.Next, i+1 {
		if i == 3 {
			break
		}
	}

	fmt.Println("\nAfter")
	deleteMiddle(m)

	n.Print()
}

func deleteMiddleNode(n *linkedlist.Node) {
	if n == nil || n.Next == nil {
		return
	}

	s, f := n, n
	var p *linkedlist.Node
	for f != nil && f.Next != nil {
		p = s
		s = s.Next
		f = f.Next.Next
	}

	p.Next = s.Next

	s = nil

}

func deleteMiddle(n *linkedlist.Node) bool {
	if n == nil || n.Next == nil {
		return false
	}
	next := n.Next
	n.Data = next.Data
	n.Next = next.Next

	next = nil

	return true
}
