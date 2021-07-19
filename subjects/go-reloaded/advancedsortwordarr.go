package student

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	size := len(a)
	for i := 0; i < size; i++ {
		selectedIdx := i
		for j := i + 1; j < size; j++ {
			if f(a[selectedIdx], a[j]) > 0 {
				selectedIdx = j
			}
		}
		a[selectedIdx], a[i] = a[i], a[selectedIdx]
	}
}

func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}
