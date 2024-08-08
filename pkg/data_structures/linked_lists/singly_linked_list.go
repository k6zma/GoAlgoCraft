package linked_lists

import "fmt"

// NodeSinglyLinked represents a node in a singly linked list.
//
// Fields:
//   - value: the value stored in the node;
//   - next: a pointer to the next node in the list.
type NodeSinglyLinked[T any] struct {
	value T
	next  *NodeSinglyLinked[T]
}

// SinglyLinkedList represents a singly linked list.
//
// Fields:
//   - head: a pointer to the first node in the list;
//   - len: the number of elements in the list;
//   - equals: a function to compare equality of two elements.
type SinglyLinkedList[T any] struct {
	head   *NodeSinglyLinked[T]
	len    int
	equals func(a, b T) bool
}

// NewSinglyLinkedList creates a new singly linked list with optional initial values.
//
// Parameters:
//   - equalsFunc: a function to compare equality of two elements;
//   - values: optional initial values to populate the list.
//
// Returns a pointer to the new SinglyLinkedList.
func NewSinglyLinkedList[T any](equalsFunc func(a, b T) bool, values ...T) *SinglyLinkedList[T] {
	list := &SinglyLinkedList[T]{equals: equalsFunc}
	for _, value := range values {
		list.InsertAtEnd(value)
		list.len++
	}

	return list
}

// Len returns the number of elements in the list.
func (sll *SinglyLinkedList[T]) Len() int {
	return sll.len
}

// InsertAtBeginning inserts a new element at the beginning of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (sll *SinglyLinkedList[T]) InsertAtBeginning(elem T) {
	newNode := &NodeSinglyLinked[T]{value: elem}

	newNode.next = sll.head
	sll.head = newNode

	sll.len++
}

// InsertAtEnd inserts a new element at the end of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (sll *SinglyLinkedList[T]) InsertAtEnd(elem T) {
	newNode := &NodeSinglyLinked[T]{value: elem}

	if sll.head == nil {
		sll.head = newNode
		return
	}

	current := sll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

	sll.len++
}

// InsertAtPos inserts a new element at the specified position in the list.
//
// Parameters:
//   - elem: the element to be inserted;
//   - pos: the position to insert the element at.
//
// Returns an error if the position is invalid.
func (sll *SinglyLinkedList[T]) InsertAtPos(elem T, pos int) error {
	if pos < 0 || pos > sll.len {
		return ErrInvalidPos
	}

	newNode := &NodeSinglyLinked[T]{value: elem}

	if pos == 0 {
		newNode.next = sll.head
		sll.head = newNode
	} else {
		current := sll.head
		for i := 0; i < pos-1; i++ {
			if current == nil {
				return ErrPosOutOfRange
			}

			current = current.next
		}

		newNode.next = current.next
		current.next = newNode
	}

	sll.len++

	return nil
}

// RemoveFirstNode removes the first element from the list.
//
// Returns an error if the list is empty.
func (sll *SinglyLinkedList[T]) RemoveFirstNode() error {
	if sll.head == nil {
		return ErrListEmpty
	}

	sll.head = sll.head.next
	sll.len--

	return nil
}

// RemoveLastNode removes the last element from the list.
//
// Returns an error if the list is empty.
func (sll *SinglyLinkedList[T]) RemoveLastNode() error {
	if sll.head == nil {
		return ErrListEmpty
	}

	if sll.head.next == nil {
		sll.head = nil
		sll.len--
		return nil
	}

	current := sll.head
	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
	sll.len--

	return nil
}

// RemoveNodeAtPosition removes the element at the specified position from the list.
//
// Parameters:
//   - position: the position of the element to be removed.
//
// Returns an error if the position is invalid or the list is empty.
func (sll *SinglyLinkedList[T]) RemoveNodeAtPosition(position int) error {
	if position < 0 || sll.head == nil {
		return ErrInvalidPos
	}

	if position == 0 {
		sll.head = sll.head.next
		sll.len--

		return nil
	}

	current := sll.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil || current.next == nil {
		return ErrPosOutOfRange
	}

	current.next = current.next.next

	sll.len--

	return nil
}

// FindNode searches for the first occurrence of the specified value in the list.
//
// Parameters:
//   - value: the value to search for.
//
// Returns the position of the element and an error if the value is not found.
func (sll *SinglyLinkedList[T]) FindNode(value T) (int, error) {
	current := sll.head
	index := 0

	for current != nil {
		if sll.equals(current.value, value) {
			return index, nil
		}

		current = current.next
		index++
	}

	return -1, ErrValueNotFound
}

// Reverse reverses the linked list.
func (sll *SinglyLinkedList[T]) Reverse() {
	sll.head = reverse(sll.head, nil)
}

// reverse is a helper function to recursively reverse the linked list.
//
// Parameters:
//   - current: the current node being processed;
//   - prev: the previous node in the reversed list.
//
// Returns the new head of the reversed list.
func reverse[T any](current, prev *NodeSinglyLinked[T]) *NodeSinglyLinked[T] {
	if current == nil {
		return prev
	}

	next := current.next
	current.next = prev

	return reverse(next, current)
}

// PrintList prints the elements of the list to the standard output.
func (sll *SinglyLinkedList[T]) PrintList() {
	current := sll.head

	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}

	fmt.Println()
}
