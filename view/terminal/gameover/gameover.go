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
	emmiter utils.Emmiter
}

func New(screen tcell.Screen, screenParams view.ScreenParams) view.GameOver {
	s := gameOver{
		emmiter: utils.NewEmmiter(screen, screenParams),
	}
	return &s
}

func (s *gameOver) Activate(score int) {
	s.emmiter.Clear()

	gameOverStyle := tcell.StyleDefault.
		Background(colors.Background).
		Foreground(tcell.ColorRed)
	style := tcell.StyleDefault.
		Foreground(tcell.ColorCadetBlue.TrueColor()).
		Background(colors.Background)

	s.emmiter.EmitCenteredText(style, fonts.Small, 2, fmt.Sprintf("SCORE: %d", score))
	s.emmiter.EmitCenteredText(gameOverStyle, fonts.Small, 9, "GAME OVER")
	s.emmiter.EmitCenteredText(style, fonts.Small, 16, "PRESS ENTER")

	s.emmiter.Show()
}
