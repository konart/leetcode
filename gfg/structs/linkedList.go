package structs

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) SetHead(node *Node) {
	if list.head == nil {
		list.head = node
	} else {
		temp := list.head
		list.head = node
		list.head.next = temp
	}

}

func (list LinkedList) GetHead(node *Node) *Node {
	return list.head
}