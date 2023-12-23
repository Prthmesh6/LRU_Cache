package main

import (
	Cache "LRU_cache/Cache/LRU"
	"fmt"
)

func main() {
	lruCache := Cache.New[int, string](2)
	// lruCache := new(lru.LRU[int, string])
	// lruCache2 := new(lru.LRU[string, string])

	// lruCache2.Set("Hello", "world")

	lruCache.Set(1, "Hello")
	lruCache.Set(2, "world")
	lruCache.Set(3, "world")
	lruCache.Set(4, "world")
	lruCache.Set(5, "world")
	fmt.Println(lruCache.Hashmap)
	fmt.Println(lruCache.DoublyLinkedList.Front().Value)
	fmt.Println(lruCache.DoublyLinkedList.Back().Value)
	for i := lruCache.DoublyLinkedList.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
	}

}
