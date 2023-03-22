package main

import "testing"

// Table Driven Test
func Test_commonChild(t *testing.T) {
	testCases := []struct {
		description string
		s1          string
		s2          string
		ans         int32
	}{
		{
			description: "Sample TestCase 1",
			s1:          "HARRY",
			s2:          "SALLY",
			ans:         2,
		},
		{
			description: "Sample TestCase 2",
			s1:          "AA",
			s2:          "BB",
			ans:         0,
		},
		{
			description: "Sample TestCase 3",
			s1:          "SHINCHAN",
			s2:          "NOHARAAA",
			ans:         3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := commonChild(testCase.s1, testCase.s2)
			if result != testCase.ans {
				t.Errorf("result: %d answer: %d", result, testCase.ans)
			}
		})
	}
}
