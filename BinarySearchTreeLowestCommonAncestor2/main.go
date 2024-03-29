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

func lca(root *Node, v1, v2 int32) int32 {
	if v1 > v2 {
		v1, v2 = v2, v1
	}

	for {
		if v1 < root.Data && v2 < root.Data {
			root = root.Left
		} else if v1 > root.Data && v2 > root.Data {
			root = root.Right
		} else {
			return root.Data
		}
	}

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

	root := genBinaryTree(arr)
	result := lca(root, v1, v2)

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
