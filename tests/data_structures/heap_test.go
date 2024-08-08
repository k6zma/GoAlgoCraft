package data_structures_test

import (
	"reflect"
	"testing"

	"github.com/k6zma/GoAlgoCraft/pkg/data_structures/heap"
	"github.com/stretchr/testify/assert"
)

type testLenHeap struct {
	testName string
	data     interface{}
	expected int
}

type testBuildingHeap struct {
	testName          string
	data              interface{}
	expectedData      interface{}
	expectedBuildFlag bool
}

type testPushingHeap struct {
	testName       string
	pushingElement interface{}
	originalData   interface{}
	expectedData   interface{}
}

type testPoppingHeap struct {
	testName string
	data     interface{}
	expected interface{}
}

// lessSlice compares two slices and returns true if the first is "less" than the second.
//
// It compares the slices' lengths first; if they differ, it returns true if the first slice is shorter.
//
// If lengths are equal, it compares corresponding elements by type (int, int32, int64, float32, float64, string) and returns true if any element in the first slice is less than the corresponding element in the second slice.
//
// Panics on unsupported types.
func lessSlice(a, b interface{}) bool {
	sliceA := reflect.ValueOf(a)
	sliceB := reflect.ValueOf(b)

	if sliceA.Len() != sliceB.Len() {
		return sliceA.Len() < sliceB.Len()
	}

	for i := 0; i < sliceA.Len(); i++ {
		valA := sliceA.Index(i).Interface()
		valB := sliceB.Index(i).Interface()

		switch vA := valA.(type) {
		case int:
			if vB, ok := valB.(int); ok {
				if vA < vB {
					return true
				} else if vA > vB {
					return false
				}
			}
		case int32:
			if vB, ok := valB.(int32); ok {
				if vA < vB {
					return true
				} else if vA > vB {
					return false
				}
			}
		case int64:
			if vB, ok := valB.(int64); ok {
				if vA < vB {
					return true
				} else if vA > vB {
					return false
				}
			}
		case float32:
			if vB, ok := valB.(float32); ok {
				if vA < vB {
					return true
				} else if vA > vB {
					return false
				}
			}
		case float64:
			if vB, ok := valB.(float64); ok {
				if vA < vB {
					return true
				} else if vA > vB {
					return false
				}
			}
		case string:
			if vB, ok := valB.(string); ok {
				if vA < vB {
					return true
				} else if vA > vB {
					return false
				}
			}
		default:
			panic("unsupported type")
		}
	}

	return false
}

// equalsSlice compares two slices and returns true if they are identical.
//
// It first checks if the slices have the same length. If not, it returns false. If they do,
//
// it compares corresponding elements by type (int, int32, int64, float32, float64, string) for equality.
//
// Returns false if any pair of elements differs or if the types are unsupported.
//
// Panics on unsupported types.
func equalsSlice(a, b interface{}) bool {
	sliceA := reflect.ValueOf(a)
	sliceB := reflect.ValueOf(b)

	if sliceA.Len() != sliceB.Len() {
		return false
	}

	for i := 0; i < sliceA.Len(); i++ {
		valA := sliceA.Index(i).Interface()
		valB := sliceB.Index(i).Interface()

		switch vA := valA.(type) {
		case int:
			if vB, ok := valB.(int); !ok || vA != vB {
				return false
			}
		case int32:
			if vB, ok := valB.(int32); !ok || vA != vB {
				return false
			}
		case int64:
			if vB, ok := valB.(int64); !ok || vA != vB {
				return false
			}
		case float32:
			if vB, ok := valB.(float32); !ok || vA != vB {
				return false
			}
		case float64:
			if vB, ok := valB.(float64); !ok || vA != vB {
				return false
			}
		case string:
			if vB, ok := valB.(string); !ok || vA != vB {
				return false
			}
		default:
			panic("unsupported type")
		}
	}

	return true
}

func TestHeapLen(t *testing.T) {
	tests := []testLenHeap{
		{
			testName: "Test with no elements of type int",
			data:     []int{},
			expected: 0,
		},
		{
			testName: "Test with 11 elements of type int",
			data:     []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: 11,
		},
		{
			testName: "Test with 50 elements of type int",
			data:     []int{7, 90, 35, 80, 87, 23, 1, 13, 29, 57, 1, 14, 23, 22, 76, 82, 50, 23, 37, 34, 65, 60, 3, 52, 24, 59, 27, 14, 59, 43, 17, 66, 32, 51, 45, 19, 74, 46, 31, 3, 2, 31, 16, 25, 53, 92, 6, 2, 14, 43},
			expected: 50,
		},
		{
			testName: "Test with no elements of type int32",
			data:     []int32{},
			expected: 0,
		},
		{
			testName: "Test with 11 elements of type int32",
			data:     []int32{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: 11,
		},
		{
			testName: "Test with 50 elements of type int32",
			data:     []int32{7, 90, 35, 80, 87, 23, 1, 13, 29, 57, 1, 14, 23, 22, 76, 82, 50, 23, 37, 34, 65, 60, 3, 52, 24, 59, 27, 14, 59, 43, 17, 66, 32, 51, 45, 19, 74, 46, 31, 3, 2, 31, 16, 25, 53, 92, 6, 2, 14, 43},
			expected: 50,
		},
		{
			testName: "Test with no elements of type int64",
			data:     []int64{},
			expected: 0,
		},
		{
			testName: "Test with 11 elements of type int64",
			data:     []int64{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: 11,
		},
		{
			testName: "Test with 50 elements of type int64",
			data:     []int64{7, 90, 35, 80, 87, 23, 1, 13, 29, 57, 1, 14, 23, 22, 76, 82, 50, 23, 37, 34, 65, 60, 3, 52, 24, 59, 27, 14, 59, 43, 17, 66, 32, 51, 45, 19, 74, 46, 31, 3, 2, 31, 16, 25, 53, 92, 6, 2, 14, 43},
			expected: 50,
		},

		{
			testName: "Test with no elements of type float32",
			data:     []float32{},
			expected: 0,
		},
		{
			testName: "Test with 11 elements of type float32",
			data:     []float32{3.14, 1.59, 4.26, 1.38, 5.97, 9.32, 2.71, 6.18, 5.55, 3.14, 5.89},
			expected: 11,
		},
		{
			testName: "Test with 50 elements of type float32",
			data:     []float32{7.01, 90.11, 35.21, 80.13, 87.14, 23.15, 1.16, 13.17, 29.18, 57.19, 1.21, 14.22, 23.23, 22.24, 76.25, 82.26, 50.27, 23.28, 37.29, 34.31, 65.32, 60.33, 3.34, 52.35, 24.36, 59.37, 27.38, 14.39, 59.41, 43.42, 17.43, 66.44, 32.45, 51.46, 45.47, 19.48, 74.49, 46.51, 31.52, 3.53, 2.54, 31.55, 16.56, 25.57, 53.58, 92.59, 6.61, 2.62, 14.63, 43.64},
			expected: 50,
		},
		{
			testName: "Test with no elements of type float64",
			data:     []float64{},
			expected: 0,
		},
		{
			testName: "Test with 11 elements of type float64",
			data:     []float64{3.14, 1.59, 4.26, 1.38, 5.97, 9.32, 2.71, 6.18, 5.55, 3.14, 5.89},
			expected: 11,
		},
		{
			testName: "Test with 50 elements of type float64",
			data:     []float64{7.01, 90.11, 35.21, 80.13, 87.14, 23.15, 1.16, 13.17, 29.18, 57.19, 1.21, 14.22, 23.23, 22.24, 76.25, 82.26, 50.27, 23.28, 37.29, 34.31, 65.32, 60.33, 3.34, 52.35, 24.36, 59.37, 27.38, 14.39, 59.41, 43.42, 17.43, 66.44, 32.45, 51.46, 45.47, 19.48, 74.49, 46.51, 3.53, 2.54, 31.55, 16.56, 25.57, 53.58, 92.59, 6.61, 2.62, 14.63, 43.64, 57.89},
			expected: 50,
		},

		{
			testName: "Test with no elements of type string",
			data:     []string{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type string",
			data: []string{
				"apple",
				"banana",
				"cherry",
				"date",
				"elderberry",
			},
			expected: 5,
		},
		{
			testName: "Test with 10 elements of type string",
			data: []string{
				"fig",
				"grape",
				"honeydew",
				"kiwi",
				"lemon",
				"mango",
				"nectarine",
				"orange",
				"papaya",
				"quince",
			},
			expected: 10,
		},
		{
			testName: "Test with 20 elements of type string",
			data: []string{
				"raspberry",
				"strawberry",
				"tangerine",
				"ugly fruit",
				"valencia orange",
				"watermelon",
				"xigua",
				"yellow watermelon",
				"zucchini",
				"apple pie",
				"blueberry muffin",
				"cantaloupe",
				"dragon fruit",
				"elderflower",
				"fig newton",
				"guava",
				"honeycrisp apple",
				"indian fig",
				"jackfruit",
				"kiwifruit",
			},
			expected: 20,
		},

		{
			testName: "Test with no elements of type []int",
			data:     [][]int{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type []int",
			data: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expected: 5,
		},
		{
			testName: "Test with no elements of type []int32",
			data:     [][]int32{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type []int32",
			data: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expected: 5,
		},
		{
			testName: "Test with no elements of type []int64",
			data:     [][]int64{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type []int64",
			data: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expected: 5,
		},
		{
			testName: "Test with no elements of type []float32",
			data:     [][]float32{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type []float32",
			data: [][]float32{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expected: 5,
		},
		{
			testName: "Test with no elements of type []float64",
			data:     [][]float64{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type []float64",
			data: [][]float64{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expected: 5,
		},
		{
			testName: "Test with no elements of type []string",
			data:     [][]string{},
			expected: 0,
		},
		{
			testName: "Test with 5 elements of type []string",
			data: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"grape", "honeydew", "kiwi"},
				{"lemon", "mango", "nectarine"},
				{"orange", "papaya", "quince"},
			},
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.data.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.Heap[int]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.Heap[int32]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.Heap[int64]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.Heap[float32]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.Heap[float64]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.Heap[string]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.Heap[[]int]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.Heap[[]int32]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.Heap[[]int64]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.Heap[[]float32]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.Heap[[]float64]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.Heap[[]string]{Data: v, Less: less, Equals: equals}
				assert.Equal(t, test.expected, h.Len())
			}
		})
	}
}

func TestMaxHeapBuild(t *testing.T) {
	tests := []testBuildingHeap{
		{
			testName:          "Test with 15 int elements for max-heap",
			data:              []int{76, 44, 2, 22, 16, 46, 17, 23, 62, 55, 56, 98, 80, 16, 11},
			expectedData:      []int{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 int32 elements for max-heap",
			data:              []int32{76, 44, 2, 22, 16, 46, 17, 23, 62, 55, 56, 98, 80, 16, 11},
			expectedData:      []int32{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 int64 elements for max-heap",
			data:              []int64{76, 44, 2, 22, 16, 46, 17, 23, 62, 55, 56, 98, 80, 16, 11},
			expectedData:      []int64{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 float32 elements for max-heap",
			data:              []float32{76.5, 44.2, 2.7, 22.1, 16.9, 46.3, 17.8, 23.4, 62.5, 55.6, 56.2, 98.1, 80.9, 16.4, 11.7},
			expectedData:      []float32{98.1, 62.5, 80.9, 44.2, 56.2, 76.5, 17.8, 23.4, 22.1, 55.6, 16.9, 46.3, 2.7, 16.4, 11.7},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 float64 elements for max-heap",
			data:              []float64{76.5, 44.2, 2.7, 22.1, 16.9, 46.3, 17.8, 23.4, 62.5, 55.6, 56.2, 98.1, 80.9, 16.4, 11.7},
			expectedData:      []float64{98.1, 62.5, 80.9, 44.2, 56.2, 76.5, 17.8, 23.4, 22.1, 55.6, 16.9, 46.3, 2.7, 16.4, 11.7},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 15 string elements for max-heap",
			data: []string{
				"banana",
				"apple",
				"kiwi",
				"date",
				"elderberry",
				"fig",
				"grape",
				"cherry",
				"lemon",
				"mango",
				"nectarine",
				"orange",
				"pear",
				"plum",
				"raspberry",
			},
			expectedData: []string{
				"raspberry",
				"nectarine",
				"plum",
				"lemon",
				"mango",
				"pear",
				"kiwi",
				"cherry",
				"date",
				"apple",
				"elderberry",
				"orange",
				"fig",
				"banana",
				"grape",
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []int for max-heap",
			data: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedData: [][]int{
				{13, 14, 15},
				{10, 11, 12},
				{7, 8, 9},
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []int32 for max-heap",
			data: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedData: [][]int32{
				{13, 14, 15},
				{10, 11, 12},
				{7, 8, 9},
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []int64 for max-heap",
			data: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedData: [][]int64{
				{13, 14, 15},
				{10, 11, 12},
				{7, 8, 9},
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []float32 for max-heap",
			data: [][]float32{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expectedData: [][]float32{
				{13.13, 14.14, 15.15},
				{10.10, 11.11, 12.12},
				{7.7, 8.8, 9.9},
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []float64 for max-heap",
			data: [][]float64{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expectedData: [][]float64{
				{13.13, 14.14, 15.15},
				{10.10, 11.11, 12.12},
				{7.7, 8.8, 9.9},
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []string for max-heap",
			data: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"grape", "honeydew", "kiwi"},
				{"lemon", "mango", "nectarine"},
				{"orange", "papaya", "quince"},
			},
			expectedData: [][]string{
				{"orange", "papaya", "quince"},
				{"lemon", "mango", "nectarine"},
				{"grape", "honeydew", "kiwi"},
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
			},
			expectedBuildFlag: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.data.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			}
		})
	}
}

func TestMinHeapBuild(t *testing.T) {
	tests := []testBuildingHeap{
		{
			testName:          "Test with 15 int elements for min-heap",
			data:              []int{76, 44, 2, 22, 16, 46, 17, 23, 62, 55, 56, 98, 80, 16, 11},
			expectedData:      []int{2, 16, 11, 22, 44, 46, 16, 23, 62, 55, 56, 98, 80, 76, 17},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 int32 elements for min-heap",
			data:              []int32{76, 44, 2, 22, 16, 46, 17, 23, 62, 55, 56, 98, 80, 16, 11},
			expectedData:      []int32{2, 16, 11, 22, 44, 46, 16, 23, 62, 55, 56, 98, 80, 76, 17},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 int64 elements for min-heap",
			data:              []int64{76, 44, 2, 22, 16, 46, 17, 23, 62, 55, 56, 98, 80, 16, 11},
			expectedData:      []int64{2, 16, 11, 22, 44, 46, 16, 23, 62, 55, 56, 98, 80, 76, 17},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 float32 elements for min-heap",
			data:              []float32{76.5, 44.2, 2.7, 22.1, 16.9, 46.3, 17.8, 23.4, 62.5, 55.6, 56.2, 98.1, 80.9, 16.4, 11.7},
			expectedData:      []float32{2.7, 16.9, 11.7, 22.1, 44.2, 46.3, 16.4, 23.4, 62.5, 55.6, 56.2, 98.1, 80.9, 76.5, 17.8},
			expectedBuildFlag: true,
		},
		{
			testName:          "Test with 15 float64 elements for min-heap",
			data:              []float64{76.5, 44.2, 2.7, 22.1, 16.9, 46.3, 17.8, 23.4, 62.5, 55.6, 56.2, 98.1, 80.9, 16.4, 11.7},
			expectedData:      []float64{2.7, 16.9, 11.7, 22.1, 44.2, 46.3, 16.4, 23.4, 62.5, 55.6, 56.2, 98.1, 80.9, 76.5, 17.8},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 15 string elements for min-heap",
			data: []string{
				"banana",
				"apple",
				"kiwi",
				"date",
				"elderberry",
				"fig",
				"grape",
				"cherry",
				"lemon",
				"mango",
				"nectarine",
				"orange",
				"pear",
				"plum",
				"raspberry",
			},
			expectedData: []string{
				"apple",
				"banana",
				"fig",
				"cherry",
				"elderberry",
				"kiwi", "grape",
				"date",
				"lemon",
				"mango",
				"nectarine",
				"orange",
				"pear",
				"plum",
				"raspberry",
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []int for min-heap",
			data: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedData: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []int32 for min-heap",
			data: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedData: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []int64 for min-heap",
			data: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedData: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
				{13, 14, 15},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []float32 for min-heap",
			data: [][]float32{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expectedData: [][]float32{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []float64 for min-heap",
			data: [][]float64{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expectedData: [][]float64{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
				{13.13, 14.14, 15.15},
			},
			expectedBuildFlag: true,
		},
		{
			testName: "Test with 5 elements of type []string for min-heap",
			data: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"grape", "honeydew", "kiwi"},
				{"lemon", "mango", "nectarine"},
				{"orange", "papaya", "quince"},
			},
			expectedData: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"grape", "honeydew", "kiwi"},
				{"lemon", "mango", "nectarine"},
				{"orange", "papaya", "quince"},
			},
			expectedBuildFlag: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.data.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				assert.Equal(t, test.expectedBuildFlag, h.IsHeap())
				assert.Equal(t, test.expectedData, h.Data)
			}
		})
	}
}

func TestMaxHeapPushing(t *testing.T) {
	tests := []testPushingHeap{
		{
			testName:       "Test with 15 int elements for max-heap with pushing 54",
			pushingElement: 54,
			originalData:   []int{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedData:   []int{98, 62, 80, 54, 56, 76, 17, 44, 22, 55, 16, 46, 2, 16, 11, 23},
		},
		{
			testName:       "Test with 15 int32 elements for max-heap with pushing 54",
			pushingElement: int32(54),
			originalData:   []int32{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedData:   []int32{98, 62, 80, 54, 56, 76, 17, 44, 22, 55, 16, 46, 2, 16, 11, 23},
		},
		{
			testName:       "Test with 15 int64 elements for max-heap with pushing 54",
			pushingElement: int64(54),
			originalData:   []int64{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedData:   []int64{98, 62, 80, 54, 56, 76, 17, 44, 22, 55, 16, 46, 2, 16, 11, 23},
		},
		{
			testName:       "Test with 15 float32 elements for max-heap with pushing 54.5",
			pushingElement: float32(54.5),
			originalData:   []float32{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expectedData:   []float32{98.0, 62.0, 80.0, 54.5, 56.0, 76.0, 17.0, 44.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0, 23.0},
		},
		{
			testName:       "Test with 15 float64 elements for max-heap with pushing 54.5",
			pushingElement: float64(54.5),
			originalData:   []float64{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expectedData:   []float64{98.0, 62.0, 80.0, 54.5, 56.0, 76.0, 17.0, 44.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0, 23.0},
		},
		{
			testName:       "Test with 5 string elements for max-heap with pushing 'pear'",
			pushingElement: "pear",
			originalData:   []string{"apple", "banana", "cherry", "date", "elderberry"},
			expectedData:   []string{"pear", "date", "elderberry", "apple", "banana", "cherry"},
		},
		{
			testName:       "Test with 3 elements of type []int for max-heap with pushing [1, 2, 3]",
			pushingElement: []int{1, 2, 3},
			originalData: [][]int{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expectedData: [][]int{
				{10, 11, 12},
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			},
		},
		{
			testName:       "Test with 3 elements of type []int32 for max-heap with pushing [1, 2, 3]",
			pushingElement: []int32{1, 2, 3},
			originalData: [][]int32{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expectedData: [][]int32{
				{10, 11, 12},
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			},
		},
		{
			testName:       "Test with 3 elements of type []int64 for max-heap with pushing [1, 2, 3]",
			pushingElement: []int64{1, 2, 3},
			originalData: [][]int64{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expectedData: [][]int64{
				{10, 11, 12},
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			},
		},
		{
			testName:       "Test with 3 elements of type []float32 for max-heap with pushing [1.1, 2.2, 3.3]",
			pushingElement: []float32{1.1, 2.2, 3.3},
			originalData: [][]float32{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expectedData: [][]float32{
				{10.10, 11.11, 12.12},
				{7.7, 8.8, 9.9},
				{4.4, 5.5, 6.6},
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName:       "Test with 3 elements of type []float64 for max-heap with pushing [1.1, 2.2, 3.3]",
			pushingElement: []float64{1.1, 2.2, 3.3},
			originalData: [][]float64{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expectedData: [][]float64{
				{10.10, 11.11, 12.12},
				{7.7, 8.8, 9.9},
				{4.4, 5.5, 6.6},
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName:       "Test with 3 elements of type []string for max-heap with pushing [\"grape\", \"honeydew\"]",
			pushingElement: []string{"grape", "honeydew"},
			originalData: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"kiwi", "lemon", "mango"},
			},
			expectedData: [][]string{
				{"kiwi", "lemon", "mango"},
				{"date", "elderberry", "fig"},
				{"apple", "banana", "cherry"},
				{"grape", "honeydew"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.originalData.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.(int))
				assert.Equal(t, test.expectedData, h.Data)
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.(int32))
				assert.Equal(t, test.expectedData, h.Data)
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.(int64))
				assert.Equal(t, test.expectedData, h.Data)
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.(float32))
				assert.Equal(t, test.expectedData, h.Data)
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.(float64))
				assert.Equal(t, test.expectedData, h.Data)
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.(string))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.([]int))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.([]int32))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.([]int64))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.([]float32))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.([]float64))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				h.Push(test.pushingElement.([]string))
				assert.Equal(t, test.expectedData, h.Data)
			}
		})
	}
}

func TestMinHeapPushing(t *testing.T) {
	tests := []testPushingHeap{
		{
			testName:       "Test with 15 int32 elements for max-heap with pushing 54",
			pushingElement: 54,
			originalData:   []int{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedData:   []int{2, 16, 11, 22, 55, 46, 16, 23, 44, 62, 56, 80, 76, 98, 17, 54},
		},
		{
			testName:       "Test with 15 int32 elements for max-heap with pushing 54",
			pushingElement: int32(54),
			originalData:   []int32{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedData:   []int32{2, 16, 11, 22, 55, 46, 16, 23, 44, 62, 56, 80, 76, 98, 17, 54},
		},
		{
			testName:       "Test with 15 int64 elements for max-heap with pushing 54",
			pushingElement: int64(54),
			originalData:   []int64{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expectedData:   []int64{2, 16, 11, 22, 55, 46, 16, 23, 44, 62, 56, 80, 76, 98, 17, 54},
		},
		{
			testName:       "Test with 15 float32 elements for max-heap with pushing 54.5",
			pushingElement: float32(54.5),
			originalData:   []float32{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expectedData:   []float32{2.0, 16.0, 11.0, 22.0, 55.0, 46.0, 16.0, 23.0, 44.0, 62.0, 56.0, 80.0, 76.0, 98.0, 17.0, 54.5},
		},
		{
			testName:       "Test with 15 float64 elements for max-heap with pushing 54.5",
			pushingElement: float64(54.5),
			originalData:   []float64{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expectedData:   []float64{2.0, 16.0, 11.0, 22.0, 55.0, 46.0, 16.0, 23.0, 44.0, 62.0, 56.0, 80.0, 76.0, 98.0, 17.0, 54.5},
		},
		{
			testName:       "Test with 5 string elements for max-heap with pushing 'pear'",
			pushingElement: "pear",
			originalData:   []string{"apple", "banana", "cherry", "date", "elderberry"},
			expectedData:   []string{"apple", "banana", "cherry", "date", "elderberry", "pear"},
		},
		{
			testName:       "Test with 3 elements of type []int for max-heap with pushing [1, 2, 3]",
			pushingElement: []int{1, 2, 3},
			originalData: [][]int{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expectedData: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{10, 11, 12},
				{7, 8, 9},
			},
		},
		{
			testName:       "Test with 3 elements of type []int32 for max-heap with pushing [1, 2, 3]",
			pushingElement: []int32{1, 2, 3},
			originalData: [][]int32{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expectedData: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{10, 11, 12},
				{7, 8, 9},
			},
		},
		{
			testName:       "Test with 3 elements of type []int64 for max-heap with pushing [1, 2, 3]",
			pushingElement: []int64{1, 2, 3},
			originalData: [][]int64{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expectedData: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{10, 11, 12},
				{7, 8, 9},
			},
		},
		{
			testName:       "Test with 3 elements of type []float32 for max-heap with pushing [1.1, 2.2, 3.3]",
			pushingElement: []float32{1.1, 2.2, 3.3},
			originalData: [][]float32{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expectedData: [][]float32{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{10.10, 11.11, 12.12},
				{7.7, 8.8, 9.9},
			},
		},
		{
			testName:       "Test with 3 elements of type []float64 for max-heap with pushing [1.1, 2.2, 3.3]",
			pushingElement: []float64{1.1, 2.2, 3.3},
			originalData: [][]float64{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expectedData: [][]float64{
				{1.1, 2.2, 3.3},
				{4.4, 5.5, 6.6},
				{10.10, 11.11, 12.12},
				{7.7, 8.8, 9.9},
			},
		},
		{
			testName:       "Test with 3 elements of type []string for max-heap with pushing [\"grape\", \"honeydew\"]",
			pushingElement: []string{"grape", "honeydew"},
			originalData: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"kiwi", "lemon", "mango"},
			},
			expectedData: [][]string{
				{"grape", "honeydew"},
				{"apple", "banana", "cherry"},
				{"kiwi", "lemon", "mango"},
				{"date", "elderberry", "fig"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.originalData.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.(int))
				assert.Equal(t, test.expectedData, h.Data)
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.(int32))
				assert.Equal(t, test.expectedData, h.Data)
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.(int64))
				assert.Equal(t, test.expectedData, h.Data)
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.(float32))
				assert.Equal(t, test.expectedData, h.Data)
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.(float64))
				assert.Equal(t, test.expectedData, h.Data)
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.(string))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.([]int))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.([]int32))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.([]int64))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.([]float32))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.([]float64))
				assert.Equal(t, test.expectedData, h.Data)
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				h.Push(test.pushingElement.([]string))
				assert.Equal(t, test.expectedData, h.Data)
			}
		})
	}
}

func TestMaxHeapPopping(t *testing.T) {
	tests := []testPoppingHeap{
		{
			testName: "Test with 15 int elements for max-heap with popping",
			data:     []int{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expected: 98,
		},
		{
			testName: "Test with 15 int32 elements for max-heap with popping",
			data:     []int32{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expected: int32(98),
		},
		{
			testName: "Test with 15 int64 elements for max-heap with popping",
			data:     []int64{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expected: int64(98),
		},
		{
			testName: "Test with 15 float32 elements for max-heap with popping",
			data:     []float32{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expected: float32(98.0),
		},
		{
			testName: "Test with 15 float64 elements for max-heap with popping",
			data:     []float64{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expected: float64(98),
		},
		{
			testName: "Test with 5 string elements for max-heap with popping",
			data:     []string{"apple", "banana", "cherry", "date", "elderberry"},
			expected: "elderberry",
		},
		{
			testName: "Test with 3 elements of type []int for max-heap with popping",
			data: [][]int{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int{10, 11, 12},
		},
		{
			testName: "Test with 3 elements of type []int32 for max-heap with popping",
			data: [][]int32{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int32{10, 11, 12},
		},
		{
			testName: "Test with 3 elements of type []int64 for max-heap with popping",
			data: [][]int64{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int64{10, 11, 12},
		},
		{
			testName: "Test with 3 elements of type []float32 for max-heap with popping",
			data: [][]float32{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expected: []float32{10.10, 11.11, 12.12},
		},
		{
			testName: "Test with 3 elements of type []float64 for max-heap with popping",
			data: [][]float64{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expected: []float64{10.10, 11.11, 12.12},
		},
		{
			testName: "Test with 3 elements of type []string for max-heap with popping",
			data: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"kiwi", "lemon", "mango"},
			},
			expected: []string{"kiwi", "lemon", "mango"},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.data.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.NewMaxHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			}
		})
	}
}

func TestMinHeapPopping(t *testing.T) {
	tests := []testPoppingHeap{
		{
			testName: "Test with 15 int elements for min-heap with popping",
			data:     []int{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expected: 2,
		},
		{
			testName: "Test with 15 int32 elements for min-heap with with popping",
			data:     []int32{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expected: int32(2),
		},
		{
			testName: "Test with 15 int64 elements for min-heap with with popping",
			data:     []int64{98, 62, 80, 44, 56, 76, 17, 23, 22, 55, 16, 46, 2, 16, 11},
			expected: int64(2),
		},
		{
			testName: "Test with 15 float32 elements for min-heap with popping",
			data:     []float32{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expected: float32(2.0),
		},
		{
			testName: "Test with 15 float64 elements for min-heap with popping",
			data:     []float64{98.0, 62.0, 80.0, 44.0, 56.0, 76.0, 17.0, 23.0, 22.0, 55.0, 16.0, 46.0, 2.0, 16.0, 11.0},
			expected: float64(2.0),
		},
		{
			testName: "Test with 5 string elements for min-heap with popping",
			data:     []string{"apple", "banana", "cherry", "date", "elderberry"},
			expected: "apple",
		},
		{
			testName: "Test with 3 elements of type []int for min-heap with popping",
			data: [][]int{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int{4, 5, 6},
		},
		{
			testName: "Test with 3 elements of type []int32 for min-heap with popping",
			data: [][]int32{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int32{4, 5, 6},
		},
		{
			testName: "Test with 3 elements of type []int64 for min-heap with popping",
			data: [][]int64{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int64{4, 5, 6},
		},
		{
			testName: "Test with 3 elements of type []float32 for min-heap with popping",
			data: [][]float32{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expected: []float32{4.4, 5.5, 6.6},
		},
		{
			testName: "Test with 3 elements of type []float64 for min-heap with popping",
			data: [][]float64{
				{4.4, 5.5, 6.6},
				{7.7, 8.8, 9.9},
				{10.10, 11.11, 12.12},
			},
			expected: []float64{4.4, 5.5, 6.6},
		},
		{
			testName: "Test with 3 elements of type []string for min-heap with popping",
			data: [][]string{
				{"apple", "banana", "cherry"},
				{"date", "elderberry", "fig"},
				{"kiwi", "lemon", "mango"},
			},
			expected: []string{"apple", "banana", "cherry"},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.data.(type) {
			case []int:
				less := func(a, b int) bool { return a < b }
				equals := func(a, b int) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []int32:
				less := func(a, b int32) bool { return a < b }
				equals := func(a, b int32) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []int64:
				less := func(a, b int64) bool { return a < b }
				equals := func(a, b int64) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []float32:
				less := func(a, b float32) bool { return a < b }
				equals := func(a, b float32) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []float64:
				less := func(a, b float64) bool { return a < b }
				equals := func(a, b float64) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case []string:
				less := func(a, b string) bool { return a < b }
				equals := func(a, b string) bool { return a == b }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]int:
				less := func(a, b []int) bool { return lessSlice(a, b) }
				equals := func(a, b []int) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]int32:
				less := func(a, b []int32) bool { return lessSlice(a, b) }
				equals := func(a, b []int32) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]int64:
				less := func(a, b []int64) bool { return lessSlice(a, b) }
				equals := func(a, b []int64) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]float32:
				less := func(a, b []float32) bool { return lessSlice(a, b) }
				equals := func(a, b []float32) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]float64:
				less := func(a, b []float64) bool { return lessSlice(a, b) }
				equals := func(a, b []float64) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			case [][]string:
				less := func(a, b []string) bool { return lessSlice(a, b) }
				equals := func(a, b []string) bool { return equalsSlice(a, b) }
				h := heap.NewMinHeap(v, less, equals)
				poppedElem, err := h.Pop()
				if err != nil {
					panic(err)
				} else {
					assert.Equal(t, test.expected, poppedElem)
				}
			}
		})
	}
}
