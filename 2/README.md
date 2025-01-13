Today I wrote a small solutiont to another problem on exercism. I feel like
there is more sophisticated solution than just checking all the cases, but
it is what it is. I like my code being easy to comprehend and 2.7ns/op on
average is pretty good result for this simple solution.

Benchmark results:
```
goos: windows
goarch: amd64
pkg: triangle
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:---|:---|:---|:---|:---|
BenchmarkKind-24|        432319639     |           2.768 ns/op     |      0 B/op   |       0 allocs/op
BenchmarkKind-24 |       442606969      |          2.873 ns/op      |     0 B/op    |      0 allocs/op
BenchmarkKind-24  |      441036051       |         2.733 ns/op       |    0 B/op     |     0 allocs/op
BenchmarkKind-24   |     443882029        |        2.758 ns/op        |   0 B/op      |    0 allocs/op
BenchmarkKind-24    |    442864238         |       2.692 ns/op         |  0 B/op       |   0 allocs/op
```
PASS
ok      triangle        8.565s
```