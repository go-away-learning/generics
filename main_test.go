package main

import (
	"testing"
)

func TestSumInts(t *testing.T) {
	want := int64(4)
	numbers := map[string]int64{
		"first":  2,
		"second": 2,
	}
	got := SumNumbers(numbers)

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func TestSumFloats(t *testing.T) {
	want := 3.50
	numbers := map[string]float64{
		"first":  1.75,
		"second": 1.75,
	}
	got := SumNumbers(numbers)

	if got != want {
		t.Errorf("got %f want %f, given %v", got, want, numbers)
	}
}
