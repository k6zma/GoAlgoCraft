package linked_lists

import "fmt"

// NodeDoublyLinked represents a node in a doubly linked list.
//
// Fields:
//   - value: the value stored in the node;
//   - next: a pointer to the next node in the list;
//   - prev: a pointer to the previous node in the list.
type NodeDoublyLinked[T any] struct {
	Value T
	Next  *NodeDoublyLinked[T]
	Prev  *NodeDoublyLinked[T]
}

// DoublyLinkedList represents a doubly linked list.
//
// Fields:
//   - head: a pointer to the first node in the list;
//   - tail: a pointer to the last node in the list;
//   - len: the number of elements in the list;
//   - equals: a function to compare equality of two elements.
type DoublyLinkedList[T any] struct {
	Head      *NodeDoublyLinked[T]
	Tail      *NodeDoublyLinked[T]
	LenOfList int
	Equals    func(a, b T) bool
}

// NewDoublyLinkedList creates a new doubly linked list with optional initial values.
//
// Parameters:
//   - equalsFunc: a function to compare equality of two elements;
//   - values: optional initial values to populate the list.
//
// Returns a pointer to the new DoublyLinkedList.
func NewDoublyLinkedList[T any](equalsFunc func(a, b T) bool) *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{Equals: equalsFunc}

	return list
}

// Len returns the number of elements in the list.
func (dll *DoublyLinkedList[T]) Len() int {
	return dll.LenOfList
}

// InsertAtBeginning inserts a new element at the beginning of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (dll *DoublyLinkedList[T]) InsertAtBeginning(elem T) {
	newNode := &NodeDoublyLinked[T]{Value: elem}

	newNode.Next = dll.Head
	if dll.Head != nil {
		dll.Head.Prev = newNode
	} else {
		dll.Tail = newNode
	}

	dll.Head = newNode

	dll.LenOfList++
}

// InsertAtEnd inserts a new element at the end of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (dll *DoublyLinkedList[T]) InsertAtEnd(elem T) {
	newNode := &NodeDoublyLinked[T]{Value: elem}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}

	dll.LenOfList++
}

// InsertAtPos inserts a new element at the specified position in the list.
//
// Parameters:
//   - elem: the element to be inserted;
//   - pos: the position to insert the element at.
//
// Returns an error if the position is invalid.
func (dll *DoublyLinkedList[T]) InsertAtPos(elem T, pos int) error {
	if pos < 0 || pos > dll.LenOfList {
		return ErrInvalidPos
	}

	newNode := &NodeDoublyLinked[T]{Value: elem}

	if pos == 0 {
		newNode.Next = dll.Head
		if dll.Head != nil {
			dll.Head.Prev = newNode
		}
		dll.Head = newNode
		if dll.LenOfList == 0 {
			dll.Tail = newNode
		}
	} else if pos == dll.LenOfList {
		newNode.Prev = dll.Tail
		if dll.Tail != nil {
			dll.Tail.Next = newNode
		}
		dll.Tail = newNode
	} else {
		current := dll.Head
		for i := 0; i < pos; i++ {
			if current == nil {
				return ErrPosOutOfRange
			}
			current = current.Next
		}

		newNode.Next = current
		newNode.Prev = current.Prev
		if current.Prev != nil {
			current.Prev.Next = newNode
		}
		current.Prev = newNode
	}

	dll.LenOfList++

	return nil
}

// RemoveFirstNode removes the first element from the list.
//
// Returns an error if the list is empty.
func (dll *DoublyLinkedList[T]) RemoveFirstNode() error {
	if dll.Head == nil {
		return ErrListEmpty
	}

	if dll.Head.Next != nil {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	} else {
		dll.Head = nil
		dll.Tail = nil
	}

	dll.LenOfList--

	return nil
}

// RemoveLastNode removes the last element from the list.
//
// Returns an error if the list is empty.
func (dll *DoublyLinkedList[T]) RemoveLastNode() error {
	if dll.Tail == nil {
		return ErrListEmpty
	}

	if dll.Tail.Prev != nil {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	} else {
		dll.Head = nil
		dll.Tail = nil
	}

	dll.LenOfList--

	return nil
}

// RemoveNodeAtPosition removes the element at the specified position from the list.
//
// Parameters:
//   - position: the position of the element to be removed.
//
// Returns an error if the position is invalid or the list is empty.
func (dll *DoublyLinkedList[T]) RemoveNodeAtPosition(position int) error {
	if position < 0 || position >= dll.LenOfList {
		return ErrInvalidPos
	}

	if position == 0 {
		return dll.RemoveFirstNode()
	} else if position == dll.LenOfList-1 {
		return dll.RemoveLastNode()
	}

	current := dll.Head
	for i := 0; i < position; i++ {
		current = current.Next
	}

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

	dll.LenOfList--

	return nil
}

// FindNode searches for the first occurrence of the specified value in the list.
//
// Parameters:
//   - value: the value to search for.
//
// Returns the position of the element and an error if the value is not found.
func (dll *DoublyLinkedList[T]) FindNode(value T) (int, error) {
	current := dll.Head
	index := 0

	for current != nil {
		if dll.Equals(current.Value, value) {
			return index, nil
		}
		current = current.Next
		index++
	}

	return -1, ErrValueNotFound
}

// Reverse reverses the linked list.
func (dll *DoublyLinkedList[T]) Reverse() {
	var prev *NodeDoublyLinked[T]
	current := dll.Head
	dll.Tail = current

	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}

	dll.Head = prev
}

// PrintList prints the elements of the list to the standard output.
func (dll *DoublyLinkedList[T]) PrintList() {
	current := dll.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}

	fmt.Println()
}

// LinkedListToSlice converts the linked list to a slice.
func (dll *DoublyLinkedList[T]) LinkedListToSlice() ([]T, error) {
	var result []T
	current := dll.Head

	if current == nil {
		return result, ErrListEmpty
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result, nil
}
