package sortem

func BubbleSort(l []int) []int {
	res := make([]int, len(l))
	copy(res, l)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(res)-1; i++ {
			if res[i] > res[i+1] {
				res[i], res[i+1] = res[i+1], res[i]
				swapped = true
			}
		}
	}
	return res
}
