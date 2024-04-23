package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	m       *sync.Mutex
}

func NewCache(duration time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		m:       &sync.Mutex{},
	}
	go c.reapLoop(duration)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	val, ok := c.entries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		c.m.Lock()
		for k, v := range c.entries {
			if v.createdAt.Add(duration).Before(time.Now()) {
				delete(c.entries, k)
			}
		}
		c.m.Unlock()
	}

}
