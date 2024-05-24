package models

const (
	ErrorScreen = (iota - 1)
	LoginScreen
)

type Screen struct {
	Id   int
	Data interface{}
}

func NewScreen() *Screen {
	return &Screen{
		Id:   LoginScreen,
		Data: nil,
	}
}

func (s Screen) SetScreenStep(newScreenId int) {
	s.Id = newScreenId
}
