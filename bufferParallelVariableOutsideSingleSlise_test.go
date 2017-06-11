package main_test

import (
	"math"
	"runtime"
	"sync"
	"testing"
)

func TestParallelBufferVarOutSingleSlise2(t *testing.T) {
	if !isSameSingle(mmParallelBufferVarOutSingleSlise2) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOutSingleSlise2(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOutSingleSlise2(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOutSingleSlise2 - with 2 buffers
func mmParallelBufferVarOutSingleSlise2(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
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
			var sum0, sum1 float64
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
					buffer0[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer1[j] = (*A)[(i*amountBuffers+1)+j*n]
				}
				for j := 0; j < n; j++ {
					sum0 = 0.0
					sum1 = 0.0
					// Create a pointer on matrix [B]
					b := (*B)[j*n : j*n+n] // <---- Mark
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * b[k] // <---- Mark
						sum1 += buffer1[k] * b[k] // <---- Mark
					}
					(*C)[(i*amountBuffers+0)+j*n] = sum0
					(*C)[(i*amountBuffers+1)+j*n] = sum1
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOutSingleSlise4(t *testing.T) {
	if !isSameSingle(mmParallelBufferVarOutSingleSlise4) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOutSingleSlise4(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOutSingleSlise4(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOutSingleSlise4 - with 4 buffers
func mmParallelBufferVarOutSingleSlise4(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
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
					buffer0[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer1[j] = (*A)[(i*amountBuffers+1)+j*n]
					buffer2[j] = (*A)[(i*amountBuffers+2)+j*n]
					buffer3[j] = (*A)[(i*amountBuffers+3)+j*n]
				}
				for j := 0; j < n; j++ {
					sum00 = 0.0
					sum01 = 0.0
					sum02 = 0.0
					sum03 = 0.0
					b := (*B)[j*n : j*n+n]
					for k := 0; k < n; k++ {
						sum00 += buffer0[k] * b[k]
						sum01 += buffer1[k] * b[k]
						sum02 += buffer2[k] * b[k]
						sum03 += buffer3[k] * b[k]

					}
					(*C)[(i*amountBuffers+0)+j*n] = sum00
					(*C)[(i*amountBuffers+1)+j*n] = sum01
					(*C)[(i*amountBuffers+2)+j*n] = sum02
					(*C)[(i*amountBuffers+3)+j*n] = sum03
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOutSingleSlise8(t *testing.T) {
	if !isSameSingle(mmParallelBufferVarOutSingleSlise8) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOutSingleSlise8(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOutSingleSlise8(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOutSingleSlise8 - with 8 buffers
func mmParallelBufferVarOutSingleSlise8(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
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
					buffer0[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer1[j] = (*A)[(i*amountBuffers+1)+j*n]
					buffer2[j] = (*A)[(i*amountBuffers+2)+j*n]
					buffer3[j] = (*A)[(i*amountBuffers+3)+j*n]
					buffer4[j] = (*A)[(i*amountBuffers+4)+j*n]
					buffer5[j] = (*A)[(i*amountBuffers+5)+j*n]
					buffer6[j] = (*A)[(i*amountBuffers+6)+j*n]
					buffer7[j] = (*A)[(i*amountBuffers+7)+j*n]
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
					b := (*B)[j*n : j*n+n]
					for k := 0; k < n; k++ {
						sum00 += buffer0[k] * b[k] // (*B)[k+j*n]
						sum01 += buffer1[k] * b[k] // (*B)[k+j*n]
						sum02 += buffer2[k] * b[k] // (*B)[k+j*n]
						sum03 += buffer3[k] * b[k] // (*B)[k+j*n]
						sum04 += buffer4[k] * b[k] // (*B)[k+j*n]
						sum05 += buffer5[k] * b[k] // (*B)[k+j*n]
						sum06 += buffer6[k] * b[k] // (*B)[k+j*n]
						sum07 += buffer7[k] * b[k] // (*B)[k+j*n]
					}
					(*C)[(i*amountBuffers+0)+j*n] = sum00
					(*C)[(i*amountBuffers+1)+j*n] = sum01
					(*C)[(i*amountBuffers+2)+j*n] = sum02
					(*C)[(i*amountBuffers+3)+j*n] = sum03
					(*C)[(i*amountBuffers+4)+j*n] = sum04
					(*C)[(i*amountBuffers+5)+j*n] = sum05
					(*C)[(i*amountBuffers+6)+j*n] = sum06
					(*C)[(i*amountBuffers+7)+j*n] = sum07
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOutSingleSlise16(t *testing.T) {
	if !isSameSingle(mmParallelBufferVarOutSingleSlise16) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOutSingleSlise16(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOutSingleSlise16(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOutSingleSlise16 - with 16 buffers
func mmParallelBufferVarOutSingleSlise16(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
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
					buffer00[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer01[j] = (*A)[(i*amountBuffers+1)+j*n]
					buffer02[j] = (*A)[(i*amountBuffers+2)+j*n]
					buffer03[j] = (*A)[(i*amountBuffers+3)+j*n]
					buffer04[j] = (*A)[(i*amountBuffers+4)+j*n]
					buffer05[j] = (*A)[(i*amountBuffers+5)+j*n]
					buffer06[j] = (*A)[(i*amountBuffers+6)+j*n]
					buffer07[j] = (*A)[(i*amountBuffers+7)+j*n]
					buffer08[j] = (*A)[(i*amountBuffers+8)+j*n]
					buffer09[j] = (*A)[(i*amountBuffers+9)+j*n]
					buffer10[j] = (*A)[(i*amountBuffers+10)+j*n]
					buffer11[j] = (*A)[(i*amountBuffers+11)+j*n]
					buffer12[j] = (*A)[(i*amountBuffers+12)+j*n]
					buffer13[j] = (*A)[(i*amountBuffers+13)+j*n]
					buffer14[j] = (*A)[(i*amountBuffers+14)+j*n]
					buffer15[j] = (*A)[(i*amountBuffers+15)+j*n]
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
					b := (*B)[j*n : j*n+n]
					for k := 0; k < n; k++ {
						sum00 += buffer00[k] * b[k] // (*B)[k+j*n]
						sum01 += buffer01[k] * b[k] // (*B)[k+j*n]
						sum02 += buffer02[k] * b[k] // (*B)[k+j*n]
						sum03 += buffer03[k] * b[k] // (*B)[k+j*n]
						sum04 += buffer04[k] * b[k] // (*B)[k+j*n]
						sum05 += buffer05[k] * b[k] // (*B)[k+j*n]
						sum06 += buffer06[k] * b[k] // (*B)[k+j*n]
						sum07 += buffer07[k] * b[k] // (*B)[k+j*n]
						sum08 += buffer08[k] * b[k] // (*B)[k+j*n]
						sum09 += buffer09[k] * b[k] // (*B)[k+j*n]
						sum10 += buffer10[k] * b[k] // (*B)[k+j*n]
						sum11 += buffer11[k] * b[k] // (*B)[k+j*n]
						sum12 += buffer12[k] * b[k] // (*B)[k+j*n]
						sum13 += buffer13[k] * b[k] // (*B)[k+j*n]
						sum14 += buffer14[k] * b[k] // (*B)[k+j*n]
						sum15 += buffer15[k] * b[k] // (*B)[k+j*n]
					}
					(*C)[(i*amountBuffers+0)+j*n] = sum00
					(*C)[(i*amountBuffers+1)+j*n] = sum01
					(*C)[(i*amountBuffers+2)+j*n] = sum02
					(*C)[(i*amountBuffers+3)+j*n] = sum03
					(*C)[(i*amountBuffers+4)+j*n] = sum04
					(*C)[(i*amountBuffers+5)+j*n] = sum05
					(*C)[(i*amountBuffers+6)+j*n] = sum06
					(*C)[(i*amountBuffers+7)+j*n] = sum07
					(*C)[(i*amountBuffers+8)+j*n] = sum08
					(*C)[(i*amountBuffers+9)+j*n] = sum09
					(*C)[(i*amountBuffers+10)+j*n] = sum10
					(*C)[(i*amountBuffers+11)+j*n] = sum11
					(*C)[(i*amountBuffers+12)+j*n] = sum12
					(*C)[(i*amountBuffers+13)+j*n] = sum13
					(*C)[(i*amountBuffers+14)+j*n] = sum14
					(*C)[(i*amountBuffers+15)+j*n] = sum15
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOutSingleSlise32(t *testing.T) {
	if !isSameSingle(mmParallelBufferVarOutSingleSlise32) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOutSingleSlise32(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOutSingleSlise32(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOutSingleSlise32 - with 32 buffers
func mmParallelBufferVarOutSingleSlise32(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
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
					buffer00[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer01[j] = (*A)[(i*amountBuffers+1)+j*n]
					buffer02[j] = (*A)[(i*amountBuffers+2)+j*n]
					buffer03[j] = (*A)[(i*amountBuffers+3)+j*n]
					buffer04[j] = (*A)[(i*amountBuffers+4)+j*n]
					buffer05[j] = (*A)[(i*amountBuffers+5)+j*n]
					buffer06[j] = (*A)[(i*amountBuffers+6)+j*n]
					buffer07[j] = (*A)[(i*amountBuffers+7)+j*n]
					buffer08[j] = (*A)[(i*amountBuffers+8)+j*n]
					buffer09[j] = (*A)[(i*amountBuffers+9)+j*n]
					buffer10[j] = (*A)[(i*amountBuffers+10)+j*n]
					buffer11[j] = (*A)[(i*amountBuffers+11)+j*n]
					buffer12[j] = (*A)[(i*amountBuffers+12)+j*n]
					buffer13[j] = (*A)[(i*amountBuffers+13)+j*n]
					buffer14[j] = (*A)[(i*amountBuffers+14)+j*n]
					buffer15[j] = (*A)[(i*amountBuffers+15)+j*n]
					buffer16[j] = (*A)[(i*amountBuffers+16)+j*n]
					buffer17[j] = (*A)[(i*amountBuffers+17)+j*n]
					buffer18[j] = (*A)[(i*amountBuffers+18)+j*n]
					buffer19[j] = (*A)[(i*amountBuffers+19)+j*n]
					buffer20[j] = (*A)[(i*amountBuffers+20)+j*n]
					buffer21[j] = (*A)[(i*amountBuffers+21)+j*n]
					buffer22[j] = (*A)[(i*amountBuffers+22)+j*n]
					buffer23[j] = (*A)[(i*amountBuffers+23)+j*n]
					buffer24[j] = (*A)[(i*amountBuffers+24)+j*n]
					buffer25[j] = (*A)[(i*amountBuffers+25)+j*n]
					buffer26[j] = (*A)[(i*amountBuffers+26)+j*n]
					buffer27[j] = (*A)[(i*amountBuffers+27)+j*n]
					buffer28[j] = (*A)[(i*amountBuffers+28)+j*n]
					buffer29[j] = (*A)[(i*amountBuffers+29)+j*n]
					buffer30[j] = (*A)[(i*amountBuffers+30)+j*n]
					buffer31[j] = (*A)[(i*amountBuffers+31)+j*n]
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
					b := (*B)[j*n : j*n+n]
					for k := 0; k < n; k++ {
						sum00 += buffer00[k] * b[k] // (*B)[k+j*n]
						sum01 += buffer01[k] * b[k] // (*B)[k+j*n]
						sum02 += buffer02[k] * b[k] // (*B)[k+j*n]
						sum03 += buffer03[k] * b[k] // (*B)[k+j*n]
						sum04 += buffer04[k] * b[k] // (*B)[k+j*n]
						sum05 += buffer05[k] * b[k] // (*B)[k+j*n]
						sum06 += buffer06[k] * b[k] // (*B)[k+j*n]
						sum07 += buffer07[k] * b[k] // (*B)[k+j*n]
						sum08 += buffer08[k] * b[k] // (*B)[k+j*n]
						sum09 += buffer09[k] * b[k] // (*B)[k+j*n]
						sum10 += buffer10[k] * b[k] // (*B)[k+j*n]
						sum11 += buffer11[k] * b[k] // (*B)[k+j*n]
						sum12 += buffer12[k] * b[k] // (*B)[k+j*n]
						sum13 += buffer13[k] * b[k] // (*B)[k+j*n]
						sum14 += buffer14[k] * b[k] // (*B)[k+j*n]
						sum15 += buffer15[k] * b[k] // (*B)[k+j*n]
						sum16 += buffer16[k] * b[k] // (*B)[k+j*n]
						sum17 += buffer17[k] * b[k] // (*B)[k+j*n]
						sum18 += buffer18[k] * b[k] // (*B)[k+j*n]
						sum19 += buffer19[k] * b[k] // (*B)[k+j*n]
						sum20 += buffer20[k] * b[k] // (*B)[k+j*n]
						sum21 += buffer21[k] * b[k] // (*B)[k+j*n]
						sum22 += buffer22[k] * b[k] // (*B)[k+j*n]
						sum23 += buffer23[k] * b[k] // (*B)[k+j*n]
						sum24 += buffer24[k] * b[k] // (*B)[k+j*n]
						sum25 += buffer25[k] * b[k] // (*B)[k+j*n]
						sum26 += buffer26[k] * b[k] // (*B)[k+j*n]
						sum27 += buffer27[k] * b[k] // (*B)[k+j*n]
						sum28 += buffer28[k] * b[k] // (*B)[k+j*n]
						sum29 += buffer29[k] * b[k] // (*B)[k+j*n]
						sum30 += buffer30[k] * b[k] // (*B)[k+j*n]
						sum31 += buffer31[k] * b[k] // (*B)[k+j*n]
					}
					(*C)[(i*amountBuffers+0)+j*n] = sum00
					(*C)[(i*amountBuffers+1)+j*n] = sum01
					(*C)[(i*amountBuffers+2)+j*n] = sum02
					(*C)[(i*amountBuffers+3)+j*n] = sum03
					(*C)[(i*amountBuffers+4)+j*n] = sum04
					(*C)[(i*amountBuffers+5)+j*n] = sum05
					(*C)[(i*amountBuffers+6)+j*n] = sum06
					(*C)[(i*amountBuffers+7)+j*n] = sum07
					(*C)[(i*amountBuffers+8)+j*n] = sum08
					(*C)[(i*amountBuffers+9)+j*n] = sum09
					(*C)[(i*amountBuffers+10)+j*n] = sum10
					(*C)[(i*amountBuffers+11)+j*n] = sum11
					(*C)[(i*amountBuffers+12)+j*n] = sum12
					(*C)[(i*amountBuffers+13)+j*n] = sum13
					(*C)[(i*amountBuffers+14)+j*n] = sum14
					(*C)[(i*amountBuffers+15)+j*n] = sum15
					(*C)[(i*amountBuffers+16)+j*n] = sum16
					(*C)[(i*amountBuffers+17)+j*n] = sum17
					(*C)[(i*amountBuffers+18)+j*n] = sum18
					(*C)[(i*amountBuffers+19)+j*n] = sum19
					(*C)[(i*amountBuffers+20)+j*n] = sum20
					(*C)[(i*amountBuffers+21)+j*n] = sum21
					(*C)[(i*amountBuffers+22)+j*n] = sum22
					(*C)[(i*amountBuffers+23)+j*n] = sum23
					(*C)[(i*amountBuffers+24)+j*n] = sum24
					(*C)[(i*amountBuffers+25)+j*n] = sum25
					(*C)[(i*amountBuffers+26)+j*n] = sum26
					(*C)[(i*amountBuffers+27)+j*n] = sum27
					(*C)[(i*amountBuffers+28)+j*n] = sum28
					(*C)[(i*amountBuffers+29)+j*n] = sum29
					(*C)[(i*amountBuffers+30)+j*n] = sum30
					(*C)[(i*amountBuffers+31)+j*n] = sum31
				}
			}
		}(t)
	}
	wg.Wait()
}

func TestParallelBufferVarOutSingleSlise64(t *testing.T) {
	if !isSameSingle(mmParallelBufferVarOutSingleSlise64) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkParallelBufferVarOutSingleSlise64(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmParallelBufferVarOutSingleSlise64(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmParallelBufferVarOutSingleSlise64 - with 64 buffers
func mmParallelBufferVarOutSingleSlise64(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
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
					buffer00[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer01[j] = (*A)[(i*amountBuffers+1)+j*n]
					buffer02[j] = (*A)[(i*amountBuffers+2)+j*n]
					buffer03[j] = (*A)[(i*amountBuffers+3)+j*n]
					buffer04[j] = (*A)[(i*amountBuffers+4)+j*n]
					buffer05[j] = (*A)[(i*amountBuffers+5)+j*n]
					buffer06[j] = (*A)[(i*amountBuffers+6)+j*n]
					buffer07[j] = (*A)[(i*amountBuffers+7)+j*n]
					buffer08[j] = (*A)[(i*amountBuffers+8)+j*n]
					buffer09[j] = (*A)[(i*amountBuffers+9)+j*n]
					buffer10[j] = (*A)[(i*amountBuffers+10)+j*n]
					buffer11[j] = (*A)[(i*amountBuffers+11)+j*n]
					buffer12[j] = (*A)[(i*amountBuffers+12)+j*n]
					buffer13[j] = (*A)[(i*amountBuffers+13)+j*n]
					buffer14[j] = (*A)[(i*amountBuffers+14)+j*n]
					buffer15[j] = (*A)[(i*amountBuffers+15)+j*n]
					buffer16[j] = (*A)[(i*amountBuffers+16)+j*n]
					buffer17[j] = (*A)[(i*amountBuffers+17)+j*n]
					buffer18[j] = (*A)[(i*amountBuffers+18)+j*n]
					buffer19[j] = (*A)[(i*amountBuffers+19)+j*n]
					buffer20[j] = (*A)[(i*amountBuffers+20)+j*n]
					buffer21[j] = (*A)[(i*amountBuffers+21)+j*n]
					buffer22[j] = (*A)[(i*amountBuffers+22)+j*n]
					buffer23[j] = (*A)[(i*amountBuffers+23)+j*n]
					buffer24[j] = (*A)[(i*amountBuffers+24)+j*n]
					buffer25[j] = (*A)[(i*amountBuffers+25)+j*n]
					buffer26[j] = (*A)[(i*amountBuffers+26)+j*n]
					buffer27[j] = (*A)[(i*amountBuffers+27)+j*n]
					buffer28[j] = (*A)[(i*amountBuffers+28)+j*n]
					buffer29[j] = (*A)[(i*amountBuffers+29)+j*n]
					buffer30[j] = (*A)[(i*amountBuffers+30)+j*n]
					buffer31[j] = (*A)[(i*amountBuffers+31)+j*n]
					buffer32[j] = (*A)[(i*amountBuffers+32)+j*n]
					buffer33[j] = (*A)[(i*amountBuffers+33)+j*n]
					buffer34[j] = (*A)[(i*amountBuffers+34)+j*n]
					buffer35[j] = (*A)[(i*amountBuffers+35)+j*n]
					buffer36[j] = (*A)[(i*amountBuffers+36)+j*n]
					buffer37[j] = (*A)[(i*amountBuffers+37)+j*n]
					buffer38[j] = (*A)[(i*amountBuffers+38)+j*n]
					buffer39[j] = (*A)[(i*amountBuffers+39)+j*n]
					buffer40[j] = (*A)[(i*amountBuffers+40)+j*n]
					buffer41[j] = (*A)[(i*amountBuffers+41)+j*n]
					buffer42[j] = (*A)[(i*amountBuffers+42)+j*n]
					buffer43[j] = (*A)[(i*amountBuffers+43)+j*n]
					buffer44[j] = (*A)[(i*amountBuffers+44)+j*n]
					buffer45[j] = (*A)[(i*amountBuffers+45)+j*n]
					buffer46[j] = (*A)[(i*amountBuffers+46)+j*n]
					buffer47[j] = (*A)[(i*amountBuffers+47)+j*n]
					buffer48[j] = (*A)[(i*amountBuffers+48)+j*n]
					buffer49[j] = (*A)[(i*amountBuffers+49)+j*n]
					buffer50[j] = (*A)[(i*amountBuffers+50)+j*n]
					buffer51[j] = (*A)[(i*amountBuffers+51)+j*n]
					buffer52[j] = (*A)[(i*amountBuffers+52)+j*n]
					buffer53[j] = (*A)[(i*amountBuffers+53)+j*n]
					buffer54[j] = (*A)[(i*amountBuffers+54)+j*n]
					buffer55[j] = (*A)[(i*amountBuffers+55)+j*n]
					buffer56[j] = (*A)[(i*amountBuffers+56)+j*n]
					buffer57[j] = (*A)[(i*amountBuffers+57)+j*n]
					buffer58[j] = (*A)[(i*amountBuffers+58)+j*n]
					buffer59[j] = (*A)[(i*amountBuffers+59)+j*n]
					buffer60[j] = (*A)[(i*amountBuffers+60)+j*n]
					buffer61[j] = (*A)[(i*amountBuffers+61)+j*n]
					buffer62[j] = (*A)[(i*amountBuffers+62)+j*n]
					buffer63[j] = (*A)[(i*amountBuffers+63)+j*n]
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
					b := (*B)[j*n : j*n+n]
					for k := 0; k < n; k++ {
						sum00 += buffer00[k] * b[k] // (*B)[k+j*n]
						sum01 += buffer01[k] * b[k] // (*B)[k+j*n]
						sum02 += buffer02[k] * b[k] // (*B)[k+j*n]
						sum03 += buffer03[k] * b[k] // (*B)[k+j*n]
						sum04 += buffer04[k] * b[k] // (*B)[k+j*n]
						sum05 += buffer05[k] * b[k] // (*B)[k+j*n]
						sum06 += buffer06[k] * b[k] // (*B)[k+j*n]
						sum07 += buffer07[k] * b[k] // (*B)[k+j*n]
						sum08 += buffer08[k] * b[k] // (*B)[k+j*n]
						sum09 += buffer09[k] * b[k] // (*B)[k+j*n]
						sum10 += buffer10[k] * b[k] // (*B)[k+j*n]
						sum11 += buffer11[k] * b[k] // (*B)[k+j*n]
						sum12 += buffer12[k] * b[k] // (*B)[k+j*n]
						sum13 += buffer13[k] * b[k] // (*B)[k+j*n]
						sum14 += buffer14[k] * b[k] // (*B)[k+j*n]
						sum15 += buffer15[k] * b[k] // (*B)[k+j*n]
						sum16 += buffer16[k] * b[k] // (*B)[k+j*n]
						sum17 += buffer17[k] * b[k] // (*B)[k+j*n]
						sum18 += buffer18[k] * b[k] // (*B)[k+j*n]
						sum19 += buffer19[k] * b[k] // (*B)[k+j*n]
						sum20 += buffer20[k] * b[k] // (*B)[k+j*n]
						sum21 += buffer21[k] * b[k] // (*B)[k+j*n]
						sum22 += buffer22[k] * b[k] // (*B)[k+j*n]
						sum23 += buffer23[k] * b[k] // (*B)[k+j*n]
						sum24 += buffer24[k] * b[k] // (*B)[k+j*n]
						sum25 += buffer25[k] * b[k] // (*B)[k+j*n]
						sum26 += buffer26[k] * b[k] // (*B)[k+j*n]
						sum27 += buffer27[k] * b[k] // (*B)[k+j*n]
						sum28 += buffer28[k] * b[k] // (*B)[k+j*n]
						sum29 += buffer29[k] * b[k] // (*B)[k+j*n]
						sum30 += buffer30[k] * b[k] // (*B)[k+j*n]
						sum31 += buffer31[k] * b[k] // (*B)[k+j*n]
						sum32 += buffer32[k] * b[k] // (*B)[k+j*n]
						sum33 += buffer33[k] * b[k] // (*B)[k+j*n]
						sum34 += buffer34[k] * b[k] // (*B)[k+j*n]
						sum35 += buffer35[k] * b[k] // (*B)[k+j*n]
						sum36 += buffer36[k] * b[k] // (*B)[k+j*n]
						sum37 += buffer37[k] * b[k] // (*B)[k+j*n]
						sum38 += buffer38[k] * b[k] // (*B)[k+j*n]
						sum39 += buffer39[k] * b[k] // (*B)[k+j*n]
						sum40 += buffer40[k] * b[k] // (*B)[k+j*n]
						sum41 += buffer41[k] * b[k] // (*B)[k+j*n]
						sum42 += buffer42[k] * b[k] // (*B)[k+j*n]
						sum43 += buffer43[k] * b[k] // (*B)[k+j*n]
						sum44 += buffer44[k] * b[k] // (*B)[k+j*n]
						sum45 += buffer45[k] * b[k] // (*B)[k+j*n]
						sum46 += buffer46[k] * b[k] // (*B)[k+j*n]
						sum47 += buffer47[k] * b[k] // (*B)[k+j*n]
						sum48 += buffer48[k] * b[k] // (*B)[k+j*n]
						sum49 += buffer49[k] * b[k] // (*B)[k+j*n]
						sum50 += buffer50[k] * b[k] // (*B)[k+j*n]
						sum51 += buffer51[k] * b[k] // (*B)[k+j*n]
						sum52 += buffer52[k] * b[k] // (*B)[k+j*n]
						sum53 += buffer53[k] * b[k] // (*B)[k+j*n]
						sum54 += buffer54[k] * b[k] // (*B)[k+j*n]
						sum55 += buffer55[k] * b[k] // (*B)[k+j*n]
						sum56 += buffer56[k] * b[k] // (*B)[k+j*n]
						sum57 += buffer57[k] * b[k] // (*B)[k+j*n]
						sum58 += buffer58[k] * b[k] // (*B)[k+j*n]
						sum59 += buffer59[k] * b[k] // (*B)[k+j*n]
						sum60 += buffer60[k] * b[k] // (*B)[k+j*n]
						sum61 += buffer61[k] * b[k] // (*B)[k+j*n]
						sum62 += buffer62[k] * b[k] // (*B)[k+j*n]
						sum63 += buffer63[k] * b[k] // (*B)[k+j*n]
					}
					(*C)[(i*amountBuffers+0)+j*n] = sum00
					(*C)[(i*amountBuffers+1)+j*n] = sum01
					(*C)[(i*amountBuffers+2)+j*n] = sum02
					(*C)[(i*amountBuffers+3)+j*n] = sum03
					(*C)[(i*amountBuffers+4)+j*n] = sum04
					(*C)[(i*amountBuffers+5)+j*n] = sum05
					(*C)[(i*amountBuffers+6)+j*n] = sum06
					(*C)[(i*amountBuffers+7)+j*n] = sum07
					(*C)[(i*amountBuffers+8)+j*n] = sum08
					(*C)[(i*amountBuffers+9)+j*n] = sum09
					(*C)[(i*amountBuffers+10)+j*n] = sum10
					(*C)[(i*amountBuffers+11)+j*n] = sum11
					(*C)[(i*amountBuffers+12)+j*n] = sum12
					(*C)[(i*amountBuffers+13)+j*n] = sum13
					(*C)[(i*amountBuffers+14)+j*n] = sum14
					(*C)[(i*amountBuffers+15)+j*n] = sum15
					(*C)[(i*amountBuffers+16)+j*n] = sum16
					(*C)[(i*amountBuffers+17)+j*n] = sum17
					(*C)[(i*amountBuffers+18)+j*n] = sum18
					(*C)[(i*amountBuffers+19)+j*n] = sum19
					(*C)[(i*amountBuffers+20)+j*n] = sum20
					(*C)[(i*amountBuffers+21)+j*n] = sum21
					(*C)[(i*amountBuffers+22)+j*n] = sum22
					(*C)[(i*amountBuffers+23)+j*n] = sum23
					(*C)[(i*amountBuffers+24)+j*n] = sum24
					(*C)[(i*amountBuffers+25)+j*n] = sum25
					(*C)[(i*amountBuffers+26)+j*n] = sum26
					(*C)[(i*amountBuffers+27)+j*n] = sum27
					(*C)[(i*amountBuffers+28)+j*n] = sum28
					(*C)[(i*amountBuffers+29)+j*n] = sum29
					(*C)[(i*amountBuffers+30)+j*n] = sum30
					(*C)[(i*amountBuffers+31)+j*n] = sum31
					(*C)[(i*amountBuffers+32)+j*n] = sum32
					(*C)[(i*amountBuffers+33)+j*n] = sum33
					(*C)[(i*amountBuffers+34)+j*n] = sum34
					(*C)[(i*amountBuffers+35)+j*n] = sum35
					(*C)[(i*amountBuffers+36)+j*n] = sum36
					(*C)[(i*amountBuffers+37)+j*n] = sum37
					(*C)[(i*amountBuffers+38)+j*n] = sum38
					(*C)[(i*amountBuffers+39)+j*n] = sum39
					(*C)[(i*amountBuffers+40)+j*n] = sum40
					(*C)[(i*amountBuffers+41)+j*n] = sum41
					(*C)[(i*amountBuffers+42)+j*n] = sum42
					(*C)[(i*amountBuffers+43)+j*n] = sum43
					(*C)[(i*amountBuffers+44)+j*n] = sum44
					(*C)[(i*amountBuffers+45)+j*n] = sum45
					(*C)[(i*amountBuffers+46)+j*n] = sum46
					(*C)[(i*amountBuffers+47)+j*n] = sum47
					(*C)[(i*amountBuffers+48)+j*n] = sum48
					(*C)[(i*amountBuffers+49)+j*n] = sum49
					(*C)[(i*amountBuffers+50)+j*n] = sum50
					(*C)[(i*amountBuffers+51)+j*n] = sum51
					(*C)[(i*amountBuffers+52)+j*n] = sum52
					(*C)[(i*amountBuffers+53)+j*n] = sum53
					(*C)[(i*amountBuffers+54)+j*n] = sum54
					(*C)[(i*amountBuffers+55)+j*n] = sum55
					(*C)[(i*amountBuffers+56)+j*n] = sum56
					(*C)[(i*amountBuffers+57)+j*n] = sum57
					(*C)[(i*amountBuffers+58)+j*n] = sum58
					(*C)[(i*amountBuffers+59)+j*n] = sum59
					(*C)[(i*amountBuffers+60)+j*n] = sum60
					(*C)[(i*amountBuffers+61)+j*n] = sum61
					(*C)[(i*amountBuffers+62)+j*n] = sum62
					(*C)[(i*amountBuffers+63)+j*n] = sum63
				}
			}
		}(t)
	}
	wg.Wait()
}
