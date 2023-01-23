package main

import (
	"reflect"
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
			llist:       genLlist([]int32{1, 2, 4}),
			data:        3,
			answer:      genLlist([]int32{1, 2, 3, 4}),
		},
		{
			description: "TestCase 1",
			llist:       genLlist([]int32{1, 2, 3, 10}),
			data:        5,
			answer:      genLlist([]int32{1, 2, 3, 5, 10}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := sortedInsert(testCase.llist, testCase.data)
			if !reflect.DeepEqual(result, testCase.answer) {
				t.Errorf("result: %v answer: %v", result, testCase.answer)
			}
		})
	}
}

func genLlist(dataArr []int32) *DoublyLinkedListNode {
	doublyLinkedList := DoublyLinkedList{}

	for _, v := range dataArr {
		doublyLinkedList.insertNodeIntoDoublyLinkedList(v)
	}
	return doublyLinkedList.head
}
