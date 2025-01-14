# Day 14 

## List Ops

Today I finished List Operations problem from Exercism using golang and
honestly this problem is the good one among the other problems because it
enables your thinking, of course if you actually follow the guidelines.
I saw a lot of people were doing it using built in functions, which you are not
supposed to do for this problem, and their solution performes faster than mine,
probably because I have hot mess here when it comes to memory allocation,
especially in case of Append() it feels like I allocate too much within the
function body and when i call of the function inside other functions, so I am
going to attempt to optimize it in oreder to squeeze out some performance and
further expand my memory management skills.

### Benchmark results

```shell
goos: darwin
goarch: arm64
pkg: listops
cpu: Apple M1
```

|Benchmark Title|Number of iterations|Average time per operation|Number of bytes allocated per operation|Number of memory allocations per operation
|:---|:---|:---|:---|:---
|BenchmarkFold/empty_list-8              |478793122                |2.193 ns/op           |0 B/op          |0 allocs/op
|BenchmarkFold/direction_independent_function_applied_to_non-empty_list-8                |182336690                |6.580 ns/op           |0 B/op          |0 allocs/op
|BenchmarkFold/direction_dependent_function_applied_to_non-empty_list-8                  |448657321                |2.669 ns/op           |0 B/op          |0 allocs/op
|BenchmarkFold/empty_list#01-8                                                           |544229410                |2.195 ns/op           |0 B/op          |0 allocs/op
|BenchmarkFold/direction_independent_function_applied_to_non-empty_list#01-8             |100000000               |10.69 ns/op            |0 B/op          |0 allocs/op
|BenchmarkFold/direction_dependent_function_applied_to_non-empty_list#01-8               |189176812                |6.520 ns/op           |0 B/op          |0 allocs/op
|BenchmarkFilter/empty_list-8                                                            |275606024                |3.817 ns/op           |0 B/op          |0 allocs/op
|BenchmarkFilter/non-empty_list-8                                                        |10378071               |114.8 ns/op            |96 B/op          |4 allocs/op
|BenchmarkLength/empty_list-8                                                            |1000000000               |0.6642 ns/op          |0 B/op          |0 allocs/op
|BenchmarkLength/non-empty_list-8                                                        |470886001                |2.584 ns/op           |0 B/op          |0 allocs/op
|BenchmarkMap/empty_list-8                                                               |542859100                |2.202 ns/op           |0 B/op          |0 allocs/op
|BenchmarkMap/non-empty_list-8                                                           |63224445                |18.62 ns/op            |0 B/op          |0 allocs/op
|BenchmarkReverse/empty_list-8                                                           |316047682                |3.819 ns/op           |0 B/op          |0 allocs/op
|BenchmarkReverse/non-empty_list-8                                                       |51780213                |22.12 ns/op           |32 B/op          |1 allocs/op
|BenchmarkAppend/empty_list-8                                                            |191017354                |6.281 ns/op           |0 B/op          |0 allocs/op
|BenchmarkAppend/empty_list_to_list-8                                                    |34280286                |34.29 ns/op           |32 B/op          |1 allocs/op
|BenchmarkAppend/non-empty_lists-8                                                       |23410351                |47.15 ns/op           |48 B/op          |1 allocs/op
|BenchmarkConcat/empty_list-8                                                            |587063443                |2.099 ns/op           |0 B/op          |0 allocs/op
|BenchmarkConcat/list_of_lists-8                                                          |8830458               |136.4 ns/op            |96 B/op          |3 allocs/op

```shell
PASS
ok      listops 26.626s
```
