package main

import (
	"container/list"
	"fmt"
)

/*

LRUCache Design

Key-value pirs


Evict lru - item is used move it in front; when max capacity reached remove the back of a list( linked list)

Add elements -

Retreive Value by key  - O(1) -> hashmap

Node
 -key
 -Value

Cache
  - Queue which is a doubly LinkedList of Nodes
	  - Most recently used is moved to front of queue
	  - lru deleted in the back when reached to capacity
  - HashMap of Keys to Nodes
  - capacity

  - lock mechanism for read/write (TODO)

Insert  (Key, Value)
Get (key) Value

FindLRU
UpdateLRU
RemoveLRU

*/

type Node struct {
	Key   string
	Value string //interface{}
}

type Cache struct {
	Capacity int
	Items    map[string]*list.Element //map{key]*node
	NodeList *list.List               // doubly list of *nodes
}

func NewCache(capacity int) *Cache {
	c := Cache{Capacity: capacity}
	c.Items = make(map[string]*list.Element)
	c.NodeList = list.New()

	return &c
}

func (c *Cache) Put(key, value string) {

	// check if it exists - update
	// also promote it
	if e, ok := c.Items[key]; ok {
		e.Value.(*Node).Value = value // update value
		c.NodeList.MoveToFront(e)
		return
	}
	// add Node
	//check capacity before adding
	if c.NodeList.Len() == c.Capacity { // remove least recently used from back of queue
		d := c.NodeList.Back()
		delete(c.Items, d.Value.(*Node).Key)
		c.NodeList.Remove(d)
	}
	// add Node to front of list
	e := c.NodeList.PushFront(&Node{Key: key, Value: value})
	c.Items[key] = e
}

func (c *Cache) Get(key string) (string, bool) {

	// find and move to front
	if e, ok := c.Items[key]; ok {
		c.NodeList.MoveToFront(e)
		return e.Value.(*Node).Value, true
	}
	return "", false
}

func (c *Cache) Remove(key string) bool {
	if e, ok := c.Items[key]; ok {
		delete(c.Items, e.Value.(*Node).Key)
		c.NodeList.Remove(e)
	}
	return false

}

func main() {
	c := NewCache(3)
	c.Put("joe", "soccer")
	c.Put("jane", "tennis")
	c.Put("bess", "rugby")
	if val, ok := c.Get("jane"); ok {
		fmt.Println("jane:", val)
	}
	c.Put("john", "baseball")
	c.Put("jill", "soccer")
	if val, ok := c.Get("joe"); ok {
		fmt.Println("joe:", val)
	} else {
		fmt.Println("joe not found")
	}
	if val, ok := c.Get("bess"); ok {
		fmt.Println("bess:", val)
	} else {
		fmt.Println("bess not found")
	}
	if val, ok := c.Get("jane"); ok {
		fmt.Println("jane:", val)
	}
}
