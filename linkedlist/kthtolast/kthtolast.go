package main

import "github.com/mariadesouza/algorithms/linkedlist"

// find kth to last element of singly linked list
func main() {
	n := linkedlist.New(0)
	n.Append(3)
	n.Append(6)
	n.Append(3)
	n.Append(9)
	n.Append(3)
	n.Append(6)
	n.Append(9)
}

func kthToLast(head *linkedlist.Node, k int) *linkedlist.Node {
	p1, p2 := head, head

	/* Move p1 k nodes to the list*/
	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	//
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
