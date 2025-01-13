# Day 11 

## Sieve 

Today I have completed Sieve problem from Exercism. There is not anything specific about it since I just googled
how the algorithm works. But, what makes my solution different, is the fact that I tweaked some array and memory
management related things, which improved performance of my solution by 13% time per operation wise, and by aroud
35% in means of Number of bytes allocated per operation and number of allocations per operation, which I consider
a very good result.

### Benchmark results

```shell
goos: windows
goarch: amd64
pkg: sieve
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkSieve-24         |707880              |1687 ns/op            |3968 B/op         |16 allocs/op
|BenchmarkSieve-24         |713112              |1677 ns/op            |3968 B/op         |16 allocs/op
|BenchmarkSieve-24         |723540              |1669 ns/op            |3968 B/op         |16 allocs/op
|BenchmarkSieve-24         |721352              |1692 ns/op            |3968 B/op         |16 allocs/op
|BenchmarkSieve-24         |712702              |1696 ns/op            |3968 B/op         |16 allocs/op
```shell
PASS
ok      sieve   7.185s
```