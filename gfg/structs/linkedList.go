package structs

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (list LinkedList) Length() int {
	return list.length
}

func GenerateFromSlice(s []int) *LinkedList {
	list := &LinkedList{}
	for _, d := range s {
		list.InsertEnd(d)
	}
	return list
}

func (list *LinkedList) SetHead(node *Node) {
	if list.head == nil {
		list.head = node
		list.length = 1
		for {
			list.length++
			if node.Next == nil {
				break
			}
			node = node.Next
		}
	} else {
		node.Next = list.head
		list.head = node
		list.length++
	}
}

func (list *LinkedList) InsertFirst(i int) {
	node := &Node{Data: i}
	node.Next = list.head
	list.head = node
	list.length++
}

func (list *LinkedList) InsertAfter(prev *Node, i int) {
	node := &Node{Data: i}
	node.Next = prev.Next
	prev.Next = node
	list.length++
}

func (list *LinkedList) InsertEnd(i int) {
	n := &Node{Data: i}
	if list.head == nil {
		list.head = n
		list.length++
		return
	}
	node := list.GetHead()
	for {
		if node.Next == nil {
			node.Next = n
			list.length++
			break
		} else {
			node = node.Next
		}
	}

}

func (list LinkedList) GetHead() *Node {
	return list.head
}

func (list *LinkedList) Delete(i int) {
	node := list.GetHead()
	temp := node
	for {
		if node != nil && node.Data == i {
			if list.head == node {
				list.head = node.Next
			} else {
				temp.Next = node.Next
			}
			list.length--
			break
		}
		temp = node
		node = node.Next
	}
}

func (list *LinkedList) DeleteAt(i int) {
	if i == 0 {
		list.head = list.head.Next
		list.length--
	} else {
		node := list.head.Next
		temp := list.head
		for x := 1; x <= i; x++ {
			if x == i {
				temp.Next = node.Next
				list.length--
			}
			temp = node
			node = node.Next
		}
	}
}

func (list LinkedList) PrintList() {
	node := list.GetHead()
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

func (list LinkedList) Search(i int) bool {
	switch {
	case list.head == nil:
		return false
	case list.head.Data == i:
		return true
	default:
		node := list.head
		for {
			if node == nil {
				return false
			} else if node.Data == i {
				return true
			} else {
				node = node.Next
			}
		}
	}
	return false
}

func (list *LinkedList) Swap(i, j int) {
	if i == j {
		return
	}
	node := list.head
	var iTemp, jTemp, iNode, jNode *Node

	for {
		if node == nil {
			break
		}
		if node.Data == i {
			iNode = node
		}
		if node.Data == j {
			jNode = node
		}
		if iNode == nil {
			iTemp = node
		}
		if jNode == nil {
			jTemp = node
		}
		node = node.Next
	}
	if iNode == nil || jNode == nil {
		return
	}
	if iTemp != nil {
		iTemp.Next = jNode
	} else {
		list.head = jNode
	}
	if jTemp != nil {
		jTemp.Next = iNode
	} else {
		list.head = iNode
	}
	temp := iNode.Next
	iNode.Next = jNode.Next
	jNode.Next = temp
}

func (list *LinkedList) Destroy() {

}