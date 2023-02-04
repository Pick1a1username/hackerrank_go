package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

const (
	MaxBribe = 2
)

func minimumBribes(q []int32) {
	// Write your code here
	num := minimumBribesImpl(q)
	if num < 0 {
		fmt.Println("Too chaotic")
	} else {
		fmt.Printf("%d\n", num)
	}
}

func minimumBribesImpl(q []int32) int {
	allCount := 0

	// read from end.
	for i := int32(len(q)) - 1; i >= 0; i-- {
		// if the person's not supposed to be here, try to make him back.
		if q[i] != i+1 {
			currCount := 0
			for {
				// check whether there's a person behind.
				if i < int32(len(q))-1 {
					// if there's a person behind, go behind of him.
					err := getBack(&q, int(i))

					// if the person is at the end of the queue, stop moving.
					if err != nil {
						break
					}
					currCount++
					// if the person have to move more than twice, it's too chaotic.
					if currCount > MaxBribe {
						return -1
					}
				} else {
					// if the person at the end of queue, give up.
					break
				}
				// if the person is at the right position, get out of this loop.
				if q[i] == i+1 {
					break
				}
			}
			allCount += currCount
		}
	}

	// return
	return allCount
}

func getBack(q *[]int32, p int) error {
	if p >= len(*q)-1 {
		return errors.New("cannot get back anymore")
	}
	(*q)[p], (*q)[p+1] = (*q)[p+1], (*q)[p]
	return nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
