package main

import (
	"GolangDSA/dstr"
	"fmt"
)

func main() {
	list := dstr.SinglyLinkedList{}
	list.InsertAtHead(5).InsertAtHead(15)
	list.Print()

	fmt.Printf("Length %d\n", list.Length)
}
