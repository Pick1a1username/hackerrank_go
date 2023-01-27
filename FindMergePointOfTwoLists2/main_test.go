package main

import (
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

func genLlist(dataArr []int32) *SinglyLinkedList {
	singlyLinkedList := SinglyLinkedList{}

	for _, v := range dataArr {
		singlyLinkedList.insertNodeIntoSinglyLinkedList(v)
	}
	return &singlyLinkedList
}
