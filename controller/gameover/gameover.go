package gameover

import (
	"context"
	"snake/controller"
	"snake/input"
	"snake/view"
)

func Run(ctx context.Context,
	score int,
	scene view.GameOver,
	in <-chan input.Input,
	startScene controller.ActivateStartScene) {
	scene.Activate(score)
	for {
		select {
		case <-ctx.Done():
			return
		case key := <-in:
			if key == input.EnterKey {
				startScene()
				return
			}
		}
	}
}
