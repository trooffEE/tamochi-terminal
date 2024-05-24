package ui

import "tamochi-terminal/ui/internal/models"

type State struct {
	user     models.User
	bubbles  models.Bubbles
	screen   models.Screen
	Terminal models.Terminal
}

func newTerminalState() State {
	return State{
		user:     *models.NewUser(),
		bubbles:  *models.NewBubbles(),
		screen:   *models.NewScreen(),
		Terminal: *models.NewTerminal(),
	}
}
