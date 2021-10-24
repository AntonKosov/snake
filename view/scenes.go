package view

type Start interface {
	Activate()
}

type Game interface {
	Activate(fieldWidth, fieldHeight int)
	DrawFrame()
	OutputSnakeHead(x, y int)
	OutputSnakeBody(x, y int)
	OutputField(x, y int)
	OutputFruit(x, y int)
	OutputScore(s int)
	GetFieldSize() (width, height int)
}

type GameOver interface {
	Activate(score int)
}
