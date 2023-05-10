package main

import "testing"

func Test_isValid(t *testing.T) {
	testCases := []struct {
		description string
		s           string
		answer      string
	}{
		{
			description: "TestCase 1",
			s:           "aabbcd",
			answer:      "NO",
		},
		{
			description: "TestCase 2",
			s:           "aabbccddeefghi",
			answer:      "NO",
		},
		{
			description: "TestCase 3",
			s:           "abcdefghhgfedecba",
			answer:      "YES",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := isValid(testCase.s)
			if result != testCase.answer {
				t.Errorf("result: %s answer: %s", result, testCase.answer)
			}
		})
	}
}
