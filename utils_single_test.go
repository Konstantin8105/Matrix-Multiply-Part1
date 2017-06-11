package main_test

import "math"

// getenerateMatrix - generate the matrix for test
func generateMatrixSingle() (A, B, C []float64) {
	// size of matrix
	n := 1024

	// initialization
	A = make([]float64, n*n)
	B = make([]float64, n*n)
	C = make([]float64, n*n)

	// defaul values can be any
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A[i+j*n] = 4.0 * float64(i+9+j)    //* rand.Float64() * float64(j-i+n*2)
			B[i+j*n] = 4.0 * float64(11-i-2*j) //* rand.Float64() * float64(j-i+n*2)
		}
	}
	return
}

// isSame - function for check algorithm of matrix multiplication
// compare result with simple and slow classic algortithm
func isSameSingle(f func(a, b, c *[]float64)) bool {
	A, B, C := generateMatrixSingle()
	f(&A, &B, &C)

	// For avoid waiting of correctnes
	// of algorithm result
	n := int(math.Sqrt(float64(len(A))))
	validC := make([]float64, n*n)
	validAlgorithmSingle(&A, &B, &validC)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if validC[i+j*n] != C[i+j*n] {
				return false
			}
		}
	}
	return true
}

func validAlgorithmSingle(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+j*n] += (*A)[i+k*n] * (*B)[k+j*n]
			}
		}
	}
}
