package sorts

// Problem Statement
// Given unsorted array: [4, 6, 3, 9, 0, 5]
// Expected output: [0, 3, 4, 5, 6, 9]

// selectionSort sorts array by selecting the lowest element and swap it with the first element
// repeat for the second lowest element, etc
func selectionSort(in []int32) []int32 {
	// first loop represents the current index/position we're at, e.g.: replacing the first element with the lowest element
	for i, dat := range in[:len(in)-1] {

		// second loop finds the 'current' lowest element
		currLowestValue := dat
		currLowestIdx := i
		for j, low := range in[i+1:] { // in[i+1:] indicates that we only iterate through the subsequent slice after the current index

			if low < currLowestValue {
				currLowestValue = low
				currLowestIdx = i + 1 + j // currLowestIdx follows the i+1 formula since it is refering to the index of the main array
			}
		}

		in[i], in[currLowestIdx] = in[currLowestIdx], in[i]
	}

	return in
}
