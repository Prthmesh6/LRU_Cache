package main

import (
	"fmt"

	Cache "github.com/Prthmesh6/lru_cache/Cache/LRU"
)

func main() {
	//initialise LRU cache by providing key-value data types and maxcapacity
	maxCapacity := 2
	lruCache := Cache.New[int, string](maxCapacity)

	//You can also get lruCache like below,
	// lruCache := new(lru.LRU[int, string])
	// lruCache2 := new(lru.LRU[string, string])
	// lruCache2.Set("Hello", "world")

	lruCache.Set(1, "Hello")
	lruCache.Set(2, "world")
	lruCache.Set(3, "world")
	lruCache.Set(4, "world")
	lruCache.Set(5, "world")

	//This will be handled gracefully as 2 is not a present key, it is overwritten by 5
	fmt.Println(lruCache.Get(2))

	// key 5 have value as "world" so it will be returned
	value, err := lruCache.Get(5)
	if err != nil {
		fmt.Println("Eror occured while getting value :- ", err)
	}
	fmt.Println(value)

}
