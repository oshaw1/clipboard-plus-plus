package hotkeys

import (
	"clipboardplusplus/internal/clipboard"
	"fmt"

	hook "github.com/robotn/gohook"
)

type HotkeyHandler struct {
	clipManager clipboard.Manager
}

func NewHandler(clipManager clipboard.Manager) *HotkeyHandler {
	return &HotkeyHandler{
		clipManager: clipManager,
	}
}

func (h *HotkeyHandler) Start() error {
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		// Only process numpad keys (97-105 for 1-9)
		if e.Rawcode >= 97 && e.Rawcode <= 105 {
			keyNum := int(e.Rawcode - 96) // Convert to 1-9
			h.handleNumKey(keyNum, e.Mask)
		}
	})
	go hook.Process(hook.Start())
	return nil
}

func (h *HotkeyHandler) handleNumKey(keyNum int, mask uint16) {
	const (
		LEFT_CONTROL  = 1 << 1 // Mask: 2
		RIGHT_CONTROL = 1 << 5 // Mask: 32
	)

	if mask == RIGHT_CONTROL {
		if err := h.clipManager.CopyToSlot(keyNum); err != nil {
			fmt.Printf("Error copying to slot %d: %v\n", keyNum, err)
		} else {
			fmt.Printf("Copied to slot %d\n", keyNum)
		}
	} else if mask == LEFT_CONTROL {
		if err := h.clipManager.PasteFromSlot(keyNum); err != nil {
			fmt.Printf("Error pasting from slot %d: %v\n", keyNum, err)
		} else {
			fmt.Printf("Pasted from slot %d\n", keyNum)
		}
	}
}

func (h *HotkeyHandler) Stop() error {
	hook.End()
	return nil
}
