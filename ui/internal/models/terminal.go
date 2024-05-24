package models

import (
	"github.com/charmbracelet/lipgloss"
	"tamochi-terminal/ui/internal/style"
)

type Terminal struct {
	Width, Height int
}

func NewTerminal() *Terminal {
	return &Terminal{
		Width:  1024,
		Height: 720,
	}
}

func (s *Terminal) PaintTerminal() lipgloss.Style {
	return style.Terminal(s.Width, s.Height)
}
