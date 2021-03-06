package list

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

func NewList(x int) *LinkedList {
	structList := node{
		data: x,
	}
	structPointers := LinkedList{
		head: &structList,
		tail: &structList,
	}
	return &structPointers
}

func (p *LinkedList) Add(x int) {
	structList := node{
		data: x,
	}
	p.tail.next = &structList
	p.tail = &structList
}

func (p *LinkedList) Remove(x int) error {
	f := p.head
	prevF := p.head
	for f.data != x && f.next != nil {
		prevF = f
		f = f.next
	}
	if p.head == f {
		p.head = p.tail
	} else {
		if x == f.data {
			prevF.next = f.next
		} else {
			return fmt.Errorf("no values to remove")
		}
	}
	return nil
}

func (p *LinkedList) String() string {
	point := p.head
	x := make([]int, 0)
	for point != nil {
		x = append(x, point.data)
		point = point.next
	}
	return fmt.Sprint(x)
}
