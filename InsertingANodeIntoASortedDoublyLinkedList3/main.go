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
	newNode := &DoublyLinkedListNode{}
	newNode.data = data

	if llist == nil {
		return newNode
	}
	if data < llist.data {
		newNode.next = llist
		llist.prev = newNode
		return newNode
	}

	var n1 *DoublyLinkedListNode = nil
	var n2 *DoublyLinkedListNode = llist
	for n2 != nil && n2.data < data {
		n1 = n2
		n2 = n2.next
	}
	if n2 == nil {
		n1.next = newNode
		newNode.prev = n1
	} else {
		n1.next = newNode
		n2.prev = newNode
		newNode.prev = n1
		newNode.next = n2
	}
	return llist
}

func main() {}
