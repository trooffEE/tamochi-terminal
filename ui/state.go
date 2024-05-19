package ui

import "tamochi-terminal/ui/internal/model"

type state struct {
	user    model.User
	bubbles model.Bubbles
}

func newTerminalState() state {
	return state{
		user:    model.NewUser(),
		bubbles: model.NewBubbles(),
	}
}
