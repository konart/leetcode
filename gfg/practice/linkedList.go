package practice

import (
	. "github.com/konart/leetcode/gfg/structs"
	"fmt"
)

func LinkedListTraversal() {
	node1 := &Node{1, nil}
	node2 := &Node{2, nil}
	list := &LinkedList{}
	list.SetHead(node1)
	node1.Next = node2
	list.PrintList()
}

func LinkedListInsert() {
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	list := &LinkedList{}
	list.SetHead(node1)
	node1.Next = node2
	list.SetHead(&Node{11, nil})
	list.InsertFirst(11)
	list.InsertAfter(node1,12)
	list.InsertEnd(22)
	list.PrintList()
	fmt.Printf("Length: %d \n", list.Length())
	//list.Delete(22)
	list.DeleteAt(2)
	list.PrintList()
	fmt.Printf("Length: %d \n", list.Length())
	x := 33
	fmt.Printf("Number %d is %t", x, list.Search(x))
}

func CompareLists(a,b *LinkedList) int {
	if a.Length() > b.Length() {
		return 1
	} else if a.Length() < b.Length() {
		return -1
	}
	nodeA := a.GetHead()
	nodeB := b.GetHead()
	for nodeA != nil {
		if nodeA.Data > nodeB.Data {
			return 1
		} else if nodeA.Data < nodeB.Data {
			return -1
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return 0
}

func AddTwoNumbers(a,b *LinkedList) (*LinkedList, error) {
	sumList := &LinkedList{}


	return sumList, nil
}