package linked_lists

import "fmt"

// NodeSinglyLinked represents a node in a singly linked list.
//
// Fields:
//   - value: the value stored in the node;
//   - next: a pointer to the next node in the list.
type NodeSinglyLinked[T any] struct {
	Value T
	Next  *NodeSinglyLinked[T]
}

// SinglyLinkedList represents a singly linked list.
//
// Fields:
//   - head: a pointer to the first node in the list;
//   - len: the number of elements in the list;
//   - equals: a function to compare equality of two elements.
type SinglyLinkedList[T any] struct {
	Head      *NodeSinglyLinked[T]
	LenOfList int
	Equals    func(a, b T) bool
}

// NewSinglyLinkedList creates a new singly linked list with optional initial values.
//
// Parameters:
//   - equalsFunc: a function to compare equality of two elements;
//   - values: optional initial values to populate the list.
//
// Returns a pointer to the new SinglyLinkedList.
func NewSinglyLinkedList[T any](equalsFunc func(a, b T) bool) *SinglyLinkedList[T] {
	list := &SinglyLinkedList[T]{Equals: equalsFunc}

	return list
}

// Len returns the number of elements in the list.
func (sll *SinglyLinkedList[T]) Len() int {
	return sll.LenOfList
}

// HeadOfList returns the pointer to the head of the list.
func (sll *SinglyLinkedList[T]) HeadOfList() *NodeSinglyLinked[T] {
	return sll.Head
}

// NextNode returns the pointer to the next node.
func (node *NodeSinglyLinked[T]) NextNode() *NodeSinglyLinked[T] {
	return node.Next
}

// ValueOfNode returns value of the node.
func (node *NodeSinglyLinked[T]) ValueOfNode() T {
	return node.Value
}

// InsertAtBeginning inserts a new element at the beginning of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (sll *SinglyLinkedList[T]) InsertAtBeginning(elem T) {
	newNode := &NodeSinglyLinked[T]{Value: elem}

	newNode.Next = sll.Head
	sll.Head = newNode

	sll.LenOfList++
}

// InsertAtEnd inserts a new element at the end of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (sll *SinglyLinkedList[T]) InsertAtEnd(elem T) {
	newNode := &NodeSinglyLinked[T]{Value: elem}

	if sll.Head == nil {
		sll.Head = newNode
	} else {
		current := sll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	sll.LenOfList++
}

// InsertAtPos inserts a new element at the specified position in the list.
//
// Parameters:
//   - elem: the element to be inserted;
//   - pos: the position to insert the element at.
//
// Returns an error if the position is invalid.
func (sll *SinglyLinkedList[T]) InsertAtPos(elem T, pos int) error {
	if pos < 0 || pos > sll.LenOfList {
		return ErrInvalidPos
	}

	newNode := &NodeSinglyLinked[T]{Value: elem}

	if pos == 0 {
		newNode.Next = sll.Head
		sll.Head = newNode
	} else {
		current := sll.Head
		for i := 0; i < pos-1; i++ {
			if current == nil {
				return ErrPosOutOfRange
			}

			current = current.Next
		}

		newNode.Next = current.Next
		current.Next = newNode
	}

	sll.LenOfList++

	return nil
}

// RemoveFirstNode removes the first element from the list.
//
// Returns an error if the list is empty.
func (sll *SinglyLinkedList[T]) RemoveFirstNode() error {
	if sll.Head == nil {
		return ErrListEmpty
	}

	sll.Head = sll.Head.Next
	sll.LenOfList--

	return nil
}

// RemoveLastNode removes the last element from the list.
//
// Returns an error if the list is empty.
func (sll *SinglyLinkedList[T]) RemoveLastNode() error {
	if sll.Head == nil {
		return ErrListEmpty
	}

	if sll.Head.Next == nil {
		sll.Head = nil
		sll.LenOfList--
		return nil
	}

	current := sll.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	sll.LenOfList--

	return nil
}

// RemoveNodeAtPosition removes the element at the specified position from the list.
//
// Parameters:
//   - position: the position of the element to be removed.
//
// Returns an error if the position is invalid or the list is empty.
func (sll *SinglyLinkedList[T]) RemoveNodeAtPosition(position int) error {
	if position < 0 || sll.Head == nil {
		return ErrInvalidPos
	}

	if position == 0 {
		sll.Head = sll.Head.Next
		sll.LenOfList--

		return nil
	}

	current := sll.Head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil || current.Next == nil {
		return ErrPosOutOfRange
	}

	current.Next = current.Next.Next

	sll.LenOfList--

	return nil
}

// FindNode searches for the first occurrence of the specified value in the list.
//
// Parameters:
//   - value: the value to search for.
//
// Returns the position of the element and an error if the value is not found.
func (sll *SinglyLinkedList[T]) FindNode(value T) (int, error) {
	current := sll.Head
	index := 0

	for current != nil {
		if sll.Equals(current.Value, value) {
			return index, nil
		}

		current = current.Next
		index++
	}

	return -1, ErrValueNotFound
}

// Reverse reverses the linked list.
func (sll *SinglyLinkedList[T]) Reverse() {
	sll.Head = reverse(sll.Head, nil)
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

	next := current.Next
	current.Next = prev

	return reverse(next, current)
}

// PrintList prints the elements of the list to the standard output.
func (sll *SinglyLinkedList[T]) PrintList() {
	current := sll.Head

	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}

	fmt.Println()
}

// LinkedListToSlice converts the linked list to a slice.
func (sll *SinglyLinkedList[T]) LinkedListToSlice() ([]T, error) {
	var result []T
	current := sll.Head

	if current == nil {
		return result, ErrListEmpty
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result, nil
}
