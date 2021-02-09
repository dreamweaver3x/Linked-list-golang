package main

import (
	"fmt"
	"list/internal/linkedlist"
)


func main() {
	pointer := list.NewList(1)
	pointer.Add(2)
	pointer.Remove(1)
	pointer.Add(3)
	pointer.Add(4)
	fmt.Println(pointer)

}