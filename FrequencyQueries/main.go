package main

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	// prepare an array for return
	frequencies := []int32{}

	// prepare an array for data structure
	data := []int32{}

	// loop by queries
	for i := 0; i < len(queries); i++ {
		op := queries[i][0]
		v := queries[i][1]
		// switch for operations
		switch op {
		case 1:
			// 1 x: Insert x in your data structure.
			data = append(data, v)
		case 2:
			// 2 y: Delete one occurence of y from your data structure, if present.
			for j := 0; j < len(data); j++ {
				if data[j] == v {
					data = removeElemArrayInt32(&data, j)
					break
				}
			}
		case 3:
			// 3 z: Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.
			// make a map for saving values checked.
			checked := map[int32]int32{}
			found := false
			// loop by data
			for j := 0; j < len(data); j++ {
				target := data[j]
				// if the value has already been checked, skip it.
				if _, ok := checked[target]; ok {
					continue
				}
				// loop by data again.
				for k := 0; k < len(data); k++ {
					// if the value is equal to the value we are looking for, count it.
					if data[k] == target {
						checked[target]++
						if checked[target] > v {
							break
						}
					}
				}
				// if the counter is equal to v, add 1 to frequencies and break.
				if checked[target] == v {
					frequencies = append(frequencies, 1)
					found = true
					break
				}
			}
			// if there's no value found, add 0 to frequencies
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

func removeElemArrayInt32(arr *[]int32, i int) []int32 {
	tmpLeftQ := append([]int32{}, (*arr)[:i]...)
	tmpRightQ := append([]int32{}, (*arr)[i+1:]...)
	return append(tmpLeftQ, tmpRightQ...)
}
