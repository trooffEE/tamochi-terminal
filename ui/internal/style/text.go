package style

import "github.com/charmbracelet/lipgloss"

func newText() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(TEXT_COLOR__DEFAULT).Copy()
}

func Heading(text interface{}) string {
	heading := newText().Bold(true).MarginBottom(2)

	switch value := text.(type) {
	case string:
		return heading.Render(value)
	case interface{}:
		return ""
	}
}

func Paragraph() lipgloss.Style {
	return newText().MarginBottom(1)
}

func Span() lipgloss.Style {
	return newText().Inline(true)
}
