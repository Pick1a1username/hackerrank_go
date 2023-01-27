package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_findMergeNode(t *testing.T) {
	testCases := []struct {
		description string
		list1       *SinglyLinkedList
		list2       *SinglyLinkedList
		mergeIndex  int32
		answer      int32
	}{
		{
			description: "TestCase 1",
			list1:       genLlist([]int32{1, 2, 3}),
			list2:       genLlist([]int32{1}),
			mergeIndex:  1,
			answer:      2,
		},
		{
			description: "TestCase 2",
			list1:       genLlist([]int32{1, 2, 3}),
			list2:       genLlist([]int32{1}),
			mergeIndex:  2,
			answer:      3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			testCase.list2.tail.next = findNodeByIndex(testCase.list1.head, int(testCase.mergeIndex))
			result := findMergeNode(testCase.list1.head, testCase.list2.head)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}

func Test_reverse(t *testing.T) {

	testCases := []struct {
		description string
		llist       *SinglyLinkedListNode
		answer      *SinglyLinkedListNode
	}{
		{
			description: "TestCase 1",
			llist:       genLlist([]int32{1, 2, 3, 4}).head,
			answer:      genLlist([]int32{4, 3, 2, 1}).head,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := reverse(testCase.llist)
			if !equalSinglyLinkedList(result, testCase.answer) {
				t.Errorf("result: [ %s ] answer: [ %s ]", llistToString(result), llistToString(testCase.answer))
			}
		})
	}
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

func genLlist(dataArr []int32) *SinglyLinkedList {
	singlyLinkedList := SinglyLinkedList{}

	for _, v := range dataArr {
		singlyLinkedList.insertNodeIntoSinglyLinkedList(v)
	}
	return &singlyLinkedList
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
