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
		}
	}{
		{
			intSlice: []int{1, 2, 5, 4},
			expected: struct {
				slice     []int
				swapCount int
			}{slice: []int{1, 2, 5}, swapCount: 1},
		},
		{
			intSlice: []int{4, 3, 2, 1},
			expected: struct {
				slice     []int
				swapCount int
			}{slice: []int{1, 2, 3, 4}, swapCount: 0},
		},
		{
			intSlice: []int{1, 2, 1, 2},
			expected: struct {
				slice     []int
				swapCount int
			}{slice: []int{1, 2, 2}, swapCount: 1},
		},
		{
			intSlice: []int{87, 90, 92, 95, 96, 93},
			expected: struct {
				slice     []int
				swapCount int
			}{slice: []int{87, 90, 92, 95, 96}, swapCount: 1},
		},
	}

	for _, v := range cases {

		actualSlice, swap := sortFinalCount(v.intSlice)

		if !reflect.DeepEqual(actualSlice, v.expected.slice) || swap != v.expected.swapCount {
			t.Errorf("Actual doesn't match expected \nActual:\t%v\t%d\nExpected: %v\t%d\n", actualSlice, swap, v.expected.slice, v.expected.swapCount)
		}

	}

}

func TestCopySlice(t *testing.T) {
	cases := []struct {
		intSlice []int
		expected []int
	}{
		{
			intSlice: []int{1, 2, 4},
			expected: []int{1, 2, 4},
		},
	}

	for _, v := range cases {
		actual := copySlice(v.intSlice)
		acptr := &actual
		exptr := &v.expected

		if !reflect.DeepEqual(actual, v.expected) && &acptr == &exptr {
			t.Errorf("Not a deep copy\nActual:\t%v\t%v\nExpected:\t%v\t%v", actual, &acptr, v.expected, &exptr)
		}
	}
}
