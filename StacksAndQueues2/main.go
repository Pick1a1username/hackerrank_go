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
	Put  int32 = 1
	Pop  int32 = 2
	Peek int32 = 3
)

type Query struct {
	Type  int32
	Value int64
}

// https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks/editorial?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=stacks-queues
type MyQueue struct {
	Fifo []int64
	Lifo []int64
}

func (q *MyQueue) replenish() {
	if len(q.Fifo) == 0 {
		for i := len(q.Lifo) - 1; i >= 0; i-- {
			q.Fifo = append(q.Fifo, q.Lifo[i])
		}
		q.Lifo = []int64{}
	}
}

func (q *MyQueue) Peek() int64 {
	q.replenish()
	return q.Fifo[len(q.Fifo)-1]
}

func (q *MyQueue) Pop() {
	q.replenish()
	q.Fifo = q.Fifo[0 : len(q.Fifo)-1]
}

func (q *MyQueue) Put(v int64) {
	q.Lifo = append(q.Lifo, v)
}

func processQueries(queries []Query) []int64 {
	myQueue := MyQueue{}
	frontVals := []int64{}

	for _, q := range queries {
		switch q.Type {
		case Put:
			myQueue.Put(q.Value)
		case Pop:
			myQueue.Pop()
		case Peek:
			frontVals = append(frontVals, myQueue.Peek())
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
