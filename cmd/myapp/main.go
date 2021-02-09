package main

import (

	"list/internal/linkedlist"
)


func main() {
pointer := list.NewList(1)
list.AddList(2, pointer)
list.AddList(3, pointer)
list.WriteList(*pointer)
list.RemoveFromList(2, pointer)
list.WriteList(*pointer)
}