package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"tamochi-terminal/ui"
)

func main() {
	app := ui.NewTerminalApp()

	if _, err := app.Run(); err != nil {
		_, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Printf("Check debug errors: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
