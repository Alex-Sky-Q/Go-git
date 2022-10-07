package main

import "testing"

func TestCorrectSent(t *testing.T) {
	type TestCase struct {
		inp string
		exp string
	}
	testCases := []TestCase{
		{"h", "H."},
		{"H", "H."},
		{"H.", "H."},
		{".H", ".H."},
		{".h", ".H."},
		{"h a", "H a."},
		{"h a!", "H a!"},
		{"h- a", "H- a."},
		{"h.- a", "H.- a."},
		{"h! A", "H! A."},
		{"h a?", "H a?"},
		{"h a? yes. ok", "H a? Yes. Ok."},
		{"h a? yes.. ok", "H a? Yes.. Ok."},
		{"hello world. with pleasure", "Hello world. With pleasure."},
		{"hi.  see you", "Hi.  See you."},
		{"hi.see you", "Hi.see you."},
		{"hi.see you.", "Hi.see you."},
		{"hi.see you.!", "Hi.see you.!"},
		{"hi.see you. ", "Hi.see you. ."},
		{"hi.see you. .", "Hi.see you. ."},
		{" ", " "},
		{"-", "-."},
		{".", "."},
		{"1", "1."},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run("TC", func(t *testing.T) {
			act := CorrectSent(tc.inp)
			if act != tc.exp {
				t.Fatalf("Exp: %s, Actual: %s", tc.exp, act)
			}
		})
	}

}

func TestSliceMean(t *testing.T) {
	type TCInt struct {
		in  []int
		exp float64
	}
	testCasesInt := []TCInt{
		{[]int{1, 2}, 1.5},
		{[]int{3, 7}, 5},
	}

	type TCFloat struct {
		in  []float64
		exp float64
	}
	testCasesFloat := []TCFloat{
		{[]float64{1, 2.5}, 1.75},
		{[]float64{3.5, 7.5}, 5.5},
	}

	for _, tc := range testCasesInt {
		t.Run("TC", func(t *testing.T) {
			got := SliceMean(tc.in)
			if got != tc.exp {
				t.Fatalf("Got: %v. Want: %v", got, tc.exp)
			}
		})
	}

	for _, tc := range testCasesFloat {
		t.Run("TC", func(t *testing.T) {
			got := SliceMean(tc.in)
			if got != tc.exp {
				t.Fatalf("Got: %v. Want: %v", got, tc.exp)
			}
		})
	}

}

func TestMinMaxDiff(t *testing.T) {
	type TC struct {
		in  []float64
		exp float64
	}
	testCases := []TC{
		{[]float64{1, 2, 3}, 2},
	}

	for _, tc := range testCases {
		t.Run("TC", func(t *testing.T) {
			got := MinMaxDiff(tc.in)
			if got != tc.exp {
				t.Fatalf("Got: %v. Want: %v", got, tc.exp)
			}
		})
	}
}
