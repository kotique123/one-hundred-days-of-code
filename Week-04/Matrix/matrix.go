package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.

type Matrix [][]int

var ErrEmptyRow = errors.New("empty row")
var ErrInvalidMatrixValue = errors.New("invalid matrix value")
var ErrUnevenMatrixSize = errors.New("uneven matrix size")

func New(s string) (Matrix, error) {
	splitByNewLines := strings.Split(s, "\n")

	var SliceOfStrings = func(slice []string) ([][]string, error) {
		result := make([][]string, len(slice))
		for i := 0; i < len(result); i++ {
			if slice[i] == "" {
				return nil, ErrEmptyRow
			}
			result[i] = make([]string, len(slice[i]))
		}
		for i := 0; i < len(result); i++ {
			result[i] = strings.Fields(slice[i])
		}
		return result, nil
	}

	var sliceOfInts = func(slice [][]string) ([][]int, error) {
		int_results := make([][]int, len(slice))
		for i := 0; i < len(int_results); i++ {
			int_results[i] = make([]int, len(slice[i]))
		}
		for i := 0; i < len(int_results); i++ {
			for j := 0; j < len(int_results[i]); j++ {
				value, err := strconv.Atoi(slice[i][j])
				if err == nil {
					int_results[i][j] = value
				} else {
					return [][]int{}, ErrInvalidMatrixValue
				}
			}
		}
		return int_results, nil
	}
	var evenRows = func(rowsAndCols [][]int) bool {
		status := false
		if len(rowsAndCols) > 1 {
			for i := 1; i < len(rowsAndCols); i++ {
				if len(rowsAndCols[i]) == len(rowsAndCols[i-1]) {
					status = true
				}
			}
		} else {
			return true
		}
		return status
	}
	stringSlice, status := SliceOfStrings(splitByNewLines)
	if status != nil {
		return nil, status
	}
	intSlice, err := sliceOfInts(stringSlice)
	if err != nil {
		return nil, err
	}
	if !(evenRows(intSlice)) {
		return nil, ErrUnevenMatrixSize
	}
	return intSlice, err
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return [][]int{}
	}

	rows := len(m)
	cols := len(m[0])

	// Create a 2D slice to store the result
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			result[col][row] = m[row][col]
		}
	}

	return result
}

func (m Matrix) Rows() [][]int {
	rows := make(Matrix, len(m))
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]int, len(m[i]))
	}
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}
	return rows
}
func (m Matrix) Set(row, col, val int) bool {
	if row >= 0 && col >= 0 {
		if row < len(m) && col < len(m[row]) {
			m[row][col] = val
			return true
		}
	}
	return false
}
