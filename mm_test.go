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

func BenchmarkSimpleInside(b *testing.B) {
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
					C[i][j] += A[i][k] * B[k][j]
				}
			}
		}
		// Finish of algorithm
	}
}

// isSame - function for check algorithm of matrix multiplication
// compare result with simple and slow classic algortithm
func isSame(f func(a, b, c *[][]float64)) bool {
	A, B, C := generateMatrix()
	f(&A, &B, &C)
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			if sum != C[i][j] {
				return false
			}
		}
	}
	return true
}

func TestSimple(t *testing.T) {
	if !isSame(mmSimple) {
		t.Errorf("Algorithm is not correct")
	}
}

func mmSimple(A, B, C *[][]float64) {
	n := len(*A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i][j] += (*A)[i][k] * (*B)[k][j]
			}
		}
	}
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
		mmSimple(&A, &B, &C)
		// Finish of algorithm
	}
}

func TestBuffer1(t *testing.T) {
	if !isSame(mmSimple) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer1(b *testing.B) {
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
		mmBuffer1(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer1 - added one buffer
func mmBuffer1(A, B, C *[][]float64) {
	n := len(*A)
	buffer := make([]float64, n, n) // Create buffer
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			buffer[j] = (*A)[i][j] // Put in buffer row of matrix [A]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i][j] += buffer[k] * (*B)[k][j]
			}
		}
	}
}
