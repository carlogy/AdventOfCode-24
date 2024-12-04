package day2

import (
	"reflect"
	"testing"
)

func TestSortFunc(t *testing.T) {
	cases := []struct {
		intSlice []int
		expected struct {
			slice     []int
			swapCount int
			dupExists bool
		}
	}{
		{
			intSlice: []int{1, 2, 5, 4},
			expected: struct {
				slice     []int
				swapCount int
				dupExists bool
			}{slice: []int{1, 2, 4, 5}, swapCount: 1, dupExists: false},
		},
		{
			intSlice: []int{4, 3, 2, 1},
			expected: struct {
				slice     []int
				swapCount int
				dupExists bool
			}{slice: []int{1, 2, 3, 4}, swapCount: 3, dupExists: false},
		},
	}

	for _, v := range cases {

		actualSlice, swap, dup := sortFinalCount(v.intSlice)

		if !reflect.DeepEqual(actualSlice, v.expected.slice) || swap != v.expected.swapCount || dup != v.expected.dupExists {
			t.Errorf("Actual doesn't match expected \nActual: %v\t%d\t%t\nExpected: %v\t%d\t%t\n", actualSlice, swap, dup, v.expected.slice, v.expected.swapCount, v.expected.dupExists)
		}

	}

}
