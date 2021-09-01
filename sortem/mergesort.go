package sortem

func MergeSort(arr []int) []int {
	res := make([]int, len(arr))

	// base case
	if len(arr) <= 1 {
		return arr
	}

	// recursive case
	pivot := len(arr) / 2
	left := MergeSort(arr[:pivot])
	right := MergeSort(arr[pivot:])

	l, r := 0, 0

	for i := 0; i < len(res); i++ {
		if l < len(left) && r < len(right) {
			if left[l] < right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
		} else {
			if l < len(left) {
				res[i] = left[l]
				l++
			}
			if r < len(right) {
				res[i] = right[r]
				r++
			}
		}
	}

	return res
}
