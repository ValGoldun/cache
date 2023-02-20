package cache

import (
	"log"
	"sync"
	"time"
)

type scheduler[K comparable] struct {
	queue map[K]cleaner
	sync.Mutex
}

func (s *scheduler[K]) add(key K, f cleaner) {
	s.Lock()
	defer s.Unlock()

	s.queue[key] = f
}

func (s *scheduler[K]) delete(key K) {
	s.Lock()
	defer s.Unlock()

	delete(s.queue, key)
}

func (s *scheduler[K]) schedule() {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		for key, f := range s.queue {
			ok := f()
			log.Println(ok)
			if ok {
				s.delete(key)
			}
		}
	}

}
