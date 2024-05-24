package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"tamochi-terminal/ui/internal/style"
)

func (s State) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.SetWindowTitle("Tamochi App"),
	)
}

func (s State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.Terminal.Width, s.Terminal.Height = msg.Width, msg.Height
		return s, nil
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return s, tea.Quit
		}
	}

	s.bubbles.TextInput.Component, cmd = s.bubbles.TextInput.Component.Update(msg)

	return s, cmd
}

func (s State) View() string {
	return s.Terminal.PaintTerminal().Render(
		style.Heading("123")

		)
}

func NewTerminalApp() *tea.Program {
	return tea.NewProgram(newTerminalState())
}
