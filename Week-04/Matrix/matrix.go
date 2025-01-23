package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

var (
	ErrEmptyRow           = errors.New("empty row")
	ErrInvalidMatrixValue = errors.New("invalid matrix value")
	ErrUnevenMatrixSize   = errors.New("uneven matrix size")
)

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	if len(rows) == 0 {
		return nil, ErrEmptyRow
	}

	var matrix Matrix
	for _, row := range rows {
		if row == "" {
			return nil, ErrEmptyRow
		}

		fields := strings.Fields(row)
		intRow := make([]int, len(fields))

		for i, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				return nil, ErrInvalidMatrixValue
			}
			intRow[i] = val
		}

		if len(matrix) > 0 && len(matrix[0]) != len(intRow) {
			return nil, ErrUnevenMatrixSize
		}

		matrix = append(matrix, intRow)
	}

	return matrix, nil
}

// Cols returns the transpose of the matrix without affecting the original.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}

	rows, cols := len(m), len(m[0])
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = m[i][j]
		}
	}

	return result
}

// Rows returns a copy of the matrix rows.
func (m Matrix) Rows() [][]int {
	result := make(Matrix, len(m))
	for i, row := range m {
		result[i] = append([]int(nil), row...) // Copy row to ensure immutability.
	}
	return result
}

// Set updates the value at (row, col) if within bounds.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[row]) {
		return false
	}
	m[row][col] = val
	return true
}
