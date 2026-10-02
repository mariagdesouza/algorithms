package main

import (
	"fmt"
	"strings"
)

// assume you hae a isSubstring
// Give s1 and s2 check if s2 is a rotation of s1 using one call to isSubstring
// e.g. waterbottle is a rotation of erbottlewat
// waterbottle + waterbottle = > wat(erbottlewat)erbottle
func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	if strings.Index(s1+s1, s2) >= 0 {
		fmt.Println(s2, "is a rotation of", s1)
	} else {
		fmt.Println(s2, "is not a rotation of", s1)
	}
}
