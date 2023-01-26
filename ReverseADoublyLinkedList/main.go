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

/*
 * Complete the 'reverse' function below.
 *
 * The function is expected to return an INTEGER_DOUBLY_LINKED_LIST.
 * The function accepts INTEGER_DOUBLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
	// Write your code here
	// I have no idea why I wrote something like this.
	if llist == nil {
		return nil
	}
	// Get the length of the list.
	llistLen := lenLinkedList(llist)

	// Bubble sort
	currHead := llist
	for i := 0; i < llistLen; i++ {
		currNode := currHead
		for j := 0; j < llistLen-1; j++ {
			// if the current node is bigger than the next node.
			if currNode.data > currNode.next.data {
				swapSelfNext(currNode, currNode.next)
			}
			if currNode.prev == nil {
				currHead = currNode
			}
			currNode = currNode.next
		}
	}

	return currHead
}

func swapSelfNext(s, n *DoublyLinkedListNode) {
	if s == nil || n == nil {
		panic("parameters should not be nil")
	}
	// swap s and n nodes.
	s.data, n.data = n.data, s.data
	s.prev, n.prev = n.prev, s.prev
	s.next, n.next = n.next, s.next
}

func lenLinkedList(llist *DoublyLinkedListNode) int {
	// count nodes.
	count := 0
	currNode := llist
	for currNode != nil {
		count += 1
		currNode = currNode.next
	}
	return count
}
