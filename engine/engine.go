package engine

import (
	"context"
	"snake/controller/game"
	"snake/controller/gameover"
	"snake/controller/start"
	"snake/input"
	"sync"
)

type eng struct {
	controllerInput chan input.Input
	sceneFactory    SceneFactory

	closingWaitGroup sync.WaitGroup
	closeSignal      chan struct{}
}

func Run(sceneFactory SceneFactory, in input.InputController) error {
	ctx, cancel := context.WithCancel(context.Background())
	e := eng{
		controllerInput: make(chan input.Input),
		sceneFactory:    sceneFactory,
		closeSignal:     make(chan struct{}),
	}
	defer func() {
		cancel()
		e.closingWaitGroup.Wait()
	}()

	e.activateStartScene(ctx)
	e.runInputHandler(in.Input())

	select {
	case <-ctx.Done():
		close(e.closeSignal)
		return ctx.Err()
	case err := <-in.Error():
		close(e.closeSignal)
		return err
	case <-e.closeSignal:
		return nil
	}
}

func (e *eng) runInputHandler(in <-chan input.Input) {
	e.closingWaitGroup.Add(1)
	go func() {
		defer e.closingWaitGroup.Done()
		for {
			select {
			case key := <-in:
				switch key {
				case input.QuitKey:
					close(e.closeSignal)
					return
				default:
					e.controllerInput <- key
				}
			case <-e.closeSignal:
				return
			}
		}
	}()
}

func (e *eng) activateStartScene(ctx context.Context) {
	e.closingWaitGroup.Add(1)
	go func() {
		defer e.closingWaitGroup.Done()
		scene := e.sceneFactory.CreateStartScene()
		start.Run(
			ctx,
			scene,
			e.controllerInput,
			func() { e.activateGameScene(ctx) },
		)
	}()
}

func (e *eng) activateGameScene(ctx context.Context) {
	e.closingWaitGroup.Add(1)
	go func() {
		defer e.closingWaitGroup.Done()
		scene := e.sceneFactory.CreateGameScene()
		w, h := scene.GetFieldSize()
		game.Run(ctx, w, h, scene, e.controllerInput,
			func(score int) { e.activateGameOver(ctx, score) })
	}()
}

func (e *eng) activateGameOver(ctx context.Context, score int) {
	e.closingWaitGroup.Add(1)
	go func() {
		defer e.closingWaitGroup.Done()
		scene := e.sceneFactory.CreateGameOverScene()
		gameover.Run(ctx, score, scene, e.controllerInput,
			func() { e.activateStartScene(ctx) })
	}()
}
