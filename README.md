# CACHE LRU

## Introduction

Cache LRU is a cache evict policy to delete the Least Recently Used cache entry, used when the cache is at the maximun capacity and we're trying to add a new entry.

![cache-lru-idea](./assets/cache-lru-idea)

## Class diagram

![class diagram](./assets/class-diagram.png)

## Unit tests

```
go test ./... -cover
?       cache-lru/cmd   [no test files]
ok      cache-lru/cache (cached)        coverage: 100.0% of statements
```

## Benchmark

Multiple scenarios were ran, from 100 to 1000000 cache limits, and we still got pretty good performance.



```
go test -benchmem -bench BenchmarkCacheLRU
goos: windows
goarch: amd64
pkg: cache-lru/cache
cpu: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
BenchmarkCacheLRU/100-12                 9780078               106.1 ns/op            32 B/op          1 allocs/op
BenchmarkCacheLRU/1000-12               10609807               107.6 ns/op            32 B/op          1 allocs/op
BenchmarkCacheLRU/10000-12              11154852               108.7 ns/op            32 B/op          1 allocs/op
BenchmarkCacheLRU/100000-12             10704116               106.7 ns/op            32 B/op          1 allocs/op
BenchmarkCacheLRU/1000000-12            11428690               105.3 ns/op            32 B/op          1 allocs/op
PASS
ok      cache-lru/cache 8.106s
```

Let's analyze the results:

* BenchmarkCacheLRU/100-12: 9,780,078 iterations per operation with an average time of 106.1 nanoseconds. During each operation, 32 bytes were allocated in 1 allocation.

* BenchmarkCacheLRU/1000-12: 10,609,807 iterations per operation with an average time of 107.6 nanoseconds. Similar to the previous case, 32 bytes were allocated in 1 allocation during each operation.

* BenchmarkCacheLRU/10000-12: 11,154,852 iterations per operation with an average time of 108.7 nanoseconds. The memory allocation pattern remains the same: 32 bytes in 1 allocation per operation.

* BenchmarkCacheLRU/100000-12: 10,704,116 iterations per operation with an average time of 106.7 nanoseconds. Memory allocation remains consistent at 32 bytes in 1 allocation per operation.

* BenchmarkCacheLRU/1000000-12: 11,428,690 iterations per operation with an average time of 105.3 nanoseconds. Similar to the other cases, 32 bytes were allocated in 1 allocation during each operation.

## Disclaimer

This is a simple challenge, made to practice with use of different data structure i'm used to.

Is just for fun!
