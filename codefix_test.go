package codefix

import (
	"testing"
)

func TestSwap(t *testing.T) {
	testCases := []struct {
		name                 string
		inputA, inputB       string
		expectedA, expectedB string
	}{
		{"swap 1", "world", "hello", "hello", "world"},
		{"swap 2", "second", "first", "first", "second"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultA, resultB := swap(tc.inputA, tc.inputB)
			if resultA != tc.expectedA || resultB != tc.expectedB {
				t.Errorf("%s(%v, %v) = (%v, %v); want (%v, %v)", tc.name, tc.inputA, tc.inputB, resultA, resultB, tc.expectedA, tc.expectedB)
			}
		})
	}
}

func TestUpdateStruct(t *testing.T) {
	s := MyStruct{Val: 10}
	s.UpdateStruct()

	if s.Val != 20 {
		t.Errorf("UpdateStruct got %v; want %v", s.Val, 20)
	}
}
