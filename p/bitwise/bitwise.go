package bitwise

import (
	"strings"
)

func SetBit(n byte, pos uint) byte {
	n |= (1 << pos)
	return n
}
func ClearBit(n byte, pos uint) byte {
	mask := byte(^(1 << pos))
	n &= mask
	return n
}
func HasBit(n byte, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func PrintBitLegend() string {
	return "8 7 6 5 4 3 2 1"
}

func PrintBits(b byte) string {

	bitstates := []string{"0", "0", "0", "0", "0", "0", "0", "0"}
	var pos, n uint = 7, 0

	for ; ; pos-- {
		if true == HasBit(b, pos) {
			bitstates[n] = "1"
		}
		n++
		if pos == 0 {
			break
		}

	}
	return strings.Join(bitstates[0:], " ")
}
