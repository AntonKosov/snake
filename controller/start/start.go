package start

import (
	"context"
	"snake/controller"
	"snake/input"
	"snake/view"
)

func Run(ctx context.Context,
	scene view.Start,
	in <-chan input.Input,
	activateGame controller.ActivateGameScene) {
	scene.Activate()
	for {
		select {
		case <-ctx.Done():
			return
		case key := <-in:
			if key == input.EnterKey {
				activateGame()
				return
			}
		}
	}
}
