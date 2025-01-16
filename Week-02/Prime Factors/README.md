# Day 13

## Prime Factors

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: phonenumber
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkPrimeFactors-8           |335511              |3480 ns/op             |784 B/op         |11 allocs/op
|BenchmarkPrimeFactors-8           |333807              |3471 ns/op             |784 B/op         |11 allocs/op
|BenchmarkPrimeFactors-8           |330691              |3474 ns/op             |784 B/op         |11 allocs/op
|BenchmarkPrimeFactors-8           |338770              |3598 ns/op             |784 B/op         |11 allocs/op
|BenchmarkPrimeFactors-8           |336673              |3466 ns/op             |784 B/op         |:w11 allocs/op

```shell
PASS
ok      prime   6.956s
```
