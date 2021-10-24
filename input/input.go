package input

type Input string

const (
	EscapeKey = "escape"
	QuitKey   = "quit"
	EnterKey  = "enter"
	UpKey     = "up"
	LeftKey   = "left"
	DownKey   = "down"
	RightKey  = "right"
)

type InputController interface {
	Input() <-chan Input
}
