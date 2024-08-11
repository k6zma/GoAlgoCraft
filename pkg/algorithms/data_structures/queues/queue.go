package queues

import (
	"fmt"
	"github.com/k6zma/GoAlgoCraft/pkg/algorithms/data_structures/linked_lists"
)

// Queue represents a queue based on a singly linked list.
//
// Fields:
//   - Head: a pointer to the first element in the queue;
//   - Tail: a pointer to the last element in the queue;
//   - LenOfQueue: the number of elements in the queue;
//   - Equals: a function to compare two elements for equality.
type Queue[T any] struct {
	Head       *linked_lists.NodeSinglyLinked[T]
	Tail       *linked_lists.NodeSinglyLinked[T]
	LenOfQueue int
	Equals     func(a, b T) bool
}

// NewQueue creates a new queue with a specified equality function.
//
// Parameters:
//   - equalsFunc: a function to compare two elements for equality.
//
// Returns a pointer to the new Queue.
func NewQueue[T any](equalsFunc func(a, b T) bool) *Queue[T] {
	queue := &Queue[T]{
		Equals: equalsFunc,
	}

	return queue
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return q.LenOfQueue
}

// HeadOfQueue returns a pointer to the first element in the queue.
func (q *Queue[T]) HeadOfQueue() *linked_lists.NodeSinglyLinked[T] {
	return q.Head
}

// TailOfQueue returns a pointer to the last element in the queue.
func (q *Queue[T]) TailOfQueue() *linked_lists.NodeSinglyLinked[T] {
	return q.Head
}

// Push adds a new element to the end of the queue.
//
// Parameters:
//   - elem: the element to be added to the queue.
func (q *Queue[T]) Push(elem T) {
	newNode := &linked_lists.NodeSinglyLinked[T]{
		Value: elem,
	}

	if q.LenOfQueue == 0 {
		q.Head = newNode
		q.Tail = newNode

		return
	}

	q.Tail.Next = newNode
	q.Tail = newNode

	q.LenOfQueue++
}

// Pop removes the element from the front of the queue and returns its value.
//
// Returns the value of the first element and an error if the queue is empty.
func (q *Queue[T]) Pop() (T, error) {
	var result T
	if q.LenOfQueue == 0 {
		return result, ErrQueueEmpty
	}

	tempNode := q.Head
	result = tempNode.Value
	q.Head = q.Head.Next

	q.LenOfQueue--

	return result, nil
}

// RemoveElemAtPos removes the element at the specified position in the queue.
//
// Parameters:
//   - pos: the position of the element to be removed.
//
// Returns an error if the position is invalid.
func (q *Queue[T]) RemoveElemAtPos(pos int) error {
	if pos < 0 || pos >= q.LenOfQueue {
		return ErrInvalidPos
	}

	if pos == 0 {
		q.Head = q.Head.Next

		if q.LenOfQueue == 1 {
			q.Tail = nil
		}

		q.LenOfQueue--

		return nil
	}

	prevNode := q.Head
	for i := 0; i < pos-1; i++ {
		prevNode = prevNode.Next
	}

	nodeToRemove := prevNode.Next
	prevNode.Next = nodeToRemove.Next

	if nodeToRemove == q.Tail {
		q.Tail = prevNode
	}

	q.LenOfQueue--

	return nil
}

// FindElem searches for an element in the queue and returns its position.
//
// Parameters:
//   - elem: the element to search for.
//
// Returns the position of the element and an error if the element is not found.
func (q *Queue[T]) FindElem(elem T) (int, error) {
	index := 0
	current := q.Head

	for current != nil {
		if q.Equals(current.Value, elem) {
			return index, nil
		}

		current = current.Next
		index++
	}

	return -1, ErrElemNotFound
}

// Reverse reverses the queue.
func (q *Queue[T]) Reverse() {
	if q.Head == nil || q.Head.Next == nil {
		return
	}

	q.Head, q.Tail = reverseQueue(q.Head, nil)
}

// reverseQueue is a helper function to recursively reverse the queue.
//
// Parameters:
//   - current: the current node being processed;
//   - prev: the previous node in the reversed queue.
//
// Returns the new head and tail of the reversed queue.
func reverseQueue[T any](current, prev *linked_lists.NodeSinglyLinked[T]) (*linked_lists.NodeSinglyLinked[T], *linked_lists.NodeSinglyLinked[T]) {
	if current == nil {
		return prev, prev
	}

	next := current.Next
	current.Next = prev

	if next == nil {
		return current, current
	}

	head, tail := reverseQueue(next, current)
	return head, tail
}

// PrintQueue prints the elements of the queue to the standard output.
func (q *Queue[T]) PrintQueue() {
	current := q.Head

	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}

	fmt.Println()
}

// QueueToSlice converts the queue to a slice.
//
// Returns a slice of the queue elements and an error if the queue is empty.
func (q *Queue[T]) QueueToSlice() ([]T, error) {
	var result []T
	current := q.Head

	if current == nil {
		return result, ErrQueueEmpty
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result, nil
}
