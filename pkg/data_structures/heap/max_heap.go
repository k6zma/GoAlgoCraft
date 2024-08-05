package heap

// NewMaxHeap initializes the max-heap with the provided data slice and sets the less function
//
// 1. Initialize the heap with the provided data slice.
//
//  2. Define the less function to maintain the max-heap property.
//     Example: func less(a, b int) bool { return a < b };
//
//  3. Define the equals function to check for equality.
//     Example: func equals(a, b int) bool { return a == b };
//
// 4. Build the heap to ensure the max-heap property is maintained.
func NewMaxHeap[T any](data []T, less func(a, b T) bool, equals func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		Data:   data,
		Less:   func(a, b T) bool { return less(b, a) },
		Equals: equals,
	}
	h.BuildHeap()

	return h
}
