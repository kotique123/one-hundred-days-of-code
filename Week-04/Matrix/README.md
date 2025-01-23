# Matrix

Today I have finally developed a solution to this trivial problem. 2D arrays aren't
something that is crazily complicated, but it took me quite some time to figure it out
since the way how 2D arrays are implemented in Go lang are different from Java, my first
programming language to which I had a somewhat broad exposure, so it affected the way how I
precept it. Other than that, this solution is pretty straight forward, and I can't really
add anyting about it.

P.S I appologize for not publishing anything for a long time. I had to work on academic
projects of mine so I could not really dedicate any reasonable amount of time to meaningfuly
solve the problem.

## Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: matrix
cpu: Apple M1
```

### Generation 1

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkNew-8           |1782790               |640.5 ns/op          |1488 B/op         |15 allocs/op
|BenchmarkNew-8           |1872801               |636.7 ns/op          |1488 B/op         |15 allocs/op
|BenchmarkRows-8         |10938470               |108.5 ns/op           |192 B/op          |5 allocs/op
|BenchmarkRows-8         |11111260               |111.1 ns/op           |192 B/op          |5 allocs/op
|BenchmarkCols-8          |7938183               |142.4 ns/op           |288 B/op          |6 allocs/op
|BenchmarkCols-8          |8452570               |137.4 ns/op           |288 B/op          |6 allocs/op

```shell
PASS
ok      matrix  22.605s
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkNew-8           |2174440               |508.6 ns/op           |744 B/op         |12 allocs/op
|BenchmarkNew-8           |2413344               |503.4 ns/op           |744 B/op         |12 allocs/op
|BenchmarkRows-8          |9678108               |115.1 ns/op           |192 B/op          |5 allocs/op
|BenchmarkRows-8         |10022906               |119.2 ns/op           |192 B/op          |5 allocs/op
|BenchmarkCols-8          |8865831               |134.9 ns/op           |288 B/op          |6 allocs/op
|BenchmarkCols-8          |8884830               |139.1 ns/op           |288 B/op          |6 allocs/op

```shell
PASS
ok      matrix  8.864s
```
