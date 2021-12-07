package main

import "testing"

func TestBitmaskCriteria(t *testing.T) {

	d := []int{
		0b_001111011000,
		0b_110110011111,
		0b_111100110100,
		0b_110000010101,
		0b_001110110000,
		0b_001110110000,
	}

	one := 1 << 11
	zero := 0
	result := bitmaskCriteria(d, 11, true)

	if result != one {
		t.Fatalf("%012b is not equal to %012b", result, one)
	}

	result = bitmaskCriteria(d, 11, false)
	if result != zero {
		t.Fatalf("%012b is not equal to %012b", result, zero)
	}

	one = 1 << 8
	result = bitmaskCriteria(d, 8, true)
	if result != one {
		t.Fatalf("%012b is not equal to %012b", result, one)
	}
	result = bitmaskCriteria(d, 8, false)
	if result != zero {
		t.Fatalf("%012b is not equal to %013b", result, zero)
	}

	one = 1 << 4
	result = bitmaskCriteria(d, 4, false)
	if result != one {
		t.Fatalf("%012b is not equal to %012b", result, one)
	}
	result = bitmaskCriteria(d, 4, true)
	if result != one {
		t.Fatalf("%012b is not equal to %012b", result, one)
	}
}
