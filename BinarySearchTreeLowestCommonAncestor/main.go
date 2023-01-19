package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Data  int32
	Left  *Node
	Right *Node
}

func insert(root *Node, data int32) *Node {
	if root == nil {
		return &Node{
			Data:  data,
			Left:  nil,
			Right: nil,
		}
	}

	var cur *Node

	if data <= root.Data {
		cur = insert(root.Left, data)
		root.Left = cur
	} else {
		cur = insert(root.Right, data)
		root.Right = cur
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

// If the last element is not target data,
// it means target data don't exist.
func getParentData(root *Node, data int32) []int32 {
	// Find the node of data while collecting node data.
	result := []int32{root.Data}
	if root.Data == data {
		return result
	}
	if root.Left != nil && data <= root.Data {
		return append(result, getParentData(root.Left, data)...)
	}
	if root.Right != nil && data > root.Data {
		return append(result, getParentData(root.Right, data)...)
	}
	return result
}

func lca(root *Node, v1, v2 int32) int32 {
	// Get parent nodes.
	v1Parents := getParentData(root, v1)
	v2Parents := getParentData(root, v2)

	// Get common nodes.
	sort.Slice(v1Parents, func(i, j int) bool { return v1Parents[i] > v1Parents[j] })
	sort.Slice(v2Parents, func(i, j int) bool { return v2Parents[i] > v2Parents[j] })

	// Get the minimum node.
	for i := 0; i < len(v1Parents); i++ {
		_, err := findIdxByValInt32(&v2Parents, v1Parents[i])
		if err == nil {
			return v1Parents[i]
		}
	}
	// Return.
	return -1
}

func findIdxByValInt32(slPtr *[]int32, v int32) (int32, error) {
	for i := 0; i < len(*slPtr); i++ {
		if (*slPtr)[i] == v {
			return int32(i), nil
		}
	}
	return -1, errors.New("not found")
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

	vTemps := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	v1Tmp, err := strconv.ParseInt(vTemps[0], 10, 64)
	checkError(err)
	v1 := int32(v1Tmp)
	v2Tmp, err := strconv.ParseInt(vTemps[1], 10, 64)
	checkError(err)
	v2 := int32(v2Tmp)

	root := Node{
		Data: arr[0],
	}
	for i := 1; i < int(n); i++ {
		insert(&root, arr[i])
	}
	result := lca(&root, v1, v2)

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
