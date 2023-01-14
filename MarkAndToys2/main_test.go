package main

import (
	"testing"
)

func Test_maximumToysImpl_1(t *testing.T) {
	prices := []int32{1, 2, 3, 4}
	k := int32(7)
	answer := int32(3)

	result := maximumToysImpl(prices, k)

	if result != answer {
		t.Errorf("result: %d answer : %d", result, answer)
	}
}

func Test_maximumToysImpl_2(t *testing.T) {
	prices := []int32{1, 12, 5, 111, 200, 1000, 10}
	k := int32(30)
	answer := int32(4)

	result := maximumToysImpl(prices, k)

	if result != answer {
		t.Errorf("result: %d answer : %d", result, answer)
	}
}
