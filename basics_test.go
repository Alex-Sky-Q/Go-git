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
