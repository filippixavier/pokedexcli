package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		store: make(map[string]cacheEntry),
		mu:    sync.Mutex{},
	}

	go func() {
		ticker := time.NewTicker(interval)

		for range ticker.C {
			cache.reapLoop(interval)
		}
	}()

	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	entry := cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	cache.store[key] = entry
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) (val []byte, exist bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.store[key]
	if !ok {
		return val, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	cache.mu.Lock()

	for key, entry := range cache.store {
		if time.Since(entry.createdAt) > interval {
			delete(cache.store, key)
		}
	}

	cache.mu.Unlock()
}
