package cache

import "time"

func NewWithTTL[K comparable, V any](ttl time.Duration) *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]entity[V]),
		ttl:  ttl,
	}
}

func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.data[key] = entity[V]{value: value, deadline: time.Now().Add(ttl)}
}
