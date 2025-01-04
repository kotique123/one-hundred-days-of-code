Today I wrote a soltuion for complex_numbers.go problem from Exercism. There
weren't any benchmarks provided, so I wrote my own, hopefully it makes sense.
```
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