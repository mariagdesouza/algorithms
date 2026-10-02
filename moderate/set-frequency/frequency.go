package main

import "fmt"

/*  Give a list of transactions

Calculate frequency of all possible item sets

E.g.

str[] = "Apple Mango Orange Mango Guava Guava Mango"

Output : frequency of Apple is : 1
         frequency of Mango is : 3
         frequency of Orange is : 1
         frequency of Guava is : 2


		 Given a list of transactions,
		 How can we calculate the  frequency counts of all possible item-sets?
		 For example,

		 Input: ID Purchased
		 Items 1 apple, banana, lemon
			   2 banana, berry, lemon, orange
			   3 banana, berry, lemon


		Output: Itemset Frequency
				apple, banana 1
				apple, lemon 1
				banana, berry 2
				banana, lemon 3 ...
				apple, banana, lemon 1
				banana, berry, lemon 2 ...
				banana, berry, lemon, orange 1



	   apple banana lemon  berry orange
	t1  1      1      1      0    0
	t2	0	   1      1      1     1
	t3	0	   1      1      1     0

		apple+banana banana+lemon lemon+berry
    t1  	1			  1		     0
    t2      0             1
    t3      0             1


https://en.wikipedia.org/wiki/Apriori_algorithm





n is the number of transactions

n
m max items
*/

func main() {

	sets := make(map[int][]string)

	sets[1] = []string{"apple", "banana", "lemon"} // 2^3 subsets
	sets[2] = []string{"banana", "berry", "lemon", "orange"}
	sets[3] = []string{"banana", "berry", "lemon"}

	m := make(map[int64]int)
	hashset := make(map[int64]string)

	for _, v := range sets {
		//findSubsets(v, m)
		findSubsets(v, 0, m, hashset)
	}

	for k, v := range hashset {
		fmt.Print(v)
		fmt.Print(" :", m[k])
		fmt.Println()
	}

}

//O(mn) Space:O(n+n)
func findSubsets(set []string, i int, m map[int64]int, hashset map[int64]string) {

	if i >= len(set) {
		return
	}

	key := set[i]

	for j := i + 1; j < len(set); j++ { //O(m*m-1* m-2..)
		key += " " + set[j]
		addToKey(key, m, hashset)
		if j-i > 1 {
			addToKey(set[i]+" "+set[j], m, hashset)
		}

	}

	findSubsets(set, i+1, m, hashset) //O(n)

}

func addToKey(key string, m map[int64]int, hashset map[int64]string) {
	h := hash(key)
	if _, ok := hashset[h]; !ok {
		hashset[h] = key
	}

	if _, ok := m[h]; ok {
		m[h]++
	} else {
		m[h] = 1
	}
}

func hash(a string) int64 {
	p := int64(0)
	for _, v := range a {
		i := int64(v)
		p += (i << 5) - i
	}

	return p
}
