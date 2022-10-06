package main

import "testing"

func TestCorrectSent(t *testing.T) {
	type TestCase struct {
		inp string
		exp string
	}
	testCases := []TestCase{
		{"h", "H."},
		{"h a", "H a."},
		{"h a!", "H a!"},
		{"h! A", "H! A."},
		{"h a?", "H a?"},
		{"hello world. with pleasure", "Hello world. With pleasure."},
		{"hi.  see you", "Hi.  See you."},
		{"hi.see you", "Hi.see you."},
		{"hi.see you.", "Hi.see you."},
		{"hi.see you..", "Hi.see you.."},
		{"hi.see you. ", "Hi.see you. ."},
		{"hi.see you. .", "Hi.see you. ."},
		{" ", " "},
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
