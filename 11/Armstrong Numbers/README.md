# Day 10 #

## Armstrong Numbers ##

I've solved another problem today. It wasn't difficult at all since I already 
did similar stuff. I tried two different approaches, one using recursion and
the other one used just loop. I wanted to practially compare performance 
implications of two different approaches and not surprisingly, loop-based
approach turned to be better due to logical implementation of recursion,
so as a resuly by using loops i sawed around 15 percents in means of number of
bytes allocated and number of memory allocations, and just a little bit in
means of average time per operation. I've left two functions inside the
`armstrong_numbers.go' file and you are free to see performance difference by
modifying the source code.

### Benchmark results ###

```shell
goos: windows
goarch: amd64
pkg: armstrong
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
BenchmarkIsNumber-24             2010850               598.7 ns/op           264 B/op          8 allocs/op
BenchmarkIsNumber-24             2012874               597.9 ns/op           264 B/op          8 allocs/op
BenchmarkIsNumber-24             1979844               593.8 ns/op           264 B/op          8 allocs/op
BenchmarkIsNumber-24             2018468               593.9 ns/op           264 B/op          8 allocs/op
BenchmarkIsNumber-24             2011494               594.3 ns/op           264 B/op          8 allocs/op
```shell
PASS
ok      armstrong       9.114s
```