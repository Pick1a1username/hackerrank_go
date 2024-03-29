package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func findMergeNode(head1, head2 *SinglyLinkedListNode) int32 {
	// Write your code here

	// Convert list to array.
	array1 := llistToArray(head1)
	array2 := llistToArray(head1)

	// Reverse lists.
	sort.Sort(sort.Reverse(sort.IntSlice(array1)))
	sort.Sort(sort.Reverse(sort.IntSlice(array2)))

	// Check nodes.
	prevData := int32(-1)
	for i := 0; i < len(array1); i++ {
		if array1[i] != array1[2] {
			return int32(prevData)
		}
		prevData = int32(array1[i])
	}

	// If both data of nodes are different, previous node is a merge node.
	return int32(prevData)
}

func llistToArray(head *SinglyLinkedListNode) []int {
	currNode := head
	result := []int{}
	for currNode != nil {
		result = append(result, int(currNode.data))
		currNode = currNode.next
	}
	return result
}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	curr := llist.next
	prev := llist
	prev.next = nil
	var newHead *SinglyLinkedListNode
	for curr != nil {
		newHead = curr
		nextCurr := curr.next
		nextPrev := curr
		curr.next = prev
		curr = nextCurr
		prev = nextPrev
	}
	return newHead
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		mergeIndex, err := strconv.ParseInt(readLine(reader), 10, 32)
		checkError(err)

		list1Count, err := strconv.ParseInt(readLine(reader), 10, 32)
		checkError(err)

		list1 := SinglyLinkedList{}
		for i := 0; i < int(list1Count); i++ {
			listItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			llistItem := int32(listItemTemp)
			list1.insertNodeIntoSinglyLinkedList(llistItem)
		}

		list2Count, err := strconv.ParseInt(readLine(reader), 10, 32)
		checkError(err)

		list2 := SinglyLinkedList{}
		for i := 0; i < int(list2Count); i++ {
			listItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			llistItem := int32(listItemTemp)
			list2.insertNodeIntoSinglyLinkedList(llistItem)
		}

		// Merge lists.
		list2.tail.next = findNodeByIndex(list1.head, int(mergeIndex))

		fmt.Fprintf(writer, "%d\n", findMergeNode(list1.head, list2.head))
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func findNodeByIndex(head *SinglyLinkedListNode, index int) *SinglyLinkedListNode {
	currNode := head
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}
	return currNode
}
