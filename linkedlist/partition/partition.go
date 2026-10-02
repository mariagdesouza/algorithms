package main

import "github.com/mariadesouza/algorithms/linkedlist"

// partition a linked list around a value x
// all nodes less than x should come before
// all nodes greater than x should come after

func main() {
	n := linkedlist.New(0)
	n.Append(3)
	n.Append(6)
	n.Append(5)
	n.Append(11)
	n.Append(7)
	n.Append(8)
	n.Append(5)
	n.Append(2)

	partitionStable(n, 5)
}

func partitionStable(n *linkedlist.Node, x int) {

	if n == nil {
		return
	}

	//m := linkedlist.New(x) // this is first element

	var l, lhead, m, mhead, h *linkedlist.Node

	for s := n; s != nil; s = s.Next {

		nn := &linkedlist.Node{Data: s.Data.(int), Next: nil}

		if s.Data.(int) < x {
			if l == nil {
				l = nn
				lhead = l
			} else {
				l.Next = nn
				l = l.Next
			}
		} else if s.Data.(int) > x {
			if h == nil {
				h = nn
			} else {
				h.Append(s.Data.(int))
			}
		} else {
			if m == nil {
				m = linkedlist.New(s.Data.(int))
				mhead = m
			} else {
				m.Next = nn
				m = m.Next
			}
		}
	}

	l.Next = mhead
	m.Next = h

	lhead.Print()
}
