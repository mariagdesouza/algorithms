package main

import (
	"container/heap"
	"fmt"

	"github.com/mariadesouza/algorithms/priorityqueue"
)

func main() {
	items := map[string]int{
		"banana": 3, "peach": 5, "apple": 2, "pear": 4,
	}

	//create a priority queue
	pq := make(priorityqueue.PriorityQueue, len(items))
	index := 0
	for value, priority := range items {
		pq[index] = &priorityqueue.Item{Value: value, Priority: priority, Index: index}
		index++
	}
	heap.Init(&pq)

	//Insert a new Item and then modify

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*priorityqueue.Item)
		fmt.Println(item.Priority, ":", item.Value)
	}

}
