# Matrix Multiplication

At the base of article is performance research of matrix multiplication.
Let`s take a few input data:

- Multiplication matrix: [A]*[B] = [C], where [A], [B], [C] - square matrix
- Size of each matrix is 2048 x 2048
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

[ 2 3 ]

[ 4 5 ]

Now, you see the square matrix with 2 rows and 2 columns.
