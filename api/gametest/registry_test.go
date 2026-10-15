package gametest

import (
	"sync"
	"testing"
)

func TestRegistry(t *testing.T) {
	var handlers registry[func()]
	called := false
	first := handlers.add(func() { called = true })
	handlers.get(first)()
	if !called {
		t.Fatal("registered callback was not invoked")
	}
	handlers.remove(first)
	second := handlers.add(func() {})
	if first == second {
		t.Fatal("removed callback ID was reused")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for a removed callback")
		}
	}()
	handlers.get(first)
}

func TestRegistryReentrant(t *testing.T) {
	var handlers registry[func()]
	id := handlers.add(func() {
		nested := handlers.add(func() {})
		handlers.get(nested)()
		handlers.remove(nested)
	})
	handlers.get(id)()
}

func TestRegistryConcurrent(t *testing.T) {
	var handlers registry[func()]
	var workers sync.WaitGroup
	for range 32 {
		workers.Go(func() {
			id := handlers.add(func() {})
			handlers.get(id)()
			handlers.remove(id)
		})
	}
	workers.Wait()
	if len(handlers.handlers) != 0 {
		t.Fatal("removed callbacks remain registered")
	}
}

func TestRegistryExhausted(t *testing.T) {
	handlers := registry[func()]{next: ^uint32(0)}
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic instead of reusing callback IDs")
		}
	}()
	handlers.add(func() {})
}
