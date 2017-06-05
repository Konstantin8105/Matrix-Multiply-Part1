package main_test

import (
	"math/rand"
	"testing"
)

// getenerateMatrix - generate the matrix for test
func generateMatrix() (A, B, C [][]float64) {
	// size of matrix
	n := 1024

	// initialization
	A = make([][]float64, n)
	B = make([][]float64, n)
	C = make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		B[i] = make([]float64, n)
		C[i] = make([]float64, n)
	}

	// defaul values can be any
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A[i][j] = 4.0 * rand.Float64() * float64(j-i+n*2)
			B[i][j] = 4.0 * rand.Float64() * float64(j-i+n*2)
		}
	}
	return
}

func BenchmarkSimple(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		// We know - the size of matrix is same for all
		// matrix. So, we can only one variable of size
		n := len(A)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					C[i][j] = A[i][k] * B[k][j]
				}
			}
		}
		// Finish of algorithm
	}
}
