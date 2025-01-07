Benchmark results:

```
goos: darwin
goarch: arm64
pkg: protein
cpu: Apple M1
```

|Benchmark Title            |Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---|
BenchmarkCodon-8  |      66459360|                17.70 ns/op|             0 B/op|          0 allocs/op|
BenchmarkCodon-8  |      66159596|                17.69 ns/op|             0 B/op|          0 allocs/op|
BenchmarkCodon-8  |      66664195|                17.69 ns/op|             0 B/op|          0 allocs/op|
BenchmarkCodon-8  |      66616240|                17.68 ns/op|             0 B/op|          0 allocs/op|
BenchmarkCodon-8  |      66326895|                17.75 ns/op|             0 B/op|          0 allocs/op|
BenchmarkProtein-8|       3823500|                308.8 ns/op|           384 B/op|         11 allocs/op|
BenchmarkProtein-8|       3876584|                309.0 ns/op|           384 B/op|         11 allocs/op|
BenchmarkProtein-8|       3866944|                309.7 ns/op|           384 B/op|         11 allocs/op|
BenchmarkProtein-8|       3876843|                309.5 ns/op|           384 B/op|         11 allocs/op|
BenchmarkProtein-8|       3874700|                316.3 ns/op|           384 B/op|         11 allocs/op|

```
PASS
ok      protein 14.339s
```
