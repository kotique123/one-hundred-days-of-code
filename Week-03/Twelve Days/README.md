# Day 16

## Twelve Days

Today I've solved yet another easy problem. It is the last easy program that is
available on Exercism for Go lang. By the end of this poriton of easy problems, I
can tell that I feel very good about all of them. I've solved some of them before I
started this challenge so by the end of this portion I felt like just by looking at
problem I already know the solution, or at least approach to it. I enjoyed writing
my own benchmarks and comments with accordance to Godoc format, and I really
enjoy doing this and documenting my solutions and its performance.

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: twelve
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkVerse-8          |297706              |4019 ns/op           |13264 B/op         |88 allocs/op
|BenchmarkVerse-8          |292642              |4013 ns/op           |13264 B/op         |88 allocs/op
|BenchmarkVerse-8          |285434              |3998 ns/op           |13264 B/op         |88 allocs/op
|BenchmarkSong-8           |215997              |5510 ns/op           |26512 B/op        |100 allocs/op
|BenchmarkSong-8           |212336              |5470 ns/op           |26512 B/op        |100 allocs/op
|BenchmarkSong-8           |217429              |5527 ns/op           |26512 B/op        |100 allocs/op

```shell
PASS
ok      twelve  8.319s
```
