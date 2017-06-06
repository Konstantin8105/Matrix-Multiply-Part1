package main_test

import (
	"testing"
)

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
		mmSimple(&A, &B, &C)
		// Finish of algorithm
	}
}
