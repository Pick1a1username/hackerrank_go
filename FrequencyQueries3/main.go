package main

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	// https://www.hackerrank.com/challenges/frequency-queries/editorial?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
	// https://www.hackerrank.com/challenges/frequency-queries/forum/comments/1134807
	// prepare an array for return
	frequenciesResult := []int32{}

	// prepare a map for data structure
	data := map[int32]int32{}
	dataFrequencies := map[int32]int32{}

	// loop by queries
	for i := 0; i < len(queries); i++ {
		op := queries[i][0]
		v := queries[i][1]
		// switch for operations
		switch op {
		case 1:
			// 1 x: Insert x in your data structure.
			// we don't need to take care of dataFrequencies[0]
			dataFrequencies[data[v]]--
			data[v]++
			dataFrequencies[data[v]]++
		case 2:
			// 2 y: Delete one occurence of y from your data structure, if present.
			if data[v] == 0 {
				continue
			}
			dataFrequencies[data[v]]--
			data[v]--
			dataFrequencies[data[v]]++
		case 3:
			// 3 z: Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.
			if dataFrequencies[v] > 0 {
				frequenciesResult = append(frequenciesResult, 1)
			} else {
				frequenciesResult = append(frequenciesResult, 0)
			}
		default:
			// else: panic
			panic("invalid operation")
		}
	}
	// return
	return frequenciesResult
}
