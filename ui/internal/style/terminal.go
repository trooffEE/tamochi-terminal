package style

import "github.com/charmbracelet/lipgloss"

func Terminal(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(10, 5, 5, 4).
		Background(BACKGROUND_COLOR__DEFAULT).
		Width(width).
		Height(height)
}
