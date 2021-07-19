package student

// Pow By Recursion
func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else if power < 0 {
		return 0
	}
	nb *= RecursivePower(nb, power-1)
	return nb
}