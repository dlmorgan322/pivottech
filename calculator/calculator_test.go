package calculator_test

import (
	"testing"

	"github.com/dlmorgan322/pivottech/calculator"
)

//test

func TestCalculator(t *testing.T) {

	tests := map[string]struct {
		x, y, want int
		op         func(int, int) int
		err        error
	}{
		"Add":  {x: 1, y: 2, want: 3, op: calculator.Add},
		"Sub":  {x: 2, y: 1, want: 1, op: calculator.Sub},
		"Mult": {x: 2, y: 3, want: 6, op: calculator.Mult},
		"Div":  {x: 6, y: 2, want: 3, op: calculator.Div},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.op != nil {
				got := test.op(test.x, test.y)
				if got != test.want {
					t.Errorf("got %d, want %d", got, test.want)
				}
				return
			}

		})

	}
}

func TestPow(t *testing.T) {

	powTest := map[string]struct {
		x, y, want float64
		op         func(float64, float64) float64
		err        error
	}{
		"Pow": {x: 6, y: 2, want: 36, op: calculator.Pow},
	}
	for name, test := range powTest {
		t.Run(name, func(t *testing.T) {
			if test.op != nil {
				got := test.op(test.x, test.y)
				if got != test.want {
					t.Errorf("got %d, want %d", got, test.want)
				}
				return
			}
		})
	}
}
