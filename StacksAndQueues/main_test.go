package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_processQueries_testcase15(b *testing.B) {
	_, arr, err := readTestDataInput("testcase15_input_private.txt")
	require.Nil(b, err)
	expected, err := readTestDataOutput("testcase15_output_private.txt")
	require.Nil(b, err)

	actual := processQueries(arr)

	if !reflect.DeepEqual(actual, expected) {
		b.Error("unexpected actual")
	}

}
