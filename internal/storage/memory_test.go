package storage

import (
	"fmt"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	tests := []struct {
		name    string
		slot    int
		content string
	}{
		{"store and retrieve normal text", 1, "test content"},
		{"store and retrieve empty string", 2, ""},
		{"store and retrieve with special chars", 3, "test\n\t!@#$%^&*()"},
		{"store in slot 0", 0, "slot zero content"},
		{"store in high number slot", 9, "slot nine content"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMemoryStorage()

			// Test Set
			err := ms.Set(tt.slot, tt.content)
			if err != nil {
				t.Errorf("Set() error = %v", err)
			}

			// Test Get
			got, err := ms.Get(tt.slot)
			if err != nil {
				t.Errorf("Get() error = %v", err)
			}
			if got != tt.content {
				t.Errorf("Get() = %v, want %v", got, tt.content)
			}
		})
	}
}

func TestMemoryStorage_GetNonExistent(t *testing.T) {
	ms := NewMemoryStorage()

	_, err := ms.Get(1)
	if err == nil {
		t.Error("Get() expected error for non-existent slot")
	}
}

func TestMemoryStorage_Concurrent(t *testing.T) {
	ms := NewMemoryStorage()
	const goroutines = 10

	done := make(chan bool, goroutines*2)

	for i := 0; i < goroutines; i++ {
		go func(slot int) {
			err := ms.Set(slot, fmt.Sprintf("content%d", slot))
			if err != nil {
				t.Errorf("Concurrent Set() error = %v", err)
			}
			done <- true
		}(i)

		go func(slot int) {
			_, err := ms.Get(slot)
			_ = err
			done <- true
		}(i)
	}

	// wait for all goroutines to finish
	for i := 0; i < goroutines*2; i++ {
		<-done
	}
}

func TestMemoryStorage_Overwrite(t *testing.T) {
	ms := NewMemoryStorage()
	slot := 1

	// first write
	content1 := "first content"
	err := ms.Set(slot, content1)
	if err != nil {
		t.Errorf("First Set() error = %v", err)
	}

	// Verify first write
	got, err := ms.Get(slot)
	if err != nil {
		t.Errorf("First Get() error = %v", err)
	}
	if got != content1 {
		t.Errorf("First Get() = %v, want %v", got, content1)
	}

	// Overwrite
	content2 := "second content"
	err = ms.Set(slot, content2)
	if err != nil {
		t.Errorf("Second Set() error = %v", err)
	}

	// Verify overwrite
	got, err = ms.Get(slot)
	if err != nil {
		t.Errorf("Second Get() error = %v", err)
	}
	if got != content2 {
		t.Errorf("Second Get() = %v, want %v", got, content2)
	}
}
