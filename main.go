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
	
	if val, ok:=c.Hash[str]; ok{
		node= c.Remove(val)
	}else{
		node=&Node{Val: str}
	}
	c.Add(node)
}

func (c *Cache) Remove(n *Node) *Node{
	fmt.Printf("remove: %s\n", n.Val)
	left:=n.Left
	right:=n.Right
	right.Left=left
	left.Right=right
	c.Queue.Length-=1;
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) *Node{
	fmt.Printf("add %s\n", n.Val)
	tmp:=c.Queue.Head.Right

	c.Queue.Head.Right=n
	
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
