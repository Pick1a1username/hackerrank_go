package main

import (
	"testing"
)

func Test_sortedInsert(t *testing.T) {

	testCases := []struct {
		description string
		llist       *DoublyLinkedListNode
		data        int32
		answer      *DoublyLinkedListNode
	}{
		{
			description: "TestCase 1",
			llist:       genLlist([]int32{3, 5, 7}),
			data:        1,
			answer:      genLlist([]int32{1, 3, 5, 7}),
		},
		{
			description: "TestCase 2",
			llist:       genLlist([]int32{3, 5, 7}),
			data:        4,
			answer:      genLlist([]int32{3, 4, 5, 7}),
		},
		{
			description: "TestCase 3",
			llist:       genLlist([]int32{3, 5, 7}),
			data:        6,
			answer:      genLlist([]int32{3, 5, 6, 7}),
		},
		{
			description: "TestCase 4",
			llist:       genLlist([]int32{3, 5, 7}),
			data:        8,
			answer:      genLlist([]int32{3, 5, 7, 8}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := sortedInsert(testCase.llist, testCase.data)
			if !equalDoublyLinkedList(result, testCase.answer) {
				t.Errorf("result: %v answer: %v", result, testCase.answer)
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
