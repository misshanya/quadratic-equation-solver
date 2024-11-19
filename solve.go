package main

import "math"

func solve(a, b, c float64) (*float64, *float64) {
	D := math.Pow(b, 2) - 4*a*c
	switch {
	case D > 0:
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		return &x1, &x2
	case D == 0:
		x1 := -b / (2 * a)
		return &x1, nil
	default:
		return nil, nil
	}
}
