package memorycache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	memory map[keys]any
	mu     sync.RWMutex
}

// Create new cache object
func (c *Cache) New() *Cache {
	return &Cache{make(map[keys]any), *new(sync.RWMutex)}
}

// Set or change value in memory
func (c *Cache) Set(key keys, val any, ttl time.Duration) {
	c.mu.Lock()
	c.memory[key] = val
	c.mu.Unlock()

	go func() {
		time.Sleep(ttl)
		c.mu.Lock()
		delete(c.memory, key)
		c.mu.Unlock()
	}()
}

// Get value from memory
func (c *Cache) Get(key keys) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !key.In(c.memory) {
		return nil, errors.New("invalid key")
	}
	return c.memory[key], nil
}

func (c *Cache) Delete(key keys) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !key.In(c.memory) {
		return errors.New("invalid key")
	}
	delete(c.memory, key)
	return nil
}
