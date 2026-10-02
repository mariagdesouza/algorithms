package linkedlist

import "fmt"

type Node struct {
	Next *Node
	Data interface{}
}

func New(d interface{}) *Node {
	nn := Node{Data: d, Next: nil}
	return &nn
}

//Append : Adds a node to linked list
func (n *Node) Append(d interface{}) {
	nn := Node{Data: d, Next: nil}

	if n == nil {
		n = &nn
	} else {
		m := n
		for m.Next != nil {
			m = m.Next
		}
		m.Next = &nn
	}
}

// Delete : Deletes the first node that matches data
func (n *Node) Delete(d interface{}) {

	if n == nil {
		return
	}

	head := n

	if head.Data == d { // head node and it matches
		n = head.Next
		return
	}

	for ; head != nil; head = head.Next {
		if head.Next.Data == d {
			head.Next = head.Next.Next
			break
		}
	}

}

//Append : Adds a node to linked list
func (n *Node) Print() {
	for m := n; m != nil; m = m.Next {
		fmt.Printf("%v ", m.Data)
	}
}
