package storage

import (
	"fmt"
	"sync"
)

type MemoryStorage struct {
	slots map[int]string
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		slots: make(map[int]string),
	}
}

func (ms *MemoryStorage) Get(slot int) (string, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	content, exists := ms.slots[slot]
	if !exists {
		return "", fmt.Errorf("no content in slot %d", slot)
	}
	return content, nil
}

func (ms *MemoryStorage) Set(slot int, content string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.slots[slot] = content
	return nil
}
