package main

import (
	"container/list"
	"sync"
	"time"
)

type Node struct {
	Key        string
	Value      interface{}
	Expiration int64
}

// LRUCache : least recently used cache
// doubleList is a double linked that makes  add/remove O(1)
// the cache is a hash map of keys to elements makes get O(1)
type LRUCache struct {
	capacity   int
	ttlSeconds time.Duration // in seconds
	items      map[string]*list.Element
	doubleList *list.List
	lock       *sync.RWMutex
}

func main() {

}
