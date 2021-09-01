package sortem

func SelectionSort(l []int) []int {
	res := make([]int, len(l))
	copy(res, l)
	for i := range res {
		lowest_value_ix := i
		for j := i + 1; j < len(res); j++ {

			if res[j] < res[lowest_value_ix] {
				lowest_value_ix = j
			}
		}
		res[i], res[lowest_value_ix] = res[lowest_value_ix], res[i]
	}
	return res
}
