package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func checkMagazine(magazine []string, note []string) {
	// Write your code here
	fmt.Printf("%s\n", checkMagazineImpl(magazine, note))
}

func checkMagazineImpl(magazine []string, note []string) string {
	// put words in magazine into a map.
	words := map[string]int{}
	for _, w := range magazine {
		words[w] += 1
	}

	// find words for note.
	for _, nw := range note {
		// if c, ok := words[nw]; ok {
		// 	if c == 0 {
		// 		return "No"
		// 	}
		// 	words[nw] -= 1
		// } else {
		// 	return "No"
		// }
		if words[nw] == 0 {
			return "No"
		} else {
			words[nw] -= 1
		}
	}

	return "Yes"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
