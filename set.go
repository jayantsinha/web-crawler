package main

import "sync"

type Set struct {
	internal  map[string]bool
	sync.RWMutex
}

// New create a new map instance
func (s *Set) New() *Set {
	s.internal = make(map[string]bool)
	return s
}

// Add adds a key into the map
func (s *Set) Add(e string) {
	s.Lock()
	s.internal[e] = true
	s.Unlock()
}

func (s *Set) Contains(e string) bool {
	s.RLock()
	defer func() {
		s.RUnlock()
	}()
	if len(s.internal) == 0 {
		return false
	}
	return s.internal[e]
}

func (s *Set) Size() int {
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	return len(s.internal)
}