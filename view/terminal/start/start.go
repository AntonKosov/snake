package start

import (
	"snake/view"
	"snake/view/terminal/colors"
	"snake/view/terminal/fonts"
	"snake/view/terminal/utils"

	tcell "github.com/gdamore/tcell/v2"
)

type startScene struct {
	emmiter utils.Emmiter
}

func New(screen tcell.Screen, screenParams view.ScreenParams) view.Start {
	s := startScene{
		emmiter: utils.NewEmmiter(screen, screenParams),
	}

	return &s
}

func (s *startScene) Activate() {
	s.emmiter.Clear()

	style := tcell.StyleDefault.Foreground(tcell.ColorGreenYellow).Background(colors.Background)

	s.emmiter.EmitCenteredText(style, fonts.Big, 5, "SNAKE")
	s.emmiter.EmitCenteredText(style, fonts.Small, 13, "PRESS ENTER")

	s.emmiter.Show()
}
