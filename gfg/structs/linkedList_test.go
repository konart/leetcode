package structs

import (
	"testing"
)

func TestGenerateFromSlice(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{},
	}

	for c, tc := range testCases {
		list := GenerateFromSlice(tc)
		node := list.head
		if list == nil {
			t.Fatalf("Test case %d: linked list was not created!", c+1)
		}
		if node == nil && len(tc) > 0 {
			t.Fatalf("Test case %d: no head node found, but test case is not empty!", c+1)
		}
		for i, e := range tc {
			if i < len(tc) - 1 && node.Next == nil {
				t.Fatalf("Test case %d: linked list length is less than that of a test case!", c+1)
			}
			if i == len(tc) - 1 && node.Next != nil {
				t.Fatalf("Test case %d: linked list length is greater than that of a test case!", c+1)
			}
			if node.Data != e {
				t.Errorf("Test case %d: expected: %d, got %d instead", c+1, e, node.Data)
			}
			node = node.Next
		}
	}
}

func TestLinkedList_Swap(t *testing.T) {
	testCases := [][][]int{
		{{1, 2, 3, 4, 5}, {3, 2, 1, 4, 5}}, //  1 is a head
		{{2, 1, 3, 4, 5}, {2, 3, 1, 4, 5}}, //  1 and 3 are adjacent
		{{2, 1, 4, 3, 5}, {2, 3, 4, 1, 5}}, //  1 and 3 are not adjacent
		{{1, 2, 4, 5, 6}, {1, 2, 4, 5, 6}}, //  1 is a head and there is no 3
		{{3, 4, 5, 6, 7}, {3, 4, 5, 6, 7}}, //  3 is a head and there is no 1
		{{1, 2, 4, 5, 3}, {3, 2, 4, 5, 1}}, //  1 and 3 are head and last elements
		{{2, 1, 4, 5, 3}, {2, 3, 4, 5, 1}}, //  3 is the last element
		{{2, 4, 3, 5, 1}, {2, 4, 1, 5, 3}}, //  1 is the last element
		{{2, 0, 4, 5, 6}, {2, 0, 4, 5, 6}}, //  neither 1 or 3 are in the slice
	}

	for c, tc := range testCases {
		list := GenerateFromSlice(tc[0])
		list.Swap(1, 3)
		node := list.head
		for _, e := range tc[1] {
			if node.Data != e {
				t.Errorf("Test case %d: expected: %d, got %d instead", c+1, e, node.Data)
			}
			node = node.Next
		}
	}
}

func TestLinkedList_Destroy(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{},
	}

	for c, tc := range testCases {
		list := GenerateFromSlice(tc)
		node := list.head
		var arr []*Node
		for {
			if node == nil {
				break
			}
			arr = append(arr, node)
			node = node.Next
		}
		list.Destroy()
		for _, e := range arr {
			if e.Next != nil {
				t.Errorf("Test case %d: expected: Next == nil", c+1)
			}
			if e.Data != 0 {
				t.Errorf("Test case %d: expected: Data == 0, got %d instead", c+1, e.Data)
			}
		}
	}
}
