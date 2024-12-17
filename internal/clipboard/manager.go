package clipboard

import (
	"clipboardplusplus/internal/storage"
	"fmt"

	"github.com/atotto/clipboard"
)

type ClipboardManager struct {
	store storage.Storer
}

func NewManager(store storage.Storer) *ClipboardManager {
	return &ClipboardManager{
		store: store,
	}
}

func (cm *ClipboardManager) CopyToSlot(slot int) error {
	text, err := clipboard.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read clipboard: %w", err)
	}

	if err := cm.store.Set(slot, text); err != nil {
		return fmt.Errorf("failed to store clipboard content: %w", err)
	}

	fmt.Printf("Copied to slot %d: %s\n", slot, truncateString(text, 50))
	return nil
}

func (cm *ClipboardManager) PasteFromSlot(slot int) error {
	text, err := cm.store.Get(slot)
	if err != nil {
		return fmt.Errorf("failed to retrieve content: %w", err)
	}

	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	fmt.Printf("Pasted from slot %d: %s\n", slot, truncateString(text, 50))
	return nil
}

func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}
