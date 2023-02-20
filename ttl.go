package cache

import "time"

func NewWithTTL[K comparable, V any](ttl time.Duration) *Cache[K, V] {
	c := New[K, V]()

	c.ttl = ttl

	return c
}

func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	var deadline = time.Now().Add(ttl)

	c.scheduler.add(key, c.newGC(key, deadline))

	c.data[key] = entity[V]{value: value, deadline: deadline}
}
