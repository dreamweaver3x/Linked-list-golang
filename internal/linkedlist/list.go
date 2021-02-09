package list

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
}


func NewList (x int) *Linkedlist{
	structList := Node{
		data: x,
	}
	structPointers := Linkedlist{
		head: &structList,
		tail: &structList,
	}
	return  &structPointers
}

func (p *Linkedlist) Add (x int){
	structList := Node {
		data: x,
	}
	p.tail.next = &structList
	p.tail = &structList

}


func (p *Linkedlist) Remove (x int){
	f := p.head
	prevF := p.head
	for f.data != x && f.next != nil {
		prevF = f
		f = f.next
	}
	if p.head == f {
		p.head = p.tail
	} else {
		if x == f.data{
		prevF.next = f.next
	} else{
		fmt.Println("no value in the List")
	}
	}
}

func (p Linkedlist) Write(){
	point := p.head
	for point != nil {
		fmt.Println(point.data)
		point = point.next
	}
}

