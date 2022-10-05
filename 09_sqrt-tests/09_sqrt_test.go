package _9_sqrt_tests

import "testing"

type TestCase struct {
	value  float64
	expVal float64
}

func TestSqrt(t *testing.T) {
	val := sqrt(4)
	if val != 2 {
		t.Fatalf("Actual value: %v, Expected value: 2", val)
	}
}

func TestMany(t *testing.T) {
	testCases := []TestCase{
		{4, 2},
		{9, 3},
	}
	for _, tc := range testCases {
		t.Run("name", func(t *testing.T) {
			val := sqrt(tc.value)
			if val != tc.expVal {
				t.Fatalf("Actual value: %v, Expected value: %v", val, tc.expVal)
			}
		})
	}
}
