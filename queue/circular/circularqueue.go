package circularqueue

import "errors"

// Implement a circular queue using an array - circular array

type CircularQueue struct {
	items      []int
	max        int
	head, tail int
}

func Init(max int) *CircularQueue {
	if max < 1 {
		return nil
	}
	var c CircularQueue
	c.items = make([]int, max)
	return &c
}

func (c *CircularQueue) IsEmpty() bool {
	if c.head == c.tail {
		return true
	}
	return false
}

func (c *CircularQueue) IsFull() bool {

	if (c.tail+1)%c.max == c.head {
		return true
	}
	return false
}

func (c *CircularQueue) Enqueue(e int) error {
	if c.IsFull() {
		return errors.New("failed enqueue: queue is full")
	}
	//add to tail
	c.tail = (c.tail + 1) % c.max
	c.items[c.tail] = e
	return nil
}

func (c *CircularQueue) Dequeue() (int, error) {
	if c.IsEmpty() {
		return 0, errors.New("failed Dequeue: queue is empty")
	}
	//remoe from head - add head
	e := c.items[c.head]
	c.head = (c.head + 1) % c.max
	return e, nil
}
