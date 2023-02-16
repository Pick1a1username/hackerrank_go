package main

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	// prepare an array for return
	frequencies := []int32{}

	// prepare a map for data structure
	data := map[int32]int32{}

	// loop by queries
	for i := 0; i < len(queries); i++ {
		op := queries[i][0]
		v := queries[i][1]
		// switch for operations
		switch op {
		case 1:
			// 1 x: Insert x in your data structure.
			data[v]++
		case 2:
			// 2 y: Delete one occurence of y from your data structure, if present.
			if _, ok := data[v]; ok {
				if data[v] > 0 {
					data[v]--
				}
			}
		case 3:
			// 3 z: Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.
			found := false

			for _, dataV := range data {
				if dataV == v {
					frequencies = append(frequencies, 1)
					found = true
					break
				}
			}

			if !found {
				frequencies = append(frequencies, 0)
			}
		default:
			// else: panic
			panic("invalid operation")
		}
	}
	// return
	return frequencies
}
