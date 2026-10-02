package main

import (
	"container/list"
	"fmt"
)

/*
Hashmap

1. Its a key value store
2. need a O(1) lookup
3. need a hashing function for the key
4. need a way to have buckets so you can store multiple values in case of collisions.
5. Need an equality function to check for keys

store a key - [hashvalue]*list.List


Hash Table supports following operations in Θ(1) time.
1) Search
2) Insert
3) Delete


//BST advantages over Hashmap

* The data in BST is sorted.
* no collisions..

*/

type Node struct {
	Key   string
	Value string
}

type Hashmap struct {
	Size   int
	Bucket []*list.List
}

func NewHashmap(size int) *Hashmap {
	h := new(Hashmap)
	h.Size = size
	h.Bucket = make([]*list.List, size)
	return h
}

func (h *Hashmap) getHashIndex(key string) int {

	p := 0
	for _, c := range key {
		i := int(c)
		p += (i << 5) - i // 31 is an odd prime and  i *31 is alwyas equal to (i << 5)- i
	}
	return (p % h.Size)
}

func (h *Hashmap) Put(key, value string) {
	i := h.getHashIndex(key)
	// Check if exists

	// new index
	if h.Bucket[i] == nil { // make list
		n := &Node{key, value}
		l := list.New() // create list
		l.PushBack(n)
		h.Bucket[i] = l
		return
	}

	//it exists - check
	l := h.Bucket[i]
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(*Node).Key == key { //if key exists in list and update
			e.Value.(*Node).Value = value
			return
		}
	} // not found insert
	n := &Node{key, value}
	l.PushBack(n)

}

func (h *Hashmap) Get(key string) (string, bool) {
	i := h.getHashIndex(key)

	if h.Bucket[i] == nil { // make list
		return "", false
	}

	l := h.Bucket[i]
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(*Node).Key == key {
			return e.Value.(*Node).Value, true
		}
	}

	return "", false
}

func main() {

	h := NewHashmap(256)

	h.Put("joe", "soccer")
	h.Put("jane", "baseball")

	h.Put("pat", "rugby")

	if val, ok := h.Get("joe"); ok {
		fmt.Println("Joe likes:", val)
	}
	if val, ok := h.Get("jane"); ok {
		fmt.Println("Jane likes:", val)
	}

	h.Put("joe", "tennis")
	if val, ok := h.Get("joe"); ok {
		fmt.Println("Joe changed his mind:", val)
	}

}
