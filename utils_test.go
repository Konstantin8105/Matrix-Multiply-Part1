package main_test

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
			A[i][j] = 4.0 * float64(i+9+j)    //* rand.Float64() * float64(j-i+n*2)
			B[i][j] = 4.0 * float64(11-i-2*j) //* rand.Float64() * float64(j-i+n*2)
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

func validAlgorithm(A, B, C *[][]float64) {
	n := len(*A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i][j] += (*A)[i][k] * (*B)[k][j]
			}
		}
	}
}
