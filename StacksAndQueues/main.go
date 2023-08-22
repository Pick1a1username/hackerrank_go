package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Enqueue       int32 = 1
	Dequeue       int32 = 2
	GetFrontValue int32 = 3
)

type Query struct {
	Type  int32
	Value int64
}

type MyQueue struct {
	Queue []int64
}

func (q *MyQueue) Enqueue(v int64) {
	q.Queue = append([]int64{v}, q.Queue...)
}

func (q *MyQueue) Dequeue() {
	q.Queue = q.Queue[0 : len(q.Queue)-1]
}

func (q *MyQueue) GetFrontValue() int64 {
	return q.Queue[len(q.Queue)-1]
}

func processQueries(queries []Query) []int64 {
	myQueue := MyQueue{}
	frontVals := []int64{}

	for _, q := range queries {
		switch q.Type {
		case Enqueue:
			myQueue.Enqueue(q.Value)
		case Dequeue:
			myQueue.Dequeue()
		case GetFrontValue:
			frontVals = append(frontVals, myQueue.GetFrontValue())
		}
	}

	return frontVals
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []Query

	for i := 0; i < int(n); i++ {
		lineTemp := strings.TrimSpace(readLine(reader))
		lineTempSplit := strings.Fields(lineTemp)
		typeTemp, err := strconv.ParseInt(strings.TrimSpace(lineTempSplit[0]), 10, 64)
		checkError(err)
		queryTemp := Query{Type: int32(typeTemp)}
		if len(lineTempSplit) == 2 {
			valTemp, err := strconv.ParseInt(strings.TrimSpace(lineTempSplit[1]), 10, 64)
			checkError(err)
			queryTemp.Value = valTemp
		}

		arr = append(arr, queryTemp)
	}

	result := processQueries(arr)

	for _, v := range result {
		fmt.Fprintf(writer, "%d\n", v)
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
