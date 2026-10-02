package main

import (
	"github.com/mariadesouza/algorithms/linkedlist"
)

// You have 2 numbers represented by a linked list, where each node  contains a single digit
// Digits are in reverse order, 1s, tens etc
// write a function that adds the two numbers and returns the sum as a linked list

/*
 6 1 7
 4 9 5


7->1->6
5->9->2

sum

*/

func main() {

	n1 := linkedlist.New(7)
	n1.Append(1)
	n1.Append(6)

	n2 := linkedlist.New(5)
	n2.Append(9)
	n2.Append(4)

	sumhead := add(n1, n2)

	sumhead.Print()
}

func add(n1, n2 *linkedlist.Node) *linkedlist.Node {

	var sumhead, sumtail *linkedlist.Node
	var carry int
	for n1p, n2p := n1, n2; n1p != nil || n2p != nil; {
		if n1p == nil && n2p == nil {
			break
		}
		var digit1, digit2, sum int
		if n1p != nil {
			digit1 = n1p.Data.(int)
			n1p = n1p.Next
		}
		if n2p != nil {
			digit2 = n2p.Data.(int)
			n2p = n2p.Next
		}

		sum = (digit1 + digit2 + carry) % 10
		carry = (digit1 + digit2 + carry) / 10
		//fmt.Println(sum, carry)
		nn := &linkedlist.Node{Data: sum, Next: nil}
		if sumhead == nil {
			sumhead = nn
			sumtail = sumhead
		} else {
			sumtail.Next = nn
			sumtail = sumtail.Next
		}

	}

	if carry > 0 {
		nn := &linkedlist.Node{Data: carry, Next: nil}
		sumtail.Next = nn
		sumtail = sumtail.Next
	}

	return sumhead
}
