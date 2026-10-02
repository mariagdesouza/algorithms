package stack

type Stack []int

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
