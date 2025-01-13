Today I wrote a solution for rotational-cipher problem from exercism.org

Benchmark results:
```
goos: windows
goarch: amd64
pkg: rotationalcipher
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:--------------------------|:-------------------|:-------------------------|:--------------------------------------|:-----------------------------------------|
BenchmarkRotationalCipher-24|459908              |2403 ns/op                |2128 B/op                              |183 allocs/op
BenchmarkRotationalCipher-24|477294              |2413 ns/op                |2128 B/op                              |183 allocs/op
BenchmarkRotationalCipher-24|486903              |2410 ns/op                |2128 B/op                              |183 allocs/op
BenchmarkRotationalCipher-24|482421              |2445 ns/op                |2128 B/op                              |183 allocs/op
BenchmarkRotationalCipher-24|500431              |2427 ns/op                |2128 B/op                              |183 allocs/op

```
PASS
ok      rotationalcipher        6.865s
```
