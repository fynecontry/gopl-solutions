// Variation of popcount package which uses a loop instead of a single expression.
package popcountv1

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var bitcount byte
	for i := 0; i < 8; i++ {
		bitcount += pc[byte(x>>(i*8))]
	}
	return int(bitcount)
}
