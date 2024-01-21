package main

import "testing"

func Test_findShortest(t *testing.T) {
	testCases := []struct {
		name       string
		graphNodes int32
		graphFrom  []int32
		graphTo    []int32
		ids        []int64
		val        int32
		expected   int32
	}{
		{
			name:       "Sample Input 0",
			graphNodes: 4,
			graphFrom:  []int32{1, 1, 4},
			graphTo:    []int32{2, 3, 2},
			ids:        []int64{1, 2, 1, 1},
			val:        1,
			expected:   1,
		},
		{
			name:       "Sample Input 1",
			graphNodes: 4,
			graphFrom:  []int32{1, 1, 4},
			graphTo:    []int32{2, 3, 2},
			ids:        []int64{1, 2, 3, 4},
			val:        2,
			expected:   -1,
		},
		{
			name:       "Sample Input 2",
			graphNodes: 5,
			graphFrom:  []int32{1, 1, 2, 3},
			graphTo:    []int32{2, 3, 4, 5},
			ids:        []int64{1, 2, 3, 3, 2},
			val:        2,
			expected:   3,
		},
	}

	for _, tt := range testCases {
		tt := tt
		actual := findShortest(tt.graphNodes, tt.graphFrom, tt.graphTo, tt.ids, tt.val)
		if actual != tt.expected {
			t.Errorf("name: %s expected: %d actual: %d", tt.name, tt.expected, actual)
		}
	}
}
