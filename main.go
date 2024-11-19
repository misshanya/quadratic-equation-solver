package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseArgs(args []string) ([]float64, error) {
	var result []float64
	for _, arg := range args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
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
