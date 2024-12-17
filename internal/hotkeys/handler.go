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
		fmt.Printf("Key event: Mask=%v, Keycode=%v, Rawcode=%v\n",
			e.Mask, e.Keycode, e.Rawcode)

		// Check for numpad keys and handle modifier masks
		if e.Rawcode >= 87 && e.Rawcode <= 89 { // Numpad 1-3
			keyNum := int(e.Rawcode - 86)
			h.handleNumKey(keyNum, e.Mask)
		} else if e.Rawcode >= 83 && e.Rawcode <= 85 { // Numpad 4-6
			keyNum := int(e.Rawcode - 79)
			h.handleNumKey(keyNum, e.Mask)
		} else if e.Rawcode >= 79 && e.Rawcode <= 81 { // Numpad 7-9
			keyNum := int(e.Rawcode - 72)
			h.handleNumKey(keyNum, e.Mask)
		}
	})

	go hook.Process(hook.Start())
	return nil
}

func (h *HotkeyHandler) handleNumKey(keyNum int, mask uint16) {
	const (
		MOD_CONTROL = 1 << 1
		MOD_ALT     = 1 << 3
	)

	isCtrl := mask&MOD_CONTROL != 0
	isAlt := mask&MOD_ALT != 0

	if isCtrl && isAlt {
		if err := h.clipManager.CopyToSlot(keyNum); err != nil {
			fmt.Printf("Error copying to slot %d: %v\n", keyNum, err)
		}
	} else if isCtrl && !isAlt {
		if err := h.clipManager.PasteFromSlot(keyNum); err != nil {
			fmt.Printf("Error pasting from slot %d: %v\n", keyNum, err)
		}
	}
}

func (h *HotkeyHandler) Stop() error {
	hook.End()
	return nil
}
