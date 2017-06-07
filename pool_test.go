package main_test

import (
	"runtime"
	"sync"
	"testing"
)

/*
func TestPool(t *testing.T) {
	if !isSame(mmPool) {
		t.Errorf("Algorithm is not correct")
	}
}
*/
func BenchmarkPool(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmPool(&A, &B, &C)
		// Finish of algorithm
	}
}

type task struct {
	bufferA00 []float64
	bufferA01 []float64
	bufferA02 []float64
	bufferA03 []float64
	rowA      int
	bufferB   []float64
	columnB   int
}

type pool struct {
	ch   chan *task
	size int
}

func newPool(size int) *pool {
	// Found amount allowable parallelism
	threads := runtime.GOMAXPROCS(0)
	if threads > runtime.NumCPU() {
		threads = runtime.NumCPU()
	}
	return &pool{
		ch:   make(chan *task, threads),
		size: size,
	}
}

func (p *pool) get() (t *task) {
	select {
	case t = <-p.ch:
	//	fmt.Printf("+")
	default:
		t = new(task)
		//t.bufferA = make([]float64, n)
		(*t).bufferB = make([]float64, p.size)
	}
	return
}

func (p *pool) put(t *task) {
	select {
	case p.ch <- t:
	default:
		//	fmt.Printf("-")
	}
}

func (p *pool) close() {
	close(p.ch)
}

// mmPool - added one buffer
func mmPool(A, B, C *[][]float64) {
	n := len(*A)
	// Found amount allowable parallelism
	threads := runtime.GOMAXPROCS(0)
	if threads > runtime.NumCPU() {
		threads = runtime.NumCPU()
	}
	// Create workgroup
	var wg sync.WaitGroup

	pl := newPool(n)
	chanTask := make(chan *task)

	type message struct {
		i, j int
		v    float64
	}
	messageChan := make(chan message, n*threads)

	go func() {
		for m := range messageChan {
			(*C)[m.i][m.j] = m.v
		}
	}()

	for t := 0; t < threads; t++ {
		// Add one goroutine in workgroup
		wg.Add(1)
		go func() {
			// Change waitgroup after work done
			defer wg.Done()
			var sum00 float64
			var sum01 float64
			var sum02 float64
			var sum03 float64
			for c := range chanTask {
				sum00 = 0.0
				sum01 = 0.0
				sum02 = 0.0
				sum03 = 0.0
				for i := 0; i < n; i++ {
					sum00 += c.bufferA00[i] * c.bufferB[i]
					sum01 += c.bufferA01[i] * c.bufferB[i]
					sum02 += c.bufferA02[i] * c.bufferB[i]
					sum03 += c.bufferA03[i] * c.bufferB[i]
				}
				messageChan <- message{c.rowA + 0, c.columnB, sum00}
				messageChan <- message{c.rowA + 1, c.columnB, sum01}
				messageChan <- message{c.rowA + 2, c.columnB, sum02}
				messageChan <- message{c.rowA + 3, c.columnB, sum03}
				(*pl).put(c)
			}
		}()
	}

	// Create workgroup
	var wg2 sync.WaitGroup
	for t := 0; t < threads; t++ {
		wg2.Add(1)
		go func(t int) {
			// Change waitgroup after work done
			defer wg2.Done()
			// loop
			for i := t * 4; i < n; i += threads * 4 {
				for j := 0; j < n; j++ {
					t := (*pl).get()
					t.rowA = i
					t.columnB = j
					t.bufferA00 = (*A)[i+0]
					t.bufferA01 = (*A)[i+1]
					t.bufferA02 = (*A)[i+2]
					t.bufferA03 = (*A)[i+3]
					for p := 0; p < n; p++ {
						t.bufferB[p] = (*B)[p][j]
					}
					chanTask <- t
				}
			}
		}(t)
	}
	wg2.Wait()
	close(chanTask)
	wg.Wait()
	(*pl).close()
	close(messageChan)
}
