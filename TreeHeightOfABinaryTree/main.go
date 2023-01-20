package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Data  int32
	Left  *Node
	Right *Node
}

func getHeight(root *Node) int32 {
	// If root is nil, return -1
	if root == nil {
		return -1
	}
	// Get left height.
	leftHeight := 1 + getHeight(root.Left)
	// Get right height.
	rightHeight := 1 + getHeight(root.Right)
	// return higher value.
	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

func insert(root *Node, data int32) *Node {
	if root == nil {
		return &Node{
			Data:  data,
			Left:  nil,
			Right: nil,
		}
	}

	if data <= root.Data {
		root.Left = insert(root.Left, data)
	} else {
		root.Right = insert(root.Right, data)
	}
	return root
}

func genBinaryTree(arr []int32) *Node {
	root := Node{Data: arr[0]}

	for i := 1; i < len(arr); i++ {
		insert(&root, arr[i])
	}

	return &root
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	root := genBinaryTree(arr)
	result := getHeight(root)

	fmt.Fprintf(writer, "%d\n", result)

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
