package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu *sync.Mutex
	cache map[string]cacheEntry
	timeOut time.Duration
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		mu: &sync.Mutex{},
		cache: map[string]cacheEntry{},
	}
	go newCache.readLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
	c.cache[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cache, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cache.val, true
}

func (c *Cache) readLoop(interval time.Duration) {
	TickCh := time.NewTicker(interval)
	defer TickCh.Stop()
	for cTime := range TickCh.C{
		c.mu.Lock()
		for cKey, cEntry:= range c.cache {
			if cTime.Sub(cEntry.createdAt) >= interval {
					delete(c.cache, cKey)
			}
		}
		c.mu.Unlock()
	}
}
