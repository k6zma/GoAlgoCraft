package queues

import (
	"fmt"

	"github.com/k6zma/GoAlgoCraft/pkg/algorithms/data_structures/heap"
)

// PriorityQueue represents a priority queue data structure.
//
// Fields:
//
//   - Heap: the underlying heap that manages the priority queue elements;
//
//   - Comparator: a function to compare the priorities of two elements;
//
//   - Equals: a function to compare two elements for equality.
type PriorityQueue[T any] struct {
	HeapData   *heap.Heap[T]
	Comparator func(a, b T) bool
	Equals     func(a, b T) bool
}

// NewPriorityQueue creates a new priority queue with the specified comparator and equality functions.
//
// Parameters:
//   - data: an initial slice of elements to populate the priority queue;
//   - comparator: a function to define the order of elements (e.g., for a min-heap or max-heap);
//   - equals: a function to determine if two elements are equal.
//
// Returns:
//   - A pointer to the new PriorityQueue.
func NewPriorityQueue[T any](data []T, comparator func(a, b T) bool, equals func(a, b T) bool) *PriorityQueue[T] {
	h := heap.NewHeap(data, comparator, equals)
	return &PriorityQueue[T]{
		HeapData:   h,
		Comparator: comparator,
		Equals:     equals,
	}
}

// Len returns the number of elements in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return pq.HeapData.Len()
}

// Push adds a new element to the priority queue.
func (pq *PriorityQueue[T]) Push(elem T) {
	pq.HeapData.Push(elem)
}

// Pop removes and returns the element with the highest priority from the priority queue.
//
// Returns the value of the element with the highest priority and an error if the queue is empty.
func (pq *PriorityQueue[T]) Pop() (T, error) {
	return pq.HeapData.Pop()
}

// Peek returns the element with the highest priority without removing it.
//
// Returns the value of the element with the highest priority and an error if the queue is empty.
func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.Len() == 0 {
		var zero T
		return zero, ErrPriorityQueueEmpty
	}
	return pq.HeapData.Data[0], nil
}

// RemoveElemAtPos removes the element at the specified position in the priority queue.
//
// Parameters:
//   - pos: the position of the element to be removed.
//
// Returns an error if the position is invalid.
func (pq *PriorityQueue[T]) RemoveElemAtPos(pos int) error {
	if pos < 0 || pos >= pq.Len() {
		return ErrInvalidPosPriorityQueue
	}

	// Swap the element to be removed with the last element
	pq.HeapData.Data[pos] = pq.HeapData.Data[pq.Len()-1]
	pq.HeapData.Data = pq.HeapData.Data[:pq.Len()-1]

	// Restore heap properties
	pq.HeapData.BuildHeap()

	return nil
}

// FindElem searches for an element in the priority queue and returns its position.
//
// Parameters:
//   - elem: the element to search for.
//
// Returns the position of the element and an error if the element is not found.
func (pq *PriorityQueue[T]) FindElem(elem T) (int, error) {
	for i, value := range pq.HeapData.Data {
		if pq.Equals(value, elem) {
			return i, nil
		}
	}
	return -1, ErrElemNotFoundPriorityQueue
}

// GetElemAtPos returns the element at the specified position in the priority queue.
//
// Parameters:
//   - pos: the zero-based position of the element to retrieve.
//
// Returns the element at the specified position and an error if the position is invalid.
func (pq *PriorityQueue[T]) GetElemAtPos(pos int) (T, error) {
	var result T

	if pos < 0 || pos >= pq.Len() {
		return result, ErrInvalidPosPriorityQueue
	}

	result = pq.HeapData.Data[pos]
	return result, nil
}

// Reverse reverses the order of elements in the priority queue.
//
// This method reverses the slice of elements in the heap.
func (pq *PriorityQueue[T]) Reverse() {
	data := pq.HeapData.Data
	for i, j := 0, pq.Len()-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// PrintQueue prints the elements of the priority queue to the standard output.
func (pq *PriorityQueue[T]) PrintQueue() {
	err := pq.HeapData.PrintHeap()
	if err != nil {
		fmt.Println(err)
	}
}

// PriorityQueueToSlice converts the priority queue to a slice.
//
// Returns a slice of the priority queue elements.
func (pq *PriorityQueue[T]) PriorityQueueToSlice() []T {
	return pq.HeapData.Data
}
