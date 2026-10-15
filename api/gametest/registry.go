package gametest

import (
	"fmt"
	"sync"
)

type registry[Handler any] struct {
	mutex    sync.RWMutex
	next     uint32
	handlers map[uint32]Handler
}

func (r *registry[Handler]) add(handler Handler) uint32 {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.next == ^uint32(0) {
		panic("gametest: callback IDs exhausted")
	}
	r.next++
	if r.handlers == nil {
		r.handlers = make(map[uint32]Handler)
	}
	r.handlers[r.next] = handler
	return r.next
}

func (r *registry[Handler]) get(id uint32) Handler {
	r.mutex.RLock()
	handler, exists := r.handlers[id]
	r.mutex.RUnlock()
	if !exists {
		panic(fmt.Sprintf("gametest: unknown callback ID %d", id))
	}
	return handler
}

func (r *registry[Handler]) remove(id uint32) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.handlers, id)
}
