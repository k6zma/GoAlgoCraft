package heap

// NewMinHeap initializes the min-heap.
//
// 1. Initialize the heap with the provided data slice.
//
//  2. Define the less function to maintain the min-heap property.
//     Example: func less(a, b int) bool { return a > b };
//
//  3. Define the equals function to check for equality.
//     Example: func equals(a, b int) bool { return a == b };
//
// 4. Build the heap to ensure the min-heap property is maintained.
func NewMinHeap[T any](data []T, less func(a, b T) bool, equals func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		Data:   data,
		Less:   less,
		Equals: equals,
	}
	h.BuildHeap()

	return h
}
