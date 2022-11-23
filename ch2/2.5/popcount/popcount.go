package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(n uint64) int {
	count := 0
	for n != 0 {
		count += 1
		n = n & (n - 1)
	}
	return count
}
