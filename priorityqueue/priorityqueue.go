package priorityqueue

// every item has a priority associated with it
// elements are dequeued in order of priority
// if two elements have same priority they are served in order

type Item struct {
	Value    string
	Priority int
	Index    int // position on the heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

//insert O(1) - at the end
func (pq *PriorityQueue) Push(x interface{}) {
	//create it on the heap
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

//removeMax - //O(n)
func (pq *PriorityQueue) Pop() interface{} {

	n := len(*pq)
	item := (*pq)[n-1]   // last element
	item.Index = -1      // for safety
	*pq = (*pq)[0 : n-1] // remove last item
	return item          //return removed element
}

//deleteHighestPriority - O(n)
