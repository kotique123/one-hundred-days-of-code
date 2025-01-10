# Day 10

## Simple Cipher

Today, I've solved another problem. It wasn't that complex despite taking almost 200 lines of code, which are mostly boilerplate.
I feel like I can rewrite this whole thing to make it more visually appealing and easy to understand, but it is not really worth it
since I don't see room for optimizing the code overall to make it faster. Probably removing regual expressions may decrease
number of allocations or time per operation. This one did not take that much time, since I have mastered the concepts that were
required to solve it even thou it looks massive.

### Benchmark results

```shell
goos: windows
goarch: amd64
pkg: cipher
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkEncodeCaesar-24         |2227010               |536.8 ns/op           |118 B/op          |7 allocs/op
|BenchmarkDecodeCaesar-24         |2936643               |408.2 ns/op           |118 B/op          |7 allocs/op
|BenchmarkNewShift-24            |73130598                |16.38 ns/op            |0 B/op          |0 allocs/op
|BenchmarkEncodeShift-24           |932270              |1212 ns/op             |638 B/op         |13 allocs/op
|BenchmarkDecodeShift-24          |1235685               |968.1 ns/op           |590 B/op         |13 allocs/op
|BenchmarkNewVigenere-24           |181284              |6557 ns/op           |11891 B/op        |170 allocs/op
|BenchmarkEncVigenere-24           |554848              |2162 ns/op            |2188 B/op         |11 allocs/op
|BenchmarkDecVigenere-24           |574495              |2063 ns/op            |2188 B/op         |11 allocs/op
```shell
PASS
ok      cipher  11.687s
```
