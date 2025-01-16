# Day 15

## Bottle Song

Nothing particular about it. Just string concatanation. Pretty straight forward.
However, there is some room for improvement in means of precomputing
switch case statements and placing them in a slice.

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: bottlesong
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
BenchmarkRecite-8         625731              1889 ns/op            4656 B/op         64 allocs/op
BenchmarkRecite-8         584026              1913 ns/op            4656 B/op         64 allocs/op
BenchmarkRecite-8         616335              1916 ns/op            4656 B/op         64 allocs/op
BenchmarkRecite-8         629553              1906 ns/op            4656 B/op         64 allocs/op
BenchmarkRecite-8         605251              1917 ns/op            4656 B/op         64 allocs/op

```shell
PASS
ok      bottlesong      6.886s
```
