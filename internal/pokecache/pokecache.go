package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mutex    *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time 
	val       []byte 
}

// Cache Constructor
func NewCache(interval time.Duration) Cache {

	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		mutex: new(sync.RWMutex),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, value []byte) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cacheMap == nil {
		c.cacheMap = make(map[string]cacheEntry)
	}

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.cacheMap == nil {
		return nil, false
	}

	entry, exists := c.cacheMap[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {

	ticker := time.NewTicker(c.interval)
    for range ticker.C {
        c.mutex.Lock()
		for key, value := range c.cacheMap {
			currentTime := time.Now()
			difference := currentTime.Sub(value.createdAt)
			if difference > c.interval {
				delete(c.cacheMap, key)
			}
		}
		c.mutex.Unlock()
    }
}
