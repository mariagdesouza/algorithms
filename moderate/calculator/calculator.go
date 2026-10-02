package main

import "unicode"

/*

Given an arithmetic expression, of positive integers and - + / *   compute the result

Note: no parenthesis
Input 2*3+5/6*3+15
Output: 23.5

Solution:

Have to compute based on order of operations  precedent order

Have to provision for expression parenthesis?

division, multiplication, addition, subtraction

- Could use two stacks - one for numbers, one for operators

*/

type Stack []interface{} //LIFO

func (s *Stack) Push(a interface{}) {
	*s = append(*s, a)
}

func (s *Stack) Pop() interface{} {
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *Stack) Top() interface{} {
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func compute(expression string) float64 {

	var result float64

	//var numbers Stack
	//var operators Stack
	var elements Stack

	var currentOperator rune
	for _, c := range expression {

		if c == ')' { ///special case .. need to pop till "("
			if elements.Len() > 1 { // ensure there are elememts
				n1 := elements.Pop().(rune)
				ex := ""
				for n1 != '(' {
					ex += string(n1)
					result = compute(ex)
					elements.Push(result)
				}
			}
		} else if c == '(' || c == '+' || c == '-' {
			elements.Push(c)
		} else if c == '*' || c == '/' 
			currentOperator = c
		} else {
			if unicode.IsNumber(c) {
				if elements.Len() > 1 {
					n1 := elements.Pop().(int)
					n2 := int(c - '0')
					r := applyOperator(n1, n2, currentOperator)
					elements.Push(r)
				}
			}
		}
	}
	//todo - pop all and compute
	return result
}

func applyOperator(n1, n2 int, operator rune) int {
	switch operator {
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	}

	return 0
}

func main() {

}
