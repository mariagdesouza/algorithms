package main

import "fmt"

/*

reading a strema of integers
lookup the rank of a number c (the number of values less than or equal to value of x not including x)


track(x int)

func getRankofNumber(x int){

}
5,1,4,4,5,9,7,13,3

getRankofNumber(1) = 0

getRankofNumber(3) = 1

getRankofNumber(4) = 2


TODO: work on this
*/

type BinaryTreeNode struct {
	Item        int
	Rank        int
	Left, Right *BinaryTreeNode
}

func NewNode(item int) *BinaryTreeNode {
	return &BinaryTreeNode{Item: item, Left: nil, Right: nil}
}

func InsertNodes(root *BinaryTreeNode, items ...int) {

	for _, i := range items {
		InsertNode(root, i)
	}
}

func InsertNode(root *BinaryTreeNode, item int) {

	t := root

	if item <= t.Item {
		if t.Left == nil {
			t.Left = NewNode(item)
			return
		}
		fmt.Println(t.Item, "is now ranked", t.Rank)
		t.Rank++
		InsertNode(t.Left, item)
	} else if item > t.Item {
		if t.Right == nil {
			t.Right = NewNode(item)
			fmt.Println(item, "adding", t.Rank)
			t.Right.Rank += t.Rank + 1
			fmt.Println(item, "is now ranked", t.Rank, "=", t.Right.Rank)
			return
		}
		InsertNode(t.Right, item)
	}
}

func getRank(root *BinaryTreeNode, x int) int {
	if root == nil {
		return 0
	}

	if x == root.Item {
		fmt.Println("rank of", x, root.Rank)
		return root.Rank
	}
	if x > root.Item {
		fmt.Println("greater", root.Rank)
		return getRank(root.Right, x)
	} else {
		return getRank(root.Left, x)
	}

}

// left->node->right
func inorder(root *BinaryTreeNode) {

	if root == nil {
		return
	}
	inorder(root.Left)

	fmt.Println(root.Item)

	inorder(root.Right)

}

func main() {
	root := NewNode(5)
	InsertNodes(root, 1, 4, 4, 5, 9, 7, 13, 3)
	//inorder(root)
	fmt.Println("4", getRank(root, 4))
	fmt.Println("9", getRank(root, 9))
	fmt.Println("13", getRank(root, 13))
}
