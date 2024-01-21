package dstr

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	head   *Node
	Length int
}

func (l *SinglyLinkedList) Search(val int) bool {
	cur := l.head
	exists := false

	for cur != nil {
		if cur.value == val {
			exists = true
		}
		cur = cur.next
	}

	return exists
}

func (l *SinglyLinkedList) InsertAtHead(val int) *SinglyLinkedList {
	n := &Node{value: val}

	if l.head != nil {
		n.next = l.head
	}
	l.head = n
	l.Length++
	return l
}

func (l *SinglyLinkedList) InsertAtTail(val int) *SinglyLinkedList {
	n := &Node{value: val, next: nil}
	if l.head == nil {
		l.head = n
		return l
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = n
	l.Length++
	return l
}

func (l *SinglyLinkedList) DeleteAtHead() bool {
	if l.head == nil {
		return false
	}

	l.head = l.head.next
	l.Length--
	return true
}

func (l *SinglyLinkedList) DeleteByVal(val int) bool {
	deleted := false

	if l.head == nil {
		return false
	}

	var prev *Node
	cur := l.head

	if cur.value == val {
		l.head = l.head.next
		deleted = true
		l.Length--
		return deleted
	}

	prev = cur
	cur = cur.next

	for cur != nil {
		if cur.value == val {
			prev.next = cur.next
			deleted = true
			l.Length--
			break
		}

		prev = cur
		cur = cur.next
	}

	return deleted
}

func (l *SinglyLinkedList) Print() {
	cur := l.head

	if cur == nil {
		fmt.Println("List empty")
		return
	}

	for cur != nil {
		fmt.Printf("%d -> ", cur.value)
		cur = cur.next
	}
	fmt.Println("nil")
}
