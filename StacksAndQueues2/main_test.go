package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processQueries_testcase15(t *testing.T) {
	_, _, err := readTestDataInput("testcase15_input_private.txt")
	assert.Nil(t, err)
	_, err = readTestDataOutput("testcase15_output_private.txt")
	assert.Nil(t, err)
}
