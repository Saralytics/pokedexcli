package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cache[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) error {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		timeAgo := time.Now().Add(-interval)
		for k, v := range c.cache {
			if v.createdAt.Before(timeAgo) {
				delete(c.cache, k)
			}
		}
	}
	return nil
}
