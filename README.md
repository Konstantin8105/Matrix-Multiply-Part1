# Matrix Multiplication

At the base of article is performance research of matrix multiplication.
Let's take a few input data:

- Multiplication matrix: [A]*[B] = [C], where [A], [B], [C] - square matrix
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

	threads := 2                                      // Is it looks strange, because I can more then 2 processors
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
=== RUN   TestSimple
--- PASS: TestSimple (23.48s)
=== RUN   TestBuffer1
--- PASS: TestBuffer1 (24.23s)
=== RUN   TestBuffer2
--- PASS: TestBuffer2 (18.21s)
=== RUN   TestBuffer4
--- PASS: TestBuffer4 (17.93s)
=== RUN   TestBuffer8
--- PASS: TestBuffer8 (16.60s)
=== RUN   TestBuffer16
--- PASS: TestBuffer16 (16.27s)
=== RUN   TestBuffer32
--- PASS: TestBuffer32 (15.96s)
=== RUN   TestBuffer64
--- PASS: TestBuffer64 (17.10s)
BenchmarkSimple-8                      1        12527000000 ns/op          0 B/op          0 allocs/op
BenchmarkBuffer1-8                     1        11497000000 ns/op       8192 B/op          1 allocs/op
BenchmarkBuffer2-8                     1        8480000000 ns/op       16384 B/op          2 allocs/op
BenchmarkBuffer4-8                     1        6039000000 ns/op       32768 B/op          4 allocs/op
BenchmarkBuffer8-8                     1        5945000000 ns/op       65536 B/op          8 allocs/op
BenchmarkBuffer16-8                    1        5959000000 ns/op      131072 B/op         16 allocs/op
BenchmarkBuffer32-8                    1        6484512700 ns/op      262144 B/op         32 allocs/op
BenchmarkBuffer64-8                    1        6637663700 ns/op      524288 B/op         64 allocs/op
PASS
ok      github.com/Konstantin8105/MatrixMultiply        225.049s
```
Like we see, we will have the optimal solution between 4 and 64 buffers. May be [42](https://en.wikipedia.org/wiki/The_Hitchhiker%27s_Guide_to_the_Galaxy).

And now our algorithm is faster at 12.5/6.03 = 2.07 times

We still use only one core of processor. So, let's create a parallel algorithm.

# More buffers + more cores = more faster.





------
#TODO

add more tests
add test for one single matrix
add tests for deep matrix

