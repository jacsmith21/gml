package mat

import (
	"fmt"

	"github.com/jacsmith21/gnn/vec"
	"strconv"
)

// Make initializes a matrix with i rows and j cols
func Make(i, j int) Matrix {
	vecs := make([]vec.Vector, j)
	for k := 0; k < j; k++ {
		vecs[k] = vec.Make(i)
	}

	return Matrix{vecs}
}

// InitRows initializes a Matrix with given row vectors
func InitRows(rows ...vec.Vector) Matrix {
	mat := Matrix{rows}
	mat.Transpose()
	return mat
}

// InitCols initializes a Matrix with given col vectors
func InitCols(cols ...vec.Vector) Matrix {
	return Matrix{cols}
}

// Copy returns a copy of the given Matrix
func Copy(from Matrix) Matrix {
	cols := make([]vec.Vector, len(from.cols))
	for i, c := range from.cols {
		cols[i] = vec.Copy(c)
	}

	return Matrix{cols}
}

// Slice returns a slice of the Matrix
func Slice(m Matrix, from, to int) Matrix {
	return Matrix{m.cols[from:to]}
}

// Mul performs the matrix multiplication m1 x m2
func Mul(m1, m2 Matrix) Matrix {
	if m1.ColCount() != m2.RowCount() {
		panic(fmt.Sprintf("m1 column count does not match m2 row count: %v vs. %v", m1.ColCount(), m2.RowCount()))
	}

	res := Make(m1.RowCount(), m2.ColCount())
	for i := 0; i < m1.RowCount(); i++ {
		for j := 0; j < m2.ColCount(); j++ {
			sum := 0.
			for index := 0; index < m1.ColCount(); index++ {
				sum += m1.At(i, index) * m2.At(index, j)
			}

			res.Set(i, j, sum)
		}
	}

	return res
}

// SumCols sums the columns of the given matrix and returns the sum vector
func SumCols(m Matrix) vec.Vector {
	sums := vec.Make(m.RowCount())
	for i := 0; i < m.RowCount(); i++ {
		sum := 0.
		for j := 0; j < m.ColCount(); j++ {
			sum += m.At(i, j)
		}
		sums.Set(i, sum)
	}
	return sums
}

// Sub subtracts the two matrices and returns the result
func Sub(m1, m2 Matrix) Matrix {
	copy := Copy(m1)
	copy.Sub(m2)
	return copy
}

// Transpose copies and transposes the given matrix
func Transpose(m Matrix) Matrix {
	copy := Copy(m)
	copy.Transpose()
	return copy
}

// InitFromStringArray initilizes an array from a 2D array of strings
func InitFromStringArray(s [][]string) Matrix {
	if len(s) == 0 {
		return Make(0, 0)
	}

	m := Make(len(s), len(s[0]))
	for i := 0; i < m.RowCount(); i++ {
		for j := 0; j < m.ColCount(); j++ {
			f, err := strconv.ParseFloat(s[i][j], 64)
			if err != nil {
				panic(fmt.Sprintf("cannot convert %v to a float", s[i][j]))
			}

			m.Set(i, j, f)
		}
	}

	return m
}
