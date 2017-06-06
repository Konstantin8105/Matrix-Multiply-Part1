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

/*
func TestBuffer1(t *testing.T) {
	if !isSame(mmBuffer1) {
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

func TestBuffer2(t *testing.T) {
	if !isSame(mmBuffer2) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer2(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer2(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer2 - added two buffers
func mmBuffer2(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
	buffer0 := make([]float64, n, n)
	buffer1 := make([]float64, n, n)
	// Now, we use (i+=2), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 2 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer0[j] = (*A)[i+0][j]
			buffer1[j] = (*A)[i+1][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer0[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer1[k] * (*B)[k][j]
			}
		}
	}
}

func TestBuffer4(t *testing.T) {
	if !isSame(mmBuffer4) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer4(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer4(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer4 - added 4 buffers
func mmBuffer4(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
	buffer0 := make([]float64, n, n)
	buffer1 := make([]float64, n, n)
	buffer2 := make([]float64, n, n)
	buffer3 := make([]float64, n, n)
	// Now, we use (i+=4), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 4 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer0[j] = (*A)[i+0][j]
			buffer1[j] = (*A)[i+1][j]
			buffer2[j] = (*A)[i+2][j]
			buffer3[j] = (*A)[i+3][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer0[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer1[k] * (*B)[k][j]
				(*C)[i+2][j] += buffer2[k] * (*B)[k][j]
				(*C)[i+3][j] += buffer3[k] * (*B)[k][j]
			}
		}
	}
}

func TestBuffer8(t *testing.T) {
	if !isSame(mmBuffer8) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer8(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer8(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer8 - added 8 buffers
func mmBuffer8(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
	buffer0 := make([]float64, n, n)
	buffer1 := make([]float64, n, n)
	buffer2 := make([]float64, n, n)
	buffer3 := make([]float64, n, n)
	buffer4 := make([]float64, n, n)
	buffer5 := make([]float64, n, n)
	buffer6 := make([]float64, n, n)
	buffer7 := make([]float64, n, n)
	// Now, we use (i+=8), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 8 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer0[j] = (*A)[i+0][j]
			buffer1[j] = (*A)[i+1][j]
			buffer2[j] = (*A)[i+2][j]
			buffer3[j] = (*A)[i+3][j]
			buffer4[j] = (*A)[i+4][j]
			buffer5[j] = (*A)[i+5][j]
			buffer6[j] = (*A)[i+6][j]
			buffer7[j] = (*A)[i+7][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer0[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer1[k] * (*B)[k][j]
				(*C)[i+2][j] += buffer2[k] * (*B)[k][j]
				(*C)[i+3][j] += buffer3[k] * (*B)[k][j]
				(*C)[i+4][j] += buffer4[k] * (*B)[k][j]
				(*C)[i+5][j] += buffer5[k] * (*B)[k][j]
				(*C)[i+6][j] += buffer6[k] * (*B)[k][j]
				(*C)[i+7][j] += buffer7[k] * (*B)[k][j]
			}
		}
	}
}

func TestBuffer16(t *testing.T) {
	if !isSame(mmBuffer16) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer16(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer16(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer16 - added 16 buffers
func mmBuffer16(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
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
	// Now, we use (i+=16), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 16 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer00[j] = (*A)[i+0][j]
			buffer01[j] = (*A)[i+1][j]
			buffer02[j] = (*A)[i+2][j]
			buffer03[j] = (*A)[i+3][j]
			buffer04[j] = (*A)[i+4][j]
			buffer05[j] = (*A)[i+5][j]
			buffer06[j] = (*A)[i+6][j]
			buffer07[j] = (*A)[i+7][j]
			buffer08[j] = (*A)[i+8][j]
			buffer09[j] = (*A)[i+9][j]
			buffer10[j] = (*A)[i+10][j]
			buffer11[j] = (*A)[i+11][j]
			buffer12[j] = (*A)[i+12][j]
			buffer13[j] = (*A)[i+13][j]
			buffer14[j] = (*A)[i+14][j]
			buffer15[j] = (*A)[i+15][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer00[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer01[k] * (*B)[k][j]
				(*C)[i+2][j] += buffer02[k] * (*B)[k][j]
				(*C)[i+3][j] += buffer03[k] * (*B)[k][j]
				(*C)[i+4][j] += buffer04[k] * (*B)[k][j]
				(*C)[i+5][j] += buffer05[k] * (*B)[k][j]
				(*C)[i+6][j] += buffer06[k] * (*B)[k][j]
				(*C)[i+7][j] += buffer07[k] * (*B)[k][j]
				(*C)[i+8][j] += buffer08[k] * (*B)[k][j]
				(*C)[i+9][j] += buffer09[k] * (*B)[k][j]
				(*C)[i+10][j] += buffer10[k] * (*B)[k][j]
				(*C)[i+11][j] += buffer11[k] * (*B)[k][j]
				(*C)[i+12][j] += buffer12[k] * (*B)[k][j]
				(*C)[i+13][j] += buffer13[k] * (*B)[k][j]
				(*C)[i+14][j] += buffer14[k] * (*B)[k][j]
				(*C)[i+15][j] += buffer15[k] * (*B)[k][j]
			}
		}
	}
}

func TestBuffer32(t *testing.T) {
	if !isSame(mmBuffer32) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer32(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer32(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer32 - added 32 buffers
func mmBuffer32(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
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
	buffer16 := make([]float64, n, n)
	buffer17 := make([]float64, n, n)
	buffer18 := make([]float64, n, n)
	buffer19 := make([]float64, n, n)
	buffer20 := make([]float64, n, n)
	buffer21 := make([]float64, n, n)
	buffer22 := make([]float64, n, n)
	buffer23 := make([]float64, n, n)
	buffer24 := make([]float64, n, n)
	buffer25 := make([]float64, n, n)
	buffer26 := make([]float64, n, n)
	buffer27 := make([]float64, n, n)
	buffer28 := make([]float64, n, n)
	buffer29 := make([]float64, n, n)
	buffer30 := make([]float64, n, n)
	buffer31 := make([]float64, n, n)
	// Now, we use (i+=32), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 32 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer00[j] = (*A)[i+0][j]
			buffer01[j] = (*A)[i+1][j]
			buffer02[j] = (*A)[i+2][j]
			buffer03[j] = (*A)[i+3][j]
			buffer04[j] = (*A)[i+4][j]
			buffer05[j] = (*A)[i+5][j]
			buffer06[j] = (*A)[i+6][j]
			buffer07[j] = (*A)[i+7][j]
			buffer08[j] = (*A)[i+8][j]
			buffer09[j] = (*A)[i+9][j]
			buffer10[j] = (*A)[i+10][j]
			buffer11[j] = (*A)[i+11][j]
			buffer12[j] = (*A)[i+12][j]
			buffer13[j] = (*A)[i+13][j]
			buffer14[j] = (*A)[i+14][j]
			buffer15[j] = (*A)[i+15][j]
			buffer16[j] = (*A)[i+16][j]
			buffer17[j] = (*A)[i+17][j]
			buffer18[j] = (*A)[i+18][j]
			buffer19[j] = (*A)[i+19][j]
			buffer20[j] = (*A)[i+20][j]
			buffer21[j] = (*A)[i+21][j]
			buffer22[j] = (*A)[i+22][j]
			buffer23[j] = (*A)[i+23][j]
			buffer24[j] = (*A)[i+24][j]
			buffer25[j] = (*A)[i+25][j]
			buffer26[j] = (*A)[i+26][j]
			buffer27[j] = (*A)[i+27][j]
			buffer28[j] = (*A)[i+28][j]
			buffer29[j] = (*A)[i+29][j]
			buffer30[j] = (*A)[i+30][j]
			buffer31[j] = (*A)[i+31][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer00[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer01[k] * (*B)[k][j]
				(*C)[i+2][j] += buffer02[k] * (*B)[k][j]
				(*C)[i+3][j] += buffer03[k] * (*B)[k][j]
				(*C)[i+4][j] += buffer04[k] * (*B)[k][j]
				(*C)[i+5][j] += buffer05[k] * (*B)[k][j]
				(*C)[i+6][j] += buffer06[k] * (*B)[k][j]
				(*C)[i+7][j] += buffer07[k] * (*B)[k][j]
				(*C)[i+8][j] += buffer08[k] * (*B)[k][j]
				(*C)[i+9][j] += buffer09[k] * (*B)[k][j]
				(*C)[i+10][j] += buffer10[k] * (*B)[k][j]
				(*C)[i+11][j] += buffer11[k] * (*B)[k][j]
				(*C)[i+12][j] += buffer12[k] * (*B)[k][j]
				(*C)[i+13][j] += buffer13[k] * (*B)[k][j]
				(*C)[i+14][j] += buffer14[k] * (*B)[k][j]
				(*C)[i+15][j] += buffer15[k] * (*B)[k][j]
				(*C)[i+16][j] += buffer16[k] * (*B)[k][j]
				(*C)[i+17][j] += buffer17[k] * (*B)[k][j]
				(*C)[i+18][j] += buffer18[k] * (*B)[k][j]
				(*C)[i+19][j] += buffer19[k] * (*B)[k][j]
				(*C)[i+20][j] += buffer20[k] * (*B)[k][j]
				(*C)[i+21][j] += buffer21[k] * (*B)[k][j]
				(*C)[i+22][j] += buffer22[k] * (*B)[k][j]
				(*C)[i+23][j] += buffer23[k] * (*B)[k][j]
				(*C)[i+24][j] += buffer24[k] * (*B)[k][j]
				(*C)[i+25][j] += buffer25[k] * (*B)[k][j]
				(*C)[i+26][j] += buffer26[k] * (*B)[k][j]
				(*C)[i+27][j] += buffer27[k] * (*B)[k][j]
				(*C)[i+28][j] += buffer28[k] * (*B)[k][j]
				(*C)[i+29][j] += buffer29[k] * (*B)[k][j]
				(*C)[i+30][j] += buffer30[k] * (*B)[k][j]
				(*C)[i+31][j] += buffer31[k] * (*B)[k][j]
			}
		}
	}
}

func TestBuffer64(t *testing.T) {
	if !isSame(mmBuffer64) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer64(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer64(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer64 - added 64 buffers
func mmBuffer64(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
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
	buffer16 := make([]float64, n, n)
	buffer17 := make([]float64, n, n)
	buffer18 := make([]float64, n, n)
	buffer19 := make([]float64, n, n)
	buffer20 := make([]float64, n, n)
	buffer21 := make([]float64, n, n)
	buffer22 := make([]float64, n, n)
	buffer23 := make([]float64, n, n)
	buffer24 := make([]float64, n, n)
	buffer25 := make([]float64, n, n)
	buffer26 := make([]float64, n, n)
	buffer27 := make([]float64, n, n)
	buffer28 := make([]float64, n, n)
	buffer29 := make([]float64, n, n)
	buffer30 := make([]float64, n, n)
	buffer31 := make([]float64, n, n)
	buffer32 := make([]float64, n, n)
	buffer33 := make([]float64, n, n)
	buffer34 := make([]float64, n, n)
	buffer35 := make([]float64, n, n)
	buffer36 := make([]float64, n, n)
	buffer37 := make([]float64, n, n)
	buffer38 := make([]float64, n, n)
	buffer39 := make([]float64, n, n)
	buffer40 := make([]float64, n, n)
	buffer41 := make([]float64, n, n)
	buffer42 := make([]float64, n, n)
	buffer43 := make([]float64, n, n)
	buffer44 := make([]float64, n, n)
	buffer45 := make([]float64, n, n)
	buffer46 := make([]float64, n, n)
	buffer47 := make([]float64, n, n)
	buffer48 := make([]float64, n, n)
	buffer49 := make([]float64, n, n)
	buffer50 := make([]float64, n, n)
	buffer51 := make([]float64, n, n)
	buffer52 := make([]float64, n, n)
	buffer53 := make([]float64, n, n)
	buffer54 := make([]float64, n, n)
	buffer55 := make([]float64, n, n)
	buffer56 := make([]float64, n, n)
	buffer57 := make([]float64, n, n)
	buffer58 := make([]float64, n, n)
	buffer59 := make([]float64, n, n)
	buffer60 := make([]float64, n, n)
	buffer61 := make([]float64, n, n)
	buffer62 := make([]float64, n, n)
	buffer63 := make([]float64, n, n)
	// Now, we use (i+=64), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 64 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer00[j] = (*A)[i+0][j]
			buffer01[j] = (*A)[i+1][j]
			buffer02[j] = (*A)[i+2][j]
			buffer03[j] = (*A)[i+3][j]
			buffer04[j] = (*A)[i+4][j]
			buffer05[j] = (*A)[i+5][j]
			buffer06[j] = (*A)[i+6][j]
			buffer07[j] = (*A)[i+7][j]
			buffer08[j] = (*A)[i+8][j]
			buffer09[j] = (*A)[i+9][j]
			buffer10[j] = (*A)[i+10][j]
			buffer11[j] = (*A)[i+11][j]
			buffer12[j] = (*A)[i+12][j]
			buffer13[j] = (*A)[i+13][j]
			buffer14[j] = (*A)[i+14][j]
			buffer15[j] = (*A)[i+15][j]
			buffer16[j] = (*A)[i+16][j]
			buffer17[j] = (*A)[i+17][j]
			buffer18[j] = (*A)[i+18][j]
			buffer19[j] = (*A)[i+19][j]
			buffer20[j] = (*A)[i+20][j]
			buffer21[j] = (*A)[i+21][j]
			buffer22[j] = (*A)[i+22][j]
			buffer23[j] = (*A)[i+23][j]
			buffer24[j] = (*A)[i+24][j]
			buffer25[j] = (*A)[i+25][j]
			buffer26[j] = (*A)[i+26][j]
			buffer27[j] = (*A)[i+27][j]
			buffer28[j] = (*A)[i+28][j]
			buffer29[j] = (*A)[i+29][j]
			buffer30[j] = (*A)[i+30][j]
			buffer31[j] = (*A)[i+31][j]
			buffer32[j] = (*A)[i+32][j]
			buffer33[j] = (*A)[i+33][j]
			buffer34[j] = (*A)[i+34][j]
			buffer35[j] = (*A)[i+35][j]
			buffer36[j] = (*A)[i+36][j]
			buffer37[j] = (*A)[i+37][j]
			buffer38[j] = (*A)[i+38][j]
			buffer39[j] = (*A)[i+39][j]
			buffer40[j] = (*A)[i+40][j]
			buffer41[j] = (*A)[i+41][j]
			buffer42[j] = (*A)[i+42][j]
			buffer43[j] = (*A)[i+43][j]
			buffer44[j] = (*A)[i+44][j]
			buffer45[j] = (*A)[i+45][j]
			buffer46[j] = (*A)[i+46][j]
			buffer47[j] = (*A)[i+47][j]
			buffer48[j] = (*A)[i+48][j]
			buffer49[j] = (*A)[i+49][j]
			buffer50[j] = (*A)[i+50][j]
			buffer51[j] = (*A)[i+51][j]
			buffer52[j] = (*A)[i+52][j]
			buffer53[j] = (*A)[i+53][j]
			buffer54[j] = (*A)[i+54][j]
			buffer55[j] = (*A)[i+55][j]
			buffer56[j] = (*A)[i+56][j]
			buffer57[j] = (*A)[i+57][j]
			buffer58[j] = (*A)[i+58][j]
			buffer59[j] = (*A)[i+59][j]
			buffer60[j] = (*A)[i+60][j]
			buffer61[j] = (*A)[i+61][j]
			buffer62[j] = (*A)[i+62][j]
			buffer63[j] = (*A)[i+63][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer00[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer01[k] * (*B)[k][j]
				(*C)[i+2][j] += buffer02[k] * (*B)[k][j]
				(*C)[i+3][j] += buffer03[k] * (*B)[k][j]
				(*C)[i+4][j] += buffer04[k] * (*B)[k][j]
				(*C)[i+5][j] += buffer05[k] * (*B)[k][j]
				(*C)[i+6][j] += buffer06[k] * (*B)[k][j]
				(*C)[i+7][j] += buffer07[k] * (*B)[k][j]
				(*C)[i+8][j] += buffer08[k] * (*B)[k][j]
				(*C)[i+9][j] += buffer09[k] * (*B)[k][j]
				(*C)[i+10][j] += buffer10[k] * (*B)[k][j]
				(*C)[i+11][j] += buffer11[k] * (*B)[k][j]
				(*C)[i+12][j] += buffer12[k] * (*B)[k][j]
				(*C)[i+13][j] += buffer13[k] * (*B)[k][j]
				(*C)[i+14][j] += buffer14[k] * (*B)[k][j]
				(*C)[i+15][j] += buffer15[k] * (*B)[k][j]
				(*C)[i+16][j] += buffer16[k] * (*B)[k][j]
				(*C)[i+17][j] += buffer17[k] * (*B)[k][j]
				(*C)[i+18][j] += buffer18[k] * (*B)[k][j]
				(*C)[i+19][j] += buffer19[k] * (*B)[k][j]
				(*C)[i+20][j] += buffer20[k] * (*B)[k][j]
				(*C)[i+21][j] += buffer21[k] * (*B)[k][j]
				(*C)[i+22][j] += buffer22[k] * (*B)[k][j]
				(*C)[i+23][j] += buffer23[k] * (*B)[k][j]
				(*C)[i+24][j] += buffer24[k] * (*B)[k][j]
				(*C)[i+25][j] += buffer25[k] * (*B)[k][j]
				(*C)[i+26][j] += buffer26[k] * (*B)[k][j]
				(*C)[i+27][j] += buffer27[k] * (*B)[k][j]
				(*C)[i+28][j] += buffer28[k] * (*B)[k][j]
				(*C)[i+29][j] += buffer29[k] * (*B)[k][j]
				(*C)[i+30][j] += buffer30[k] * (*B)[k][j]
				(*C)[i+31][j] += buffer31[k] * (*B)[k][j]
				(*C)[i+32][j] += buffer32[k] * (*B)[k][j]
				(*C)[i+33][j] += buffer33[k] * (*B)[k][j]
				(*C)[i+34][j] += buffer34[k] * (*B)[k][j]
				(*C)[i+35][j] += buffer35[k] * (*B)[k][j]
				(*C)[i+36][j] += buffer36[k] * (*B)[k][j]
				(*C)[i+37][j] += buffer37[k] * (*B)[k][j]
				(*C)[i+38][j] += buffer38[k] * (*B)[k][j]
				(*C)[i+39][j] += buffer39[k] * (*B)[k][j]
				(*C)[i+40][j] += buffer40[k] * (*B)[k][j]
				(*C)[i+41][j] += buffer41[k] * (*B)[k][j]
				(*C)[i+42][j] += buffer42[k] * (*B)[k][j]
				(*C)[i+43][j] += buffer43[k] * (*B)[k][j]
				(*C)[i+44][j] += buffer44[k] * (*B)[k][j]
				(*C)[i+45][j] += buffer45[k] * (*B)[k][j]
				(*C)[i+46][j] += buffer46[k] * (*B)[k][j]
				(*C)[i+47][j] += buffer47[k] * (*B)[k][j]
				(*C)[i+48][j] += buffer48[k] * (*B)[k][j]
				(*C)[i+49][j] += buffer49[k] * (*B)[k][j]
				(*C)[i+50][j] += buffer50[k] * (*B)[k][j]
				(*C)[i+51][j] += buffer51[k] * (*B)[k][j]
				(*C)[i+52][j] += buffer52[k] * (*B)[k][j]
				(*C)[i+53][j] += buffer53[k] * (*B)[k][j]
				(*C)[i+54][j] += buffer54[k] * (*B)[k][j]
				(*C)[i+55][j] += buffer55[k] * (*B)[k][j]
				(*C)[i+56][j] += buffer56[k] * (*B)[k][j]
				(*C)[i+57][j] += buffer57[k] * (*B)[k][j]
				(*C)[i+58][j] += buffer58[k] * (*B)[k][j]
				(*C)[i+59][j] += buffer59[k] * (*B)[k][j]
				(*C)[i+60][j] += buffer60[k] * (*B)[k][j]
				(*C)[i+61][j] += buffer61[k] * (*B)[k][j]
				(*C)[i+62][j] += buffer62[k] * (*B)[k][j]
				(*C)[i+63][j] += buffer63[k] * (*B)[k][j]
			}
		}
	}
}
*/
