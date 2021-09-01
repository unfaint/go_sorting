package sortem

func InsertionSort(l []int) []int {
	res := make([]int, len(l))
	copy(res, l)
	if len(l) < 2 {
		return res
	}

	for i := 1; i < len(res); i++ {
		j, item_to_insert := i-1, res[i]
		for ; (j >= 0) && (res[j] > item_to_insert); j-- {
			res[j+1] = res[j]
		}
		res[j+1] = item_to_insert
	}

	return res
}
