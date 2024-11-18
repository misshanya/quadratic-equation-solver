package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Number interface {
	int | float64
}

func parseArgs(args []string) ([]int, error) {
	var result []int
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func power[T Number](base, exponent T) T {
	return T(math.Pow(float64(base), float64(exponent)))
}

func squareRoot[T Number](value T) T {
	return T(math.Sqrt(float64(value)))
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments. 3 arguments (a, b, c) are excepted")
		return
	}

	args, err := parseArgs(os.Args[1:4])
	if err != nil {
		fmt.Println("Arguments parsing error")
		return
	}

	a, b, c := args[0], args[1], args[2]

	x1, x2 := solve(a, b, c)

	switch {
	case x1 != nil && x2 != nil:
		fmt.Printf("X1: %v\nX2: %v\n", *x1, *x2)
	case x1 != nil:
		fmt.Printf("X: %v\n", *x1)
	default:
		fmt.Println("No roots")
	}
}

func solve[T Number](a, b, c T) (*T, *T) {
	D := power(b, 2) - 4*a*c
	switch {
	case D > 0:
		x1 := (-b + squareRoot(D)) / (2 * a)
		x2 := (-b - squareRoot(D)) / (2 * a)
		return &x1, &x2
	case D == 0:
		x1 := -b / (2 * a)
		return &x1, nil
	}
	return nil, nil
}
