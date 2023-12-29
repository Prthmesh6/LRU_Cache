package main

import (
	"context"
	"fmt"

	Cache "github.com/Prthmesh6/lru_cache/Cache/LRU"
)

func main() {
	//initialise LRU cache by providing key-value data types and maxcapacity
	maxCapacity := 2
	lruCache := Cache.NewLruCache[int, string](context.Background(), maxCapacity)

	//You can also get lruCache like below, you just need to define the
	//data types for key and value, so that particular object of lrucache
	//will accept only those data types. Example :-
	//lruCache := new(lru.LRU[int, string])
	//lruCache2 := new(lru.LRU[string, string])
	//lruCache2.Set("Hello", "world")

	//adding the value string for key int
	lruCache.Set(context.Background(), 1, "Hello")
	lruCache.Set(context.Background(), 2, "world")
	lruCache.Set(context.Background(), 3, "world")
	lruCache.Set(context.Background(), 4, "world")
	lruCache.Set(context.Background(), 5, "world")

	//This will be handled gracefully as 2 is not a present key, it is overwritten by 5
	fmt.Println(lruCache.Get(context.Background(), 2))

	// key 5 have value as "world" so it will be returned
	value, err := lruCache.Get(context.Background(), 5)
	if err != nil {
		fmt.Println("Eror occured while getting value :- ", err)
	}
	fmt.Println(value)

}
