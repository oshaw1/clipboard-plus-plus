# Clipboard++
A multi-slot clipboard manager written in Go

## Usage

Download the latest built binary from: https://github.com/oshaw1/clipboard-plus-plus/releases

To build / run use a command such as `go build cmd/clipboardplusplus/main.go` or `go run cmd/clipboardplusplus/main.go`

To copy something using the tool first copy to your clipboard (Ctrl+C) then copy to a slot Right-Ctrl+Numpad[1-9]

Right-Ctrl+Numpad[1-9] to copy current clipboard to slot
Left-Ctrl+Numpad[1-9] to paste from slot
Press Ctrl+C to exit

## Packages used:
- **robotgo/gohook**: Keyboard event handling and system automation
- **atotto/clipboard**: Clipboard management
