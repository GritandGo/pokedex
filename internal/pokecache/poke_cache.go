package pokecache
import (
	"time"
	"sync"
)



type Cache struct {
	cacheList map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}


type cacheEntry struct {
	createdAt time.Time
	val []byte 
	}




func NewCache(interval time.Duration) *Cache {
	cache := make(map[string]cacheEntry)

	c := Cache{
		cacheList: cache,
		interval: interval,
	}
	go c.reapLoop()
	return &c
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.cacheList[key] = entry
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheList[key]
	return entry.val, ok
}


func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cacheList {
			start := entry.createdAt
			elapsed := time.Since(start)
			if elapsed > c.interval {
				delete(c.cacheList, key)
			}


			}

			c.mu.Unlock()
		}

		
	}	
	
