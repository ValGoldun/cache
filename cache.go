package cache

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	data      map[K]entity[V]
	ttl       time.Duration
	scheduler *scheduler[K]
	sync.RWMutex
}

type entity[V any] struct {
	value    V
	deadline time.Time
}

func New[K comparable, V any]() *Cache[K, V] {
	c := &Cache[K, V]{
		data: make(map[K]entity[V]),
		scheduler: &scheduler[K]{
			queue: make(map[K]cleaner),
		},
	}

	go c.scheduler.schedule()

	return c
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.Lock()
	defer c.Unlock()

	var deadline time.Time
	if c.ttl == 0 {
		deadline = time.Time{}
	} else {
		deadline = time.Now().Add(c.ttl)

		c.scheduler.add(key, c.newCleaner(key, deadline))
	}

	c.data[key] = entity[V]{value: value, deadline: deadline}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.RLock()
	defer c.RUnlock()

	e, exist := c.data[key]
	if !exist {
		return e.value, false
	}

	if !e.deadline.IsZero() && time.Now().After(e.deadline) {
		c.Delete(key)

		return e.value, false
	}

	return e.value, true
}

func (c *Cache[K, V]) Delete(key K) {
	c.Lock()
	defer c.Unlock()

	delete(c.data, key)
}
