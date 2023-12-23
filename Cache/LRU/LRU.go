package Cache

import (
	"container/list"
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

type lru[T1 comparable, T2 comparable] struct {
	Capacity         int
	Hashmap          *map[T1]*list.Element
	DoublyLinkedList *list.List
}

func New[T1 comparable, T2 comparable](capacity int) *lru[T1, T2] {
	hashmap := make(map[T1]*list.Element)
	return &lru[T1, T2]{
		Capacity:         capacity,
		Hashmap:          &hashmap,
		DoublyLinkedList: list.New(),
	}
}

type InsertNode[T1 comparable, T2 comparable] struct {
	key T1
	val T2
}

func (l *lru[T1, T2]) Set(key T1, element T2) {
	mutex.Lock()
	defer mutex.Unlock()

	insertNode := getInsertNode(key, element)
	//check whether value is already present in LRU
	if valuePresent, ok := (*l.Hashmap)[key]; ok {
		valuePresent.Value = insertNode
		l.DoublyLinkedList.MoveToFront(valuePresent)
	} else {
		(*l.Hashmap)[key] = l.DoublyLinkedList.PushFront(insertNode)
		//remove old element if capacity exceeds
		if len(*l.Hashmap) > l.Capacity {
			removedElement := l.DoublyLinkedList.Remove(l.DoublyLinkedList.Back())
			removedElement1, ok := (removedElement).(*InsertNode[T1, T2])

			if ok {
				keyToDelete := removedElement1.key
				delete(*l.Hashmap, keyToDelete)
			}
		}
	}
}

func (l *lru[T1, T2]) Get(element T1) (result T2, err error) {
	mutex.Lock()
	defer mutex.Unlock()
	val, ok := (*l.Hashmap)[element]
	if ok {
		l.DoublyLinkedList.MoveToFront(val)
	}
	ans := (*l.Hashmap)[element]
	if ans == nil {
		return result, fmt.Errorf("key %v is not present in cache", element)
	}
	return ans.Value.(*InsertNode[T1, T2]).val, nil
}

func (l lru[T1, T2]) GetCapacity() int {
	return l.Capacity
}

func getInsertNode[T1 comparable, T2 comparable](key T1, value T2) *InsertNode[T1, T2] {
	return &InsertNode[T1, T2]{
		key: key,
		val: value,
	}
}
