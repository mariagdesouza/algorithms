package main

// check if a singly linked list is a palindrome

import (
	"fmt"

	"github.com/mariadesouza/algorithms/linkedlist"
)

func main() {

	n := linkedlist.New(0)
	n.Append('m')
	n.Append('o')
	n.Append('m')

	// can use map and if exists remove
	for m := n; m != nil; m = m.Next {
		fmt.Printf("%d:", m.Data)
	}

	var a []rune

	var slowPtr, fastPtr *linkedlist.Node
	slowPtr = n
	fastPtr = n

	for slowPtr != nil && fastPtr != nil {
		fmt.Println(slowPtr.Data.(rune))
		a = append(a, slowPtr.Data.(rune))
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	fmt.Println(a)
	index := len(a) - 1
	isPalindrome := true
	for slowPtr != nil && index >= 0 {
		if slowPtr.Data == a[index] {
			index--
		} else {
			isPalindrome = false
			break
		}
	}

	fmt.Println(isPalindrome)
	//slowPtr is now at mid

}
