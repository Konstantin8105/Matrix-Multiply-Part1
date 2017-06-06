package main_test

import (
	"runtime"
	"sync"
	"testing"
)

func TestParallelBufferVarOut2(t *testing.T) {
	if !isSame(mmParallelBufferVarOut2) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOut2(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOut2(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOut2 - with 2 buffers
func mmParallelBufferVarOut2(A, B, C *[][]float64) {
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
			var sum0, sum1 float64 // <---- Mark
			// Create buffers
			amountBuffers := 2
			buffer0 := make([]float64, n, n)
			buffer1 := make([]float64, n, n)
			// Calculate amount of calculation part
			// for that goroutine
			amountParts := n / amountBuffers
			for i := init; i < amountParts; i += threads {
				for j := 0; j < n; j++ {
					// Put in buffer row of matrix [A]
					buffer0[j] = (*A)[i*amountBuffers+0][j]
					buffer1[j] = (*A)[i*amountBuffers+1][j]
				}
				for j := 0; j < n; j++ {
					sum0 = 0.0 // <---- Mark
					sum1 = 0.0 // <---- Mark
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j] // <---- Mark
						sum1 += buffer1[k] * (*B)[k][j] // <---- Mark
					}
					(*C)[i*amountBuffers+0][j] = sum0 // <---- Mark
					(*C)[i*amountBuffers+1][j] = sum1 // <---- Mark
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOut4(t *testing.T) {
	if !isSame(mmParallelBufferVarOut4) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOut4(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOut4(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOut4 - with 4 buffers
func mmParallelBufferVarOut4(A, B, C *[][]float64) {
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
			var sum00, sum01, sum02, sum03 float64
			// Create buffers
			amountBuffers := 4
			buffer0 := make([]float64, n, n)
			buffer1 := make([]float64, n, n)
			buffer2 := make([]float64, n, n)
			buffer3 := make([]float64, n, n)
			// Calculate amount of calculation part
			// for that goroutine
			amountParts := n / amountBuffers
			for i := init; i < amountParts; i += threads {
				for j := 0; j < n; j++ {
					// Put in buffer row of matrix [A]
					buffer0[j] = (*A)[i*amountBuffers+0][j]
					buffer1[j] = (*A)[i*amountBuffers+1][j]
					buffer2[j] = (*A)[i*amountBuffers+2][j]
					buffer3[j] = (*A)[i*amountBuffers+3][j]
				}
				for j := 0; j < n; j++ {
					sum00 = 0.0
					sum01 = 0.0
					sum02 = 0.0
					sum03 = 0.0
					for k := 0; k < n; k++ {
						sum00 += buffer0[k] * (*B)[k][j]
						sum01 += buffer1[k] * (*B)[k][j]
						sum02 += buffer2[k] * (*B)[k][j]
						sum03 += buffer3[k] * (*B)[k][j]
					}
					(*C)[i*amountBuffers+0][j] = sum00
					(*C)[i*amountBuffers+1][j] = sum01
					(*C)[i*amountBuffers+2][j] = sum02
					(*C)[i*amountBuffers+3][j] = sum03
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOut8(t *testing.T) {
	if !isSame(mmParallelBufferVarOut8) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOut8(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOut8(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOut8 - with 8 buffers
func mmParallelBufferVarOut8(A, B, C *[][]float64) {
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
			var sum00, sum01, sum02, sum03, sum04, sum05, sum06, sum07 float64
			// Create buffers
			amountBuffers := 8
			buffer0 := make([]float64, n, n)
			buffer1 := make([]float64, n, n)
			buffer2 := make([]float64, n, n)
			buffer3 := make([]float64, n, n)
			buffer4 := make([]float64, n, n)
			buffer5 := make([]float64, n, n)
			buffer6 := make([]float64, n, n)
			buffer7 := make([]float64, n, n)
			// Calculate amount of calculation part
			// for that goroutine
			amountParts := n / amountBuffers
			for i := init; i < amountParts; i += threads {
				for j := 0; j < n; j++ {
					// Put in buffer row of matrix [A]
					buffer0[j] = (*A)[i*amountBuffers+0][j]
					buffer1[j] = (*A)[i*amountBuffers+1][j]
					buffer2[j] = (*A)[i*amountBuffers+2][j]
					buffer3[j] = (*A)[i*amountBuffers+3][j]
					buffer4[j] = (*A)[i*amountBuffers+4][j]
					buffer5[j] = (*A)[i*amountBuffers+5][j]
					buffer6[j] = (*A)[i*amountBuffers+6][j]
					buffer7[j] = (*A)[i*amountBuffers+7][j]
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
					for k := 0; k < n; k++ {
						sum00 += buffer0[k] * (*B)[k][j]
						sum01 += buffer1[k] * (*B)[k][j]
						sum02 += buffer2[k] * (*B)[k][j]
						sum03 += buffer3[k] * (*B)[k][j]
						sum04 += buffer4[k] * (*B)[k][j]
						sum05 += buffer5[k] * (*B)[k][j]
						sum06 += buffer6[k] * (*B)[k][j]
						sum07 += buffer7[k] * (*B)[k][j]
					}
					(*C)[i*amountBuffers+0][j] = sum00
					(*C)[i*amountBuffers+1][j] = sum01
					(*C)[i*amountBuffers+2][j] = sum02
					(*C)[i*amountBuffers+3][j] = sum03
					(*C)[i*amountBuffers+4][j] = sum04
					(*C)[i*amountBuffers+5][j] = sum05
					(*C)[i*amountBuffers+6][j] = sum06
					(*C)[i*amountBuffers+7][j] = sum07
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOut16(t *testing.T) {
	if !isSame(mmParallelBufferVarOut16) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOut16(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOut16(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOut16 - with 16 buffers
func mmParallelBufferVarOut16(A, B, C *[][]float64) {
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

func TestParallelBufferVarOut32(t *testing.T) {
	if !isSame(mmParallelBufferVarOut32) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOut32(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOut32(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOut32 - with 32 buffers
func mmParallelBufferVarOut32(A, B, C *[][]float64) {
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
			var sum10, sum11, sum12, sum13, sum14, sum15, sum16, sum17, sum18, sum19 float64
			var sum20, sum21, sum22, sum23, sum24, sum25, sum26, sum27, sum28, sum29 float64
			var sum30, sum31 float64
			// Create buffers
			amountBuffers := 32
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
					buffer16[j] = (*A)[i*amountBuffers+16][j]
					buffer17[j] = (*A)[i*amountBuffers+17][j]
					buffer18[j] = (*A)[i*amountBuffers+18][j]
					buffer19[j] = (*A)[i*amountBuffers+19][j]
					buffer20[j] = (*A)[i*amountBuffers+20][j]
					buffer21[j] = (*A)[i*amountBuffers+21][j]
					buffer22[j] = (*A)[i*amountBuffers+22][j]
					buffer23[j] = (*A)[i*amountBuffers+23][j]
					buffer24[j] = (*A)[i*amountBuffers+24][j]
					buffer25[j] = (*A)[i*amountBuffers+25][j]
					buffer26[j] = (*A)[i*amountBuffers+26][j]
					buffer27[j] = (*A)[i*amountBuffers+27][j]
					buffer28[j] = (*A)[i*amountBuffers+28][j]
					buffer29[j] = (*A)[i*amountBuffers+29][j]
					buffer30[j] = (*A)[i*amountBuffers+30][j]
					buffer31[j] = (*A)[i*amountBuffers+31][j]
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
					sum16 = 0.0
					sum17 = 0.0
					sum18 = 0.0
					sum19 = 0.0
					sum20 = 0.0
					sum21 = 0.0
					sum22 = 0.0
					sum23 = 0.0
					sum24 = 0.0
					sum25 = 0.0
					sum26 = 0.0
					sum27 = 0.0
					sum28 = 0.0
					sum29 = 0.0
					sum30 = 0.0
					sum31 = 0.0
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
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
						sum20 += buffer20[k] * (*B)[k][j]
						sum21 += buffer21[k] * (*B)[k][j]
						sum22 += buffer22[k] * (*B)[k][j]
						sum23 += buffer23[k] * (*B)[k][j]
						sum24 += buffer24[k] * (*B)[k][j]
						sum25 += buffer25[k] * (*B)[k][j]
						sum26 += buffer26[k] * (*B)[k][j]
						sum27 += buffer27[k] * (*B)[k][j]
						sum28 += buffer28[k] * (*B)[k][j]
						sum29 += buffer29[k] * (*B)[k][j]
						sum30 += buffer30[k] * (*B)[k][j]
						sum31 += buffer31[k] * (*B)[k][j]
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
					(*C)[i*amountBuffers+16][j] = sum16
					(*C)[i*amountBuffers+17][j] = sum17
					(*C)[i*amountBuffers+18][j] = sum18
					(*C)[i*amountBuffers+19][j] = sum19
					(*C)[i*amountBuffers+20][j] = sum20
					(*C)[i*amountBuffers+21][j] = sum21
					(*C)[i*amountBuffers+22][j] = sum22
					(*C)[i*amountBuffers+23][j] = sum23
					(*C)[i*amountBuffers+24][j] = sum24
					(*C)[i*amountBuffers+25][j] = sum25
					(*C)[i*amountBuffers+26][j] = sum26
					(*C)[i*amountBuffers+27][j] = sum27
					(*C)[i*amountBuffers+28][j] = sum28
					(*C)[i*amountBuffers+29][j] = sum29
					(*C)[i*amountBuffers+30][j] = sum30
					(*C)[i*amountBuffers+31][j] = sum31
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOut64(t *testing.T) {
	if !isSame(mmParallelBufferVarOut64) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOut64(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOut64(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOut64 - with 64 buffers
func mmParallelBufferVarOut64(A, B, C *[][]float64) {
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
			var sum10, sum11, sum12, sum13, sum14, sum15, sum16, sum17, sum18, sum19 float64
			var sum20, sum21, sum22, sum23, sum24, sum25, sum26, sum27, sum28, sum29 float64
			var sum30, sum31, sum32, sum33, sum34, sum35, sum36, sum37, sum38, sum39 float64
			var sum40, sum41, sum42, sum43, sum44, sum45, sum46, sum47, sum48, sum49 float64
			var sum50, sum51, sum52, sum53, sum54, sum55, sum56, sum57, sum58, sum59 float64
			var sum60, sum61, sum62, sum63 float64
			// Create buffers
			amountBuffers := 64
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
					buffer16[j] = (*A)[i*amountBuffers+16][j]
					buffer17[j] = (*A)[i*amountBuffers+17][j]
					buffer18[j] = (*A)[i*amountBuffers+18][j]
					buffer19[j] = (*A)[i*amountBuffers+19][j]
					buffer20[j] = (*A)[i*amountBuffers+20][j]
					buffer21[j] = (*A)[i*amountBuffers+21][j]
					buffer22[j] = (*A)[i*amountBuffers+22][j]
					buffer23[j] = (*A)[i*amountBuffers+23][j]
					buffer24[j] = (*A)[i*amountBuffers+24][j]
					buffer25[j] = (*A)[i*amountBuffers+25][j]
					buffer26[j] = (*A)[i*amountBuffers+26][j]
					buffer27[j] = (*A)[i*amountBuffers+27][j]
					buffer28[j] = (*A)[i*amountBuffers+28][j]
					buffer29[j] = (*A)[i*amountBuffers+29][j]
					buffer30[j] = (*A)[i*amountBuffers+30][j]
					buffer31[j] = (*A)[i*amountBuffers+31][j]
					buffer32[j] = (*A)[i*amountBuffers+32][j]
					buffer33[j] = (*A)[i*amountBuffers+33][j]
					buffer34[j] = (*A)[i*amountBuffers+34][j]
					buffer35[j] = (*A)[i*amountBuffers+35][j]
					buffer36[j] = (*A)[i*amountBuffers+36][j]
					buffer37[j] = (*A)[i*amountBuffers+37][j]
					buffer38[j] = (*A)[i*amountBuffers+38][j]
					buffer39[j] = (*A)[i*amountBuffers+39][j]
					buffer40[j] = (*A)[i*amountBuffers+40][j]
					buffer41[j] = (*A)[i*amountBuffers+41][j]
					buffer42[j] = (*A)[i*amountBuffers+42][j]
					buffer43[j] = (*A)[i*amountBuffers+43][j]
					buffer44[j] = (*A)[i*amountBuffers+44][j]
					buffer45[j] = (*A)[i*amountBuffers+45][j]
					buffer46[j] = (*A)[i*amountBuffers+46][j]
					buffer47[j] = (*A)[i*amountBuffers+47][j]
					buffer48[j] = (*A)[i*amountBuffers+48][j]
					buffer49[j] = (*A)[i*amountBuffers+49][j]
					buffer50[j] = (*A)[i*amountBuffers+50][j]
					buffer51[j] = (*A)[i*amountBuffers+51][j]
					buffer52[j] = (*A)[i*amountBuffers+52][j]
					buffer53[j] = (*A)[i*amountBuffers+53][j]
					buffer54[j] = (*A)[i*amountBuffers+54][j]
					buffer55[j] = (*A)[i*amountBuffers+55][j]
					buffer56[j] = (*A)[i*amountBuffers+56][j]
					buffer57[j] = (*A)[i*amountBuffers+57][j]
					buffer58[j] = (*A)[i*amountBuffers+58][j]
					buffer59[j] = (*A)[i*amountBuffers+59][j]
					buffer60[j] = (*A)[i*amountBuffers+60][j]
					buffer61[j] = (*A)[i*amountBuffers+61][j]
					buffer62[j] = (*A)[i*amountBuffers+62][j]
					buffer63[j] = (*A)[i*amountBuffers+63][j]
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
					sum16 = 0.0
					sum17 = 0.0
					sum18 = 0.0
					sum19 = 0.0
					sum20 = 0.0
					sum21 = 0.0
					sum22 = 0.0
					sum23 = 0.0
					sum24 = 0.0
					sum25 = 0.0
					sum26 = 0.0
					sum27 = 0.0
					sum28 = 0.0
					sum29 = 0.0
					sum30 = 0.0
					sum31 = 0.0
					sum32 = 0.0
					sum33 = 0.0
					sum34 = 0.0
					sum35 = 0.0
					sum36 = 0.0
					sum37 = 0.0
					sum38 = 0.0
					sum39 = 0.0
					sum40 = 0.0
					sum41 = 0.0
					sum42 = 0.0
					sum43 = 0.0
					sum44 = 0.0
					sum45 = 0.0
					sum46 = 0.0
					sum47 = 0.0
					sum48 = 0.0
					sum49 = 0.0
					sum50 = 0.0
					sum51 = 0.0
					sum52 = 0.0
					sum53 = 0.0
					sum54 = 0.0
					sum55 = 0.0
					sum56 = 0.0
					sum57 = 0.0
					sum58 = 0.0
					sum59 = 0.0
					sum60 = 0.0
					sum61 = 0.0
					sum62 = 0.0
					sum63 = 0.0
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
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
						sum20 += buffer20[k] * (*B)[k][j]
						sum21 += buffer21[k] * (*B)[k][j]
						sum22 += buffer22[k] * (*B)[k][j]
						sum23 += buffer23[k] * (*B)[k][j]
						sum24 += buffer24[k] * (*B)[k][j]
						sum25 += buffer25[k] * (*B)[k][j]
						sum26 += buffer26[k] * (*B)[k][j]
						sum27 += buffer27[k] * (*B)[k][j]
						sum28 += buffer28[k] * (*B)[k][j]
						sum29 += buffer29[k] * (*B)[k][j]
						sum30 += buffer30[k] * (*B)[k][j]
						sum31 += buffer31[k] * (*B)[k][j]
						sum32 += buffer32[k] * (*B)[k][j]
						sum33 += buffer33[k] * (*B)[k][j]
						sum34 += buffer34[k] * (*B)[k][j]
						sum35 += buffer35[k] * (*B)[k][j]
						sum36 += buffer36[k] * (*B)[k][j]
						sum37 += buffer37[k] * (*B)[k][j]
						sum38 += buffer38[k] * (*B)[k][j]
						sum39 += buffer39[k] * (*B)[k][j]
						sum40 += buffer40[k] * (*B)[k][j]
						sum41 += buffer41[k] * (*B)[k][j]
						sum42 += buffer42[k] * (*B)[k][j]
						sum43 += buffer43[k] * (*B)[k][j]
						sum44 += buffer44[k] * (*B)[k][j]
						sum45 += buffer45[k] * (*B)[k][j]
						sum46 += buffer46[k] * (*B)[k][j]
						sum47 += buffer47[k] * (*B)[k][j]
						sum48 += buffer48[k] * (*B)[k][j]
						sum49 += buffer49[k] * (*B)[k][j]
						sum50 += buffer50[k] * (*B)[k][j]
						sum51 += buffer51[k] * (*B)[k][j]
						sum52 += buffer52[k] * (*B)[k][j]
						sum53 += buffer53[k] * (*B)[k][j]
						sum54 += buffer54[k] * (*B)[k][j]
						sum55 += buffer55[k] * (*B)[k][j]
						sum56 += buffer56[k] * (*B)[k][j]
						sum57 += buffer57[k] * (*B)[k][j]
						sum58 += buffer58[k] * (*B)[k][j]
						sum59 += buffer59[k] * (*B)[k][j]
						sum60 += buffer60[k] * (*B)[k][j]
						sum61 += buffer61[k] * (*B)[k][j]
						sum62 += buffer62[k] * (*B)[k][j]
						sum63 += buffer63[k] * (*B)[k][j]
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
					(*C)[i*amountBuffers+16][j] = sum16
					(*C)[i*amountBuffers+17][j] = sum17
					(*C)[i*amountBuffers+18][j] = sum18
					(*C)[i*amountBuffers+19][j] = sum19
					(*C)[i*amountBuffers+20][j] = sum20
					(*C)[i*amountBuffers+21][j] = sum21
					(*C)[i*amountBuffers+22][j] = sum22
					(*C)[i*amountBuffers+23][j] = sum23
					(*C)[i*amountBuffers+24][j] = sum24
					(*C)[i*amountBuffers+25][j] = sum25
					(*C)[i*amountBuffers+26][j] = sum26
					(*C)[i*amountBuffers+27][j] = sum27
					(*C)[i*amountBuffers+28][j] = sum28
					(*C)[i*amountBuffers+29][j] = sum29
					(*C)[i*amountBuffers+30][j] = sum30
					(*C)[i*amountBuffers+31][j] = sum31
					(*C)[i*amountBuffers+32][j] = sum32
					(*C)[i*amountBuffers+33][j] = sum33
					(*C)[i*amountBuffers+34][j] = sum34
					(*C)[i*amountBuffers+35][j] = sum35
					(*C)[i*amountBuffers+36][j] = sum36
					(*C)[i*amountBuffers+37][j] = sum37
					(*C)[i*amountBuffers+38][j] = sum38
					(*C)[i*amountBuffers+39][j] = sum39
					(*C)[i*amountBuffers+40][j] = sum40
					(*C)[i*amountBuffers+41][j] = sum41
					(*C)[i*amountBuffers+42][j] = sum42
					(*C)[i*amountBuffers+43][j] = sum43
					(*C)[i*amountBuffers+44][j] = sum44
					(*C)[i*amountBuffers+45][j] = sum45
					(*C)[i*amountBuffers+46][j] = sum46
					(*C)[i*amountBuffers+47][j] = sum47
					(*C)[i*amountBuffers+48][j] = sum48
					(*C)[i*amountBuffers+49][j] = sum49
					(*C)[i*amountBuffers+50][j] = sum50
					(*C)[i*amountBuffers+51][j] = sum51
					(*C)[i*amountBuffers+52][j] = sum52
					(*C)[i*amountBuffers+53][j] = sum53
					(*C)[i*amountBuffers+54][j] = sum54
					(*C)[i*amountBuffers+55][j] = sum55
					(*C)[i*amountBuffers+56][j] = sum56
					(*C)[i*amountBuffers+57][j] = sum57
					(*C)[i*amountBuffers+58][j] = sum58
					(*C)[i*amountBuffers+59][j] = sum59
					(*C)[i*amountBuffers+60][j] = sum60
					(*C)[i*amountBuffers+61][j] = sum61
					(*C)[i*amountBuffers+62][j] = sum62
					(*C)[i*amountBuffers+63][j] = sum63
				}
			}
		}(t)
	}
	wg.Wait()
}
