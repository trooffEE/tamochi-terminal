package model

import "github.com/charmbracelet/bubbles/textinput"

type Bubbles struct {
	textInput Component[textinput.Model]
}

type Component[T any] struct {
	component T
	err       error
}

func NewTextInput() Component[textinput.Model] {
	ti := textinput.New()
	ti.Placeholder = "Test"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Component[textinput.Model]{
		component: ti,
		err:       nil,
	}
}

func NewBubbles() Bubbles {
	return Bubbles{
		textInput: NewTextInput(),
	}
}
