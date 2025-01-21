# Day 10

## Sum of Multiples

Today I have solved the sum of multiples problem in Golang. Honestly, I don't like
my approach at all since it is too point blank and basically bruteforcing. I tried to
optimize it by preallocating some memory for the slice so that append will be more
efficient, yet it is still atrociously costly and slow. I've reasearched different
approaches and there are plenty of those that are way faster than mine since
they don't have to sort and remove duplicate elements from a slice of multiples.
Therefore they don't need to store multiples in a slice and so they just add values,
which makes their approach way faster. Anyways, I don't feel good about just copying
their work, so here I post my solution which I don't like in anyways. However, let it
be it. I am aiming to solve one more problem today later on(I am writing this around 1am).

### Benchmark results ###

```shell
goos: darwin
goarch: arm64
pkg: summultiples
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkSumMultiples-8             |3512            |341174 ns/op          |364449 B/op         |17 allocs/op
|BenchmarkSumMultiples-8             |3380            |343087 ns/op          |364450 B/op         |17 allocs/op
|BenchmarkSumMultiples-8             |3452            |346530 ns/op          |364450 B/op         |17 allocs/op
|BenchmarkSumMultiples-8             |3391            |341496 ns/op          |364449 B/op         |17 allocs/op
|BenchmarkSumMultiples-8             |3481            |341536 ns/op          |364450 B/op         |17 allocs/op

```shell
PASS
ok      summultiples    6.958s
```
