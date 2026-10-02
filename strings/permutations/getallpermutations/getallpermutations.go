package main

import "fmt"

func main() {
	str1 := "aaaa"

	var memo []string
	memo = getallpermutation(str1, "", memo)
	fmt.Println(memo)
	fmt.Println(len(memo))

	getallpermutations(str1)
}

//O(n!)
// duplicate combinations especially with recurring characters
func getallpermutation(str, prefix string, memo []string) []string {

	if len(str) == 0 {
		memo = append(memo, prefix)
	} else {
		for i := 0; i < len(str); i++ {
			rem := str[0:i] + str[i+1:]
			memo = getallpermutation(rem, prefix+string(str[i]), memo) //ac, t // c, ta //  ,cat
		}
	}
	return memo
}

func getallpermutations(str string) {
	freqMap := buildFrequencyMap(str)
	var memo []string
	memo = printPermutationsRec(freqMap, "", len(str), memo)
	fmt.Println(memo)
	fmt.Println(len(memo))

}

// accounting for duplicates can  improve this -  use a map for each character wiht a count
func buildFrequencyMap(str string) map[rune]int {
	frequency := make(map[rune]int)
	for _, c := range str {
		if n, ok := frequency[c]; ok {
			frequency[c] = n + 1
		} else {
			frequency[c] = 1
		}
	}
	return frequency
}

func printPermutationsRec(frequencyMap map[rune]int, prefix string, remaining int, memo []string) []string {

	//fmt.Println(prefix)
	if remaining == 0 {
		memo = append(memo, prefix)
		return memo
	}
	for c, count := range frequencyMap {
		if count > 0 {
			frequencyMap[c] = count - 1
			memo = printPermutationsRec(frequencyMap, prefix+string(c), remaining-1, memo)
			frequencyMap[c] = count
		}
	}
	return memo
}
