package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	inp []int
	exp []int
}

func TestBubbleSort(t *testing.T) {
	TCS := []TestCase{
		{[]int{1, 3, 6}, []int{1, 3, 6}},
		{[]int{6, 3, 1}, []int{1, 3, 6}},
		{[]int{-1, 6, 3}, []int{-1, 3, 6}},
		{[]int{1, 6, -3}, []int{-3, 1, 6}},
	}

	for _, TC := range TCS {
		t.Run("Test Case", func(t *testing.T) {
			BubbleSort(TC.inp)
			if !reflect.DeepEqual(TC.inp, TC.exp) {
				t.Fatalf("Slices do not match")
			}
		})
	}
}
