package main

import "fmt"

func main() {
	str := "hello"
	fmt.Println("reverse of", str, "is", reverse([]rune(str)))

}

func reverse(str []rune) string {
	if len(str) == 0 {
		return string(str)
	}

	n := len(str) - 1
	i := 0

	for i < n {
		str[i], str[n] = str[n], str[i]
		i++
		n--
	}
	return string(str)
}
