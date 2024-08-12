package queues

import (
	"fmt"
	"github.com/k6zma/GoAlgoCraft/pkg/algorithms/data_structures/linked_lists"
)

// Deque represents a double-ended queue based on a doubly linked list.
//
// Fields:
//   - Head: a pointer to the first element in the deque;
//   - Tail: a pointer to the last element in the deque;
//   - LenOfDeque: the number of elements in the deque;
//   - Equals: a function to compare two elements for equality.
type Deque[T any] struct {
	Head       *linked_lists.NodeDoublyLinked[T]
	Tail       *linked_lists.NodeDoublyLinked[T]
	LenOfDeque int
	Equals     func(a, b T) bool
}

// NewDeque creates a new deque with a specified equality function.
//
// Parameters:
//   - equalsFunc: a function to compare two elements for equality.
//
// Returns a pointer to the new Deque.
func NewDeque[T any](equalsFunc func(a, b T) bool) *Deque[T] {
	deque := &Deque[T]{
		Equals: equalsFunc,
	}

	return deque
}

// Len returns the number of elements in the deque.
func (d *Deque[T]) Len() int {
	return d.LenOfDeque
}

// HeadOfDeque returns a pointer to the first element in the deque.
func (d *Deque[T]) HeadOfDeque() *linked_lists.NodeDoublyLinked[T] {
	return d.Head
}

// TailOfDeque returns a pointer to the last element in the deque.
func (d *Deque[T]) TailOfDeque() *linked_lists.NodeDoublyLinked[T] {
	return d.Tail
}

// PushAtBegin adds a new element to the front of the deque.
//
// Parameters:
//   - elem: the element to be added to the front of the deque.
func (d *Deque[T]) PushAtBegin(elem T) {
	newNode := &linked_lists.NodeDoublyLinked[T]{
		Value: elem,
	}

	if d.LenOfDeque == 0 {
		d.Head = newNode
		d.Tail = newNode
	} else {
		newNode.Next = d.Head
		d.Head.Prev = newNode
		d.Head = newNode
	}

	d.LenOfDeque++
}

// PushAtEnd adds a new element to the end of the deque.
//
// Parameters:
//   - elem: the element to be added to the end of the deque.
func (d *Deque[T]) PushAtEnd(elem T) {
	newNode := &linked_lists.NodeDoublyLinked[T]{
		Value: elem,
	}

	if d.LenOfDeque == 0 {
		d.Head = newNode
		d.Tail = newNode
	} else {
		newNode.Prev = d.Tail
		d.Tail.Next = newNode
		d.Tail = newNode
	}

	d.LenOfDeque++
}

// PopBegin removes the element from the front of the deque and returns its value.
//
// Returns the value of the first element and an error if the deque is empty.
func (d *Deque[T]) PopBegin() (T, error) {
	var result T
	if d.LenOfDeque == 0 {
		return result, ErrDequeEmpty
	}

	result = d.Head.Value

	if d.Head == d.Tail {
		d.Head = nil
		d.Tail = nil
	} else {
		d.Head = d.Head.Next
		d.Head.Prev = nil
	}

	d.LenOfDeque--

	return result, nil
}

// PopEnd removes the element from the end of the deque and returns its value.
//
// Returns the value of the last element and an error if the deque is empty.
func (d *Deque[T]) PopEnd() (T, error) {
	var result T
	if d.LenOfDeque == 0 {
		return result, ErrDequeEmpty
	}

	result = d.Tail.Value
	if d.Head == d.Tail {
		d.Head = nil
		d.Tail = nil
	} else {
		d.Tail = d.Tail.Prev
		d.Tail.Next = nil
	}

	d.LenOfDeque--

	return result, nil
}

// RemoveElemAtPos removes the element at the specified position in the deque.
//
// Parameters:
//   - pos: the position of the element to be removed.
//
// Returns an error if the position is invalid.
func (d *Deque[T]) RemoveElemAtPos(pos int) error {
	if pos < 0 || pos >= d.LenOfDeque {
		return ErrInvalidPosDeque
	}

	if pos == 0 {
		_, err := d.PopBegin()
		return err
	}

	if pos == d.LenOfDeque-1 {
		_, err := d.PopEnd()
		return err
	}

	current := d.Head
	for i := 0; i < pos; i++ {
		current = current.Next
	}

	prevNode := current.Prev
	nextNode := current.Next

	prevNode.Next = nextNode
	nextNode.Prev = prevNode

	d.LenOfDeque--

	return nil
}

// FindElem searches for an element in the deque and returns its position.
//
// Parameters:
//   - elem: the element to search for.
//
// Returns the position of the element and an error if the element is not found.
func (d *Deque[T]) FindElem(elem T) (int, error) {
	index := 0
	current := d.Head

	for current != nil {
		if d.Equals(current.Value, elem) {
			return index, nil
		}
		current = current.Next
		index++
	}

	return -1, ErrElemNotFoundDeque
}

// GetElemAtPos returns the element at the specified position in the deque.
//
// Parameters:
//   - pos: the zero-based position of the element to retrieve.
//
// Returns the element at the specified position and an error if the position is invalid.
func (d *Deque[T]) GetElemAtPos(pos int) (T, error) {
	var result T

	if pos < 0 || pos >= d.LenOfDeque {
		return result, ErrInvalidPosDeque
	}

	if pos < d.LenOfDeque/2 {
		current := d.Head
		for i := 0; i < pos; i++ {
			current = current.Next
		}
		result = current.Value
	} else {
		current := d.Tail
		for i := d.LenOfDeque - 1; i > pos; i-- {
			current = current.Prev
		}
		result = current.Value
	}

	return result, nil
}

// Reverse reverses the deque.
//
// The method traverses the deque, swapping the next and previous pointers of each node.
// After the traversal, it updates the head and tail pointers to reflect the reversed order.
func (d *Deque[T]) Reverse() {
	if d.Head == nil || d.Head.Next == nil {
		return
	}

	var temp *linked_lists.NodeDoublyLinked[T]
	current := d.Head

	for current != nil {
		temp = current.Prev
		current.Prev = current.Next
		current.Next = temp
		current = current.Prev
	}

	if temp != nil {
		d.Head, d.Tail = d.Tail, d.Head
	}
}

// PrintDeque prints the elements of the deque to the standard output.
func (d *Deque[T]) PrintDeque() {
	current := d.Head

	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}

	fmt.Println()
}

// DequeToSlice converts the deque to a slice.
//
// Returns a slice of the deque elements and an error if the deque is empty.
func (d *Deque[T]) DequeToSlice() ([]T, error) {
	var result []T
	current := d.Head

	if current == nil {
		return result, ErrDequeEmpty
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result, nil
}
