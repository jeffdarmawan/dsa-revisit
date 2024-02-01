package reverse

// Problem Statement
// Given unsorted array: [4, 6, 3, 9, 0, 5]
// Expected output: [5, 0, 9, 3, 6, 4]

func reverseExternal(in []int32) []int32 {
	arrSize := len(in)

	// initialize and assign value into a new array
	out := make([]int32, arrSize)
	for i, dat := range in {
		out[len(in)-(i+1)] = dat
	}

	return out
}

func reverseInternal(in []int32) []int32 {
	lastItter := len(in) / 2
	for i := 0; i < lastItter; i++ {
		// swap values internally by accessing its index
		in[len(in)-(i+1)], in[i] = in[i], in[len(in)-(i+1)]
	}

	return in
}
