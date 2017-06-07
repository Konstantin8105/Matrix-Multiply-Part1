# Matrix Multiplication

At the base of article is performance research of matrix multiplication.
Let's take a few input data:

- Multiplication matrix: [A] * [B] = [C], where [A], [B], [C] - square matrix
- Size of each matrix is 1024 x 1024
- Type of values: **float64**
- Matrix is dense. The algorithms for sparse matrix is outside of that article.
- Programming language: [Go](https://golang.org/)
- Laptop precossor for experiments: [Intel(R) Core(TM) i5-3230M CPU @2.6 GHz](https://ark.intel.com/ru/products/72164/Intel-Core-i5-3230M-Processor-3M-Cache-up-to-3_20-GHz-rPGA)
> The number of CPU cores : 2
>
> The number of threads: 4
>
> CPU cache : 3MB

You can see the theory in [Wikipedia](https://en.wikipedia.org/wiki/Matrix_multiplication).
Shortly:
Matrix is table of values, for example:

```
[ 2 3 ]
[ 4 5 ]
```

Now, you see the square matrix with 2 rows and 2 columns. Like you understood, values can be any.
So, a little example for matrix multiplication looks like that:

```
[ 2 3 ]  *  [ 1 6 ]  =  [ 2*1+3*9 2*6+3*8 ] =  [ 29 36 ]
[ 4 5 ]     [ 9 8 ]     [ 4*1+5*9 4*6+5*8 ]    [ 49 64 ]
```

Yes, we will optimize the classic algorithm of multiplication with O(n^3).

We can start to analyze present design of algorithm in exist libraries.

Code in project [JAMA : A Java Matrix Package](http://math.nist.gov/javanumerics/jama/)
```java
public class Matrix implements ... {
   ...
   public Matrix times (Matrix B) {
	  ...
      Matrix X = new Matrix(m,B.n);
      double[][] C = X.getArray();
      double[] Bcolj = new double[n];       // Buffer of column matrix B
      for (int j = 0; j < B.n; j++) {
         for (int k = 0; k < n; k++) {
            Bcolj[k] = B.A[k][j];           // Put in buffer values
         }
         for (int i = 0; i < m; i++) {
            double[] Arowi = A[i];
            double s = 0;
            for (int k = 0; k < n; k++) {
               s += Arowi[k]*Bcolj[k];     // Multiplication
            }
            C[i][j] = s;
         }
      }
      return X;
   }
   ...
}
```
Comment - at the future, I will show the reason - Why buffer of column for matrix B is good? But we create the better.

Code in project [go.matrix](https://github.com/skelterjohn/go.matrix/blob/daa59528eefd43623a4c8e36373a86f9eef870a2/arithmetic.go)
```go
func ParallelProduct(A, B MatrixRO) (C *DenseMatrix) {
	...
	C = Zeros(A.Rows(), B.Cols())
	in := make(chan int)
	quit := make(chan bool)

	dotRowCol := func() {
		for {
			select {
			case i := <-in:
				sums := make([]float64, B.Cols())      // Array for intermediante results
				for k := 0; k < A.Cols(); k++ {
					for j := 0; j < B.Cols(); j++ {
						sums[j] += A.Get(i, k) * B.Get(k, j)
					}
				}
				for j := 0; j < B.Cols(); j++ {
					C.Set(i, j, sums[j])
				}
			case <-quit:
				return
			}
		}
	}

	threads := 2                                      // Is it looks strange, because I have more then 2 processors
	for i := 0; i < threads; i++ {
		go dotRowCol()
	}
	for i := 0; i < A.Rows(); i++ {
		in <- i
	}
	for i := 0; i < threads; i++ {
		quit <- true
	}
	return
}
```
Comments:
- Strange, but it is true, amount of threads put in code. The best way to use actual processors on user computer
- We don't see buffer
- We see using array for intermadiante results. Now, it is not clear - it is good or not. We will see.

# Stop theory, more practic, more benchmarks

Now, we are ready for experiments.
At the first time, we look on first benchmark in detail for understood each line of code.
```go
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
```

For start benchmark test we put in command line next text:
```command line
go test -v -bench=. -benchtime=1m -benchmem
```
Flags:
>
> -v             - show all output
>
> -bench=.       - start all benchmark
>
> -benchtime=1m  - minimal time of benchmark, 1m = 1 minute
>
> -benchmem      - show amount of memory allocations

Our first benchmark result:
```command line
BenchmarkSimple-4   	       5	15305107558 ns/op	       0 B/op	       0 allocs/op
```
So, we see next: our test executed 5 times and ~15.3 sec for each multiplication and we don't allocation addition memory.

For future algorithm optimization, we have to refactoring the code for avoid mistake and minimaze the time for benchmark research. **But** the cost of that refactoring is each bechmark will be little bit slow, in our case it is now so important.

Firstly, we create a simple(slow) check function for compare results all new algorithms.
```go
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
 			// For our case it is acceptable,
			// but by default we cannot compare
			// two float values directly
			if sum != C[i][j] {
				return false
			}
		}
	}
	return true
}
```
Put our "simple, first" algorithm inside function outside of test. Like that:
```go
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
```
We see the simple, clear function with 3 input, output matrix.
Prefix of function name `mm` mean - "Matrix Multiplication".

Our test look very beatiful:
```go
func TestSimple(t *testing.T) {
	if !isSame(mmSimple) {
		t.Errorf("Algorithm is not correct")
	}
}
```
Our benchmark look is same clear:
```go
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
```

In next steps, I will hide test and benchmark code, but show only algorithm code.
It will be simple for concentrate on algorithm code.

What we see in function `mmSimple` for optimization?
> We can create a buffer of row of matrix A.
>
> The reason: we prepare memory and if all is Ok, we will
> put memory in CPU cache.
> One more attention: if all out data for calculation inside
> CPU cache, then that calculation will be calculated
> fast, because speed of memory is more fast then RAM.
>
> At the next time, we will see the way for preliminary
> garantee putting memory in CPU cache.

So, our code is:
```go
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
```

We have to change the command line, because now we check 2 algorithm at the one time.
```command line
go test -v -bench=. -benchtime=3m -benchmem
```

Result is:
```command line
=== RUN   TestSimple
--- PASS: TestSimple (36.08s)
=== RUN   TestBuffer1
--- PASS: TestBuffer1 (35.56s)
BenchmarkSimple-4    	      10	22155569004 ns/op	       8 B/op	       0 allocs/op
BenchmarkBuffer1-4   	      10	18786789917 ns/op	    8377 B/op	       1 allocs/op
PASS
ok  	github.com/Konstantin8105/MatrixMultiply	522.489s
```
What we see?
1.	Both algorithm is good(haven't bug)
2.	Algorithm with buffer little bit faster at 22.1/18.7 = 1.18 times (18%)

Like we see in example of JAMA library, the putting buffer is good way.
But it is not a clear:
* Why?
* Why only one buffer?
* Can we create the algorithm more faster?
* Can we create a parallel algorithm?

Let's start to create new experiments.

# More buffers, more faster??? And yes, and not.

Create a multiplication matrix algorithm with 2 buffers.
We can that algorithm, only if size of matrix is multiple by 2.
```go
// mmBuffer2 - added two buffer
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
```
Create a multiplication matrix algorithm with 4 buffers.
We can that algorithm, only if size of matrix is multiple by 4.
```go
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
```

Creating the multiplication matrix algorithm with 8, 16, 32, 64 buffers are not show, but you can see in code file.

Results:
```command line

=== RUN   TestBuffer1
--- PASS: TestBuffer1 (22.16s)
=== RUN   TestBuffer2
--- PASS: TestBuffer2 (20.69s)
=== RUN   TestBuffer4
--- PASS: TestBuffer4 (16.06s)
=== RUN   TestBuffer8
--- PASS: TestBuffer8 (14.71s)
=== RUN   TestBuffer16
--- PASS: TestBuffer16 (13.98s)
=== RUN   TestBuffer32
--- PASS: TestBuffer32 (10.66s)
=== RUN   TestBuffer64
--- PASS: TestBuffer64 (11.08s)
BenchmarkBuffer1-4                  	       1	25774664502 ns/op	    8192 B/op	       1 allocs/op
BenchmarkBuffer2-4                  	       1	19334469447 ns/op	   16384 B/op	       2 allocs/op
BenchmarkBuffer4-4                  	       1	14938243077 ns/op	   32768 B/op	       4 allocs/op
BenchmarkBuffer8-4                  	       1	12447383207 ns/op	   65536 B/op	       8 allocs/op
BenchmarkBuffer16-4                 	       1	11231806529 ns/op	  131072 B/op	      16 allocs/op
BenchmarkBuffer32-4                 	       1	9488617339 ns/op	  262144 B/op	      32 allocs/op
BenchmarkBuffer64-4                 	       1	9177935520 ns/op	  524288 B/op	      64 allocs/op
PASS
```
Like we see, we will have the optimal solution between 4 and 64 buffers. May be [42](https://en.wikipedia.org/wiki/The_Hitchhiker%27s_Guide_to_the_Galaxy) - we have to continue the analyzing.

And now our algorithm is faster at 22.2/9.17 = 2.42 times

We still use only one core of processor. So, let's create a parallel algorithm.

# Speed := moreBuffers + moreCores

In that article, we use programming Go with incredible simple to create parallel/concurrency.
Just, `go`.

Just for you, look on parallel version of algorithm with addition comments for clear undestood.

```go
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
```

If you feel the pain of waiting the results of benchmarks before, then now it is lost, because now it much better performance. Let's look on results of parallel algorithm with buffers.

```command line
=== RUN   TestParallelBuffer2
--- PASS: TestParallelBuffer2 (14.96s)
=== RUN   TestParallelBuffer4
--- PASS: TestParallelBuffer4 (13.56s)
=== RUN   TestParallelBuffer8
--- PASS: TestParallelBuffer8 (11.51s)
=== RUN   TestParallelBuffer16
--- PASS: TestParallelBuffer16 (9.28s)
=== RUN   TestParallelBuffer32
--- PASS: TestParallelBuffer32 (8.76s)
=== RUN   TestParallelBuffer64
--- PASS: TestParallelBuffer64 (9.30s)
BenchmarkParallelBuffer2-4    	       1	8776396393 ns/op	   67568 B/op	      16 allocs/op
BenchmarkParallelBuffer4-4    	       1	7070533991 ns/op	  132352 B/op	      21 allocs/op
BenchmarkParallelBuffer8-4    	       1	5332511460 ns/op	  263504 B/op	      38 allocs/op
BenchmarkParallelBuffer16-4   	       1	6616536928 ns/op	  525648 B/op	      70 allocs/op
BenchmarkParallelBuffer32-4   	       1	4270706396 ns/op	 1049584 B/op	     134 allocs/op
BenchmarkParallelBuffer64-4   	       1	4881902735 ns/op	 2098512 B/op	     262 allocs/op
PASS
ok  	command-line-arguments	104.808s
```
Like we see, we will have the optimal solution between 4 and 64 buffers. May be [42](https://en.wikipedia.org/wiki/The_Hitchhiker%27s_Guide_to_the_Galaxy) - we have to continue the analyzing.

Now our algorithm is faster at 22.2/4.27 = 5.22 times

# Try minimaze inialiazations of variables

At this part of article, we try to optimize algorithm.
We add new variable for summ in step for calculate the matrix [C] and first initialize outside of loop.
Look on code on parallel algoritm with 2 buffers and some initialization outside of loop.

```go
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
```
Changed 7 lines of code and added comment : `<---- Mark`

Look on results:
```command line
=== RUN   TestParallelBufferVarOut2
--- PASS: TestParallelBufferVarOut2 (6.84s)
=== RUN   TestParallelBufferVarOut4
--- PASS: TestParallelBufferVarOut4 (6.22s)
=== RUN   TestParallelBufferVarOut8
--- PASS: TestParallelBufferVarOut8 (6.69s)
=== RUN   TestParallelBufferVarOut16
--- PASS: TestParallelBufferVarOut16 (6.84s)
=== RUN   TestParallelBufferVarOut32
--- PASS: TestParallelBufferVarOut32 (5.08s)
=== RUN   TestParallelBufferVarOut64
--- PASS: TestParallelBufferVarOut64 (5.01s)
BenchmarkParallelBufferVarOut2-4    	       1	4583026466 ns/op	   66048 B/op	      12 allocs/op
BenchmarkParallelBufferVarOut4-4    	       1	3573125013 ns/op	  131104 B/op	      18 allocs/op
BenchmarkParallelBufferVarOut8-4    	       1	2216523200 ns/op	  263008 B/op	      36 allocs/op
BenchmarkParallelBufferVarOut16-4   	       1	2383133349 ns/op	  525648 B/op	      70 allocs/op
BenchmarkParallelBufferVarOut32-4   	       1	1793297208 ns/op	 1049936 B/op	     134 allocs/op
BenchmarkParallelBufferVarOut64-4   	       1	1837229377 ns/op	 2098512 B/op	     262 allocs/op
PASS
```
The optimal of amount buffers is between 16 ... 64 in our task.
Now our algorithm is faster at 22.2/1.79 = 12.4 times

# Result graph

![Graph 1](https://github.com/Konstantin8105/MatrixMultiply/blob/master/images/1.png)

In graphs, less is better.

![Graph 2](https://github.com/Konstantin8105/MatrixMultiply/blob/master/images/2.png)

In graphs, less is better.

# Create preliminary optimization formula

Let's amount all results and create some preliminary model of calculation.

Simple image of architecture:
```ascii picture
+---------------------------+  +-------------------+
| CPU                       |  | RAM               |
|  +--------+  +---------+  |==| memory            |
|  |  Core  |==| CPU     |  |  |                   |
|  +--------+  | memory  |  |==|                   |
|  +--------+  |         |  |  |                   |
|  |  Core  |==|         |  |==|                   |
|  +--------+  |         |  |  |                   |
|  +--------+  |         |  |==|                   |
|  |  Core  |==|         |  |  |                   |
|  +--------+  |         |  |==|                   |
|  +--------+  |         |  |  |                   |
|  |  Core  |==|         |  |==|                   |
|  +--------+  +---------+  |  |                   |
+---------------------------+  +-------------------+
```

Calculation amount of memory inside each goroutine:

```go
0.		go func(init int) {
1.			// Change waitgroup after work done
2.			defer wg.Done()
3.			// Inialize addition variables
4.			var sum0, sum1 float64
5.			// Create buffers
6.			amountBuffers := 2
7.			buffer0 := make([]float64, n, n)
8.			buffer1 := make([]float64, n, n)
9.			// Calculate amount of calculation part
10.			// for that goroutine
11.			amountParts := n / amountBuffers
12.			for i := init; i < amountParts; i += threads {
13.				for j := 0; j < n; j++ {
14.					// Put in buffer row of matrix [A]
15.					buffer0[j] = (*A)[i*amountBuffers+0][j]
16.					buffer1[j] = (*A)[i*amountBuffers+1][j]
17.				}
18.				for j := 0; j < n; j++ {
19.					sum0 = 0.0
20.					sum1 = 0.0
21.					for k := 0; k < n; k++ {
22.						sum0 += buffer0[k] * (*B)[k][j]
23.						sum1 += buffer1[k] * (*B)[k][j]
24.					}
25.					(*C)[i*amountBuffers+0][j] = sum0
26.					(*C)[i*amountBuffers+1][j] = sum1
27.				}
28.			}
29.		}(t)
```

Variables for calculation:

```variables
n  - amount elements in each column or rows in matrix
SF - size of float64
AB - amount of buffers
TH - amount of threads
```

We are sure: we don't know the location of data - inside CPU memory or inside RAM.

We have:

```calculation
TH * AB * SF     - memory for variables "summ"
TH * AB * SF * n - memory for buffers (vectors of matrix [A])
TH * AB * SF * n - memory for vectors of matrix [B]
TH * AB * SF     - memory for values of matrix [C]
------------------- 
Summary :

MEMORY := TH * AB * SF * ( 2 + 2 * n)
```

In our case :
```calculation
TH     = 4 processors
SF     = 64 bit = 8 bytes
n      = 2048 elements
MEMORY = 3 MB ~= 3 000 000 bytes

Optimal size of buffers:
              MEMORY                  3 000 000
AB = -------------------------- = ----------------------- = 22.8 ~= 23 buffers
       TH * SF * ( 2 + 2 * n )      4 * 8 * (2 + 2*2048)
```

We can check our formula on another [processor](http://ark.intel.com/products/65730/Intel-Xeon-Processor-E3-1240-v2-8M-Cache-3_40-GHz) with cpu cache 8MB and 8 processors.
```results
BenchmarkParallelBufferVarOut2-8               1        1510086400 ns/op      131104 B/op         18 allocs/op
BenchmarkParallelBufferVarOut4-8               2         858049050 ns/op      263256 B/op         37 allocs/op
BenchmarkParallelBufferVarOut8-8               2         566032400 ns/op      524320 B/op         66 allocs/op
BenchmarkParallelBufferVarOut16-8              2         606034700 ns/op     1049064 B/op        131 allocs/op
BenchmarkParallelBufferVarOut32-8              2         727041550 ns/op     2097288 B/op        259 allocs/op
BenchmarkParallelBufferVarOut64-8              2         730041750 ns/op     4195832 B/op        518 allocs/op

Optimal size of buffers:
              MEMORY                   8 000 000
AB = -------------------------- = ---------------------- = 30.5 ~= 31 buffers
       TH * SF * ( 2 + 2 * n )     8 * 8 * (2 + 2*2048)
```

Like we see on that 2 examples, formula is preliminary, because we still have 2 problems:
1. Processors can calculate faster then memory prepare the data.
2. Memory can prepare data faster then processor can calculate.

So, we have to create a new algorithm without that problems. But at the first, little bit about our results.

# Results

For founding the optimal solution, we use used next:
* More benchmarks, show we way for optimal solution
* Parallel algorithm is better, then one core algorithm
* Feel free to use buffers, but not too much
* Try initialize the variables outside the loop
* Try to use physical limits of hardware(cpu, ram)

The result, we create the algorithm at 12.4 times fast then naive algorithm.

------
#TODO

add note about allocation - don't affaid)
add more tests
add test for one single matrix
add tests for deep matrix
add more visual graph

don't create a buffers - transponse A
add code link in text
add https://asciinema.org/ for benchmark

worker pool for calculate, prepare the buffers
