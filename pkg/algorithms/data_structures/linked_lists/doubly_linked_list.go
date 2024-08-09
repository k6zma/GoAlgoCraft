package linked_lists

import "fmt"

// NodeDoublyLinked represents a node in a doubly linked list.
//
// Fields:
//   - value: the value stored in the node;
//   - next: a pointer to the next node in the list;
//   - prev: a pointer to the previous node in the list.
type NodeDoublyLinked[T any] struct {
	value T
	next  *NodeDoublyLinked[T]
	prev  *NodeDoublyLinked[T]
}

// DoublyLinkedList represents a doubly linked list.
//
// Fields:
//   - head: a pointer to the first node in the list;
//   - tail: a pointer to the last node in the list;
//   - len: the number of elements in the list;
//   - equals: a function to compare equality of two elements.
type DoublyLinkedList[T any] struct {
	head   *NodeDoublyLinked[T]
	tail   *NodeDoublyLinked[T]
	len    int
	equals func(a, b T) bool
}

// NewDoublyLinkedList creates a new doubly linked list with optional initial values.
//
// Parameters:
//   - equalsFunc: a function to compare equality of two elements;
//   - values: optional initial values to populate the list.
//
// Returns a pointer to the new DoublyLinkedList.
func NewDoublyLinkedList[T any](equalsFunc func(a, b T) bool) *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{equals: equalsFunc}

	return list
}

// Len returns the number of elements in the list.
func (dll *DoublyLinkedList[T]) Len() int {
	return dll.len
}

// InsertAtBeginning inserts a new element at the beginning of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (dll *DoublyLinkedList[T]) InsertAtBeginning(elem T) {
	newNode := &NodeDoublyLinked[T]{value: elem}

	newNode.next = dll.head
	if dll.head != nil {
		dll.head.prev = newNode
	} else {
		dll.tail = newNode
	}

	dll.head = newNode

	dll.len++
}

// InsertAtEnd inserts a new element at the end of the list.
//
// Parameters:
//   - elem: the element to be inserted.
func (dll *DoublyLinkedList[T]) InsertAtEnd(elem T) {
	newNode := &NodeDoublyLinked[T]{value: elem}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}

	dll.len++
}

// InsertAtPos inserts a new element at the specified position in the list.
//
// Parameters:
//   - elem: the element to be inserted;
//   - pos: the position to insert the element at.
//
// Returns an error if the position is invalid.
func (dll *DoublyLinkedList[T]) InsertAtPos(elem T, pos int) error {
	if pos < 0 || pos > dll.len {
		return ErrInvalidPos
	}

	newNode := &NodeDoublyLinked[T]{value: elem}

	if pos == 0 {
		newNode.next = dll.head
		if dll.head != nil {
			dll.head.prev = newNode
		}
		dll.head = newNode
		if dll.len == 0 {
			dll.tail = newNode
		}
	} else if pos == dll.len {
		newNode.prev = dll.tail
		if dll.tail != nil {
			dll.tail.next = newNode
		}
		dll.tail = newNode
	} else {
		current := dll.head
		for i := 0; i < pos; i++ {
			if current == nil {
				return ErrPosOutOfRange
			}
			current = current.next
		}

		newNode.next = current
		newNode.prev = current.prev
		if current.prev != nil {
			current.prev.next = newNode
		}
		current.prev = newNode
	}

	dll.len++

	return nil
}

// RemoveFirstNode removes the first element from the list.
//
// Returns an error if the list is empty.
func (dll *DoublyLinkedList[T]) RemoveFirstNode() error {
	if dll.head == nil {
		return ErrListEmpty
	}

	if dll.head.next != nil {
		dll.head = dll.head.next
		dll.head.prev = nil
	} else {
		dll.head = nil
		dll.tail = nil
	}

	dll.len--

	return nil
}

// RemoveLastNode removes the last element from the list.
//
// Returns an error if the list is empty.
func (dll *DoublyLinkedList[T]) RemoveLastNode() error {
	if dll.tail == nil {
		return ErrListEmpty
	}

	if dll.tail.prev != nil {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	} else {
		dll.head = nil
		dll.tail = nil
	}

	dll.len--

	return nil
}

// RemoveNodeAtPosition removes the element at the specified position from the list.
//
// Parameters:
//   - position: the position of the element to be removed.
//
// Returns an error if the position is invalid or the list is empty.
func (dll *DoublyLinkedList[T]) RemoveNodeAtPosition(position int) error {
	if position < 0 || position >= dll.len {
		return ErrInvalidPos
	}

	if position == 0 {
		return dll.RemoveFirstNode()
	} else if position == dll.len-1 {
		return dll.RemoveLastNode()
	}

	current := dll.head
	for i := 0; i < position; i++ {
		current = current.next
	}

	current.prev.next = current.next
	current.next.prev = current.prev

	dll.len--

	return nil
}

// FindNode searches for the first occurrence of the specified value in the list.
//
// Parameters:
//   - value: the value to search for.
//
// Returns the position of the element and an error if the value is not found.
func (dll *DoublyLinkedList[T]) FindNode(value T) (int, error) {
	current := dll.head
	index := 0

	for current != nil {
		if dll.equals(current.value, value) {
			return index, nil
		}
		current = current.next
		index++
	}

	return -1, ErrValueNotFound
}

// Reverse reverses the linked list.
func (dll *DoublyLinkedList[T]) Reverse() {
	var prev *NodeDoublyLinked[T]
	current := dll.head
	dll.tail = current

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}

	dll.head = prev
}

// PrintList prints the elements of the list to the standard output.
func (dll *DoublyLinkedList[T]) PrintList() {
	current := dll.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}

	fmt.Println()
}

// LinkedListToSlice converts the linked list to a slice.
func (dll *DoublyLinkedList[T]) LinkedListToSlice() ([]T, error) {
	var result []T
	current := dll.head

	if current == nil {
		return result, ErrListEmpty
	}

	for current != nil {
		result = append(result, current.value)
		current = current.next
	}

	return result, nil
}
