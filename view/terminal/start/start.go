package start

import (
	"snake/view"
	"snake/view/terminal/colors"
	"snake/view/terminal/fonts"
	"snake/view/terminal/utils"

	tcell "github.com/gdamore/tcell/v2"
)

type startScene struct {
	emitter utils.Emitter
}

func New(screen tcell.Screen, screenParams view.ScreenParams) view.Start {
	s := startScene{
		emitter: utils.NewEmitter(screen, screenParams),
	}

	return &s
}

func (s *startScene) Activate() {
	s.emitter.Clear()

	style := tcell.StyleDefault.Foreground(tcell.ColorGreenYellow).Background(colors.Background)

	s.emitter.EmitCenteredText(style, fonts.Big, 5, "SNAKE")
	s.emitter.EmitCenteredText(style, fonts.Small, 13, "PRESS ENTER")

	s.emitter.Show()
}
