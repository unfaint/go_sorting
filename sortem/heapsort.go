package sortem

func heapify(l []int, n int, root int) {
	left := root*2 + 1
	right := root*2 + 2
	largest := root
	if (left < n) && (l[left] > l[largest]) {
		largest = left
	}
	if (right < n) && (l[right] > l[largest]) {
		largest = right
	}

	if largest != root {
		l[root], l[largest] = l[largest], l[root]
		heapify(l, n, largest)
	}

}

func HeapSort(l []int) []int {
	res := make([]int, len(l))
	copy(res, l)
	for i := len(res); i >= 0; i-- {
		heapify(res, len(res), i)
	}

	for i := len(res) - 1; i > 0; i-- {
		res[0], res[i] = res[i], res[0]
		heapify(res, i, 0)
	}
	return res
}
