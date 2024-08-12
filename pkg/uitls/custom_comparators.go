package custom_comparators

import "reflect"

// LessSlice compares two slices and returns true if the first is "less" than the second.
//
// It compares the slices' lengths first; if they differ, it returns true if the first slice is shorter.
//
// If lengths are equal, it compares corresponding elements by type (int, int32, int64, float32, float64, string) and returns true if any element in the first slice is less than the corresponding element in the second slice.
//
// Panics on unsupported types.
func LessSlice(a, b interface{}) bool {
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

// MoreSlice compares two slices and returns true if the first is "greater" than the second.
//
// It compares the slices' lengths first; if they differ, it returns true if the first slice is longer.
//
// If lengths are equal, it compares corresponding elements by type (int, int32, int64, float32, float64, string) and returns true if any element in the first slice is greater than the corresponding element in the second slice.
//
// Panics on unsupported types.
func MoreSlice(a, b interface{}) bool {
	sliceA := reflect.ValueOf(a)
	sliceB := reflect.ValueOf(b)

	if sliceA.Len() != sliceB.Len() {
		return sliceA.Len() > sliceB.Len()
	}

	for i := 0; i < sliceA.Len(); i++ {
		valA := sliceA.Index(i).Interface()
		valB := sliceB.Index(i).Interface()

		switch vA := valA.(type) {
		case int:
			if vB, ok := valB.(int); ok {
				if vA > vB {
					return true
				} else if vA < vB {
					return false
				}
			}
		case int32:
			if vB, ok := valB.(int32); ok {
				if vA > vB {
					return true
				} else if vA < vB {
					return false
				}
			}
		case int64:
			if vB, ok := valB.(int64); ok {
				if vA > vB {
					return true
				} else if vA < vB {
					return false
				}
			}
		case float32:
			if vB, ok := valB.(float32); ok {
				if vA > vB {
					return true
				} else if vA < vB {
					return false
				}
			}
		case float64:
			if vB, ok := valB.(float64); ok {
				if vA > vB {
					return true
				} else if vA < vB {
					return false
				}
			}
		case string:
			if vB, ok := valB.(string); ok {
				if vA > vB {
					return true
				} else if vA < vB {
					return false
				}
			}
		default:
			panic("unsupported type")
		}
	}

	return false
}

// EqualsSlice compares two slices and returns true if they are identical.
//
// It first checks if the slices have the same length. If not, it returns false. If they do,
//
// it compares corresponding elements by type (int, int32, int64, float32, float64, string) for equality.
//
// Returns false if any pair of elements differs or if the types are unsupported.
//
// Panics on unsupported types.
func EqualsSlice(a, b interface{}) bool {
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
