package data_structures_test

import (
	"fmt"
	"testing"

	"github.com/k6zma/GoAlgoCraft/pkg/algorithms/data_structures/linked_lists"
	"github.com/k6zma/GoAlgoCraft/pkg/uitls"
	"github.com/stretchr/testify/assert"
)

type testInsertList struct {
	testName      string
	initialValues interface{}
	insertValue   interface{}
	expectedList  interface{}
}

type testInsertAtPositionList struct {
	testName      string
	initialValues interface{}
	insertValue   interface{}
	position      int
	expectedList  interface{}
}

type testRemoveList struct {
	testName      string
	initialValues interface{}
	expectedList  interface{}
	expectedError error
}

type testRemoveAtPositionList struct {
	testName      string
	initialValues interface{}
	position      int
	expectedList  interface{}
	expectedError error
}

type testFindNodeList struct {
	testName      string
	initialValues interface{}
	findValue     interface{}
	expectedIndex int
	expectedError error
}

type testReverseList struct {
	testName      string
	initialValues interface{}
	expectedList  interface{}
}

type testLenList struct {
	testName      string
	initialValues interface{}
	expectedLen   int
}

func TestSinglyLinkedList_InsertAtBeginning(t *testing.T) {
	tests := []testInsertList{
		{
			testName:      "Insert at begin int elem in empty singly linked list",
			initialValues: []int{},
			insertValue:   1,
			expectedList:  []int{1},
		},
		{
			testName:      "Insert at begin int elem in non-empty singly linked list",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   1,
			expectedList:  []int{1, 3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Insert at begin 5 int elem in non-empty singly linked list",
			initialValues: []int{35, 50, 69, 71, 74},
			insertValue:   []int{89, 95, 100},
			expectedList:  []int{100, 95, 89, 74, 71, 69, 50, 35},
		},

		{
			testName:      "Insert at begin int32 elem in empty singly linked list",
			initialValues: []int32{},
			insertValue:   int32(1),
			expectedList:  []int32{1},
		},
		{
			testName:      "Insert at begin int32 elem in non-empty singly linked list",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int32(1),
			expectedList:  []int32{1, 3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Insert at begin 5 int32 elem in non-empty singly linked list",
			initialValues: []int32{35, 50, 69, 71, 74},
			insertValue:   []int32{89, 95, 100},
			expectedList:  []int32{100, 95, 89, 74, 71, 69, 50, 35},
		},

		{
			testName:      "Insert at begin int64 elem in empty singly linked list",
			initialValues: []int64{},
			insertValue:   int64(1),
			expectedList:  []int64{1},
		},
		{
			testName:      "Insert at begin int64 elem in non-empty singly linked list",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int64(1),
			expectedList:  []int64{1, 3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Insert at begin 5 int64 elem in non-empty singly linked list",
			initialValues: []int64{35, 50, 69, 71, 74},
			insertValue:   []int64{89, 95, 100},
			expectedList:  []int64{100, 95, 89, 74, 71, 69, 50, 35},
		},

		{
			testName:      "Insert at begin float32 elem in empty singly linked list",
			initialValues: []float32{},
			insertValue:   float32(1.0),
			expectedList:  []float32{1.0},
		},
		{
			testName:      "Insert at begin float32 elem in non-empty singly linked list",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			insertValue:   float32(1.0),
			expectedList:  []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
		},
		{
			testName:      "Insert at begin 5 float32 elem in non-empty singly linked list",
			initialValues: []float32{35.7, 50.4, 69.9, 71.2, 74.8},
			insertValue:   []float32{89.1, 95.5, 100.3},
			expectedList:  []float32{100.3, 95.5, 89.1, 74.8, 71.2, 69.9, 50.4, 35.7},
		},

		{
			testName:      "Insert at begin float64 elem in empty singly linked list",
			initialValues: []float64{},
			insertValue:   float64(1.0),
			expectedList:  []float64{1.0},
		},
		{
			testName:      "Insert at begin float64 elem in non-empty singly linked list",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			insertValue:   float64(1.0),
			expectedList:  []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
		},
		{
			testName:      "Insert at begin 5 float64 elem in non-empty singly linked list",
			initialValues: []float64{35.7, 50.4, 69.9, 71.2, 74.8},
			insertValue:   []float64{89.1, 95.5, 100.3},
			expectedList:  []float64{100.3, 95.5, 89.1, 74.8, 71.2, 69.9, 50.4, 35.7},
		},

		{
			testName:      "Insert at begin string elem in empty singly linked list",
			initialValues: []string{},
			insertValue:   "raspberry",
			expectedList:  []string{"raspberry"},
		},
		{
			testName:      "Insert at begin string elem in non-empty singly linked list",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			insertValue:   "apple pie",
			expectedList:  []string{"apple pie", "zucchini", "yellow watermelon", "xigua", "watermelon", "valencia orange", "ugly fruit", "tangerine", "strawberry"},
		},
		{
			testName:      "Insert at begin 3 string elem in non-empty singly linked list",
			initialValues: []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton"},
			insertValue:   []string{"guava", "honeycrisp apple", "indian fig"},
			expectedList:  []string{"indian fig", "honeycrisp apple", "guava", "fig newton", "elderflower", "dragon fruit", "cantaloupe", "blueberry muffin"},
		},

		{
			testName:      "Insert at begin [][]int elem in empty singly linked list",
			initialValues: [][]int{},
			insertValue:   []int{1, 2, 3},
			expectedList: [][]int{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at begin [][]int elem in non-empty singly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: []int{1, 2, 3},
			expectedList: [][]int{
				{1, 2, 3},
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Insert at begin multiple [][]int elems in non-empty singly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Insert at begin [][]int32 elem in empty singly linked list",
			initialValues: [][]int32{},
			insertValue:   []int32{1, 2, 3},
			expectedList: [][]int32{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at begin [][]int32 elem in non-empty singly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: []int32{1, 2, 3},
			expectedList: [][]int32{
				{1, 2, 3},
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Insert at begin multiple [][]int32 elems in non-empty singly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int32{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int32{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Insert at begin [][]int64 elem in empty singly linked list",
			initialValues: [][]int64{},
			insertValue:   []int64{1, 2, 3},
			expectedList: [][]int64{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at begin [][]int64 elem in non-empty singly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: []int64{1, 2, 3},
			expectedList: [][]int64{
				{1, 2, 3},
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Insert at begin multiple [][]int64 elems in non-empty singly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int64{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int64{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Insert at begin [][]float32 elem in empty singly linked list",
			initialValues: [][]float32{},
			insertValue:   []float32{1.1, 2.2, 3.3},
			expectedList: [][]float32{
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at begin [][]float32 elem in non-empty singly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float32{1.1, 2.2, 3.3},
			expectedList: [][]float32{
				{1.1, 2.2, 3.3},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},
		{
			testName: "Insert at begin multiple [][]float32 elems in non-empty singly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: [][]float32{
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float32{
				{70.2, 80.3},
				{50.9, 60.1},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},

		{
			testName:      "Insert at begin [][]float64 elem in empty singly linked list",
			initialValues: [][]float64{},
			insertValue:   []float64{1.1, 2.2, 3.3},
			expectedList: [][]float64{
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at begin [][]float64 elem in non-empty singly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float64{1.1, 2.2, 3.3},
			expectedList: [][]float64{
				{1.1, 2.2, 3.3},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},
		{
			testName: "Insert at begin multiple [][]float64 elems in non-empty singly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: [][]float64{
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float64{
				{70.2, 80.3},
				{50.9, 60.1},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},

		{
			testName:      "Insert at begin [][]string elem in empty singly linked list",
			initialValues: [][]string{},
			insertValue:   []string{"apple", "banana"},
			expectedList: [][]string{
				{"apple", "banana"},
			},
		},
		{
			testName: "Insert at begin [][]string elem in non-empty singly linked list",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			insertValue: []string{"apple", "banana"},
			expectedList: [][]string{
				{"apple", "banana"},
				{"elderberry", "fig"},
				{"cherry", "date"},
			},
		},
		{
			testName: "Insert at begin multiple [][]string elems in non-empty singly linked list",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
			},
			insertValue: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"cherry", "date"},
				{"apple", "banana"},
				{"kiwi", "lemon"},
				{"grape", "honeydew"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case int:
					list.InsertAtBeginning(x)
				case []int:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case int32:
					list.InsertAtBeginning(x)
				case []int32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case int64:
					list.InsertAtBeginning(x)
				case []int64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case float32:
					list.InsertAtBeginning(x)
				case []float32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case float64:
					list.InsertAtBeginning(x)
				case []float64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case string:
					list.InsertAtBeginning(x)
				case []string:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []int:
					list.InsertAtBeginning(x)
				case [][]int:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []int32:
					list.InsertAtBeginning(x)
				case [][]int32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []int64:
					list.InsertAtBeginning(x)
				case [][]int64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []float32:
					list.InsertAtBeginning(x)
				case [][]float32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []float64:
					list.InsertAtBeginning(x)
				case [][]float64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []string:
					list.InsertAtBeginning(x)
				case [][]string:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_InsertAtBeginning(t *testing.T) {
	tests := []testInsertList{
		{
			testName:      "Insert at begin int elem in empty doubly linked list",
			initialValues: []int{},
			insertValue:   1,
			expectedList:  []int{1},
		},
		{
			testName:      "Insert at begin int elem in non-empty doubly linked list",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   1,
			expectedList:  []int{1, 3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Insert at begin multiple int elems in non-empty doubly linked list",
			initialValues: []int{35, 50, 69, 71, 74},
			insertValue:   []int{89, 95, 100},
			expectedList:  []int{100, 95, 89, 74, 71, 69, 50, 35},
		},

		{
			testName:      "Insert at begin int32 elem in empty doubly linked list",
			initialValues: []int32{},
			insertValue:   int32(1),
			expectedList:  []int32{1},
		},
		{
			testName:      "Insert at begin int32 elem in non-empty doubly linked list",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int32(1),
			expectedList:  []int32{1, 3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Insert at begin multiple int32 elems in non-empty doubly linked list",
			initialValues: []int32{35, 50, 69, 71, 74},
			insertValue:   []int32{89, 95, 100},
			expectedList:  []int32{100, 95, 89, 74, 71, 69, 50, 35},
		},

		{
			testName:      "Insert at begin int64 elem in empty doubly linked list",
			initialValues: []int64{},
			insertValue:   int64(1),
			expectedList:  []int64{1},
		},
		{
			testName:      "Insert at begin int64 elem in non-empty doubly linked list",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int64(1),
			expectedList:  []int64{1, 3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Insert at begin multiple int64 elems in non-empty doubly linked list",
			initialValues: []int64{35, 50, 69, 71, 74},
			insertValue:   []int64{89, 95, 100},
			expectedList:  []int64{100, 95, 89, 74, 71, 69, 50, 35},
		},

		{
			testName:      "Insert at begin float32 elem in empty doubly linked list",
			initialValues: []float32{},
			insertValue:   float32(1.5),
			expectedList:  []float32{1.5},
		},
		{
			testName:      "Insert at begin float32 elem in non-empty doubly linked list",
			initialValues: []float32{100.5, 42.1, 72.3, 5.4, 36.2, 66.7, 77.8, 79.9, 51.0, 3.1},
			insertValue:   float32(1.5),
			expectedList:  []float32{1.5, 3.1, 51.0, 79.9, 77.8, 66.7, 36.2, 5.4, 72.3, 42.1, 100.5},
		},
		{
			testName:      "Insert at begin multiple float32 elems in non-empty doubly linked list",
			initialValues: []float32{35.5, 50.6, 69.7, 71.8, 74.9},
			insertValue:   []float32{89.9, 95.1, 100.2},
			expectedList:  []float32{100.2, 95.1, 89.9, 74.9, 71.8, 69.7, 50.6, 35.5},
		},

		{
			testName:      "Insert at begin float64 elem in empty doubly linked list",
			initialValues: []float64{},
			insertValue:   float64(1.5),
			expectedList:  []float64{1.5},
		},
		{
			testName:      "Insert at begin float64 elem in non-empty doubly linked list",
			initialValues: []float64{100.5, 42.1, 72.3, 5.4, 36.2, 66.7, 77.8, 79.9, 51.0, 3.1},
			insertValue:   float64(1.5),
			expectedList:  []float64{1.5, 3.1, 51.0, 79.9, 77.8, 66.7, 36.2, 5.4, 72.3, 42.1, 100.5},
		},
		{
			testName:      "Insert at begin multiple float64 elems in non-empty doubly linked list",
			initialValues: []float64{35.5, 50.6, 69.7, 71.8, 74.9},
			insertValue:   []float64{89.9, 95.1, 100.2},
			expectedList:  []float64{100.2, 95.1, 89.9, 74.9, 71.8, 69.7, 50.6, 35.5},
		},

		{
			testName:      "Insert at begin string elem in empty doubly linked list",
			initialValues: []string{},
			insertValue:   "first",
			expectedList:  []string{"first"},
		},
		{
			testName:      "Insert at begin string elem in non-empty doubly linked list",
			initialValues: []string{"hello", "world", "this", "is", "a", "test"},
			insertValue:   "first",
			expectedList:  []string{"first", "test", "a", "is", "this", "world", "hello"},
		},
		{
			testName:      "Insert at begin multiple string elems in non-empty doubly linked list",
			initialValues: []string{"foo", "bar", "baz"},
			insertValue:   []string{"alpha", "beta", "gamma"},
			expectedList:  []string{"gamma", "beta", "alpha", "baz", "bar", "foo"},
		},

		{
			testName:      "Insert at begin []int elem in empty doubly linked list",
			initialValues: [][]int{},
			insertValue:   []int{1, 2, 3},
			expectedList: [][]int{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at begin []int elem in non-empty doubly linked list",
			initialValues: [][]int{
				{10, 20, 30},
				{40, 50, 60},
			},
			insertValue: []int{1, 2, 3},
			expectedList: [][]int{
				{1, 2, 3},
				{40, 50, 60},
				{10, 20, 30},
			},
		},
		{
			testName: "Insert at begin multiple []int elems in non-empty doubly linked list",
			initialValues: [][]int{
				{10, 20, 30},
				{40, 50, 60},
			},
			insertValue: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedList: [][]int{
				{4, 5, 6},
				{1, 2, 3},
				{40, 50, 60},
				{10, 20, 30},
			},
		},

		{
			testName:      "Insert at begin []string elem in empty singly linked list",
			initialValues: [][]string{},
			insertValue:   []string{"apple", "banana"},
			expectedList: [][]string{
				{"apple", "banana"},
			},
		},
		{
			testName: "Insert at begin []string elem in non-empty singly linked list",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			insertValue: []string{"apple", "banana"},
			expectedList: [][]string{
				{"apple", "banana"},
				{"elderberry", "fig"},
				{"cherry", "date"},
			},
		},
		{
			testName: "Insert at begin multiple []string elems in non-empty singly linked list",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
			},
			insertValue: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"cherry", "date"},
				{"apple", "banana"},
				{"kiwi", "lemon"},
				{"grape", "honeydew"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case int:
					list.InsertAtBeginning(x)
				case []int:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case int32:
					list.InsertAtBeginning(x)
				case []int32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case int64:
					list.InsertAtBeginning(x)
				case []int64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case float32:
					list.InsertAtBeginning(x)
				case []float32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case float64:
					list.InsertAtBeginning(x)
				case []float64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case string:
					list.InsertAtBeginning(x)
				case []string:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []int:
					list.InsertAtBeginning(x)
				case [][]int:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []int32:
					list.InsertAtBeginning(x)
				case [][]int32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []int64:
					list.InsertAtBeginning(x)
				case [][]int64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []float32:
					list.InsertAtBeginning(x)
				case [][]float32:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []float64:
					list.InsertAtBeginning(x)
				case [][]float64:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtBeginning(elem)
				}

				switch x := test.insertValue.(type) {
				case []string:
					list.InsertAtBeginning(x)
				case [][]string:
					for _, elem := range x {
						list.InsertAtBeginning(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_InsertAtEnd(t *testing.T) {
	tests := []testInsertList{
		{
			testName:      "Insert at end int elem in empty singly linked list",
			initialValues: []int{},
			insertValue:   1,
			expectedList:  []int{1},
		},
		{
			testName:      "Insert at end int elem in non-empty singly linked list",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   1,
			expectedList:  []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3, 1},
		},
		{
			testName:      "Insert at end 5 int elem in non-empty singly linked list",
			initialValues: []int{35, 50, 69, 71, 74},
			insertValue:   []int{89, 95, 100},
			expectedList:  []int{35, 50, 69, 71, 74, 89, 95, 100},
		},

		{
			testName:      "Insert at end int32 elem in empty singly linked list",
			initialValues: []int32{},
			insertValue:   int32(1),
			expectedList:  []int32{1},
		},
		{
			testName:      "Insert at end int32 elem in non-empty singly linked list",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int32(1),
			expectedList:  []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3, 1},
		},
		{
			testName:      "Insert at end 5 int32 elem in non-empty singly linked list",
			initialValues: []int32{35, 50, 69, 71, 74},
			insertValue:   []int32{89, 95, 100},
			expectedList:  []int32{35, 50, 69, 71, 74, 89, 95, 100},
		},

		{
			testName:      "Insert at end int64 elem in empty singly linked list",
			initialValues: []int64{},
			insertValue:   int64(1),
			expectedList:  []int64{1},
		},
		{
			testName:      "Insert at end int64 elem in non-empty singly linked list",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int64(1),
			expectedList:  []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3, 1},
		},
		{
			testName:      "Insert at end 5 int64 elem in non-empty singly linked list",
			initialValues: []int64{35, 50, 69, 71, 74},
			insertValue:   []int64{89, 95, 100},
			expectedList:  []int64{35, 50, 69, 71, 74, 89, 95, 100},
		},

		{
			testName:      "Insert at end float32 elem in empty singly linked list",
			initialValues: []float32{},
			insertValue:   float32(1.0),
			expectedList:  []float32{1.0},
		},
		{
			testName:      "Insert at end float32 elem in non-empty singly linked list",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			insertValue:   float32(1.0),
			expectedList:  []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3, 1},
		},
		{
			testName:      "Insert at end 5 float32 elem in non-empty singly linked list",
			initialValues: []float32{35.7, 50.4, 69.9, 71.2, 74.8},
			insertValue:   []float32{89.1, 95.5, 100.3},
			expectedList:  []float32{35.7, 50.4, 69.9, 71.2, 74.8, 89.1, 95.5, 100.3},
		},

		{
			testName:      "Insert at end float64 elem in empty singly linked list",
			initialValues: []float64{},
			insertValue:   float64(1.0),
			expectedList:  []float64{1.0},
		},
		{
			testName:      "Insert at end float64 elem in non-empty singly linked list",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			insertValue:   float64(1.0),
			expectedList:  []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3, 1},
		},
		{
			testName:      "Insert at end 5 float64 elem in non-empty singly linked list",
			initialValues: []float64{35.7, 50.4, 69.9, 71.2, 74.8},
			insertValue:   []float64{89.1, 95.5, 100.3},
			expectedList:  []float64{35.7, 50.4, 69.9, 71.2, 74.8, 89.1, 95.5, 100.3},
		},

		{
			testName:      "Insert at end string elem in empty singly linked list",
			initialValues: []string{},
			insertValue:   "raspberry",
			expectedList:  []string{"raspberry"},
		},
		{
			testName:      "Insert at end string elem in non-empty singly linked list",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			insertValue:   "apple pie",
			expectedList:  []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini", "apple pie"},
		},
		{
			testName:      "Insert at end 3 string elem in non-empty singly linked list",
			initialValues: []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton"},
			insertValue:   []string{"guava", "honeycrisp apple", "indian fig"},
			expectedList:  []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton", "guava", "honeycrisp apple", "indian fig"},
		},

		{
			testName:      "Insert at end []int elem in empty singly linked list",
			initialValues: [][]int{},
			insertValue:   []int{1, 2, 3},
			expectedList: [][]int{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end []int elem in non-empty singly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: []int{1, 2, 3},
			expectedList: [][]int{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end multiple []int elems in non-empty singly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
		},

		{
			testName:      "Insert at end []int32 elem in empty singly linked list",
			initialValues: [][]int32{},
			insertValue:   []int32{1, 2, 3},
			expectedList: [][]int32{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end []int32 elem in non-empty singly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: []int32{1, 2, 3},
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end multiple []int32 elems in non-empty singly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int32{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
		},

		{
			testName:      "Insert at end []int64 elem in empty singly linked list",
			initialValues: [][]int64{},
			insertValue:   []int64{1, 2, 3},
			expectedList: [][]int64{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end []int64 elem in non-empty singly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: []int64{1, 2, 3},
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end multiple []int64 elems in non-empty singly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int64{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
		},

		{
			testName:      "Insert at end []float32 elem in empty singly linked list",
			initialValues: [][]float32{},
			insertValue:   []float32{1.1, 2.2, 3.3},
			expectedList: [][]float32{
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end []float32 elem in non-empty singly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float32{1.1, 2.2, 3.3},
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end multiple []float32 elems in non-empty singly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: [][]float32{
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
		},

		{
			testName:      "Insert at end []float64 elem in empty singly linked list",
			initialValues: [][]float64{},
			insertValue:   []float64{1.1, 2.2, 3.3},
			expectedList: [][]float64{
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end []float64 elem in non-empty singly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float64{1.1, 2.2, 3.3},
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end multiple []float64 elems in non-empty singly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: [][]float64{
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
		},

		{
			testName:      "Insert at end []string elem in empty singly linked list",
			initialValues: [][]string{},
			insertValue:   []string{"apple", "banana"},
			expectedList: [][]string{
				{"apple", "banana"},
			},
		},
		{
			testName: "Insert at end []string elem in non-empty singly linked list",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			insertValue: []string{"apple", "banana"},
			expectedList: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
				{"apple", "banana"},
			},
		},
		{
			testName: "Insert at end multiple []string elems in non-empty singly linked list",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
			},
			insertValue: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case int:
					list.InsertAtEnd(x)
				case []int:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case int32:
					list.InsertAtEnd(x)
				case []int32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case int64:
					list.InsertAtEnd(x)
				case []int64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case float32:
					list.InsertAtEnd(x)
				case []float32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case float64:
					list.InsertAtEnd(x)
				case []float64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case string:
					list.InsertAtEnd(x)
				case []string:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []int:
					list.InsertAtEnd(x)
				case [][]int:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []int32:
					list.InsertAtEnd(x)
				case [][]int32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []int64:
					list.InsertAtEnd(x)
				case [][]int64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []float32:
					list.InsertAtEnd(x)
				case [][]float32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []float64:
					list.InsertAtEnd(x)
				case [][]float64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []string:
					list.InsertAtEnd(x)
				case [][]string:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_InsertAtEnd(t *testing.T) {
	tests := []testInsertList{
		{
			testName:      "Insert at end int elem in empty doubly linked list",
			initialValues: []int{},
			insertValue:   1,
			expectedList:  []int{1},
		},
		{
			testName:      "Insert at end int elem in non-empty doubly linked list",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   1,
			expectedList:  []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3, 1},
		},
		{
			testName:      "Insert at end 5 int elem in non-empty doubly linked list",
			initialValues: []int{35, 50, 69, 71, 74},
			insertValue:   []int{89, 95, 100},
			expectedList:  []int{35, 50, 69, 71, 74, 89, 95, 100},
		},

		{
			testName:      "Insert at end int32 elem in empty doubly linked list",
			initialValues: []int32{},
			insertValue:   int32(1),
			expectedList:  []int32{1},
		},
		{
			testName:      "Insert at end int32 elem in non-empty doubly linked list",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int32(1),
			expectedList:  []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3, 1},
		},
		{
			testName:      "Insert at end 5 int32 elem in non-empty doubly linked list",
			initialValues: []int32{35, 50, 69, 71, 74},
			insertValue:   []int32{89, 95, 100},
			expectedList:  []int32{35, 50, 69, 71, 74, 89, 95, 100},
		},

		{
			testName:      "Insert at end int64 elem in empty doubly linked list",
			initialValues: []int64{},
			insertValue:   int64(1),
			expectedList:  []int64{1},
		},
		{
			testName:      "Insert at end int64 elem in non-empty doubly linked list",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			insertValue:   int64(1),
			expectedList:  []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3, 1},
		},
		{
			testName:      "Insert at end 5 int64 elem in non-empty doubly linked list",
			initialValues: []int64{35, 50, 69, 71, 74},
			insertValue:   []int64{89, 95, 100},
			expectedList:  []int64{35, 50, 69, 71, 74, 89, 95, 100},
		},

		{
			testName:      "Insert at end float32 elem in empty doubly linked list",
			initialValues: []float32{},
			insertValue:   float32(1.0),
			expectedList:  []float32{1.0},
		},
		{
			testName:      "Insert at end float32 elem in non-empty doubly linked list",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			insertValue:   float32(1.0),
			expectedList:  []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3, 1},
		},
		{
			testName:      "Insert at end 5 float32 elem in non-empty doubly linked list",
			initialValues: []float32{35.7, 50.4, 69.9, 71.2, 74.8},
			insertValue:   []float32{89.1, 95.5, 100.3},
			expectedList:  []float32{35.7, 50.4, 69.9, 71.2, 74.8, 89.1, 95.5, 100.3},
		},

		{
			testName:      "Insert at end float64 elem in empty doubly linked list",
			initialValues: []float64{},
			insertValue:   float64(1.0),
			expectedList:  []float64{1.0},
		},
		{
			testName:      "Insert at end float64 elem in non-empty doubly linked list",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			insertValue:   float64(1.0),
			expectedList:  []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3, 1},
		},
		{
			testName:      "Insert at end 5 float64 elem in non-empty doubly linked list",
			initialValues: []float64{35.7, 50.4, 69.9, 71.2, 74.8},
			insertValue:   []float64{89.1, 95.5, 100.3},
			expectedList:  []float64{35.7, 50.4, 69.9, 71.2, 74.8, 89.1, 95.5, 100.3},
		},

		{
			testName:      "Insert at end string elem in empty doubly linked list",
			initialValues: []string{},
			insertValue:   "raspberry",
			expectedList:  []string{"raspberry"},
		},
		{
			testName:      "Insert at end string elem in non-empty doubly linked list",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			insertValue:   "apple pie",
			expectedList:  []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini", "apple pie"},
		},
		{
			testName:      "Insert at end 3 string elem in non-empty doubly linked list",
			initialValues: []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton"},
			insertValue:   []string{"guava", "honeycrisp apple", "indian fig"},
			expectedList:  []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton", "guava", "honeycrisp apple", "indian fig"},
		},

		{
			testName:      "Insert at end []int elem in empty doubly linked list",
			initialValues: [][]int{},
			insertValue:   []int{1, 2, 3},
			expectedList: [][]int{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end []int elem in non-empty doubly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: []int{1, 2, 3},
			expectedList: [][]int{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end multiple []int elems in non-empty doubly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
		},

		{
			testName:      "Insert at end []int32 elem in empty doubly linked list",
			initialValues: [][]int32{},
			insertValue:   []int32{1, 2, 3},
			expectedList: [][]int32{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end []int32 elem in non-empty doubly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: []int32{1, 2, 3},
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end multiple []int32 elems in non-empty doubly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int32{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
		},

		{
			testName:      "Insert at end []int64 elem in empty doubly linked list",
			initialValues: [][]int64{},
			insertValue:   []int64{1, 2, 3},
			expectedList: [][]int64{
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end []int64 elem in non-empty doubly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: []int64{1, 2, 3},
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},
		{
			testName: "Insert at end multiple []int64 elems in non-empty doubly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: [][]int64{
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
		},

		{
			testName:      "Insert at end []float32 elem in empty doubly linked list",
			initialValues: [][]float32{},
			insertValue:   []float32{1.1, 2.2, 3.3},
			expectedList: [][]float32{
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end []float32 elem in non-empty doubly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float32{1.1, 2.2, 3.3},
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end multiple []float32 elems in non-empty doubly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: [][]float32{
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
		},

		{
			testName:      "Insert at end []float64 elem in empty doubly linked list",
			initialValues: [][]float64{},
			insertValue:   []float64{1.1, 2.2, 3.3},
			expectedList: [][]float64{
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end []float64 elem in non-empty doubly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float64{1.1, 2.2, 3.3},
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},
		{
			testName: "Insert at end multiple []float64 elems in non-empty doubly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: [][]float64{
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
		},

		{
			testName:      "Insert at end []string elem in empty doubly linked list",
			initialValues: [][]string{},
			insertValue:   []string{"apple", "banana"},
			expectedList: [][]string{
				{"apple", "banana"},
			},
		},
		{
			testName: "Insert at end []string elem in non-empty doubly linked list",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			insertValue: []string{"apple", "banana"},
			expectedList: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
				{"apple", "banana"},
			},
		},
		{
			testName: "Insert at end multiple []string elems in non-empty doubly linked list",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
			},
			insertValue: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case int:
					list.InsertAtEnd(x)
				case []int:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case int32:
					list.InsertAtEnd(x)
				case []int32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case int64:
					list.InsertAtEnd(x)
				case []int64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case float32:
					list.InsertAtEnd(x)
				case []float32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case float64:
					list.InsertAtEnd(x)
				case []float64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case string:
					list.InsertAtEnd(x)
				case []string:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []int:
					list.InsertAtEnd(x)
				case [][]int:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []int32:
					list.InsertAtEnd(x)
				case [][]int32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []int64:
					list.InsertAtEnd(x)
				case [][]int64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []float32:
					list.InsertAtEnd(x)
				case [][]float32:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []float64:
					list.InsertAtEnd(x)
				case [][]float64:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				switch x := test.insertValue.(type) {
				case []string:
					list.InsertAtEnd(x)
				case [][]string:
					for _, elem := range x {
						list.InsertAtEnd(elem)
					}
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_InsertAtPos(t *testing.T) {
	tests := []testInsertAtPositionList{
		{
			testName:      "Insert int elem at position 0 in empty singly linked list",
			initialValues: []int{},
			insertValue:   1,
			position:      0,
			expectedList:  []int{1},
		},
		{
			testName:      "Insert int elem at position 2 in non-empty singly linked list",
			initialValues: []int{100, 42, 72, 5, 36},
			insertValue:   99,
			position:      2,
			expectedList:  []int{100, 42, 99, 72, 5, 36},
		},

		{
			testName:      "Insert int32 elem at position 0 in empty singly linked list",
			initialValues: []int32{},
			insertValue:   int32(1),
			position:      0,
			expectedList:  []int32{1},
		},
		{
			testName:      "Insert int32 elem at position 2 in non-empty singly linked list",
			initialValues: []int32{100, 42, 72, 5, 36},
			insertValue:   int32(99),
			position:      2,
			expectedList:  []int32{100, 42, 99, 72, 5, 36},
		},

		{
			testName:      "Insert int64 elem at position 0 in empty singly linked list",
			initialValues: []int64{},
			insertValue:   int64(1),
			position:      0,
			expectedList:  []int64{1},
		},
		{
			testName:      "Insert int64 elem at position 2 in non-empty singly linked list",
			initialValues: []int64{100, 42, 72, 5, 36},
			insertValue:   int64(99),
			position:      2,
			expectedList:  []int64{100, 42, 99, 72, 5, 36},
		},

		{
			testName:      "Insert float32 elem at position 0 in empty singly linked list",
			initialValues: []float32{},
			insertValue:   float32(1.0),
			position:      0,
			expectedList:  []float32{1.0},
		},
		{
			testName:      "Insert float32 elem at position 2 in non-empty singly linked list",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1},
			insertValue:   float32(99.9),
			position:      2,
			expectedList:  []float32{100.5, 42.3, 99.9, 72.7, 5.9, 36.1},
		},

		{
			testName:      "Insert float64 elem at position 0 in empty singly linked list",
			initialValues: []float64{},
			insertValue:   float64(1.0),
			position:      0,
			expectedList:  []float64{1.0},
		},
		{
			testName:      "Insert float64 elem at position 2 in non-empty singly linked list",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1},
			insertValue:   float64(99.9),
			position:      2,
			expectedList:  []float64{100.5, 42.3, 99.9, 72.7, 5.9, 36.1},
		},

		{
			testName:      "Insert string elem at position 0 in empty singly linked list",
			initialValues: []string{},
			insertValue:   "raspberry",
			position:      0,
			expectedList:  []string{"raspberry"},
		},
		{
			testName:      "Insert string elem at position 2 in non-empty singly linked list",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange"},
			insertValue:   "apple pie",
			position:      2,
			expectedList:  []string{"strawberry", "tangerine", "apple pie", "ugly fruit", "valencia orange"},
		},

		{
			testName:      "Insert []int elem at position 0 in empty singly linked list",
			initialValues: [][]int{},
			insertValue:   []int{1, 2, 3},
			position:      0,
			expectedList:  [][]int{{1, 2, 3}},
		},
		{
			testName: "Insert []int elem at position 2 in non-empty singly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: []int{1, 2, 3},
			position:    2,
			expectedList: [][]int{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},

		{
			testName:      "Insert []int32 elem at position 0 in empty singly linked list",
			initialValues: [][]int32{},
			insertValue:   []int32{1, 2, 3},
			position:      0,
			expectedList:  [][]int32{{1, 2, 3}},
		},
		{
			testName: "Insert []int32 elem at position 2 in non-empty singly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: []int32{1, 2, 3},
			position:    2,
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},

		{
			testName:      "Insert []int64 elem at position 0 in empty singly linked list",
			initialValues: [][]int64{},
			insertValue:   []int64{1, 2, 3},
			position:      0,
			expectedList:  [][]int64{{1, 2, 3}},
		},
		{
			testName: "Insert []int64 elem at position 2 in non-empty singly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: []int64{1, 2, 3},
			position:    2,
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},

		{
			testName:      "Insert []float32 elem at position 0 in empty singly linked list",
			initialValues: [][]float32{},
			insertValue:   []float32{1.1, 2.2, 3.3},
			position:      0,
			expectedList:  [][]float32{{1.1, 2.2, 3.3}},
		},
		{
			testName: "Insert []float32 elem at position 2 in non-empty singly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float32{1.1, 2.2, 3.3},
			position:    2,
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},

		{
			testName:      "Insert []float64 elem at position 0 in empty singly linked list",
			initialValues: [][]float64{},
			insertValue:   []float64{1.1, 2.2, 3.3},
			position:      0,
			expectedList:  [][]float64{{1.1, 2.2, 3.3}},
		},
		{
			testName: "Insert []float64 elem at position 2 in non-empty singly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float64{1.1, 2.2, 3.3},
			position:    2,
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},

		{
			testName:      "Insert []string elem at position 0 in empty singly linked list",
			initialValues: [][]string{},
			insertValue:   []string{"apple", "banana"},
			position:      0,
			expectedList:  [][]string{{"apple", "banana"}},
		},
		{
			testName: "Insert []string elem at position 2 in non-empty singly linked list",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
			},
			insertValue: []string{"apple", "banana"},
			position:    2,
			expectedList: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(int)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(int32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(int64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(float32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(float64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(string)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
					fmt.Println(elem)
				}

				x := test.insertValue.([]int)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}
				fmt.Println(result)
				//assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]int32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]int64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]float32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]float64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]string)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_InsertAtPos(t *testing.T) {
	tests := []testInsertAtPositionList{
		{
			testName:      "Insert int elem at position 0 in empty doubly linked list",
			initialValues: []int{},
			insertValue:   1,
			position:      0,
			expectedList:  []int{1},
		},
		{
			testName:      "Insert int elem at position 2 in non-empty doubly linked list",
			initialValues: []int{100, 42, 72, 5, 36},
			insertValue:   99,
			position:      2,
			expectedList:  []int{100, 42, 99, 72, 5, 36},
		},

		{
			testName:      "Insert int32 elem at position 0 in empty doubly linked list",
			initialValues: []int32{},
			insertValue:   int32(1),
			position:      0,
			expectedList:  []int32{1},
		},
		{
			testName:      "Insert int32 elem at position 2 in non-empty doubly linked list",
			initialValues: []int32{100, 42, 72, 5, 36},
			insertValue:   int32(99),
			position:      2,
			expectedList:  []int32{100, 42, 99, 72, 5, 36},
		},

		{
			testName:      "Insert int64 elem at position 0 in empty doubly linked list",
			initialValues: []int64{},
			insertValue:   int64(1),
			position:      0,
			expectedList:  []int64{1},
		},
		{
			testName:      "Insert int64 elem at position 2 in non-empty doubly linked list",
			initialValues: []int64{100, 42, 72, 5, 36},
			insertValue:   int64(99),
			position:      2,
			expectedList:  []int64{100, 42, 99, 72, 5, 36},
		},

		{
			testName:      "Insert float32 elem at position 0 in empty doubly linked list",
			initialValues: []float32{},
			insertValue:   float32(1.0),
			position:      0,
			expectedList:  []float32{1.0},
		},
		{
			testName:      "Insert float32 elem at position 2 in non-empty doubly linked list",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1},
			insertValue:   float32(99.9),
			position:      2,
			expectedList:  []float32{100.5, 42.3, 99.9, 72.7, 5.9, 36.1},
		},

		{
			testName:      "Insert float64 elem at position 0 in empty doubly linked list",
			initialValues: []float64{},
			insertValue:   float64(1.0),
			position:      0,
			expectedList:  []float64{1.0},
		},
		{
			testName:      "Insert float64 elem at position 2 in non-empty doubly linked list",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1},
			insertValue:   float64(99.9),
			position:      2,
			expectedList:  []float64{100.5, 42.3, 99.9, 72.7, 5.9, 36.1},
		},

		{
			testName:      "Insert string elem at position 0 in empty doubly linked list",
			initialValues: []string{},
			insertValue:   "raspberry",
			position:      0,
			expectedList:  []string{"raspberry"},
		},
		{
			testName:      "Insert string elem at position 2 in non-empty doubly linked list",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange"},
			insertValue:   "apple pie",
			position:      2,
			expectedList:  []string{"strawberry", "tangerine", "apple pie", "ugly fruit", "valencia orange"},
		},

		{
			testName:      "Insert []int elem at position 0 in empty doubly linked list",
			initialValues: [][]int{},
			insertValue:   []int{1, 2, 3},
			position:      0,
			expectedList:  [][]int{{1, 2, 3}},
		},
		{
			testName: "Insert []int elem at position 2 in non-empty doubly linked list",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			insertValue: []int{1, 2, 3},
			position:    2,
			expectedList: [][]int{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},

		{
			testName:      "Insert []int32 elem at position 0 in empty doubly linked list",
			initialValues: [][]int32{},
			insertValue:   []int32{1, 2, 3},
			position:      0,
			expectedList:  [][]int32{{1, 2, 3}},
		},
		{
			testName: "Insert []int32 elem at position 2 in non-empty doubly linked list",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			insertValue: []int32{1, 2, 3},
			position:    2,
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},

		{
			testName:      "Insert []int64 elem at position 0 in empty doubly linked list",
			initialValues: [][]int64{},
			insertValue:   []int64{1, 2, 3},
			position:      0,
			expectedList:  [][]int64{{1, 2, 3}},
		},
		{
			testName: "Insert []int64 elem at position 2 in non-empty doubly linked list",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			insertValue: []int64{1, 2, 3},
			position:    2,
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
				{1, 2, 3},
			},
		},

		{
			testName:      "Insert []float32 elem at position 0 in empty doubly linked list",
			initialValues: [][]float32{},
			insertValue:   []float32{1.1, 2.2, 3.3},
			position:      0,
			expectedList:  [][]float32{{1.1, 2.2, 3.3}},
		},
		{
			testName: "Insert []float32 elem at position 2 in non-empty doubly linked list",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float32{1.1, 2.2, 3.3},
			position:    2,
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},

		{
			testName:      "Insert []float64 elem at position 0 in empty doubly linked list",
			initialValues: [][]float64{},
			insertValue:   []float64{1.1, 2.2, 3.3},
			position:      0,
			expectedList:  [][]float64{{1.1, 2.2, 3.3}},
		},
		{
			testName: "Insert []float64 elem at position 2 in non-empty doubly linked list",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			insertValue: []float64{1.1, 2.2, 3.3},
			position:    2,
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{1.1, 2.2, 3.3},
			},
		},

		{
			testName:      "Insert []string elem at position 0 in empty doubly linked list",
			initialValues: [][]string{},
			insertValue:   []string{"apple", "banana"},
			position:      0,
			expectedList:  [][]string{{"apple", "banana"}},
		},
		{
			testName: "Insert []string elem at position 2 in non-empty doubly linked list",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
			},
			insertValue: []string{"apple", "banana"},
			position:    2,
			expectedList: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(int)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(int32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(int64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(float32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(float64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.(string)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]int)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]int32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]int64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]float32)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]float64)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				x := test.insertValue.([]string)

				err := list.InsertAtPos(x, test.position)
				if err != nil {
					panic(err)
				}

				result, err := list.LinkedListToSlice()
				if err != nil {
					panic(err)
				}

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_RemoveFirstNode(t *testing.T) {
	tests := []testRemoveList{
		{
			testName:      "Remove first node with int elem at in singly linked list with 0 elements",
			initialValues: []int{},
			expectedList:  []int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with int elem at in singly linked list with 1 element",
			initialValues: []int{1},
			expectedList:  []int(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with int elem at in singly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int{42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with int32 elem at in singly linked list with 0 elements",
			initialValues: []int32{},
			expectedList:  []int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with int32 elem at in singly linked list with 1 element",
			initialValues: []int32{1},
			expectedList:  []int32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with int32 elem at in singly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int32{42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with int32 elem at in singly linked list with 0 elements",
			initialValues: []int64{},
			expectedList:  []int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with int32 elem at in singly linked list with 1 element",
			initialValues: []int64{1},
			expectedList:  []int64(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with int32 elem at in singly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int64{42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with float32 elem at in singly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with float32 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with float32 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with float64 elem at in singly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with float64 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with float64 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with string elem at in singly linked list with 0 elements",
			initialValues: []string{},
			expectedList:  []string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with string elem at in singly linked list with 1 element",
			initialValues: []string{"raspberry"},
			expectedList:  []string(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with string elem at in singly linked list with 10 element",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedList:  []string{"tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []int elem at in singly linked list with 0 elements",
			initialValues: [][]int{},
			expectedList:  [][]int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []int elem at in singly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			expectedList:  [][]int(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []int elem at in singly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int{
				{30, 40},
				{50, 60},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []int32 elem at in singly linked list with 0 elements",
			initialValues: [][]int32{},
			expectedList:  [][]int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []int32 elem at in singly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			expectedList:  [][]int32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []int32 elem at in singly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int32{
				{30, 40},
				{50, 60},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []int64 elem at in singly linked list with 0 elements",
			initialValues: [][]int64{},
			expectedList:  [][]int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []int64 elem at in singly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			expectedList:  [][]int64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []int64 elem at in singly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int64{
				{30, 40},
				{50, 60},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []float32 elem at in singly linked list with 0 elements",
			initialValues: [][]float32{},
			expectedList:  [][]float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []float32 elem at in singly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			expectedList:  [][]float32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []float32 elem at in singly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float32{
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []float64 elem at in singly linked list with 0 elements",
			initialValues: [][]float64{},
			expectedList:  [][]float64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []float64 elem at in singly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			expectedList:  [][]float64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []float64 elem at in singly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float64{
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []string elem at in singly linked list with 0 elements",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []string elem at in singly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			expectedList:  [][]string(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []string elem at in singly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveFirstNode(t *testing.T) {
	tests := []testRemoveList{
		{
			testName:      "Remove first node with int elem at in doubly linked list with 0 elements",
			initialValues: []int{},
			expectedList:  []int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with int elem at in doubly linked list with 1 element",
			initialValues: []int{1},
			expectedList:  []int(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with int elem at in doubly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int{42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int32{},
			expectedList:  []int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with int32 elem at in doubly linked list with 1 element",
			initialValues: []int32{1},
			expectedList:  []int32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with int32 elem at in doubly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int32{42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int64{},
			expectedList:  []int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with int32 elem at in doubly linked list with 1 element",
			initialValues: []int64{1},
			expectedList:  []int64(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with int32 elem at in doubly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int64{42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with float32 elem at in doubly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with float32 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with float32 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with float64 elem at in doubly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with float64 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with float64 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with string elem at in doubly linked list with 0 elements",
			initialValues: []string{},
			expectedList:  []string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove first node with string elem at in doubly linked list with 1 element",
			initialValues: []string{"raspberry"},
			expectedList:  []string(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove first node with string elem at in doubly linked list with 10 element",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedList:  []string{"tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []int elem at in doubly linked list with 0 elements",
			initialValues: [][]int{},
			expectedList:  [][]int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []int elem at in doubly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			expectedList:  [][]int(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []int elem at in doubly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int{
				{30, 40},
				{50, 60},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []int32 elem at in doubly linked list with 0 elements",
			initialValues: [][]int32{},
			expectedList:  [][]int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []int32 elem at in doubly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			expectedList:  [][]int32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []int32 elem at in doubly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int32{
				{30, 40},
				{50, 60},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []int64 elem at in doubly linked list with 0 elements",
			initialValues: [][]int64{},
			expectedList:  [][]int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []int64 elem at in doubly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			expectedList:  [][]int64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []int64 elem at in doubly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int64{
				{30, 40},
				{50, 60},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []float32 elem at in doubly linked list with 0 elements",
			initialValues: [][]float32{},
			expectedList:  [][]float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []float32 elem at in doubly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			expectedList:  [][]float32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []float32 elem at in doubly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float32{
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []float64 elem at in doubly linked list with 0 elements",
			initialValues: [][]float64{},
			expectedList:  [][]float64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []float64 elem at in doubly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			expectedList:  [][]float64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []float64 elem at in doubly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float64{
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove first node with []string elem at in doubly linked list with 0 elements",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove first node with []string elem at in doubly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			expectedList:  [][]string(nil),
			expectedError: nil,
		},
		{
			testName: "Remove first node with []string elem at in doubly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveFirstNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_RemoveLastNode(t *testing.T) {
	tests := []testRemoveList{
		{
			testName:      "Remove last node with int elem at in singly linked list with 0 elements",
			initialValues: []int{},
			expectedList:  []int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with int elem at in singly linked list with 1 element",
			initialValues: []int{1},
			expectedList:  []int(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with int elem at in singly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int{100, 42, 72, 5, 36, 66, 77, 79, 51},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with int32 elem at in singly linked list with 0 elements",
			initialValues: []int32{},
			expectedList:  []int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with int32 elem at in singly linked list with 1 element",
			initialValues: []int32{1},
			expectedList:  []int32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with int32 elem at in singly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int32{100, 42, 72, 5, 36, 66, 77, 79, 51},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with int32 elem at in singly linked list with 0 elements",
			initialValues: []int64{},
			expectedList:  []int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with int32 elem at in singly linked list with 1 element",
			initialValues: []int64{1},
			expectedList:  []int64(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with int32 elem at in singly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int64{100, 42, 72, 5, 36, 66, 77, 79, 51},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with float32 elem at in singly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with float32 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with float32 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with float64 elem at in singly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with float64 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with float64 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with string elem at in singly linked list with 0 elements",
			initialValues: []string{},
			expectedList:  []string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with string elem at in singly linked list with 1 element",
			initialValues: []string{"raspberry"},
			expectedList:  []string(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with string elem at in singly linked list with 10 element",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedList:  []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon"},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []int elem at in singly linked list with 0 elements",
			initialValues: [][]int{},
			expectedList:  [][]int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []int elem at in singly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			expectedList:  [][]int(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []int elem at in singly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []int32 elem at in singly linked list with 0 elements",
			initialValues: [][]int32{},
			expectedList:  [][]int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []int32 elem at in singly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			expectedList:  [][]int32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []int32 elem at in singly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []int64 elem at in singly linked list with 0 elements",
			initialValues: [][]int64{},
			expectedList:  [][]int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []int64 elem at in singly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			expectedList:  [][]int64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []int64 elem at in singly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []float32 elem at in singly linked list with 0 elements",
			initialValues: [][]float32{},
			expectedList:  [][]float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []float32 elem at in singly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			expectedList:  [][]float32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []float32 elem at in singly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []float64 elem at in singly linked list with 0 elements",
			initialValues: [][]float64{},
			expectedList:  [][]float64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []float64 elem at in singly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			expectedList:  [][]float64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []float64 elem at in singly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []string elem at in singly linked list with 0 elements",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []string elem at in singly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			expectedList:  [][]string(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []string elem at in singly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveLastNode(t *testing.T) {
	tests := []testRemoveList{
		{
			testName:      "Remove last node with int elem at in doubly linked list with 0 elements",
			initialValues: []int{},
			expectedList:  []int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with int elem at in doubly linked list with 1 element",
			initialValues: []int{1},
			expectedList:  []int(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with int elem at in doubly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int{100, 42, 72, 5, 36, 66, 77, 79, 51},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int32{},
			expectedList:  []int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with int32 elem at in doubly linked list with 1 element",
			initialValues: []int32{1},
			expectedList:  []int32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with int32 elem at in doubly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int32{100, 42, 72, 5, 36, 66, 77, 79, 51},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int64{},
			expectedList:  []int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with int32 elem at in doubly linked list with 1 element",
			initialValues: []int64{1},
			expectedList:  []int64(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with int32 elem at in doubly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int64{100, 42, 72, 5, 36, 66, 77, 79, 51},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with float32 elem at in doubly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with float32 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with float32 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with float64 elem at in doubly linked list with 0 elements",
			initialValues: []float32{},
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with float64 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0},
			expectedList:  []float32(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with float64 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedList:  []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with string elem at in doubly linked list with 0 elements",
			initialValues: []string{},
			expectedList:  []string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName:      "Remove last node with string elem at in doubly linked list with 1 element",
			initialValues: []string{"raspberry"},
			expectedList:  []string(nil),
			expectedError: nil,
		},
		{
			testName:      "Remove last node with string elem at in doubly linked list with 10 element",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedList:  []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon"},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []int elem at in doubly linked list with 0 elements",
			initialValues: [][]int{},
			expectedList:  [][]int(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []int elem at in doubly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			expectedList:  [][]int(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []int elem at in doubly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []int32 elem at in doubly linked list with 0 elements",
			initialValues: [][]int32{},
			expectedList:  [][]int32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []int32 elem at in doubly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			expectedList:  [][]int32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []int32 elem at in doubly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []int64 elem at in doubly linked list with 0 elements",
			initialValues: [][]int64{},
			expectedList:  [][]int64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []int64 elem at in doubly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			expectedList:  [][]int64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []int64 elem at in doubly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []float32 elem at in doubly linked list with 0 elements",
			initialValues: [][]float32{},
			expectedList:  [][]float32(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []float32 elem at in doubly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			expectedList:  [][]float32(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []float32 elem at in doubly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []float64 elem at in doubly linked list with 0 elements",
			initialValues: [][]float64{},
			expectedList:  [][]float64(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []float64 elem at in doubly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			expectedList:  [][]float64(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []float64 elem at in doubly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove last node with []string elem at in doubly linked list with 0 elements",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
			expectedError: linked_lists.ErrListEmpty,
		},
		{
			testName: "Remove last node with []string elem at in doubly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			expectedList:  [][]string(nil),
			expectedError: nil,
		},
		{
			testName: "Remove last node with []string elem at in doubly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveLastNode()

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_RemoveNodeAtPos(t *testing.T) {
	tests := []testRemoveAtPositionList{
		{
			testName:      "Remove node at pos with int elem at in singly linked list with 0 elements",
			initialValues: []int{},
			position:      1,
			expectedList:  []int(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int elem at in singly linked list with 0 elements",
			initialValues: []int{1},
			position:      1,
			expectedList:  []int{1},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName:      "Remove node at pos with int elem at in singly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			position:      3,
			expectedList:  []int{100, 42, 72, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with int32 elem at in singly linked list with 0 elements",
			initialValues: []int32{},
			position:      1,
			expectedList:  []int32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int32 elem at in singly linked list with 0 elements",
			initialValues: []int32{1},
			position:      1,
			expectedList:  []int32{1},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName:      "Remove node at pos with int64 elem at in singly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			position:      3,
			expectedList:  []int64{100, 42, 72, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with int64 elem at in singly linked list with 0 elements",
			initialValues: []int64{},
			position:      1,
			expectedList:  []int64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int64 elem at in singly linked list with 0 elements",
			initialValues: []int64{1},
			position:      1,
			expectedList:  []int64{1},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName:      "Remove node at pos with int64 elem at in singly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			position:      3,
			expectedList:  []int64{100, 42, 72, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with float32 elem at in singly linked list with 0 elements",
			initialValues: []float32{},
			position:      1,
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float32 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0},
			position:      1,
			expectedList:  []float32{1.0},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName:      "Remove node at pos with float32 elem at in singly linked list with 10 elements",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			position:      3,
			expectedList:  []float32{1.0, 3.3, 51.2, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with float64 elem at in singly linked list with 0 elements",
			initialValues: []float64{},
			position:      1,
			expectedList:  []float64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float64 elem at in singly linked list with 1 element",
			initialValues: []float64{1.0},
			position:      1,
			expectedList:  []float64{1.0},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName:      "Remove node at pos with float64 elem at in singly linked list with 10 elements",
			initialValues: []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			position:      3,
			expectedList:  []float64{1.0, 3.3, 51.2, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with float64 elem at in singly linked list with 0 elements",
			initialValues: []string{},
			position:      0,
			expectedList:  []string(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float64 elem at in singly linked list with 1 element",
			initialValues: []string{"raspberry"},
			position:      1,
			expectedList:  []string{"raspberry"},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName:      "Remove node at pos with float64 elem at in singly linked list with 10 elements",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			position:      3,
			expectedList:  []string{"strawberry", "tangerine", "ugly fruit", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []int elem at in singly linked list with 0 elements",
			initialValues: [][]int{},
			position:      0,
			expectedList:  [][]int(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int elem at in singly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			position: 1,
			expectedList: [][]int{
				{30, 40},
			},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName: "Remove node at pos with []int elem at in singly linked list with 10 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			position: 2,
			expectedList: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []int32 elem at in singly linked list with 0 elements",
			initialValues: [][]int32{},
			position:      0,
			expectedList:  [][]int32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int32 elem at in singly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			position: 1,
			expectedList: [][]int32{
				{30, 40},
			},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName: "Remove node at pos with []int32 elem at in singly linked list with 10 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			position: 2,
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []int64 elem at in singly linked list with 0 elements",
			initialValues: [][]int64{},
			position:      0,
			expectedList:  [][]int64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int64 elem at in singly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			position: 1,
			expectedList: [][]int64{
				{30, 40},
			},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName: "Remove node at pos with []int64 elem at in singly linked list with 10 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			position: 2,
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []float32 elem at in singly linked list with 0 elements",
			initialValues: [][]float32{},
			position:      1,
			expectedList:  [][]float32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove last node with []float32 elem at in singly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			position: 1,
			expectedList: [][]float32{
				{10.5, 20.6},
			},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName: "Remove last node with []float32 elem at in singly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			position: 2,
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []float64 elem at in singly linked list with 0 elements",
			initialValues: [][]float64{},
			position:      1,
			expectedList:  [][]float64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove last node with []float64 elem at in singly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			position: 1,
			expectedList: [][]float64{
				{10.5, 20.6},
			},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName: "Remove last node with []float64 elem at in singly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			position: 2,
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []string elem at in singly linked list with 0 elements",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
			position:      1,
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []string elem at in singly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			position: 1,
			expectedList: [][]string{
				{"cherry", "date"},
			},
			expectedError: linked_lists.ErrPosOutOfRange,
		},
		{
			testName: "Remove node at pos with []string elem at in singly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			position: 2,
			expectedList: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveNodeAtPos(t *testing.T) {
	tests := []testRemoveAtPositionList{
		{
			testName:      "Remove node at pos with int elem at in doubly linked list with 0 elements",
			initialValues: []int{},
			position:      1,
			expectedList:  []int(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int elem at in doubly linked list with 0 elements",
			initialValues: []int{1},
			position:      1,
			expectedList:  []int{1},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int elem at in doubly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			position:      3,
			expectedList:  []int{100, 42, 72, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int32{},
			position:      1,
			expectedList:  []int32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int32{1},
			position:      1,
			expectedList:  []int32{1},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int64 elem at in doubly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			position:      3,
			expectedList:  []int64{100, 42, 72, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with int64 elem at in doubly linked list with 0 elements",
			initialValues: []int64{},
			position:      1,
			expectedList:  []int64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int64 elem at in doubly linked list with 0 elements",
			initialValues: []int64{1},
			position:      1,
			expectedList:  []int64{1},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with int64 elem at in doubly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			position:      3,
			expectedList:  []int64{100, 42, 72, 36, 66, 77, 79, 51, 3},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with float32 elem at in doubly linked list with 0 elements",
			initialValues: []float32{},
			position:      1,
			expectedList:  []float32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float32 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0},
			position:      1,
			expectedList:  []float32{1.0},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float32 elem at in doubly linked list with 10 elements",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			position:      3,
			expectedList:  []float32{1.0, 3.3, 51.2, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with float64 elem at in doubly linked list with 0 elements",
			initialValues: []float64{},
			position:      1,
			expectedList:  []float64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float64 elem at in doubly linked list with 1 element",
			initialValues: []float64{1.0},
			position:      1,
			expectedList:  []float64{1.0},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float64 elem at in doubly linked list with 10 elements",
			initialValues: []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			position:      3,
			expectedList:  []float64{1.0, 3.3, 51.2, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with float64 elem at in doubly linked list with 0 elements",
			initialValues: []string{},
			position:      0,
			expectedList:  []string(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float64 elem at in doubly linked list with 1 element",
			initialValues: []string{"raspberry"},
			position:      1,
			expectedList:  []string{"raspberry"},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName:      "Remove node at pos with float64 elem at in doubly linked list with 10 elements",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			position:      3,
			expectedList:  []string{"strawberry", "tangerine", "ugly fruit", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []int elem at in doubly linked list with 0 elements",
			initialValues: [][]int{},
			position:      0,
			expectedList:  [][]int(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int elem at in doubly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			position: 1,
			expectedList: [][]int{
				{30, 40},
			},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int elem at in doubly linked list with 10 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			position: 2,
			expectedList: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []int32 elem at in doubly linked list with 0 elements",
			initialValues: [][]int32{},
			position:      0,
			expectedList:  [][]int32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int32 elem at in doubly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			position: 1,
			expectedList: [][]int32{
				{30, 40},
			},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int32 elem at in doubly linked list with 10 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			position: 2,
			expectedList: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []int64 elem at in doubly linked list with 0 elements",
			initialValues: [][]int64{},
			position:      0,
			expectedList:  [][]int64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int64 elem at in doubly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			position: 1,
			expectedList: [][]int64{
				{30, 40},
			},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []int64 elem at in doubly linked list with 10 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			position: 2,
			expectedList: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []float32 elem at in doubly linked list with 0 elements",
			initialValues: [][]float32{},
			position:      1,
			expectedList:  [][]float32(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove last node with []float32 elem at in doubly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			position: 1,
			expectedList: [][]float32{
				{10.5, 20.6},
			},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove last node with []float32 elem at in doubly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			position: 2,
			expectedList: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []float64 elem at in doubly linked list with 0 elements",
			initialValues: [][]float64{},
			position:      1,
			expectedList:  [][]float64(nil),
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove last node with []float64 elem at in doubly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			position: 1,
			expectedList: [][]float64{
				{10.5, 20.6},
			},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove last node with []float64 elem at in doubly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			position: 2,
			expectedList: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedError: nil,
		},

		{
			testName:      "Remove node at pos with []string elem at in doubly linked list with 0 elements",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
			position:      1,
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []string elem at in doubly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			position: 1,
			expectedList: [][]string{
				{"cherry", "date"},
			},
			expectedError: linked_lists.ErrInvalidPos,
		},
		{
			testName: "Remove node at pos with []string elem at in doubly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			position: 2,
			expectedList: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				errRemove := list.RemoveNodeAtPosition(test.position)

				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_FindNode(t *testing.T) {
	tests := []testFindNodeList{
		{
			testName:      "Find node with int elem at in singly linked list with 0 elements",
			initialValues: []int{},
			findValue:     1,
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int elem at in singly linked list with 1 element",
			initialValues: []int{1},
			findValue:     1,
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with int elem at in singly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     52,
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int elem at in singly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     5,
			expectedIndex: 3,
			expectedError: nil,
		},

		{
			testName:      "Find node with int32 elem at in singly linked list with 0 elements",
			initialValues: []int32{},
			findValue:     int32(1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int32 elem at in singly linked list with 1 element",
			initialValues: []int32{1},
			findValue:     int32(1),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with int32 elem at in singly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int32(52),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int32 elem at in singly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int32(5),
			expectedIndex: 3,
			expectedError: nil,
		},

		{
			testName:      "Find node with int64 elem at in singly linked list with 0 elements",
			initialValues: []int64{},
			findValue:     int64(1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int64 elem at in singly linked list with 1 element",
			initialValues: []int64{1},
			findValue:     int64(1),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with int64 elem at in singly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int64(52),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int64 elem at in singly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int64(5),
			expectedIndex: 3,
			expectedError: nil,
		},

		{
			testName:      "Find node with float32 elem at in singly linked list with 0 elements",
			initialValues: []float32{},
			findValue:     float32(1.0),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float32 elem at in singly linked list with 1 element",
			initialValues: []float32{1.0},
			findValue:     float32(1.0),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with float32 elem at in singly linked list with 10 elements",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float32(52.1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float32 elem at in singly linked list with 10 elements",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float32(51.2),
			expectedIndex: 2,
			expectedError: nil,
		},

		{
			testName:      "Find node with float64 elem at in singly linked list with 0 elements",
			initialValues: []float64{},
			findValue:     float64(1.0),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float64 elem at in singly linked list with 1 element",
			initialValues: []float64{1.0},
			findValue:     float64(1.0),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with float64 elem at in singly linked list with 10 elements",
			initialValues: []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float64(52.1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float64 elem at in singly linked list with 10 elements",
			initialValues: []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float64(51.2),
			expectedIndex: 2,
			expectedError: nil,
		},

		{
			testName:      "Find node with string elem at in singly linked list with 0 elements",
			initialValues: []string{},
			findValue:     "raspberry",
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with string elem at in singly linked list with 1 element",
			initialValues: []string{"raspberry"},
			findValue:     "raspberry",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with string elem at in singly linked list with 10 elements",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			findValue:     "apple",
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with string elem at in singly linked list with 10 elements",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			findValue:     "tangerine",
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []int elem at in singly linked list with 0 elements",
			initialValues: [][]int{},
			findValue:     []int{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int elem at in singly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			findValue:     []int{30, 40},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []int elem at in singly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int elem at in singly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int{30, 40},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []int32 elem at in singly linked list with 0 elements",
			initialValues: [][]int32{},
			findValue:     []int32{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int32 elem at in singly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			findValue:     []int32{30, 40},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []int32 elem at in singly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int32{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int32 elem at in singly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int32{30, 40},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []int64 elem at in singly linked list with 0 elements",
			initialValues: [][]int64{},
			findValue:     []int64{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int64 elem at in singly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			findValue:     []int64{30, 40},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []int64 elem at in singly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int64{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int64 elem at in singly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int64{30, 40},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []float32 elem at in singly linked list with 0 elements",
			initialValues: [][]float32{},
			findValue:     []float32{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float32 elem at in singly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			findValue:     []float32{10.5, 20.6},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []float32 elem at in singly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float32{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float32 elem at in singly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float32{30.7, 40.8},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []float64 elem at in singly linked list with 0 elements",
			initialValues: [][]float64{},
			findValue:     []float64{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float64 elem at in singly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			findValue:     []float64{10.5, 20.6},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []float64 elem at in singly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float64{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float64 elem at in singly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float64{30.7, 40.8},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []string elem at in singly linked list with 0 elements",
			initialValues: [][]string{},
			findValue:     []string{"cherry"},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []string elem at in singly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			findValue:     []string{"cherry", "date"},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []string elem at in singly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			findValue:     []string{"raspberry", "tangerine"},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []string elem at in singly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			findValue:     []string{"apple", "banana"},
			expectedIndex: 1,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(int))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(int32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(int64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(float32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(float64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(string))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]int))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]int32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]int64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]float32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]float64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]string))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			}
		})
	}
}

func TestDoublyLinkedList_FindNode(t *testing.T) {
	tests := []testFindNodeList{
		{
			testName:      "Find node with int elem at in doubly linked list with 0 elements",
			initialValues: []int{},
			findValue:     1,
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int elem at in doubly linked list with 1 element",
			initialValues: []int{1},
			findValue:     1,
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with int elem at in doubly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     52,
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int elem at in doubly linked list with 10 elements",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     5,
			expectedIndex: 3,
			expectedError: nil,
		},

		{
			testName:      "Find node with int32 elem at in doubly linked list with 0 elements",
			initialValues: []int32{},
			findValue:     int32(1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int32 elem at in doubly linked list with 1 element",
			initialValues: []int32{1},
			findValue:     int32(1),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with int32 elem at in doubly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int32(52),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int32 elem at in doubly linked list with 10 elements",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int32(5),
			expectedIndex: 3,
			expectedError: nil,
		},

		{
			testName:      "Find node with int64 elem at in doubly linked list with 0 elements",
			initialValues: []int64{},
			findValue:     int64(1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int64 elem at in doubly linked list with 1 element",
			initialValues: []int64{1},
			findValue:     int64(1),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with int64 elem at in doubly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int64(52),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with int64 elem at in doubly linked list with 10 elements",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			findValue:     int64(5),
			expectedIndex: 3,
			expectedError: nil,
		},

		{
			testName:      "Find node with float32 elem at in doubly linked list with 0 elements",
			initialValues: []float32{},
			findValue:     float32(1.0),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float32 elem at in doubly linked list with 1 element",
			initialValues: []float32{1.0},
			findValue:     float32(1.0),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with float32 elem at in doubly linked list with 10 elements",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float32(52.1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float32 elem at in doubly linked list with 10 elements",
			initialValues: []float32{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float32(51.2),
			expectedIndex: 2,
			expectedError: nil,
		},

		{
			testName:      "Find node with float64 elem at in doubly linked list with 0 elements",
			initialValues: []float64{},
			findValue:     float64(1.0),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float64 elem at in doubly linked list with 1 element",
			initialValues: []float64{1.0},
			findValue:     float64(1.0),
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with float64 elem at in doubly linked list with 10 elements",
			initialValues: []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float64(52.1),
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with float64 elem at in doubly linked list with 10 elements",
			initialValues: []float64{1.0, 3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
			findValue:     float64(51.2),
			expectedIndex: 2,
			expectedError: nil,
		},

		{
			testName:      "Find node with string elem at in doubly linked list with 0 elements",
			initialValues: []string{},
			findValue:     "raspberry",
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with string elem at in doubly linked list with 1 element",
			initialValues: []string{"raspberry"},
			findValue:     "raspberry",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName:      "Find node with string elem at in doubly linked list with 10 elements",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			findValue:     "apple",
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName:      "Find node with string elem at in doubly linked list with 10 elements",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			findValue:     "tangerine",
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []int elem at in doubly linked list with 0 elements",
			initialValues: [][]int{},
			findValue:     []int{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int elem at in doubly linked list with 1 element",
			initialValues: [][]int{
				{30, 40},
			},
			findValue:     []int{30, 40},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []int elem at in doubly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int elem at in doubly linked list with 3 elements",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int{30, 40},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []int32 elem at in doubly linked list with 0 elements",
			initialValues: [][]int32{},
			findValue:     []int32{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int32 elem at in doubly linked list with 1 element",
			initialValues: [][]int32{
				{30, 40},
			},
			findValue:     []int32{30, 40},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []int32 elem at in doubly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int32{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int32 elem at in doubly linked list with 3 elements",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int32{30, 40},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []int64 elem at in doubly linked list with 0 elements",
			initialValues: [][]int64{},
			findValue:     []int64{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int64 elem at in doubly linked list with 1 element",
			initialValues: [][]int64{
				{30, 40},
			},
			findValue:     []int64{30, 40},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []int64 elem at in doubly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int64{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []int64 elem at in doubly linked list with 3 elements",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
			},
			findValue:     []int64{30, 40},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []float32 elem at in doubly linked list with 0 elements",
			initialValues: [][]float32{},
			findValue:     []float32{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float32 elem at in doubly linked list with 1 element",
			initialValues: [][]float32{
				{10.5, 20.6},
			},
			findValue:     []float32{10.5, 20.6},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []float32 elem at in doubly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float32{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float32 elem at in doubly linked list with 3 elements",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float32{30.7, 40.8},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []float64 elem at in doubly linked list with 0 elements",
			initialValues: [][]float64{},
			findValue:     []float64{1},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float64 elem at in doubly linked list with 1 element",
			initialValues: [][]float64{
				{10.5, 20.6},
			},
			findValue:     []float64{10.5, 20.6},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []float64 elem at in doubly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float64{228, 1337},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []float64 elem at in doubly linked list with 3 elements",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
			},
			findValue:     []float64{30.7, 40.8},
			expectedIndex: 1,
			expectedError: nil,
		},

		{
			testName:      "Find node with []string elem at in doubly linked list with 0 elements",
			initialValues: [][]string{},
			findValue:     []string{"cherry"},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []string elem at in doubly linked list with 1 element",
			initialValues: [][]string{
				{"cherry", "date"},
			},
			findValue:     []string{"cherry", "date"},
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			testName: "Find node with []string elem at in doubly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			findValue:     []string{"raspberry", "tangerine"},
			expectedIndex: -1,
			expectedError: linked_lists.ErrValueNotFound,
		},
		{
			testName: "Find node with []string elem at in doubly linked list with 3 elements",
			initialValues: [][]string{
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			findValue:     []string{"apple", "banana"},
			expectedIndex: 1,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(int))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(int32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(int64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(float32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(float64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.(string))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]int))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]int32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]int64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]float32))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]float64))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				index, errRemove := list.FindNode(test.findValue.([]string))

				assert.Equal(t, test.expectedError, errRemove)
				assert.Equal(t, test.expectedIndex, index)
			}
		})
	}
}

func TestSinglyLinkedList_ReverseList(t *testing.T) {
	tests := []testReverseList{
		{
			testName:      "Reverse empty singly linked list with int values",
			initialValues: []int{},
			expectedList:  []int(nil),
		},
		{
			testName:      "Reverse non-empty singly linked list with int values",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int{3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Reverse non-empty singly linked list with int values",
			initialValues: []int{35, 50, 69, 71, 74},
			expectedList:  []int{74, 71, 69, 50, 35},
		},

		{
			testName:      "Reverse empty singly linked list with int32 values",
			initialValues: []int32{},
			expectedList:  []int32(nil),
		},
		{
			testName:      "Reverse non-empty singly linked list with int32 values",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int32{3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Reverse non-empty singly linked list with int32 values",
			initialValues: []int32{35, 50, 69, 71, 74},
			expectedList:  []int32{74, 71, 69, 50, 35},
		},

		{
			testName:      "Reverse empty singly linked list with int64 values",
			initialValues: []int64{},
			expectedList:  []int64(nil),
		},
		{
			testName:      "Reverse non-empty singly linked list with int64 values",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int64{3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Reverse non-empty singly linked list with int64 values",
			initialValues: []int64{35, 50, 69, 71, 74},
			expectedList:  []int64{74, 71, 69, 50, 35},
		},

		{
			testName:      "Reverse empty singly linked list with float32 values",
			initialValues: []float32{},
			expectedList:  []float32(nil),
		},
		{
			testName:      "Reverse non-empty singly linked list with float32 values",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			expectedList:  []float32{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
		},
		{
			testName:      "Reverse non-empty singly linked list with float32 values",
			initialValues: []float32{35.7, 50.4, 69.9, 71.2, 74.8},
			expectedList:  []float32{74.8, 71.2, 69.9, 50.4, 35.7},
		},

		{
			testName:      "Reverse empty singly linked list with float64 values",
			initialValues: []float64{},
			expectedList:  []float64(nil),
		},
		{
			testName:      "Reverse non-empty singly linked list with float32 values",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			expectedList:  []float64{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
		},
		{
			testName:      "Reverse non-empty singly linked list with float32 values",
			initialValues: []float64{35.7, 50.4, 69.9, 71.2, 74.8},
			expectedList:  []float64{74.8, 71.2, 69.9, 50.4, 35.7},
		},

		{
			testName:      "Reverse empty singly linked list with string values",
			initialValues: []string{},
			expectedList:  []string(nil),
		},
		{
			testName:      "Reverse non-empty singly linked list with string values",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedList:  []string{"zucchini", "yellow watermelon", "xigua", "watermelon", "valencia orange", "ugly fruit", "tangerine", "strawberry"},
		},
		{
			testName:      "Reverse non-empty singly linked list with string values",
			initialValues: []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton"},
			expectedList:  []string{"fig newton", "elderflower", "dragon fruit", "cantaloupe", "blueberry muffin"},
		},

		{
			testName:      "Reverse empty singly linked list with []int values",
			initialValues: [][]int{},
			expectedList:  [][]int(nil),
		},
		{
			testName: "Reverse non-empty singly linked list with []int values",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedList: [][]int{
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Reverse non-empty singly linked list with []int values",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Reverse empty singly linked list with []int32 values",
			initialValues: [][]int32{},
			expectedList:  [][]int32(nil),
		},
		{
			testName: "Reverse non-empty singly linked list with []int32 values",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedList: [][]int32{
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Reverse non-empty singly linked list with []int32 values",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int32{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Reverse empty singly linked list with []int64 values",
			initialValues: [][]int64{},
			expectedList:  [][]int64(nil),
		},
		{
			testName: "Reverse non-empty singly linked list with []int64 values",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedList: [][]int64{
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Reverse non-empty singly linked list with []int64 values",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int64{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Reverse empty singly linked list with []float32 value",
			initialValues: [][]float32{},
			expectedList:  [][]float32(nil),
		},
		{
			testName: "Reverse non-empty singly linked list with []float32 value",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedList: [][]float32{
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},
		{
			testName: "Reverse non-empty singly linked list with []float32 value",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float32{
				{70.2, 80.3},
				{50.9, 60.1},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},

		{
			testName:      "Reverse empty singly linked list with []float64 value",
			initialValues: [][]float64{},
			expectedList:  [][]float64(nil),
		},
		{
			testName: "Reverse non-empty singly linked list with []float64 value",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedList: [][]float64{
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},
		{
			testName: "Reverse non-empty singly linked list with []float64 value",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float64{
				{70.2, 80.3},
				{50.9, 60.1},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},

		{
			testName:      "Reverse empty singly linked list with []string value",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
		},
		{
			testName: "Reverse non-empty singly linked list with []string value",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			expectedList: [][]string{
				{"elderberry", "fig"},
				{"cherry", "date"},
			},
		},
		{
			testName: "Reverse non-empty singly linked list with []string value",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"cherry", "date"},
				{"apple", "banana"},
				{"kiwi", "lemon"},
				{"grape", "honeydew"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestDoublyLinkedList_ReverseList(t *testing.T) {
	tests := []testReverseList{
		{
			testName:      "Reverse empty doubly linked list with int values",
			initialValues: []int{},
			expectedList:  []int(nil),
		},
		{
			testName:      "Reverse non-empty doubly linked list with int values",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int{3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Reverse non-empty doubly linked list with int values",
			initialValues: []int{35, 50, 69, 71, 74},
			expectedList:  []int{74, 71, 69, 50, 35},
		},

		{
			testName:      "Reverse empty doubly linked list with int32 values",
			initialValues: []int32{},
			expectedList:  []int32(nil),
		},
		{
			testName:      "Reverse non-empty doubly linked list with int32 values",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int32{3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Reverse non-empty doubly linked list with int32 values",
			initialValues: []int32{35, 50, 69, 71, 74},
			expectedList:  []int32{74, 71, 69, 50, 35},
		},

		{
			testName:      "Reverse empty doubly linked list with int64 values",
			initialValues: []int64{},
			expectedList:  []int64(nil),
		},
		{
			testName:      "Reverse non-empty doubly linked list with int64 values",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedList:  []int64{3, 51, 79, 77, 66, 36, 5, 72, 42, 100},
		},
		{
			testName:      "Reverse non-empty doubly linked list with int64 values",
			initialValues: []int64{35, 50, 69, 71, 74},
			expectedList:  []int64{74, 71, 69, 50, 35},
		},

		{
			testName:      "Reverse empty doubly linked list with float32 values",
			initialValues: []float32{},
			expectedList:  []float32(nil),
		},
		{
			testName:      "Reverse non-empty doubly linked list with float32 values",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			expectedList:  []float32{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
		},
		{
			testName:      "Reverse non-empty doubly linked list with float32 values",
			initialValues: []float32{35.7, 50.4, 69.9, 71.2, 74.8},
			expectedList:  []float32{74.8, 71.2, 69.9, 50.4, 35.7},
		},

		{
			testName:      "Reverse empty doubly linked list with float64 values",
			initialValues: []float64{},
			expectedList:  []float64(nil),
		},
		{
			testName:      "Reverse non-empty doubly linked list with float32 values",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			expectedList:  []float64{3.3, 51.2, 79.4, 77.8, 66.6, 36.1, 5.9, 72.7, 42.3, 100.5},
		},
		{
			testName:      "Reverse non-empty doubly linked list with float32 values",
			initialValues: []float64{35.7, 50.4, 69.9, 71.2, 74.8},
			expectedList:  []float64{74.8, 71.2, 69.9, 50.4, 35.7},
		},

		{
			testName:      "Reverse empty doubly linked list with string values",
			initialValues: []string{},
			expectedList:  []string(nil),
		},
		{
			testName:      "Reverse non-empty doubly linked list with string values",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedList:  []string{"zucchini", "yellow watermelon", "xigua", "watermelon", "valencia orange", "ugly fruit", "tangerine", "strawberry"},
		},
		{
			testName:      "Reverse non-empty doubly linked list with string values",
			initialValues: []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton"},
			expectedList:  []string{"fig newton", "elderflower", "dragon fruit", "cantaloupe", "blueberry muffin"},
		},

		{
			testName:      "Reverse empty doubly linked list with []int values",
			initialValues: [][]int{},
			expectedList:  [][]int(nil),
		},
		{
			testName: "Reverse non-empty doubly linked list with []int values",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedList: [][]int{
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Reverse non-empty doubly linked list with []int values",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Reverse empty doubly linked list with []int32 values",
			initialValues: [][]int32{},
			expectedList:  [][]int32(nil),
		},
		{
			testName: "Reverse non-empty doubly linked list with []int32 values",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedList: [][]int32{
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Reverse non-empty doubly linked list with []int32 values",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int32{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Reverse empty doubly linked list with []int64 values",
			initialValues: [][]int64{},
			expectedList:  [][]int64(nil),
		},
		{
			testName: "Reverse non-empty doubly linked list with []int64 values",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedList: [][]int64{
				{30, 40},
				{10, 20},
			},
		},
		{
			testName: "Reverse non-empty doubly linked list with []int64 values",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedList: [][]int64{
				{70, 80},
				{50, 60},
				{30, 40},
				{10, 20},
			},
		},

		{
			testName:      "Reverse empty doubly linked list with []float32 value",
			initialValues: [][]float32{},
			expectedList:  [][]float32(nil),
		},
		{
			testName: "Reverse non-empty doubly linked list with []float32 value",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedList: [][]float32{
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},
		{
			testName: "Reverse non-empty doubly linked list with []float32 value",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float32{
				{70.2, 80.3},
				{50.9, 60.1},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},

		{
			testName:      "Reverse empty doubly linked list with []float64 value",
			initialValues: [][]float64{},
			expectedList:  [][]float64(nil),
		},
		{
			testName: "Reverse non-empty doubly linked list with []float64 value",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedList: [][]float64{
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},
		{
			testName: "Reverse non-empty doubly linked list with []float64 value",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedList: [][]float64{
				{70.2, 80.3},
				{50.9, 60.1},
				{30.7, 40.8},
				{10.5, 20.6},
			},
		},

		{
			testName:      "Reverse empty doubly linked list with []string value",
			initialValues: [][]string{},
			expectedList:  [][]string(nil),
		},
		{
			testName: "Reverse non-empty doubly linked list with []string value",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			expectedList: [][]string{
				{"elderberry", "fig"},
				{"cherry", "date"},
			},
		},
		{
			testName: "Reverse non-empty doubly linked list with []string value",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedList: [][]string{
				{"cherry", "date"},
				{"apple", "banana"},
				{"kiwi", "lemon"},
				{"grape", "honeydew"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewDoublyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				list.Reverse()
				result, _ := list.LinkedListToSlice()

				assert.Equal(t, test.expectedList, result)
			}
		})
	}
}

func TestSinglyLinkedList_LenOfList(t *testing.T) {
	tests := []testLenList{
		{
			testName:      "Singly linked list with 0 elements of type int",
			initialValues: []int{},
			expectedLen:   0,
		},
		{
			testName:      "Singly linked list with 1 element of type int",
			initialValues: []int{1},
			expectedLen:   1,
		},
		{
			testName:      "Singly linked list with 10 elements of type int",
			initialValues: []int{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedLen:   10,
		},
		{
			testName:      "Singly linked list with 8 elements of type int",
			initialValues: []int{35, 50, 69, 71, 74, 89, 95, 100},
			expectedLen:   8,
		},

		{
			testName:      "Singly linked list with 0 elements of type int32",
			initialValues: []int32{},
			expectedLen:   0,
		},
		{
			testName:      "Singly linked list with 1 element of type int32",
			initialValues: []int32{1},
			expectedLen:   1,
		},
		{
			testName:      "Singly linked list with 10 elements of type int32",
			initialValues: []int32{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedLen:   10,
		},
		{
			testName:      "Singly linked list with 8 elements of type int32",
			initialValues: []int32{35, 50, 69, 71, 74, 89, 95, 100},
			expectedLen:   8,
		},

		{
			testName:      "Singly linked list with 0 elements of type int64",
			initialValues: []int64{},
			expectedLen:   0,
		},
		{
			testName:      "Singly linked list with 1 element of type int64",
			initialValues: []int64{1},
			expectedLen:   1,
		},
		{
			testName:      "Singly linked list with 10 elements of type int64",
			initialValues: []int64{100, 42, 72, 5, 36, 66, 77, 79, 51, 3},
			expectedLen:   10,
		},
		{
			testName:      "Singly linked list with 8 elements of type int32",
			initialValues: []int64{35, 50, 69, 71, 74, 89, 95, 100},
			expectedLen:   8,
		},

		{
			testName:      "Singly linked list with 0 elements of type float32",
			initialValues: []float32{},
			expectedLen:   0,
		},
		{
			testName:      "Singly linked list with 1 element of type float32",
			initialValues: []float32{1.0},
			expectedLen:   1,
		},
		{
			testName:      "Singly linked list with 10 elements of type float32",
			initialValues: []float32{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			expectedLen:   10,
		},
		{
			testName:      "Singly linked list with 8 elements of type float32",
			initialValues: []float32{35.7, 50.4, 69.9, 71.2, 74.8, 89.1, 95.5, 100.3},
			expectedLen:   8,
		},

		{
			testName:      "Singly linked list with 0 elements of type float64",
			initialValues: []float64{},
			expectedLen:   0,
		},
		{
			testName:      "Singly linked list with 1 element of type float64",
			initialValues: []float64{1.0},
			expectedLen:   1,
		},
		{
			testName:      "Singly linked list with 10 elements of type float64",
			initialValues: []float64{100.5, 42.3, 72.7, 5.9, 36.1, 66.6, 77.8, 79.4, 51.2, 3.3},
			expectedLen:   10,
		},
		{
			testName:      "Singly linked list with 8 elements of type float64",
			initialValues: []float64{35.7, 50.4, 69.9, 71.2, 74.8, 89.1, 95.5, 100.3},
			expectedLen:   8,
		},

		{
			testName:      "Singly linked list with 0 elements of type string",
			initialValues: []string{},
			expectedLen:   0,
		},
		{
			testName:      "Singly linked list with 1 element of type string",
			initialValues: []string{"raspberry"},
			expectedLen:   1,
		},
		{
			testName:      "Singly linked list with 8 elements of type string",
			initialValues: []string{"strawberry", "tangerine", "ugly fruit", "valencia orange", "watermelon", "xigua", "yellow watermelon", "zucchini"},
			expectedLen:   8,
		},
		{
			testName:      "Singly linked list with 8 elements of type string",
			initialValues: []string{"blueberry muffin", "cantaloupe", "dragon fruit", "elderflower", "fig newton", "guava", "honeycrisp apple", "indian fig"},
			expectedLen:   8,
		},

		{
			testName:      "Singly linked list with 0 elements of type []int",
			initialValues: [][]int{},
			expectedLen:   0,
		},
		{
			testName: "Singly linked list with 1 element of type []int",
			initialValues: [][]int{
				{1, 2},
			},
			expectedLen: 1,
		},
		{
			testName: "Singly linked list with 2 elements of type []int",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
			},
			expectedLen: 2,
		},
		{
			testName: "Singly linked list with 4 elements of type []int",
			initialValues: [][]int{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedLen: 4,
		},

		{
			testName:      "Singly linked list with 0 elements of type []int32",
			initialValues: [][]int32{},
			expectedLen:   0,
		},
		{
			testName: "Singly linked list with 1 element of type []int32",
			initialValues: [][]int32{
				{1, 2},
			},
			expectedLen: 1,
		},
		{
			testName: "Singly linked list with 2 elements of type []int32",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
			},
			expectedLen: 2,
		},
		{
			testName: "Singly linked list with 4 elements of type []int32",
			initialValues: [][]int32{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedLen: 4,
		},

		{
			testName:      "Singly linked list with 0 elements of type []int64",
			initialValues: [][]int64{},
			expectedLen:   0,
		},
		{
			testName: "Singly linked list with 1 element of type []int64",
			initialValues: [][]int64{
				{1, 2},
			},
			expectedLen: 1,
		},
		{
			testName: "Singly linked list with 2 elements of type []int64",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
			},
			expectedLen: 2,
		},
		{
			testName: "Singly linked list with 4 elements of type []int64",
			initialValues: [][]int64{
				{10, 20},
				{30, 40},
				{50, 60},
				{70, 80},
			},
			expectedLen: 4,
		},

		{
			testName:      "Singly linked list with 0 elements of type []float32",
			initialValues: [][]float32{},
			expectedLen:   0,
		},
		{
			testName: "Singly linked list with 1 element of type []float32",
			initialValues: [][]float32{
				{1, 2},
			},
			expectedLen: 1,
		},
		{
			testName: "Singly linked list with 2 elements of type []float32",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedLen: 2,
		},
		{
			testName: "Singly linked list with 4 elements of type []float32",
			initialValues: [][]float32{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedLen: 4,
		},

		{
			testName:      "Singly linked list with 0 elements of type []float64",
			initialValues: [][]float64{},
			expectedLen:   0,
		},
		{
			testName: "Singly linked list with 1 element of type []float64",
			initialValues: [][]float64{
				{1, 2},
			},
			expectedLen: 1,
		},
		{
			testName: "Singly linked list with 2 elements of type []float64",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
			},
			expectedLen: 2,
		},
		{
			testName: "Singly linked list with 4 elements of type []float64",
			initialValues: [][]float64{
				{10.5, 20.6},
				{30.7, 40.8},
				{50.9, 60.1},
				{70.2, 80.3},
			},
			expectedLen: 4,
		},

		{
			testName:      "Singly linked list with 0 elements of type []string",
			initialValues: [][]string{},
			expectedLen:   0,
		},
		{
			testName: "Singly linked list with 1 element of type []string",
			initialValues: [][]string{
				{"apple"},
			},
			expectedLen: 1,
		},
		{
			testName: "Singly linked list with 2 elements of type []string",
			initialValues: [][]string{
				{"cherry", "date"},
				{"elderberry", "fig"},
			},
			expectedLen: 2,
		},
		{
			testName: "Singly linked list with 4 elements of type []string",
			initialValues: [][]string{
				{"grape", "honeydew"},
				{"kiwi", "lemon"},
				{"apple", "banana"},
				{"cherry", "date"},
			},
			expectedLen: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			switch v := test.initialValues.(type) {
			case []int:
				equals := func(a, b int) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case []int32:
				equals := func(a, b int32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case []int64:
				equals := func(a, b int64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case []float32:
				equals := func(a, b float32) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case []float64:
				equals := func(a, b float64) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case []string:
				equals := func(a, b string) bool { return a == b }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case [][]int:
				equals := func(a, b []int) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case [][]int32:
				equals := func(a, b []int32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case [][]int64:
				equals := func(a, b []int64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case [][]float32:
				equals := func(a, b []float32) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case [][]float64:
				equals := func(a, b []float64) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			case [][]string:
				equals := func(a, b []string) bool { return custom_comparators.EqualsSlice(a, b) }
				list := linked_lists.NewSinglyLinkedList(equals)

				for _, elem := range v {
					list.InsertAtEnd(elem)
				}

				lenOfList := list.Len()

				assert.Equal(t, test.expectedLen, lenOfList)
			}
		})
	}
}
