package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

type Cluster struct {
	Serverprefix string
	Pool         []int
}

func CreateCluster(prefix string, n int) *Cluster {

	c := &Cluster{}
	c.Serverprefix = prefix
	for i := 0; i < n; i++ {
		c.Pool = append(c.Pool, i+1)
	}

	return c
}
func (c *Cluster) Allocate(n int) {
	for i := 0; i < n; i++ {
		c.Pool = append(c.Pool, allocate(c.Pool))
	}
}

func (c *Cluster) Deallocate(n int) {
	c.Pool = deallocate(c.Pool, n)
}

func (c *Cluster) String() string {
	var display string
	display = "Available " + c.Serverprefix + "Servers:"
	for i := 0; i < len(c.Pool); i++ {
		display = display + c.Serverprefix + strconv.Itoa(c.Pool[i]) + " "
	}
	return display
}

func main() {

	fmt.Println("How many api servers in use?")
	var n int
	fmt.Scan(&n)
	if n < 1 {
		log.Fatal("Insufficient servers allocated")
	}

	fmt.Println("How many sandbox servers in use?")
	var m int
	fmt.Scan(&m)
	if m < 1 {
		log.Fatal("Insufficient servers allocated")
	}

	apiCluster := CreateCluster("apibox", n)
	sandboxCluster := CreateCluster("sandbox", m)

	apiCluster.Allocate(2)
	sandboxCluster.Deallocate(2)

	fmt.Println(apiCluster)
	fmt.Println(sandboxCluster)
}

func deallocate(a []int, n int) []int {

	for i := 0; i < len(a); i++ {
		if a[i] == n {
			if i == 0 {
				a = a[1:]
			} else if i == len(a)-1 {
				a = a[0 : len(a)-1]
			} else {
				a = append(a[:i], a[i+1:]...)
			}
			break
		}
	}
	return a
}

func allocate(a []int) int {
	if len(a) == 0 {
		return 1
	}
	sort.Ints(a)
	num := -1
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			num = i + 1
			break
		}
	}
	if num == -1 {
		num = a[len(a)-1] + 1
	}
	return num
}
