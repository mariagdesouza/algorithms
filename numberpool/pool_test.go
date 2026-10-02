package main

import (
	"reflect"
	"testing"
)

func TestAllocate(t *testing.T) {

	testAllocateTable := []struct {
		allocated []int
		expected  int
	}{
		{[]int{5, 3, 1}, 2},
		{[]int{3, 2, 1}, 4},
		{[]int{}, 1},
		{[]int{6, 5, 4, 2}, 1},
	}

	for _, test := range testAllocateTable {
		result := allocate(test.allocated)
		if test.expected != result {
			t.Error("function allocate returned incorrect result. Expected:", test.expected, "Got:", result)
		}
	}

}

func TestDeallocate(t *testing.T) {

	testDeallocateTable := []struct {
		allocated       []int
		removeservernum int
		expected        []int
	}{
		{[]int{5, 3, 2, 1}, 2, []int{5, 3, 1}},
		{[]int{5, 3, 2, 1}, 1, []int{5, 3, 2}},
		{[]int{3, 2, 1}, 4, []int{3, 2, 1}},
		{[]int{}, 1, []int{}},
		{[]int{6, 5, 4, 2}, 6, []int{5, 4, 2}},
	}

	for _, test := range testDeallocateTable {
		result := deallocate(test.allocated, test.removeservernum)
		if reflect.DeepEqual(test.expected, result) == false {
			t.Error("function deallocate returned incorrect result. Expected:", test.expected, "Got:", result)
		}
	}

}
