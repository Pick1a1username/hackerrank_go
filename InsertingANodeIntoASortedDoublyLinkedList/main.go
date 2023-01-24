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
	currNode := llist
	for {
		if currNode == nil {
			newNode := DoublyLinkedListNode{
				data: data,
				prev: nil,
				next: nil,
			}
			return &newNode
		}
		if currNode.prev == nil && data < currNode.data {
			newNode := DoublyLinkedListNode{
				data: data,
				prev: nil,
				next: currNode,
			}
			currNode.prev = &newNode
			return &newNode
		}
		if currNode.next == nil && currNode.data <= data {
			newNode := DoublyLinkedListNode{
				data: data,
				prev: currNode,
				next: nil,
			}
			currNode.next = &newNode
			return llist
		}
		if data >= currNode.data && data < currNode.next.data {
			newNode := DoublyLinkedListNode{
				data: data,
				prev: currNode,
				next: currNode.next,
			}
			currNode.next = &newNode
			currNode.next.prev = &newNode
			return llist
		}
		currNode = currNode.next
	}
}

func main() {}
