package helpers

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func BuildMatrix[T Number](nRows, nCols int) [][]T {
	matrix := make([][]T, nRows)
	for i := range matrix {
		matrix[i] = make([]T, nCols)
	}

	return matrix
}

func PrintMatrix[T Number](matrix [][]T) {
	for _, row := range matrix {
		for _, elem := range row {
			fmt.Printf("%5v ", elem)
		}
		fmt.Println()
	}
}
