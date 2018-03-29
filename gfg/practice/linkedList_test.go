package practice

import (
	. "github.com/konart/leetcode/gfg/structs"
	"testing"
)



func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		x, y, sum *LinkedList
	}{
		{GenerateFromSlice([]int{5,6,3}),
		GenerateFromSlice([]int{5,6,3}),
		GenerateFromSlice([]int{1,4,0,5}),
		},
	}

	for c, tc := range testCases {
		result, err := AddTwoNumbers(tc.x, tc.y)
		if err != nil {
			t.Fatalf("Adding two number returned an error: %s", err)
		}
		if CompareLists(result, tc.sum) != 0 {
			t.Errorf("Test case %d: expected %s, got %s instead", c, tc.sum, result)
		}
	}
}

func TestCompareLists(t *testing.T) {
	testCases := []struct {
		x, y *LinkedList
		expected int
	}{
		{GenerateFromSlice([]int{5,6,1}),
			GenerateFromSlice([]int{5,6,3}),
			-1,
		},
		{GenerateFromSlice([]int{5,6,3}),
			GenerateFromSlice([]int{5,6,1}),
			1,
		},
		{GenerateFromSlice([]int{5,6,1,1}),
			GenerateFromSlice([]int{5,6,3}),
			1,
		},
		{GenerateFromSlice([]int{5,6,3}),
			GenerateFromSlice([]int{5,6,1,1}),
			-1,
		},
		{GenerateFromSlice([]int{5,6,3}),
			GenerateFromSlice([]int{5,6,3}),
			0,
		},
	}

	for c, tc := range testCases {
		result := CompareLists(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Test case %d: expectd %v, got %v instead", c, tc.expected, result)
		}
	}
}