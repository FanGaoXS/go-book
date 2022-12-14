package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(n uint64) int {
	var res byte
	for i := 0; i < 8; i++ {
		res += pc[byte(n>>(i*8))]
	}
	return int(res)
}
