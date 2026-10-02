package queue

// First In First Out
// remove least recently added

// CPU scheduling
// Disk Scheduling
type Queue []interface{}

//enqueue - insert/add to rear of queue
func (q Queue) Enqueue(a interface{}) Queue {
	return append(q, a)
}

//dequeue - remove from front of queue
func (q Queue) Dequeue() Queue {
	if len(q) < 2 {
		return nil
	}
	return q[1:]
}

//front  - return first element in the queue from front
func (q Queue) Front() interface{} {
	if len(q) > 0 {
		return q[0]
	}
	return -1
}
