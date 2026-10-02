package main

import "github.com/mariadesouza/algorithms/linkedlist"

// check if a linked list is circular
func main() {

}

// return start of loop
func checkLoop(n *linkedlist.Node) *linkedlist.Node {
	slow, fast := n, n

	for fast != nil && fast.Next != nil {
		slow = slow.Next      //1 hop
		fast = fast.Next.Next // 2 hop
		if slow == fast {     //Collision - loop
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = n           // return slow to head
	for slow != fast { //move them at same pace - where they meet is start of loop
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
