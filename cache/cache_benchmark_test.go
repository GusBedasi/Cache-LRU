package cache

import (
	"math/rand"
	"testing"
)

func BenchmarkCacheLRU(b *testing.B) {
	scenarios := []struct {
		name      string
		cacheSize int
	}{
		{
			name:      "100",
			cacheSize: 100,
		},
		{
			name:      "1000",
			cacheSize: 1000,
		},
		{
			name:      "10000",
			cacheSize: 10000,
		},
		{
			name:      "100000",
			cacheSize: 100000,
		},
		{
			name:      "1000000",
			cacheSize: 1000000,
		},
	}

	for _, scenario := range scenarios {
		b.Run(scenario.name, func(b *testing.B) {
			// Set up your cache with initial values as needed for the benchmark
			cacheLRU := CreateCacheLRU(3)

			// Fill cache
			fillCache(cacheLRU)

			// Reset the benchmark timer after setup
			b.ResetTimer()

			var result int

			for i := 0; i < b.N; i++ {
				result = cacheLRU.Get(rand.Intn(cacheLRU.capacity))
				cacheLRU.Set(rand.Intn(cacheLRU.capacity), rand.Intn(cacheLRU.capacity))
			}

			_ = result
		})
	}
}

func fillCache(cache *CacheLRU) {
	for i := 0; i <= cache.capacity; i++ {
		cache.Set(i, i)
	}
}
