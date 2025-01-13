# Day 12

## Acronym

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: acronym
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkAcronym-8        |594974              |1965 ns/op            |1272 B/op         |39 allocs/op
|BenchmarkAcronym-8        |585213              |1971 ns/op            |1272 B/op         |39 allocs/op
|BenchmarkAcronym-8        |562820              |1971 ns/op            |1272 B/op         |39 allocs/op
|BenchmarkAcronym-8        |574878              |1971 ns/op            |1272 B/op         |39 allocs/op
|BenchmarkAcronym-8        |578889              |1992 ns/op            |1272 B/op         |39 allocs/op

```shell
PASS
ok      acronym 6.600s
```
