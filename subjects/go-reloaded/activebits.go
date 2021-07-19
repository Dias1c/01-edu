package student

func ActiveBits(n int) uint {
	res := uint(0)
	for n != 0 {
		if n&1 == 1 {
			res++
		}
		n = n >> 1
	}
	return res
}
