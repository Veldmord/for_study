package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       sync.Mutex{},
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if value, ok := l.items[key]; ok {
		value.Value = cacheItem{key: key, value: value}
		l.queue.MoveToFront(value)
		return true
	}

	if l.queue.Len() >= l.capacity {
		oldItem := l.queue.Back()
		l.queue.Remove(oldItem)
		delete(l.items, oldItem.Value.(cacheItem).key)
	}

	l.items[key] = l.queue.PushFront(cacheItem{key: key, value: value})
	return false

}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if v, ok := l.items[key]; ok {
		l.queue.MoveToFront(v)
		return v.Value.(cacheItem).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = make(map[Key]*ListItem)
	l.queue = NewList()
}
