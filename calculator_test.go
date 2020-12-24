package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

var testCasesAdd = []struct {
	desc string
	a    float64
	b    []float64
	want float64
}{
	{"add simple", 2, []float64{2}, 4},
	{"add half", 2, []float64{1.5}, 3.5},
	{"add negative", 2, []float64{-3}, -1},
	{"add multi", 3, []float64{2, 1}, 6},
}

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesAdd {
		got := calculator.Add(tc.a, tc.b...)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

var testCasesSubtract = []struct {
	desc string
	a    float64
	b    []float64
	want float64
}{
	{"subtract simple", 4, []float64{2}, 2},
	{"subtract half", 2, []float64{1.5}, 0.5},
	{"subtract negative", 2, []float64{-3}, 5},
	{"subtract multi", 5, []float64{3, 5, -3}, 0},
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesSubtract {
		got := calculator.Subtract(tc.a, tc.b...)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

var testCasesMultiply = []struct {
	desc string
	a    float64
	b    []float64
	want float64
}{
	{"multiply simple", 4, []float64{2}, 8},
	{"multiply half", 2, []float64{1.75}, 3.5},
	{"multiply negative", 2, []float64{-3}, -6},
	{"multiply multi", 2, []float64{2, -2, 2, -2, 2, -2, -2, -2}, -512},
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesMultiply {
		got := calculator.Multiply(tc.a, tc.b...)
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

func TestMultiplyRand(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Float64() * 10
		b := rand.Float64() * 10
		want := a * b
		got := calculator.Multiply(a, b)
		if want != got {
			t.Errorf("Multipling a: %f with b: %f failed. Want: %f, got: %f", a, b, want, got)
		}
	}
}

var testCasesDivide = []struct {
	desc string
	a    float64
	b    []float64
	want float64
	err  bool
}{
	{"divide simple", 4, []float64{2}, 2, false},
	{"divide half", 3.5, []float64{1.75}, 2, false},
	{"divide negative", -6, []float64{-3}, 2, false},
	{"divide a zero", 0, []float64{5}, 0, false},
	{"divide by zero", 3, []float64{0}, 0, true},
	{"divide multi", 27, []float64{3, -3, 3}, -1, false},
	{"divide multi with zero", 81, []float64{5, 7, 1, 0, 3}, 0, true},
}

func TestDivide(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesDivide {
		got, err := calculator.Divide(tc.a, tc.b...)
		if err != nil {
			if tc.err != true {
				t.Error("unexpected error returned")
			}
		} else {
			if tc.err == true {
				t.Error("no error reported when expected one")
			}
		}
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

var testCasesSqrt = []struct {
	desc string
	a    float64
	want float64
	err  bool
}{
	{"sqrt simple", 4, 2, false},
	{"sqrt half", 0.25, 0.5, false},
	{"sqrt negative", -6, 0, true},
	{"sqrt a zero", 0, 0, false},
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	for _, tc := range testCasesSqrt {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			if tc.err != true {
				t.Error("no error reported when expected one")
			}
		} else {
			if tc.err == true {
				t.Error("unexpected error returned")
			}
		}
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}

var testStringHandler = []struct {
	desc  string
	input string
	want  float64
	err   bool
}{
	{"addition, no space", "2+2", 4, false},
	{"addition, with space", "6 + 2", 8, false},
	{"subtract, negative", "20-32", -12, false},
	{"multiple, float", "3.5 * 2", 7, false},
	{"divide with zero", "88/0", 0, true},
}

func TestStringHandler(t *testing.T) {
	for _, tc := range testStringHandler {
		got, err := calculator.StringHandler(tc.input)
		if err != nil {
			if tc.err != true {
				t.Error("unexpected error returned", err)
			}
		} else {
			if tc.err == true {
				t.Error("no error reported when expected one")
			}
		}
		if tc.want != got {
			t.Errorf("%s: want %f, got %f", tc.desc, tc.want, got)
		}
	}
}
