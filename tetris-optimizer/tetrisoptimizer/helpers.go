package tetrisoptimizer

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func getNearSqrtNum(num int) int {
	res := 2
	for ; res*res < num; res++ {

	}
	return res
}
