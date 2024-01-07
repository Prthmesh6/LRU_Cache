package Cache

import (
	"container/list"
	"context"
	"fmt"
	"sync"
)

type lru[T1 comparable, T2 comparable] struct {
	capacity         int
	hashmap          map[T1]*list.Element
	doublyLinkedList *list.List
	mutex            sync.Mutex
}

func NewLruCache[T1 comparable, T2 comparable](ctx context.Context, capacity int) *lru[T1, T2] {
	hashmap := make(map[T1]*list.Element)
	return &lru[T1, T2]{
		capacity:         capacity,
		hashmap:          hashmap,
		doublyLinkedList: list.New(),
		mutex:            sync.Mutex{},
	}
}

type InsertNode[T1 comparable, T2 comparable] struct {
	key T1
	val T2
}

func (l *lru[T1, T2]) Set(ctx context.Context, key T1, element T2) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	insertNode := getInsertNode(key, element)
	//check whether value is already present in LRU
	if valuePresent, ok := (l.hashmap)[key]; ok {
		valuePresent.Value = insertNode
		l.doublyLinkedList.MoveToFront(valuePresent)
	} else {
		(l.hashmap)[key] = l.doublyLinkedList.PushFront(insertNode)
		//remove old element if capacity exceeds
		if len(l.hashmap) > l.capacity {
			removedElement := l.doublyLinkedList.Remove(l.doublyLinkedList.Back())
			removedElement1, ok := (removedElement).(*InsertNode[T1, T2])

			if ok {
				keyToDelete := removedElement1.key
				delete(l.hashmap, keyToDelete)
			} else {
				return fmt.Errorf("errKeyNotAdded")
			}
		}
	}
	return nil
}

func (l *lru[T1, T2]) Get(ctx context.Context, element T1) (result T2, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	val, ok := (l.hashmap)[element]
	if ok {
		l.doublyLinkedList.MoveToFront(val)
	}
	ans := (l.hashmap)[element]
	if ans == nil {
		return result, fmt.Errorf("errKeyNotAvailable %v", element)
	}
	return ans.Value.(*InsertNode[T1, T2]).val, nil
}

func (l *lru[T1, T2]) GetCapacity(ctx context.Context) (int, error) {
	return l.capacity, nil
}

func getInsertNode[T1 comparable, T2 comparable](key T1, value T2) *InsertNode[T1, T2] {
	return &InsertNode[T1, T2]{
		key: key,
		val: value,
	}
}
