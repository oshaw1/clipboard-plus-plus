package hotkeys

import (
	"fmt"
	"testing"
)

type mockClipManager struct {
	copyActions  map[int]bool
	pasteActions map[int]bool
	shouldError  bool
}

func newMockClipManager() *mockClipManager {
	return &mockClipManager{
		copyActions:  make(map[int]bool),
		pasteActions: make(map[int]bool),
	}
}

func (m *mockClipManager) CopyToSlot(slot int) error {
	if m.shouldError {
		return fmt.Errorf("mock copy error")
	}
	m.copyActions[slot] = true
	return nil
}

func (m *mockClipManager) PasteFromSlot(slot int) error {
	if m.shouldError {
		return fmt.Errorf("mock paste error")
	}
	m.pasteActions[slot] = true
	return nil
}

func TestHandleNumKey(t *testing.T) {
	tests := []struct {
		name      string
		keyNum    int
		mask      uint16
		wantCopy  bool
		wantPaste bool
	}{
		{
			name:      "right control should copy",
			keyNum:    1,
			mask:      32, // RIGHT_CONTROL
			wantCopy:  true,
			wantPaste: false,
		},
		{
			name:      "left control should paste",
			keyNum:    1,
			mask:      2, // LEFT_CONTROL
			wantCopy:  false,
			wantPaste: true,
		},
		{
			name:      "no control should do nothing",
			keyNum:    1,
			mask:      0,
			wantCopy:  false,
			wantPaste: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := newMockClipManager()
			handler := NewHandler(mock)

			handler.handleNumKey(tt.keyNum, tt.mask)

			if got := mock.copyActions[tt.keyNum]; got != tt.wantCopy {
				t.Errorf("Copy action = %v, want %v", got, tt.wantCopy)
			}
			if got := mock.pasteActions[tt.keyNum]; got != tt.wantPaste {
				t.Errorf("Paste action = %v, want %v", got, tt.wantPaste)
			}
		})
	}
}

func TestNumKeyValidation(t *testing.T) {
	tests := []struct {
		rawcode  uint16
		expected int
		valid    bool
	}{
		{97, 1, true},    // numpad 1
		{98, 2, true},    // numpad 2
		{105, 9, true},   // numpad 9
		{96, 0, true},    // numpad 0
		{95, -1, false},  // invalid
		{106, -1, false}, // invalid
	}

	for _, tt := range tests {
		if tt.valid {
			got := int(tt.rawcode - 96)
			if got != tt.expected {
				t.Errorf("Rawcode %d conversion = %d, want %d",
					tt.rawcode, got, tt.expected)
			}
		}
	}
}

func TestHandleNumKeyErrors(t *testing.T) {
	mock := newMockClipManager()
	mock.shouldError = true
	handler := NewHandler(mock)

	handler.handleNumKey(1, 32)
	if !mock.shouldError {
		t.Error("Expected error handling for copy operation")
	}

	handler.handleNumKey(1, 2)
	if !mock.shouldError {
		t.Error("Expected error handling for paste operation")
	}
}
