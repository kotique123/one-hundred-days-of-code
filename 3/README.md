Today I wrote a solution for yet another Golang exercise on Exercism. I have 
noticed that the solutions that other people on Exercism submit include some 
type of data stucture involved. However, I suspect that it makes their 
solutions less performant than mine.

Benchmark results:
```
goos: windows
goarch: amd64
pkg: resistorcolorduo
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:---|:---|:---|:---|:---|
BenchmarkValue-24|       39246852|                29.78 ns/op|            0 B/op|          0 allocs/op|
BenchmarkValue-24 |      40437260 |               29.58 ns/op |           0 B/op |         0 allocs/op|
BenchmarkValue-24  |     38478310  |              29.62 ns/op  |          0 B/op  |        0 allocs/op|
BenchmarkValue-24   |    38946370   |             29.48 ns/op   |         0 B/op   |       0 allocs/op|
BenchmarkValue-24    |   40375219    |            29.94 ns/op    |        0 B/op    |      0 allocs/op|
```
PASS
ok      resistorcolorduo        6.149s
```