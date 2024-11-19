package main

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	name     string
	data     []float64
	excepted []float64
}{
	{"solve equation with two roots", []float64{0.5, 0.125, 0.0}, []float64{0.0, -0.25}},
	{"solve equation with one root", []float64{1.0, -4.0, 4.0}, []float64{2.0}},
	{"solve equation without roots", []float64{1.0, 2.0, 3.0}, []float64{}},
}

func TestSolve(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			excepted := tc.excepted
			data := tc.data

			got := solve(data[0], data[1], data[2])

			if !reflect.DeepEqual(excepted, got) {
				t.Errorf("excepted %v, got %v", excepted, got)
			}
		})
	}
}
