package cache

import (
	"time"
)

type cleaner func() bool

func (c *Cache[K, V]) newCleaner(key K, deadline time.Time) cleaner {
	return func() bool {
		if deadline.Before(time.Now()) {
			c.Delete(key)

			return true
		}
		return false
	}
}
