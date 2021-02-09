package main

import (
	"fmt"
	"list/pkg/linked_list"
)

func main() {
	pointer := list.NewList(1)
	pointer.Add(2)
	_ = pointer.Remove(1)
	pointer.Add(3)
	pointer.Add(4)
	fmt.Println(pointer)
}
