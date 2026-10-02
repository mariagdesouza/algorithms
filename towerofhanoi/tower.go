package main

/*
Tower of Hanoi is a mathematical puzzle. It consists of three poles and a number of disks of different sizes which can slide onto any poles. The puzzle starts with the disk in a neat stack in ascending order of size in one pole, the smallest at the top thus making a conical shape. The objective of the puzzle is to move all the disks from one pole (say ‘source pole’) to another pole (say ‘destination pole’) with the help of third pole (say auxiliary pole).

The puzzle has the following two rules:

      1. You can’t place a larger disk onto smaller disk
      2. Only one disk can be moved at a time

We’ve already discussed recursive solution for Tower of Hanoi.
 We have also seen that, for n disks, total 2n – 1 moves are required.
*/
type Stack []int

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() int {
	c := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return c
}

func main() {
	var a Stack
	a.Push(3)
	a.Push(2)
	a.Push(1)

	var b Stack
	towerOFHanoi(a, b)

}

func towerOFHanoi(src Stack, dest Stack) {

}
