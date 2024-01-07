package cache

import "testing"

func Benchmark(b *testing.B) {
	cacheLRU := CreateCacheLRU(3)

	cacheLRU.Set(1, 1)
	cacheLRU.Set(2, 2)
	cacheLRU.Set(3, 3)

	cacheLRU.Get(2)
	cacheLRU.Get(2)

	cacheLRU.Get(3)

	cacheLRU.Set(4, 4)

	cacheLRU.Get(1)
	cacheLRU.Get(4)

	cacheLRU.Get(5)
}
