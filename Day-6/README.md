Today I wrote a solution to proverb.go from Exercism. There is nothing
remarkable about it since it is pretty straight forward, however I have
noticed something interesting about how other people approach the problem.
For some reason, many of them use append keyword to append a phrase at the end
of a slice. This does not make any sense to me since the slice at the very beginning
is empty and thus there is nothing to append to. Taking a notion of that, what
I did is simply initalizing an slice of predetermined size based on length of an
input slice and then just wrote data in each individual index. As a result, my approach
is more performant and faster than other's approaches.
Benchmark results:
```
goos: windows
goarch: amd64
pkg: proverb
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:---|:---|:---|:---|:---|
BenchmarkProverb-24|      1842192|               657.4 ns/op|          1008 B/op|         22 allocs/op|
BenchmarkProverb-24|      1867722|               643.0 ns/op|          1008 B/op|         22 allocs/op|
BenchmarkProverb-24|      1865642|               665.2 ns/op|          1008 B/op|         22 allocs/op|
BenchmarkProverb-24|      1836740|               669.3 ns/op|          1008 B/op|         22 allocs/op|
BenchmarkProverb-24|      1842373|               644.3 ns/op|          1008 B/op|         22 allocs/op|

```
PASS
ok      proverb 10.597s
```