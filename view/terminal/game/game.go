package game

import (
	"fmt"
	"snake/view"
	"snake/view/terminal/colors"
	"snake/view/terminal/fonts"
	"snake/view/terminal/utils"

	tcell "github.com/gdamore/tcell/v2"
)

type gameScene struct {
	emitter utils.Emitter
	// screen       tcell.Screen
	screenParams view.ScreenParams
	fieldWidth   int
	fieldHeight  int

	// styles
	lightFieldStyle tcell.Style
	darkFieldStyle  tcell.Style
	snakeHeadStyle  tcell.Style
	snakeBodyStyle  tcell.Style
	fruitStyle      tcell.Style
	scoreStyle      tcell.Style
}

func New(screen tcell.Screen, screenParams view.ScreenParams) view.Game {
	// characters: https://www.compart.com/en/unicode/block/U+2580
	s := gameScene{
		emitter: utils.NewEmitter(screen, screenParams),
		// screen:          screen,
		screenParams:    screenParams,
		lightFieldStyle: tcell.StyleDefault.Background(tcell.ColorLightGreen),
		darkFieldStyle:  tcell.StyleDefault.Background(tcell.ColorDarkGreen),
		snakeHeadStyle:  tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorYellow),
		snakeBodyStyle:  tcell.StyleDefault.Background(tcell.ColorLightYellow),
		fruitStyle:      tcell.StyleDefault.Background(tcell.ColorRed),
		scoreStyle:      tcell.StyleDefault.Background(colors.Background).Foreground(tcell.ColorBlue),
	}
	return &s
}

func (s *gameScene) Activate(fieldWidth, fieldHeight int) {
	s.emitter.Clear()

	s.fieldWidth = fieldWidth
	s.fieldHeight = fieldHeight

	s.drawField()

	s.DrawFrame()
}

func (s *gameScene) DrawFrame() {
	s.emitter.Show()
}

func (s *gameScene) OutputSnakeHead(x, y int) {
	s.emitter.SetFieldContent(x, y, s.snakeHeadStyle, '▓')
}

func (s *gameScene) OutputSnakeBody(x, y int) {
	s.emitter.SetFieldContent(x, y, s.snakeBodyStyle, '░')
}

func (s *gameScene) OutputField(x, y int) {
	style := s.getFieldCellStyle(x, y)
	s.emitter.SetFieldContent(x, y, style, ' ')
}

func (s *gameScene) OutputFruit(x, y int) {
	s.emitter.SetFieldContent(x, y, s.fruitStyle, ' ')
}

func (s *gameScene) OutputScore(score int) {
	font := fonts.Small
	fontHeight := fonts.FontHeight(font)
	s.emitter.FillTopRows(s.scoreStyle, fontHeight)
	s.emitter.EmitCenteredText(s.scoreStyle, font, 0, fmt.Sprintf("SCORE: %d", score))
}

func (s *gameScene) GetFieldSize() (width, height int) {
	return s.screenParams.FieldWidth, s.screenParams.FieldHeight
}

func (s *gameScene) getFieldCellStyle(x, y int) tcell.Style {
	if (x+y)%2 == 0 {
		return s.lightFieldStyle
	}

	return s.darkFieldStyle
}

func (s *gameScene) drawField() {
	for x := 0; x < s.fieldWidth; x++ {
		for y := 0; y < s.fieldHeight; y++ {
			s.OutputField(x, y)
		}
	}
}
