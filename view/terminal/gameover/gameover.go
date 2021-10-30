package gameover

import (
	"fmt"
	"snake/view"
	"snake/view/terminal/colors"
	"snake/view/terminal/fonts"
	"snake/view/terminal/utils"

	tcell "github.com/gdamore/tcell/v2"
)

type gameOver struct {
	emitter utils.Emitter
}

func New(screen tcell.Screen, screenParams view.ScreenParams) view.GameOver {
	s := gameOver{
		emitter: utils.NewEmitter(screen, screenParams),
	}
	return &s
}

func (s *gameOver) Activate(score int) {
	s.emitter.Clear()

	gameOverStyle := tcell.StyleDefault.
		Background(colors.Background).
		Foreground(tcell.ColorRed)
	style := tcell.StyleDefault.
		Foreground(tcell.ColorCadetBlue.TrueColor()).
		Background(colors.Background)

	s.emitter.EmitCenteredText(style, fonts.Small, 2, fmt.Sprintf("SCORE: %d", score))
	s.emitter.EmitCenteredText(gameOverStyle, fonts.Small, 9, "GAME OVER")
	s.emitter.EmitCenteredText(style, fonts.Small, 16, "PRESS ENTER")

	s.emitter.Show()
}
