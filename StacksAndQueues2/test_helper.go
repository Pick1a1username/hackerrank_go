package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readTestDataInput(testDataPath string) (n int32, arr []Query, err error) {
	f, err := os.Open(testDataPath)
	defer f.Close()
	if err != nil {
		return 0, nil, err
	}
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	nTemp, err := strconv.ParseInt(strings.TrimSpace(fileScanner.Text()), 10, 64)
	if err != nil {
		return 0, nil, err
	}
	n = int32(nTemp)

	for i := 0; i < int(n); i++ {
		fileScanner.Scan()
		lineTemp := strings.TrimSpace(fileScanner.Text())
		lineTempSplit := strings.Fields(lineTemp)
		typeTemp, err := strconv.ParseInt(strings.TrimSpace(lineTempSplit[0]), 10, 64)
		if err != nil {
			return 0, nil, err
		}
		queryTemp := Query{Type: int32(typeTemp)}
		if len(lineTempSplit) == 2 {
			valTemp, err := strconv.ParseInt(strings.TrimSpace(lineTempSplit[1]), 10, 64)
			if err != nil {
				return 0, nil, err
			}
			queryTemp.Value = valTemp
		}

		arr = append(arr, queryTemp)
	}

	return n, arr, nil
}

func readTestDataOutput(testDataPath string) (arr []int64, err error) {
	f, err := os.Open(testDataPath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineTemp, err := strconv.ParseInt(strings.TrimSpace(fileScanner.Text()), 10, 64)
		if err != nil {
			return nil, err
		}
		arr = append(arr, lineTemp)
	}

	return arr, nil
}
