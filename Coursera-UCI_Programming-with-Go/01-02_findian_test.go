package main

import (
	"testing"
)

func TestCheckString(t *testing.T) {
	type TestCase struct {
		inp string
		exp string
	}

	tcs := []TestCase{
		{"ian", "Found!"},
		{"Ian", "Found!"},
		{"iuiygaygn", "Found!"},
		{"I d skd a efju N", "Found!"},
		{"ihhhhhn", "Not Found!"},
		{"ina", "Not Found!"},
		{"xian", "Not Found!"},
	}

	for _, tc := range tcs {
		t.Run("Test", func(t *testing.T) {
			res := CheckString(tc.inp)
			if res != tc.exp {
				t.Fatalf("Test Case failed")
			}
		})
	}

}
