package main

/*

add to divide

divide = subtract repeatedly

*/
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(addToDivide(a, b))

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func addToDivide(a, b int) (int, error) {
	if b == 0 {
		return -1, fmt.Errorf("Divide by 0")
	}

	if a < b {
		return -1, fmt.Errorf("Float division")
	}

	//a divided by b - subtract b from a (a-b) till a < b
	x, y := abs(a), abs(b)

	result := 0
	for ; x >= y; result++ {
		x = addToSubtract(x, y)
	}
	return result, nil
}

func addToSubtract(a, b int) int {

	b = -b
	fmt.Println(a, b)
	return (a + b)
}
