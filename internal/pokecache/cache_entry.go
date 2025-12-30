package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	v  map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	var newEntry = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.v[key] = newEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.v[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		tickTime := time.Now()
		c.mu.Lock()
		for key, entry := range c.v {
			lifetime := tickTime.Sub(entry.createdAt)
			if lifetime >= interval {
				delete(c.v, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	emptyMap := make(map[string]cacheEntry)

	c := Cache{
		v: emptyMap,
	}
	go c.reapLoop(interval)
	return &c
}
