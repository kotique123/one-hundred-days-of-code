# Day-9

## Largest series product

Today I've solved Largest Series Product problem. It took a little more time
than I exepected. There is nothing very remarkable about the problem itself.
However, I've optimized the solution of this problem by 50-60% in means of
average time per operation and by number of bytes per operation and number
of memory allocation by preallocating memory for a slice and then writing
data in index. In addition I allocate memory for the slice after checks for
valid span and digits slice have been completed.

### Here are the results:

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkLargestSeriesProduct-24          |689143              |1816 ns/op            |2632 B/op         |50 allocs/op
|BenchmarkLargestSeriesProduct-24          |696034              |1776 ns/op            |2632 B/op         |50 allocs/op
|BenchmarkLargestSeriesProduct-24          |561474              |1871 ns/op            |2632 B/op         |50 allocs/op
|BenchmarkLargestSeriesProduct-24          |649266              |1803 ns/op            |2632 B/op         |50 allocs/op
|BenchmarkLargestSeriesProduct-24          |624680              |1759 ns/op            |2632 B/op         |50 allocs/op

### Final benchmark results after optimizations:

```shell
goos: windows
goarch: amd64
pkg: lsproduct
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkLargestSeriesProduct-24         |1403816               |859.3 ns/op          |1073 B/op         |16 allocs/op
|BenchmarkLargestSeriesProduct-24         |1403554               |854.9 ns/op          |1073 B/op         |16 allocs/op
|BenchmarkLargestSeriesProduct-24         |1415016               |863.4 ns/op          |1073 B/op         |16 allocs/op
|BenchmarkLargestSeriesProduct-24         |1401687               |863.7 ns/op          |1073 B/op         |16 allocs/op
|BenchmarkLargestSeriesProduct-24         |1415426               |858.1 ns/op          |1073 B/op         |16 allocs/op

```shell
PASS
ok      lsproduct       10.379s
```
