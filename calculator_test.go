package calculator_test

import (
	"calculator"
	"testing"
)

var testCasesAdd = []struct {
	desc string
	a    float64
	b    float64
	want float64
}{
	{"add simple", 2, 2, 4},
	{"add half", 2, 1.5, 3.5},
	{"add negative", 2, -3, -1},
}

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesAdd {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

var testCasesSubtract = []struct {
	desc string
	a    float64
	b    float64
	want float64
}{
	{"subtract simple", 4, 2, 2},
	{"subtract half", 2, 1.5, 0.5},
	{"subtract negative", 2, -3, 5},
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesSubtract {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 36
	got := calculator.Multiply(6, 6)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
