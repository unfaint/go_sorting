package sortem

import "testing"

func CompareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAllSortings(t *testing.T) {
	functions := make(map[string]func([]int) []int)

	functions["BubbleSort"] = BubbleSort
	functions["SelectionSort"] = SelectionSort
	functions["HeapSort"] = HeapSort
	functions["InsertionSort"] = InsertionSort
	functions["MergeSort"] = MergeSort
	functions["QuickSort"] = QuickSort

	cases := []struct {
		l        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{2, 1, 1}, []int{1, 1, 2}},
		{[]int{3, 5, 2, 1, 6, 4, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for name, f := range functions {
		t.Logf("testing %v", name)
		for i, c := range cases {
			testCase := make([]int, len(c.l))
			copy(testCase, c.l)
			res := f(testCase)
			if !CompareSlices(res, c.expected) {
				t.Errorf("%v failed case #%d: %v (expected %v, got %v)", name, i, testCase, c.expected, res)
			} else {
				t.Logf("case #%d: %v - ok", i, testCase)
			}
		}
	}
}
