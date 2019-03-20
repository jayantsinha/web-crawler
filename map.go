package main

import "sync"

type Map struct {
	set  map[string]bool
	sync.RWMutex
}

// New create a new map instance
func (m *Map) New() *Map {
	m.set = make(map[string]bool)
	return m
}

// Add adds a key into the map
func (m *Map) Add(e string) {
	m.Lock()
	m.set[e] = true
	m.Unlock()
}

func (m *Map) Contains(e string) bool {
	m.Lock()
	defer func() {
		m.Unlock()
	}()
	if len(m.set) == 0 {
		return false
	}
	return m.set[e]
}