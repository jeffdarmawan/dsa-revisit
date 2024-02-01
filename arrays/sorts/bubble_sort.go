package sorts

// Problem Statement
// Given unsorted array: [4, 6, 3, 9, 0, 5]
// Expected output: [0, 3, 4, 5, 6, 9]

// bubbleSort sorts array by moving larger element to the right side of the array
// steps are repeated until all values are traversed
func bubbleSort(in []int32) []int32 {

	// completeness is measured from the back of the array
	for i := len(in) - 1; i >= 0; i-- {
		for j := range in[:i] {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}

	return in
}
