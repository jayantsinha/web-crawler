package main

import (
	"sync"
	"testing"
)

func TestMap_New(t *testing.T) {
	var m Set
	m.New()
	if len(m.internal) != 0 {
		t.Errorf("Expected length of new map to be %v, but found %v", 0, len(m.internal))
	}
}

// Testing concurrent addition of elements to map
func TestMap_Add(t *testing.T) {
	var m Set
	var wg sync.WaitGroup
	m.New()
	strseq := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for _, e := range strseq {
		wg.Add(1)

		go func(m *Set, e string) {
			defer wg.Done()
			m.Add(e)
		}(&m, e)
	}
	wg.Wait()

	if len(m.internal) != len(strseq) {
		t.Errorf("Expected size of map to be %v, found %v", len(strseq), len(m.internal))
	}

}

func TestMap_Contains(t *testing.T) {
	var m Set
	var wg sync.WaitGroup
	m.New()
	strseq := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	// adding z synchronously to test call in between threads
	m.Add("z")
	for i, e := range strseq {
		wg.Add(1)

		go func(m *Set, e string) {
			defer wg.Done()
			// Checking for contains in between threads creating a read lock on the map.
			if i%2 != 0 {
				if !m.Contains("z") {
					t.Errorf("Expects element \"z\" to be present in the map but condition returned false.")
				}
			}
			m.Add(e)
		}(&m, e)
	}
	wg.Wait()
}
