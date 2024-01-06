package main

import (
	"cache-lru/cache"
	"fmt"
)

func main() {
	cacheLRU := cache.CreateCacheLRU(3)

	cacheLRU.Set(1, 1)
	cacheLRU.Set(2, 2)
	cacheLRU.Set(3, 3)

	fmt.Println(cacheLRU.Get(2))
	fmt.Println(cacheLRU.Get(2))

	fmt.Println(cacheLRU.Get(3))

	cacheLRU.Set(4, 4)

	fmt.Println(cacheLRU.Get(1))
	fmt.Println(cacheLRU.Get(4))

	fmt.Println(cacheLRU.Get(5))
}
