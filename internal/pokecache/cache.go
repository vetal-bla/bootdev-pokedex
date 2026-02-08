package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entry    map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entry:    map[string]CacheEntry{},
		interval: interval,
	}
	go cache.reapPool()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	cache.entry[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.mu.Unlock()
	fmt.Println("ADDED CACHE")
}

func (cache *Cache) Get(key string) (val []byte, isFound bool) {
	var v []byte
	cache.mu.Lock()
	elem, ok := cache.entry[key]
	if !ok {
		cache.mu.Unlock()
		return v, false
	}
	cache.mu.Unlock()
	fmt.Println("FROM CACHE")
	return elem.val, true
}

func (cache *Cache) reapPool() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for range ticker.C {
		cache.mu.Lock()
		now := time.Now()

		for k, v := range cache.entry {
			if now.Sub(v.createdAt) > cache.interval {
				delete(cache.entry, k)
			}
		}
		cache.mu.Unlock()
	}

}
