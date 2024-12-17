package main

import (
	"fmt"
	"os"
	"os/signal"

	"clipboardplusplus/internal/clipboard"
	"clipboardplusplus/internal/hotkeys"
	"clipboardplusplus/internal/storage"
)

const version = "1.0.0"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("Clipboard++ v%s\n", version)
		return
	}

	store := storage.NewMemoryStorage()
	clipManager := clipboard.NewManager(store)
	hotkeyHandler := hotkeys.NewHandler(clipManager)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	printBanner()

	if err := hotkeyHandler.Start(); err != nil {
		fmt.Printf("Error starting Clipboard++: %v\n", err)
		os.Exit(1)
	}

	<-c
	fmt.Println("\nShutting down Clipboard++...")
	hotkeyHandler.Stop()
}

func printBanner() {
	fmt.Println(`
Clipboard++ v` + version + `
A minimal clipboard manager for power users
https://oshaw1.github.io

Right-Ctrl+[1-9] to copy current clipboard to slot
Left-Ctrl+[1-9] to paste from slot
Press Ctrl+C to exit
─────────────────────────────────────`)
}
