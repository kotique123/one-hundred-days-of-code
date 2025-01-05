Today I wrote a soltuion for complex_numbers.go problem from Exercism. There
weren't any benchmarks provided, so I wrote my own, hopefully it makes sense.
```
file: complex_numbers.go
goos: windows
goarch: amd64
pkg: complex
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:---|:---|:---|:---|:---|
BenchmarkAllOperationsTogether-24|       88633492|                12.97 ns/op|            0 B/op |         0 allocs/op|
BenchmarkAllOperationsTogether-24 |      92778722 |               13.08 ns/op |           0 B/op  |        0 allocs/op|
BenchmarkAllOperationsTogether-24  |     78601933 |              12.90 ns/op   |         0 B/op    |      0 allocs/op|
BenchmarkAllOperationsTogether-24   |    92837580  |              12.93 ns/op   |         0 B/op    |      0 allocs/op|
BenchmarkAllOperationsTogether-24    |   91725586   |             12.93 ns/op    |        0 B/op     |     0 allocs/op|
```
PASS
ok      complex 5.984s
```

So, I wrote complex_numbers_optimized.go where I rewrote some of the functions
aiming to eliminate redundant variable creation in order to improve
performance. Benchmark results showed 1% improvement which is rather an
observational error than an actual perofrmance improvement since an 
[online decompiler](https://godbolt.org/) did not show any difference between
two implementations of the funcitons. However, I will leave both versions here.
```
file: complex_numbers_optimized.go
goos: windows
goarch: amd64
pkg: complex
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:---|:---|:---|:---|:---|
BenchmarkAllOperationsTogether-24|       89801538|                12.98 ns/op|            0 B/op|          0 allocs/op|
BenchmarkAllOperationsTogether-24 |      93929786 |               12.85 ns/op |           0 B/op |         0 allocs/op|
BenchmarkAllOperationsTogether-24  |     85597506  |              12.85 ns/op  |          0 B/op  |        0 allocs/op|
BenchmarkAllOperationsTogether-24   |    90399565   |             12.82 ns/op   |         0 B/op   |       0 allocs/op|
BenchmarkAllOperationsTogether-24    |   91598857    |            12.87 ns/op    |        0 B/op    |      0 allocs/op|
BenchmarkAllOperationsTogether-24     |  95435782     |           12.86 ns/op     |       0 B/op     |     0 allocs/op|
BenchmarkAllOperationsTogether-24      | 93071594      |          12.90 ns/op      |      0 B/op      |    0 allocs/op|
BenchmarkAllOperationsTogether-24       |93750000       |         12.84 ns/op       |     0 B/op       |   0 allocs/op|
BenchmarkAllOperationsTogether-24       |89267780        |        12.84 ns/op        |    0 B/op        |  0 allocs/op|
BenchmarkAllOperationsTogether-24       |91511541         |       12.89 ns/op         |   0 B/op         | 0 allocs/op|