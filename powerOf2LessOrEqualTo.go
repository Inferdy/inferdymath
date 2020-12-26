package inferdymath

import "fmt"

// PowerOf2LessOrEqualTo returns the power of 2 closest and less or equal to x
func PowerOf2LessOrEqualTo(x int) (int, error) {
	if x < 1 {
		return 0, fmt.Errorf("math: function is undefined for non-positive %d", x)
	}

	var result int = 1

loop:
	x >>= 1
	if x == 0 {
		return result, nil
	} //else
	result <<= 1
	goto loop
}
