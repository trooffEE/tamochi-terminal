package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (s state) Init() tea.Cmd {
	return tea.Batch(
		tea.ClearScreen,
		tea.SetWindowTitle("Tamochi App"),
	)
}

func (s state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s state) View() string {
	return "Hello, kitty"
}

func NewTerminalApp() *tea.Program {
	return tea.NewProgram(newTerminalState())
}
