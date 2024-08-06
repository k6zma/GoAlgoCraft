package heap

import (
	"errors"
	"fmt"
)

var ErrHeapEmpty = errors.New("heap is empty")

// Heap represents a heap data structure.
//
// Fields:
//
//   - Data: a slice containing heap elements;
//
//   - less: a function (comparator) that should return true if the first element is less than the second.
//     Example: func less(a, b int) bool { return a < b };
//
//   - equals: a function (comparator) that determines the equality of elements.
//     Example: func equals(a, b int) bool { return a < b }.
type Heap[T any] struct {
	Data   []T
	Less   func(a, b T) bool
	Equals func(a, b T) bool
}

// Compares two elements in the heap by their indices using the Less function.
func (h *Heap[T]) heapLess(i, j int) bool {
	return h.Less(h.Data[i], h.Data[j])
}

// Checks if two elements in the heap are equal by their indices using the Equals function.
func (h *Heap[T]) heapEquals(i, j int) bool {
	return h.Equals(h.Data[i], h.Data[j])
}

// Swaps two elements in the heap by their indices.
func (h *Heap[T]) heapSwap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

// Len returns the number of elements that are in the heap.
func (h *Heap[T]) Len() int {
	return len(h.Data)
}

// heapSiftDown restores the heap properties if the value of the modified element increases.
//
// 1. If the i-th element is smaller than its children, the subtree is already a heap;
//
// 2. Otherwise, swap the i-th element with the smallest of its children;
//
// 3. Perform heapSiftDown for the swapped child;
//
// 4. Runs in O(log n) time.
func (h *Heap[T]) heapSiftDown(i int) {
	heapSize := h.Len()
	for 2*i+1 < heapSize {
		left := 2*i + 1
		right := 2*i + 2
		j := left

		if right < heapSize && h.heapLess(right, left) {
			j = right
		}
		if h.heapLess(i, j) || h.heapEquals(i, j) {
			break
		}

		h.heapSwap(i, j)
		i = j
	}
}

// heapSiftUp restores the heap properties if the value of the modified element decreases.
//
// 1. If the element is greater than its parent, the heap condition is satisfied for the entire tree;
//
// 2. Otherwise, swap the element with its parent;
//
// 3. Perform heapSiftUp for the parent;
//
// 4. The procedure runs in O(log n) time.
func (h *Heap[T]) heapSiftUp(i int) {
	for h.heapLess(i, (i-1)/2) {
		h.heapSwap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// BuildHeap builds a heap with the minimum/maximum at the root from an unordered array.
//
// 1. Perform heapSiftDown for nodes with at least one child, from (n/d) to 0;
//
// 2. This approach runs in O(n) time.
func (h *Heap[T]) BuildHeap() {
	heapSize := h.Len()
	for i := (heapSize / 2) - 1; i >= 0; i-- {
		h.heapSiftDown(i)
	}
}

// IsHeap checks whether the heap property is maintained throughout the heap.
//
// 1. Iterates through all non-leaf nodes (indices from 0 to heapSize/2 - 1);
//
// 2. For each node, checks if the node is less than its left child or right child;
//
// 3. If any node violates the heap property, returns false;
//
// 4. If all nodes satisfy the heap property, returns true.
func (h *Heap[T]) IsHeap() bool {
	heapSize := h.Len()
	for i := 0; i < heapSize/2; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < heapSize && h.Less(h.Data[left], h.Data[i]) {
			return false
		}
		if right < heapSize && h.Less(h.Data[right], h.Data[i]) {
			return false
		}
	}

	return true
}

// Push adds a new element to the heap.
//
// 1. Appends the element to the end of the heap;
//
// 2. Restores the heap property using heapSiftUp;
//
// 3. Runs in O(log n) time.
func (h *Heap[T]) Push(elem T) {
	h.Data = append(h.Data, elem)
	h.heapSiftUp(h.Len() - 1)
}

// Pop removes and returns the top element of the heap:
//
// 1. Returns an error if the heap is empty.
//
// 2. Replaces the root with the last element.
//
// 3. Restores the heap property using heapSiftDown.
// 4. Runs in O(log n) time.
func (h *Heap[T]) Pop() (T, error) {
	if h.Len() == 0 {
		var zero T
		return zero, ErrHeapEmpty
	}

	root := h.Data[0]

	lastIndex := h.Len() - 1
	h.Data[0] = h.Data[lastIndex]
	h.Data = h.Data[:lastIndex]

	h.heapSiftDown(0)

	return root, nil
}

// PrintHeap prints the elements of the heap to the standard output.
//
// If the heap is empty, it returns an ErrHeapEmpty error.
//
// The elements are printed in their heap order.
func (h *Heap[T]) PrintHeap() (err error) {
	if h.Len() == 0 {
		return ErrHeapEmpty
	}

	fmt.Printf("%v", h.Data)

	return nil
}
