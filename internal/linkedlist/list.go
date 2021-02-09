package list

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
}


func NewList (x interface{}) *Linkedlist{
	structList := Node{
		data: x,
	}
	structPointers := Linkedlist{
		head: &structList,
		tail: &structList,
	}
	return  &structPointers
}

func AddList (x interface{}, point *Linkedlist){
	structList := Node {
		data: x,
	}
	point.tail.next = &structList
	point.tail = &structList

}


func RemoveFromList (x interface{}, needHead *Linkedlist){
	p := needHead.head
	prevP := needHead.head
	for p.data != x && p.next != nil {
		prevP = p
		p = p.next
	}
	if x == p.data {
		prevP.next = p.next
	} else {
		fmt.Println("no value in the List")
	}
}

func WriteList (list Linkedlist){
	point := list.head
	for point != nil {
		fmt.Println(point.data)
		point = point.next
	}
}

