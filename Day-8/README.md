# Day 8

## Angram

Today I've started solving angram problem and it was pretty strenous. I figured out a pretty clever way of solving it without converting
the whole string to lowercase, thus optimizing performance. However, my approach could not be applied to non latin characters, so 
I was forced to convert the input into lowercase so that I could pass the tests. At the end of the day, my solution is still quite
performant, compared to approaches from others, but there is still room for improvement because some things can still be optimized.
Yesterday, I've mentioned that I should solve more than one problem a day, however, as I anticipated, some 'easy' problems may take 
way more time to solve them, compared to other 'easy' problems. Tomorrow is promising to be a free day due to weather conditions, so
I hope I will be able to finish more problems.

### Benchmark results

```shell
goos: windows
goarch: amd64
pkg: anagram
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkDetectAnagrams-24        |242913              |4449 ns/op             |640 B/op         |36 allocs/op
|BenchmarkDetectAnagrams-24        |267164              |4458 ns/op             |640 B/op         |36 allocs/op
|BenchmarkDetectAnagrams-24        |267582              |4450 ns/op             |640 B/op         |36 allocs/op
|BenchmarkDetectAnagrams-24        |264844              |4452 ns/op             |640 B/op         |36 allocs/op
|BenchmarkDetectAnagrams-24        |269941              |4436 ns/op             |640 B/op         |36 allocs/op

```shell
PASS
ok      anagram 6.188s
```
