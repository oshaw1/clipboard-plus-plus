package clipboard

import (
	"testing"
)

// mock storage for testing
type mockStorage struct {
	data map[int]string
}

func newMockStorage() *mockStorage {
	return &mockStorage{
		data: make(map[int]string),
	}
}

func (ms *mockStorage) Get(slot int) (string, error) {
	text, exists := ms.data[slot]
	if !exists {
		return "", nil
	}
	return text, nil
}

func (ms *mockStorage) Set(slot int, content string) error {
	ms.data[slot] = content
	return nil
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{"short string", "test", 10, "test"},
		{"exact length", "test", 4, "test"},
		{"needs truncating", "testing", 4, "test..."},
		{"empty string", "", 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := truncateString(tt.input, tt.length)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestStorage(t *testing.T) {
	store := newMockStorage()

	testContent := "test content"
	err := store.Set(1, testContent)
	if err != nil {
		t.Errorf("failed to set content: %v", err)
	}

	content, err := store.Get(1)
	if err != nil {
		t.Errorf("failed to get content: %v", err)
	}
	if content != testContent {
		t.Errorf("got %q, want %q", content, testContent)
	}

	content, err = store.Get(2)
	if content != "" {
		t.Errorf("expected empty string for unused slot, got %q", content)
	}
}
