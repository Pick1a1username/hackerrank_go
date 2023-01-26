package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_reverse(t *testing.T) {

	testCases := []struct {
		description string
		llist       *DoublyLinkedListNode
		answer      *DoublyLinkedListNode
	}{
		{
			description: "TestCase 1",
			llist:       genLlist([]int32{1, 2, 3, 4}),
			answer:      genLlist([]int32{4, 3, 2, 1}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := reverse(testCase.llist)
			if !equalDoublyLinkedList(result, testCase.answer) {
				t.Errorf("result: [ %s ] answer: [ %s ]", llistToString(result), llistToString(testCase.answer))
			}
		})
	}
}

func equalDoublyLinkedList(a, b *DoublyLinkedListNode) bool {
	currA := a
	currB := b
	for {
		if currA == nil && currB == nil {
			return true
		}
		if (currA == nil && currB != nil) || (currA != nil && currB == nil) {
			return false
		}
		if currA.data != currB.data {
			return false
		}
		currA = currA.next
		currB = currB.next

	}
}

func genLlist(dataArr []int32) *DoublyLinkedListNode {
	doublyLinkedList := DoublyLinkedList{}

	for _, v := range dataArr {
		doublyLinkedList.insertNodeIntoDoublyLinkedList(v)
	}
	return doublyLinkedList.head
}

func llistToString(llist *DoublyLinkedListNode) string {
	delimiter := " "
	stringList := ""
	currNode := llist
	for currNode != nil {
		stringList += fmt.Sprintf("%d", currNode.data) + delimiter
		currNode = currNode.next
	}
	stringList = strings.TrimSuffix(stringList, delimiter)
	return stringList
}
