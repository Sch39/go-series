package registry

import "sync"

type Registry struct {
	mu      sync.RWMutex
	entries map[string]int
}

func New() *Registry {
	return &Registry{entries: make(map[string]int)}
}

func (r *Registry) Put(key string, value int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries[key] = value
}

func (r *Registry) Lookup(key string) (int, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok := r.entries[key]
	return value, ok
}

func (r *Registry) Delete(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.entries, key)
}
