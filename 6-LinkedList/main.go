// Write a code to create linked list. Write a function to add  node to list. Write a function to print linked list. Write a function to edit node value at particular index.
package main

import "fmt"

// Node structure
type Node struct {
	value int
	next  *Node
}

// LinkedList structure
type LinkedList struct {
	head *Node
}

// Add a new node at the end
func (l *LinkedList) Add(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

// Print all nodes
func (l *LinkedList) Print() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.value)
		curr = curr.next
	}
	fmt.Println("nil")
}

// Edit node value at given index (0-based index)
func (l *LinkedList) EditAt(index int, newValue int) error {
	curr := l.head
	pos := 0

	for curr != nil {
		if pos == index {
			curr.value = newValue
			return nil
		}
		pos++
		curr = curr.next
	}

	return fmt.Errorf("index %d out of range", index)
}

func main() {
	list := LinkedList{}

	// Add nodes
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)

	// Print list
	fmt.Println("Original List:")
	list.Print()

	// Edit node at index 2
	err := list.EditAt(2, 99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print updated list
	fmt.Println("Updated List:")
	list.Print()
}
