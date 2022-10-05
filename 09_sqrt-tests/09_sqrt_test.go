package _9_sqrt_tests

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

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
	//testCases := []TestCase{
	//	{4, 2},
	//	{9, 3},
	//}
	var testCases []TestCase
	tcFile, _ := os.Open("testCases.csv")
	defer tcFile.Close()
	rd := csv.NewReader(tcFile)
	testcasesStr, _ := rd.ReadAll()
	for _, tcStr := range testcasesStr {
		a1, _ := strconv.ParseFloat(tcStr[0], 1)
		a2, _ := strconv.ParseFloat(tcStr[1], 1)
		tc := TestCase{a1, a2}
		testCases = append(testCases, tc)
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
