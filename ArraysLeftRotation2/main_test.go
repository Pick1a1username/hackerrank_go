package main

import (
	"reflect"
	"testing"
)

func Test_rotLeft_1(t *testing.T) {
	rotNum := int32(4)
	arr := []int32{1, 2, 3, 4, 5}
	answer := []int32{5, 1, 2, 3, 4}

	result := rotLeft(arr, rotNum)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v", result, answer)
	}
}
