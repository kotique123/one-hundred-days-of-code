# Day-7

## Protein translation

Today I've completed yet another easy exercism problem. I've compated my
approach to others' approaches and it proved to be around 50ns/op faster.
Note, that I compare it to top solutions in a list sorted by users with the most
submissions. I think I should solve two problems a day since I have 16 more easy
problems to do, and 65 more medium and hard problems left. I assume that some of
medium and hard problems may take more than a day, so at least, I want to be done
with easy stuff in advance so that I can have some time to work on complicated
problems. The fun thing is, some of the easy probelms required me more than a day
to complete.

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: protein
cpu: Apple M1
```

|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkCodon-8        |66459360                |17.70 ns/op             |0 B/op          |0 allocs/op
|BenchmarkCodon-8        |66159596                |17.69 ns/op             |0 B/op          |0 allocs/op
|BenchmarkCodon-8        |66664195                |17.69 ns/op             |0 B/op          |0 allocs/op
|BenchmarkCodon-8        |66616240                |17.68 ns/op             |0 B/op          |0 allocs/op
|BenchmarkCodon-8        |66326895                |17.75 ns/op             |0 B/op          |0 allocs/op
|BenchmarkProtein-8       |3823500                |308.8 ns/op           |384 B/op         |11 allocs/op
|BenchmarkProtein-8       |3876584                |309.0 ns/op           |384 B/op         |11 allocs/op
|BenchmarkProtein-8       |3866944                |309.7 ns/op           |384 B/op         |11 allocs/op
|BenchmarkProtein-8       |3876843                |309.5 ns/op           |384 B/op         |11 allocs/op
|BenchmarkProtein-8       |3874700                |316.3 ns/op           |384 B/op         |11 allocs/op

```shell
PASS
ok      protein 14.339s
```
