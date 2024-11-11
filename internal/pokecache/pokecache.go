package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	data, ok := c.entries[key]
	c.mu.Unlock()
	return data.val, ok
}
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for currentTime := range ticker.C {
		limit := currentTime.Add(-interval)
		c.mu.Lock()
		for key, value := range c.entries {
			if value.createdAt.Before(limit) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)

	return &newCache
}
