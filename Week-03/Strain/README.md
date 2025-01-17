# Day 16

## Strain

Today I've solved my first medium problem. I wasn't quite sure on how to appraoch it
since it required a concept previously unknown for me, so even after reading
documentation I felt lost and confused, so I consulted others solutions and it helped
a lot since it gave me an insight on how the syntax may look like and how to approach
the problem. I left a credit for his work in comments in strain.go

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: strain
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkKeep-8          |9735231               |116.8 ns/op            |96 B/op          |4 allocs/op
|BenchmarkKeep-8          |9927956               |120.6 ns/op            |96 B/op          |4 allocs/op
|BenchmarkKeep-8         |10077148               |119.9 ns/op            |96 B/op          |4 allocs/op
|BenchmarkDiscard-8       |8994262               |124.7 ns/op            |96 B/op          |4 allocs/op
|BenchmarkDiscard-8       |9738944               |125.0 ns/op            |96 B/op          |4 allocs/op
|BenchmarkDiscard-8       |9760266               |127.6 ns/op            |96 B/op          |4 allocs/op

```shell
PASS
ok      strain  10.042s
```
