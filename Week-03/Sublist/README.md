# Day 15

## Sublist

Today I've solved a subliset problem. There were two iterations in my solution that
differened in means of how neatly the program was written. The more neat version,
which is the current version, performes slower by 1ns/op than the more sphagetti
one. However, in this particular case, I will stick to the solwer version in pursue
of clean code that is easy to read since the perofrmance difference is negligeble.

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: sublist
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkSublist-8       |5954212               |171.6 ns/op             |0 B/op          |0 allocs/op
|BenchmarkSublist-8       |6976312               |171.3 ns/op             |0 B/op          |0 allocs/op
|BenchmarkSublist-8       |6983932               |171.2 ns/op             |0 B/op          |0 allocs/op
|BenchmarkSublist-8       |6984753               |171.4 ns/op             |0 B/op          |0 allocs/op
|BenchmarkSublist-8       |6973992               |171.2 ns/op             |0 B/op          |0 allocs/op

```shell
PASS
ok      sublist 6.930s
```
