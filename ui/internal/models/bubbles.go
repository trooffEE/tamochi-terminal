package models

import "github.com/charmbracelet/bubbles/textinput"

type Bubbles struct {
	TextInput Component[textinput.Model]
}

type Component[T any] struct {
	Component T
	Err       error
}

func NewTextInput() Component[textinput.Model] {
	ti := textinput.New()
	ti.Placeholder = "Test"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Component[textinput.Model]{
		Component: ti,
		Err:       nil,
	}
}

func NewBubbles() *Bubbles {
	return &Bubbles{
		TextInput: NewTextInput(),
	}
}
