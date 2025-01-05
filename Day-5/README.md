Today I wrote a solution to atbash_cipher.go and wrote some comments about my solution.
I tried to make my solution performant yet easy to comprehend, and I think that I have
successfully reached my goal. I have compared my solution performance to other solutions
and it prooved to be better than average, however, 
(this solution)[https://exercism.org/tracks/go/exercises/atbash-cipher/solutions/bobahop)]
appeared to be around 4 times more performant than mine which is not surprising since I
traverse through a string at least 4 times while a solution made by bobahop traverses
through a string once. So I copied his code and put in in comments of my solution and annotated
it because I find it amazing.
Here are the benchmarks:
```
goos: windows
goarch: amd64
pkg: atbash
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```
|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation                    
|:---|:---|:---|:---|:---|
BenchmarkAtbash-24|        791431|              1398 ns/op|            1512 B/op|         40 allocs/op|
BenchmarkAtbash-24|        831375|              1401 ns/op|            1512 B/op|         40 allocs/op|
BenchmarkAtbash-24|        832581|              1406 ns/op|            1512 B/op|         40 allocs/op|
BenchmarkAtbash-24|        853891|              1406 ns/op|            1512 B/op|         40 allocs/op|
BenchmarkAtbash-24|        847528|              1400 ns/op|            1512 B/op|         40 allocs/op|
```
PASS
ok      atbash  6.936s
```