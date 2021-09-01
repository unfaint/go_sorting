package sortem

func QuickSort(l []int) []int {
	if len(l) < 2 {
		return l
	}

	left, pivot := 0, len(l)/2

	// temporary swap pivot with the last element
	l[pivot], l[len(l)-1] = l[len(l)-1], l[pivot]
	pivotElement := l[len(l)-1]

	for i := 0; i < len(l)-1; i++ {
		if l[i] < pivotElement {
			l[i], l[left] = l[left], l[i]
			left++
		}
	}
	// swap pivot with the first largest element
	l[left], l[len(l)-1] = l[len(l)-1], l[left]

	QuickSort(l[:left])
	QuickSort(l[left+1:])

	return l
}
