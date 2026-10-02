package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*

Write  a method to print last k lines on an input fil
*/

func main() {
	err := readFile("test.txt", 2)
	if err != nil {
		log.Fatal(err)
	}

}

func readFile(filename string, k int) error {

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	strBuffer := make([]string, k)
	scanner := bufio.NewScanner(f)
	i := k - 1           //7
	for scanner.Scan() { // N lines O(N)
		strBuffer[i] = scanner.Text()
		fmt.Println(i, strBuffer[i])
		if i > 0 {
			i--
		} else {
			i = k - 1
		}
	}

	for _, str := range strBuffer { //O(k)
		fmt.Println(str)
	}

	return nil
}
