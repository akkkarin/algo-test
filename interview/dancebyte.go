package interview

import "fmt"

// Node - List Node
type Node struct {
	next  *Node
	value int
}

// List - linked list of node, unidirectional
type List struct {
	header *Node
}

// Add - add new node
func (l *List) Add(v int) {
	n := &Node{
		value: v,
		next:  l.header,
	}
	l.header = n
}

// Print - print list
func (l List) Print() {
	p := l.header
	for p != nil {
		fmt.Printf("->%v ", p.value)
		p = p.next
	}
	fmt.Println()
}

// Reverse - reverse the list
// prev and p
// tmp = p, p moves forward
// tmp.next = prev, reverse the direction of current node
// prev moves to tmp(p)
func (l *List) Reverse() {
	p := l.header
	if p == nil {
		return
	}
	var prev *Node
	for {
		if p == nil {
			l.header = prev
			return
		}
		tmp := p
		p = p.next
		tmp.next = prev
		prev = tmp
	}
}

// StepReverse -
func (l *List) StepReverse(step int) {
	if step < 0 {
		return
	}
	p := l.header
	length := 0
	for p != nil {
		p = p.next
		length++
	}
	if length <= step {
		return
	}

	start := length % step
	cur := step

	var curHeader,prev *Node
	p = l.header
	for i := 0; i < length; i++ {
		if i < start{
			p = p.next
			continue
		}
		if i == start{
			curHeader = p
		}
		tmp := p
		p = p.next
		tmp.next = 
	}
}

func gen(size int) *List {
	l := &List{}
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	return l
}
