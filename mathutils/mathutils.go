// Package mathutils contain some math related algorithms or functions
package mathutils

// GCD greatest common divisor via Euclidean algorithm
// Stolen from : https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// AbsInt returns absolute number of an integer.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
