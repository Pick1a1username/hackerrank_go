package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	// Write your code here
	newNode := &SinglyLinkedListNode{
		data: data,
	}

	if position == 0 {
		newNode.next = llist
		return newNode
	}

	curr := llist
	for i := int32(1); i < position; i++ {
		curr = curr.next
	}

	newNode.next = curr.next
	curr.next = newNode
	return llist

}
