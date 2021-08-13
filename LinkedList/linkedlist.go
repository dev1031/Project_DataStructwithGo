package linkedlist

import (
	"errors"
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

// Search returns node position with given value from linked list
func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

// Insert inserts new node at the end of  from linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// InsertAt inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value int) {
	// create a new node
	newNode := Node{}
	newNode.value = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// DeleteAt deletes node at given position from linked list
func (l *LinkedList) DeleteAt(pos int) error {
	// validate the position
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
}

// DeleteVal deletes node having given value from linked list
func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

func SinglyLinkedList() {
	l := LinkedList{}
	fmt.Println("\n************* Insert *************")
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(15)
	fmt.Println("************* Print *************")
	l.Print()

	fmt.Println("\n************* Search *************")
	fmt.Println("Position of 12 value: ", l.Search(12))
	fmt.Println("Position of 14 value: ", l.Search(14))
	fmt.Println("Position of 15 value: ", l.Search(15))
	fmt.Println("Position of 100 value: ", l.Search(100))

	fmt.Println("\n************* GetAt *************")
	fmt.Println("Get At 1st position: ", l.GetAt(0))
	fmt.Println("Get At 3rd position: ", l.GetAt(2))
	fmt.Println("Get At 4th position: ", l.GetAt(3))
	fmt.Println("Get At -5 position (head is returned): ", l.GetAt(-5))

	// fmt.Println("\n************* InsertAt *************")
	// l.InsertAt(0, 12)
	// l.InsertAt(-1, 13)
	// l.InsertAt(1, 14)
	// l.InsertAt(2, 15)
	// fmt.Println("************* Print *************")
	// l.Print()

	fmt.Println("\n************* DeleteAt *************")
	l.DeleteAt(3)
	fmt.Println("************* Print *************")
	l.Print()

	// fmt.Println("\n************* DeleteVal *************")
	// l.DeleteVal(14)
	// fmt.Println("************* Print *************")
	// l.Print()
}
