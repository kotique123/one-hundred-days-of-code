# Day 13

## Phone Number

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: phonenumber
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkNumber-8         |915006              |1274 ns/op             |288 B/op         |18 allocs/op
|BenchmarkAreaCode-8       |909860              |1280 ns/op             |288 B/op         |18 allocs/op
|BenchmarkFormat-8         |615063              |1943 ns/op             |608 B/op         |38 allocs/op

```shell
PASS
ok      phonenumber     4.302s
```
