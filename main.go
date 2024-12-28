package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: newQueue(), Hash: Hash{}}
}

func newQueue() Queue {
	head:=&Node{}
	tail:=&Node{}
	head.Right=tail
	tail.Left=head
	return Queue{Head: head, Tail: tail, Length}
}

func (c *Cache) Check(str string){
	node:=&Node{}
	
}

type Hash map[string]*Node

func main() {
	fmt.Println("Start cache")

	cache := NewCache()

	for _, word := range []string{"parrot", "avocado", "bird", "nest", "tomato", "tree"} {
		cache.Check(word)
		cache.Display()

	}
}
