package main_test

import (
	"math/rand"
	"runtime"
	"sync"
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

// isSame - function for check algorithm of matrix multiplication
// compare result with simple and slow classic algortithm
func isSame(f func(a, b, c *[][]float64)) bool {
	A, B, C := generateMatrix()
	f(&A, &B, &C)

	// For avoid waiting of correctnes
	// of algorithm result
	n := len(A)
	validC := make([][]float64, n)
	for i := 0; i < n; i++ {
		validC[i] = make([]float64, n)
	}
	validAlgorithm(&A, &B, &validC)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if validC[i][j] != C[i][j] {
				return false
			}
		}
	}
	return true
}

// validAlgorithm - with 16 buffers
func validAlgorithm(A, B, C *[][]float64) {
	n := len(*A)
	// Found amount allowable parallelism
	threads := runtime.GOMAXPROCS(0)
	if threads > runtime.NumCPU() {
		threads = runtime.NumCPU()
	}
	// Create workgroup
	var wg sync.WaitGroup
	// Run calculation in goroutines
	for t := 0; t < threads; t++ {
		// Add one goroutine in workgroup
		wg.Add(1)
		// The value "init" is a number of thread
		// that created for offset of loop
		go func(init int) {
			// Change waitgroup after work done
			defer wg.Done()
			// Inialize addition variables
			var sum00, sum01, sum02, sum03, sum04, sum05, sum06, sum07, sum08, sum09 float64
			var sum10, sum11, sum12, sum13, sum14, sum15 float64
			// Create buffers
			amountBuffers := 16
			buffer00 := make([]float64, n, n)
			buffer01 := make([]float64, n, n)
			buffer02 := make([]float64, n, n)
			buffer03 := make([]float64, n, n)
			buffer04 := make([]float64, n, n)
			buffer05 := make([]float64, n, n)
			buffer06 := make([]float64, n, n)
			buffer07 := make([]float64, n, n)
			buffer08 := make([]float64, n, n)
			buffer09 := make([]float64, n, n)
			buffer10 := make([]float64, n, n)
			buffer11 := make([]float64, n, n)
			buffer12 := make([]float64, n, n)
			buffer13 := make([]float64, n, n)
			buffer14 := make([]float64, n, n)
			buffer15 := make([]float64, n, n)
			// Calculate amount of calculation part
			// for that goroutine
			amountParts := n / amountBuffers
			for i := init; i < amountParts; i += threads {
				for j := 0; j < n; j++ {
					// Put in buffer row of matrix [A]
					buffer00[j] = (*A)[i*amountBuffers+0][j]
					buffer01[j] = (*A)[i*amountBuffers+1][j]
					buffer02[j] = (*A)[i*amountBuffers+2][j]
					buffer03[j] = (*A)[i*amountBuffers+3][j]
					buffer04[j] = (*A)[i*amountBuffers+4][j]
					buffer05[j] = (*A)[i*amountBuffers+5][j]
					buffer06[j] = (*A)[i*amountBuffers+6][j]
					buffer07[j] = (*A)[i*amountBuffers+7][j]
					buffer08[j] = (*A)[i*amountBuffers+8][j]
					buffer09[j] = (*A)[i*amountBuffers+9][j]
					buffer10[j] = (*A)[i*amountBuffers+10][j]
					buffer11[j] = (*A)[i*amountBuffers+11][j]
					buffer12[j] = (*A)[i*amountBuffers+12][j]
					buffer13[j] = (*A)[i*amountBuffers+13][j]
					buffer14[j] = (*A)[i*amountBuffers+14][j]
					buffer15[j] = (*A)[i*amountBuffers+15][j]
				}
				for j := 0; j < n; j++ {
					sum00 = 0.0
					sum01 = 0.0
					sum02 = 0.0
					sum03 = 0.0
					sum04 = 0.0
					sum05 = 0.0
					sum06 = 0.0
					sum07 = 0.0
					sum08 = 0.0
					sum09 = 0.0
					sum10 = 0.0
					sum11 = 0.0
					sum12 = 0.0
					sum13 = 0.0
					sum14 = 0.0
					sum15 = 0.0
					for k := 0; k < n; k++ {
						sum00 += buffer00[k] * (*B)[k][j]
						sum01 += buffer01[k] * (*B)[k][j]
						sum02 += buffer02[k] * (*B)[k][j]
						sum03 += buffer03[k] * (*B)[k][j]
						sum04 += buffer04[k] * (*B)[k][j]
						sum05 += buffer05[k] * (*B)[k][j]
						sum06 += buffer06[k] * (*B)[k][j]
						sum07 += buffer07[k] * (*B)[k][j]
						sum08 += buffer08[k] * (*B)[k][j]
						sum09 += buffer09[k] * (*B)[k][j]
						sum10 += buffer10[k] * (*B)[k][j]
						sum11 += buffer11[k] * (*B)[k][j]
						sum12 += buffer12[k] * (*B)[k][j]
						sum13 += buffer13[k] * (*B)[k][j]
						sum14 += buffer14[k] * (*B)[k][j]
						sum15 += buffer15[k] * (*B)[k][j]
					}
					(*C)[i*amountBuffers+0][j] = sum00
					(*C)[i*amountBuffers+1][j] = sum01
					(*C)[i*amountBuffers+2][j] = sum02
					(*C)[i*amountBuffers+3][j] = sum03
					(*C)[i*amountBuffers+4][j] = sum04
					(*C)[i*amountBuffers+5][j] = sum05
					(*C)[i*amountBuffers+6][j] = sum06
					(*C)[i*amountBuffers+7][j] = sum07
					(*C)[i*amountBuffers+8][j] = sum08
					(*C)[i*amountBuffers+9][j] = sum09
					(*C)[i*amountBuffers+10][j] = sum10
					(*C)[i*amountBuffers+11][j] = sum11
					(*C)[i*amountBuffers+12][j] = sum12
					(*C)[i*amountBuffers+13][j] = sum13
					(*C)[i*amountBuffers+14][j] = sum14
					(*C)[i*amountBuffers+15][j] = sum15
				}
			}
		}(t)
	}
	wg.Wait()
}
