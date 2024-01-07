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

```
goos: windows
goarch: amd64
pkg: cache-lru/cache
cpu: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
=== RUN   Benchmark
Benchmark
Benchmark-12
1000000000             0 B/op          0 allocs/op
PASS
ok      cache-lru/cache 2.854s
```

## Disclaimer

This is a simple challenge, made to practice with use of different data structure i'm used to.

Is just for fun!
