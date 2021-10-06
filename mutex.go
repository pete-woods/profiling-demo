package profiling_demo

import (
	"sync"
)

type AtomicVariable struct {
	mu  sync.RWMutex
	val uint64
}

func (av *AtomicVariable) Inc() {
	av.mu.Lock()
	defer av.mu.Unlock()

	av.val++
}

func (av *AtomicVariable) Get() uint64 {
	av.mu.RLock()
	defer av.mu.RUnlock()

	return av.val
}
