package main_test

import (
	"runtime"
	"sync"
	"testing"
)

func TestParallelBuffer2(t *testing.T) {
	if !isSame(mmParallelBuffer2) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBuffer2(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBuffer2(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBuffer2 - with 2 buffers
func mmParallelBuffer2(A, B, C *[][]float64) {
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
					for k := 0; k < n; k++ {
						(*C)[i*amountBuffers+0][j] += buffer0[k] * (*B)[k][j]
						(*C)[i*amountBuffers+1][j] += buffer1[k] * (*B)[k][j]
					}
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBuffer4(t *testing.T) {
	if !isSame(mmParallelBuffer4) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBuffer4(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBuffer4(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBuffer4 - with 4 buffers
func mmParallelBuffer4(A, B, C *[][]float64) {
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
					for k := 0; k < n; k++ {
						(*C)[i*amountBuffers+0][j] += buffer0[k] * (*B)[k][j]
						(*C)[i*amountBuffers+1][j] += buffer1[k] * (*B)[k][j]
						(*C)[i*amountBuffers+2][j] += buffer2[k] * (*B)[k][j]
						(*C)[i*amountBuffers+3][j] += buffer3[k] * (*B)[k][j]
					}
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBuffer8(t *testing.T) {
	if !isSame(mmParallelBuffer8) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBuffer8(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBuffer8(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBuffer8 - with 8 buffers
func mmParallelBuffer8(A, B, C *[][]float64) {
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
					for k := 0; k < n; k++ {
						(*C)[i*amountBuffers+0][j] += buffer0[k] * (*B)[k][j]
						(*C)[i*amountBuffers+1][j] += buffer1[k] * (*B)[k][j]
						(*C)[i*amountBuffers+2][j] += buffer2[k] * (*B)[k][j]
						(*C)[i*amountBuffers+3][j] += buffer3[k] * (*B)[k][j]
						(*C)[i*amountBuffers+4][j] += buffer4[k] * (*B)[k][j]
						(*C)[i*amountBuffers+5][j] += buffer5[k] * (*B)[k][j]
						(*C)[i*amountBuffers+6][j] += buffer6[k] * (*B)[k][j]
						(*C)[i*amountBuffers+7][j] += buffer7[k] * (*B)[k][j]
					}
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBuffer16(t *testing.T) {
	if !isSame(mmParallelBuffer16) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBuffer16(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBuffer16(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBuffer16 - with 16 buffers
func mmParallelBuffer16(A, B, C *[][]float64) {
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
					for k := 0; k < n; k++ {
						(*C)[i*amountBuffers+0][j] += buffer00[k] * (*B)[k][j]
						(*C)[i*amountBuffers+1][j] += buffer01[k] * (*B)[k][j]
						(*C)[i*amountBuffers+2][j] += buffer02[k] * (*B)[k][j]
						(*C)[i*amountBuffers+3][j] += buffer03[k] * (*B)[k][j]
						(*C)[i*amountBuffers+4][j] += buffer04[k] * (*B)[k][j]
						(*C)[i*amountBuffers+5][j] += buffer05[k] * (*B)[k][j]
						(*C)[i*amountBuffers+6][j] += buffer06[k] * (*B)[k][j]
						(*C)[i*amountBuffers+7][j] += buffer07[k] * (*B)[k][j]
						(*C)[i*amountBuffers+8][j] += buffer08[k] * (*B)[k][j]
						(*C)[i*amountBuffers+9][j] += buffer09[k] * (*B)[k][j]
						(*C)[i*amountBuffers+10][j] += buffer10[k] * (*B)[k][j]
						(*C)[i*amountBuffers+11][j] += buffer11[k] * (*B)[k][j]
						(*C)[i*amountBuffers+12][j] += buffer12[k] * (*B)[k][j]
						(*C)[i*amountBuffers+13][j] += buffer13[k] * (*B)[k][j]
						(*C)[i*amountBuffers+14][j] += buffer14[k] * (*B)[k][j]
						(*C)[i*amountBuffers+15][j] += buffer15[k] * (*B)[k][j]
					}
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBuffer32(t *testing.T) {
	if !isSame(mmParallelBuffer32) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBuffer32(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBuffer32(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBuffer32 - with 32 buffers
func mmParallelBuffer32(A, B, C *[][]float64) {
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
					for k := 0; k < n; k++ {
						(*C)[i*amountBuffers+0][j] += buffer00[k] * (*B)[k][j]
						(*C)[i*amountBuffers+1][j] += buffer01[k] * (*B)[k][j]
						(*C)[i*amountBuffers+2][j] += buffer02[k] * (*B)[k][j]
						(*C)[i*amountBuffers+3][j] += buffer03[k] * (*B)[k][j]
						(*C)[i*amountBuffers+4][j] += buffer04[k] * (*B)[k][j]
						(*C)[i*amountBuffers+5][j] += buffer05[k] * (*B)[k][j]
						(*C)[i*amountBuffers+6][j] += buffer06[k] * (*B)[k][j]
						(*C)[i*amountBuffers+7][j] += buffer07[k] * (*B)[k][j]
						(*C)[i*amountBuffers+8][j] += buffer08[k] * (*B)[k][j]
						(*C)[i*amountBuffers+9][j] += buffer09[k] * (*B)[k][j]
						(*C)[i*amountBuffers+10][j] += buffer10[k] * (*B)[k][j]
						(*C)[i*amountBuffers+11][j] += buffer11[k] * (*B)[k][j]
						(*C)[i*amountBuffers+12][j] += buffer12[k] * (*B)[k][j]
						(*C)[i*amountBuffers+13][j] += buffer13[k] * (*B)[k][j]
						(*C)[i*amountBuffers+14][j] += buffer14[k] * (*B)[k][j]
						(*C)[i*amountBuffers+15][j] += buffer15[k] * (*B)[k][j]
						(*C)[i*amountBuffers+16][j] += buffer16[k] * (*B)[k][j]
						(*C)[i*amountBuffers+17][j] += buffer17[k] * (*B)[k][j]
						(*C)[i*amountBuffers+18][j] += buffer18[k] * (*B)[k][j]
						(*C)[i*amountBuffers+19][j] += buffer19[k] * (*B)[k][j]
						(*C)[i*amountBuffers+20][j] += buffer20[k] * (*B)[k][j]
						(*C)[i*amountBuffers+21][j] += buffer21[k] * (*B)[k][j]
						(*C)[i*amountBuffers+22][j] += buffer22[k] * (*B)[k][j]
						(*C)[i*amountBuffers+23][j] += buffer23[k] * (*B)[k][j]
						(*C)[i*amountBuffers+24][j] += buffer24[k] * (*B)[k][j]
						(*C)[i*amountBuffers+25][j] += buffer25[k] * (*B)[k][j]
						(*C)[i*amountBuffers+26][j] += buffer26[k] * (*B)[k][j]
						(*C)[i*amountBuffers+27][j] += buffer27[k] * (*B)[k][j]
						(*C)[i*amountBuffers+28][j] += buffer28[k] * (*B)[k][j]
						(*C)[i*amountBuffers+29][j] += buffer29[k] * (*B)[k][j]
						(*C)[i*amountBuffers+30][j] += buffer30[k] * (*B)[k][j]
						(*C)[i*amountBuffers+31][j] += buffer31[k] * (*B)[k][j]
					}
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBuffer64(t *testing.T) {
	if !isSame(mmParallelBuffer64) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBuffer64(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBuffer64(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBuffer64 - with 64 buffers
func mmParallelBuffer64(A, B, C *[][]float64) {
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
					for k := 0; k < n; k++ {
						(*C)[i*amountBuffers+0][j] += buffer00[k] * (*B)[k][j]
						(*C)[i*amountBuffers+1][j] += buffer01[k] * (*B)[k][j]
						(*C)[i*amountBuffers+2][j] += buffer02[k] * (*B)[k][j]
						(*C)[i*amountBuffers+3][j] += buffer03[k] * (*B)[k][j]
						(*C)[i*amountBuffers+4][j] += buffer04[k] * (*B)[k][j]
						(*C)[i*amountBuffers+5][j] += buffer05[k] * (*B)[k][j]
						(*C)[i*amountBuffers+6][j] += buffer06[k] * (*B)[k][j]
						(*C)[i*amountBuffers+7][j] += buffer07[k] * (*B)[k][j]
						(*C)[i*amountBuffers+8][j] += buffer08[k] * (*B)[k][j]
						(*C)[i*amountBuffers+9][j] += buffer09[k] * (*B)[k][j]
						(*C)[i*amountBuffers+10][j] += buffer10[k] * (*B)[k][j]
						(*C)[i*amountBuffers+11][j] += buffer11[k] * (*B)[k][j]
						(*C)[i*amountBuffers+12][j] += buffer12[k] * (*B)[k][j]
						(*C)[i*amountBuffers+13][j] += buffer13[k] * (*B)[k][j]
						(*C)[i*amountBuffers+14][j] += buffer14[k] * (*B)[k][j]
						(*C)[i*amountBuffers+15][j] += buffer15[k] * (*B)[k][j]
						(*C)[i*amountBuffers+16][j] += buffer16[k] * (*B)[k][j]
						(*C)[i*amountBuffers+17][j] += buffer17[k] * (*B)[k][j]
						(*C)[i*amountBuffers+18][j] += buffer18[k] * (*B)[k][j]
						(*C)[i*amountBuffers+19][j] += buffer19[k] * (*B)[k][j]
						(*C)[i*amountBuffers+20][j] += buffer20[k] * (*B)[k][j]
						(*C)[i*amountBuffers+21][j] += buffer21[k] * (*B)[k][j]
						(*C)[i*amountBuffers+22][j] += buffer22[k] * (*B)[k][j]
						(*C)[i*amountBuffers+23][j] += buffer23[k] * (*B)[k][j]
						(*C)[i*amountBuffers+24][j] += buffer24[k] * (*B)[k][j]
						(*C)[i*amountBuffers+25][j] += buffer25[k] * (*B)[k][j]
						(*C)[i*amountBuffers+26][j] += buffer26[k] * (*B)[k][j]
						(*C)[i*amountBuffers+27][j] += buffer27[k] * (*B)[k][j]
						(*C)[i*amountBuffers+28][j] += buffer28[k] * (*B)[k][j]
						(*C)[i*amountBuffers+29][j] += buffer29[k] * (*B)[k][j]
						(*C)[i*amountBuffers+30][j] += buffer30[k] * (*B)[k][j]
						(*C)[i*amountBuffers+31][j] += buffer31[k] * (*B)[k][j]
						(*C)[i*amountBuffers+32][j] += buffer32[k] * (*B)[k][j]
						(*C)[i*amountBuffers+33][j] += buffer33[k] * (*B)[k][j]
						(*C)[i*amountBuffers+34][j] += buffer34[k] * (*B)[k][j]
						(*C)[i*amountBuffers+35][j] += buffer35[k] * (*B)[k][j]
						(*C)[i*amountBuffers+36][j] += buffer36[k] * (*B)[k][j]
						(*C)[i*amountBuffers+37][j] += buffer37[k] * (*B)[k][j]
						(*C)[i*amountBuffers+38][j] += buffer38[k] * (*B)[k][j]
						(*C)[i*amountBuffers+39][j] += buffer39[k] * (*B)[k][j]
						(*C)[i*amountBuffers+40][j] += buffer40[k] * (*B)[k][j]
						(*C)[i*amountBuffers+41][j] += buffer41[k] * (*B)[k][j]
						(*C)[i*amountBuffers+42][j] += buffer42[k] * (*B)[k][j]
						(*C)[i*amountBuffers+43][j] += buffer43[k] * (*B)[k][j]
						(*C)[i*amountBuffers+44][j] += buffer44[k] * (*B)[k][j]
						(*C)[i*amountBuffers+45][j] += buffer45[k] * (*B)[k][j]
						(*C)[i*amountBuffers+46][j] += buffer46[k] * (*B)[k][j]
						(*C)[i*amountBuffers+47][j] += buffer47[k] * (*B)[k][j]
						(*C)[i*amountBuffers+48][j] += buffer48[k] * (*B)[k][j]
						(*C)[i*amountBuffers+49][j] += buffer49[k] * (*B)[k][j]
						(*C)[i*amountBuffers+50][j] += buffer50[k] * (*B)[k][j]
						(*C)[i*amountBuffers+51][j] += buffer51[k] * (*B)[k][j]
						(*C)[i*amountBuffers+52][j] += buffer52[k] * (*B)[k][j]
						(*C)[i*amountBuffers+53][j] += buffer53[k] * (*B)[k][j]
						(*C)[i*amountBuffers+54][j] += buffer54[k] * (*B)[k][j]
						(*C)[i*amountBuffers+55][j] += buffer55[k] * (*B)[k][j]
						(*C)[i*amountBuffers+56][j] += buffer56[k] * (*B)[k][j]
						(*C)[i*amountBuffers+57][j] += buffer57[k] * (*B)[k][j]
						(*C)[i*amountBuffers+58][j] += buffer58[k] * (*B)[k][j]
						(*C)[i*amountBuffers+59][j] += buffer59[k] * (*B)[k][j]
						(*C)[i*amountBuffers+60][j] += buffer60[k] * (*B)[k][j]
						(*C)[i*amountBuffers+61][j] += buffer61[k] * (*B)[k][j]
						(*C)[i*amountBuffers+62][j] += buffer62[k] * (*B)[k][j]
						(*C)[i*amountBuffers+63][j] += buffer63[k] * (*B)[k][j]
					}
				}
			}
		}(t)
	}
	wg.Wait()
}
