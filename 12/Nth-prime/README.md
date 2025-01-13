# Day 12 

## Nth Prime

### Benchmark results 

```shell
goos: darwin
goarch: arm64
pkg: prime
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkNth-8               |482           |2334180 ns/op               |0 B/op          |0 allocs/op
|BenchmarkNth-8               |512           |2336114 ns/op               |0 B/op          |0 allocs/op
|BenchmarkNth-8               |513           |2328801 ns/op               |0 B/op          |0 allocs/op
|BenchmarkNth-8               |512           |2330727 ns/op               |0 B/op          |0 allocs/op
|BenchmarkNth-8               |512           |2334802 ns/op               |0 B/op          |0 allocs/op

```shell
PASS
ok      prime   7.337s
```
