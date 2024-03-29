package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	// https://stackoverflow.com/a/1597641

	len1 := countNodes(head1)
	len2 := countNodes(head2)

	curr1 := head1
	curr2 := head2

	for len1 > len2 {
		curr1 = curr1.next
		len1--
	}
	for len2 > len1 {
		curr2 = curr2.next
		len2--
	}
	for curr1 != nil {
		if curr1 == curr2 {
			return curr1.data
		}
		curr1 = curr1.next
		curr2 = curr2.next
	}
	return -1
}

func countNodes(head *SinglyLinkedListNode) int {
	curr := head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
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
