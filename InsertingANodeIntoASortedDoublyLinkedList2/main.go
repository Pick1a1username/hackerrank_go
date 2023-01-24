package main

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
	node := &DoublyLinkedListNode{
		next: nil,
		prev: nil,
		data: nodeData,
	}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = node
	} else {
		doublyLinkedList.tail.next = node
		node.prev = doublyLinkedList.tail
	}

	doublyLinkedList.tail = node
}

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	// Write your code here
	n := &DoublyLinkedListNode{}
	n.data = data
	if llist == nil {
		return n
	}
	if data <= llist.data {
		n.next = llist
		llist.prev = n
		return n
	}
	rest := sortedInsert(llist.next, data)
	llist.next = rest
	rest.prev = llist
	return llist
}

func main() {}
