package main

import (
	"testing"
	"reflect"
)

var TestData1 [][]int32 = [][]int32{
	[]int32 { -9, -9, -9, 1, 1, 1 },
	[]int32 { 0, -9, 0, 4, 3, 2 },
	[]int32 { -9, -9, -9, 1, 2, 3 },
	[]int32 { 0, 0, 8, 6, 6, 0 },
	[]int32 { 0, 0, 0, -2, 0, 0 },
	[]int32 { 0, 0, 1, 2, 4, 0 },
}

func Test_getHourGlassElems_1(t *testing.T) {
	var answerArr []int32 = []int32{ 1, 2, 3, 6, -2, 0, 0}
	resultArr, resultErr := getHourGlassElems(TestData1, 2, 3)
	if !reflect.DeepEqual(answerArr, resultArr) {
		t.Errorf("answer: %v result: %v\n", answerArr, resultArr)
	}
	if resultErr != nil {
		t.Errorf("answer: nil result: %s\n", resultErr.Error())
	}
}

func Test_hourglassSum_1(t *testing.T) {

	var answer int32 = 28
	result := hourglassSum(TestData1)
	if answer != result {
		t.Errorf("answer: %d result: %d\n", answer, result)
	}
}

