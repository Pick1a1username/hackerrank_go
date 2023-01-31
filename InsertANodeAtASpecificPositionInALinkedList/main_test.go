package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_insertNodeAtPosition(t *testing.T) {
	testCases := []struct {
		description string
		llist       *SinglyLinkedListNode
		data        int32
		position    int32
		answer      *SinglyLinkedListNode
	}{
		{
			description: "TestCase 1",
			llist:       genLlist([]int32{16, 13, 7}).head,
			data:        1,
			position:    2,
			answer:      genLlist([]int32{16, 13, 1, 7}).head,
		},
		{
			description: "TestCase 2",
			llist:       genLlist([]int32{11, 9, 19, 10, 4}).head,
			data:        20,
			position:    3,
			answer:      genLlist([]int32{11, 9, 19, 20, 10, 4}).head,
		},
		{
			description: "TestCase 3",
			llist:       genLlist([]int32{16, 13, 7}).head,
			data:        1,
			position:    0,
			answer:      genLlist([]int32{1, 16, 13, 7}).head,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := insertNodeAtPosition(testCase.llist, testCase.data, testCase.position)
			if !equalSinglyLinkedList(result, testCase.answer) {
				t.Errorf("result: %s answer: %s", llistToString(result), llistToString(testCase.answer))
			}
		})
	}
}

func genLlist(dataArr []int32) *SinglyLinkedList {
	singlyLinkedList := SinglyLinkedList{}

	for _, v := range dataArr {
		singlyLinkedList.insertNodeIntoSinglyLinkedList(v)
	}
	return &singlyLinkedList
}

func equalSinglyLinkedList(a, b *SinglyLinkedListNode) bool {
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

func llistToString(llist *SinglyLinkedListNode) string {
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
