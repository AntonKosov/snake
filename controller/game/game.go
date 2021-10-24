package game

import (
	"context"
	"math/rand"
	"snake/controller"
	"snake/input"
	"snake/utils"
	"snake/view"
	"time"
)

const (
	baseTurnDelayMs = 250
	speedUpMs       = 2
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

type game struct {
	fieldWidth       int
	fieldHeight      int
	scene            view.Game
	score            int
	activateGameOver controller.ActivateGameOverScene
	fruit            utils.Vector
	snake            *snake
}

func Run(ctx context.Context,
	fieldWidth, fieldHeight int,
	scene view.Game,
	in <-chan input.Input,
	activateGameOver controller.ActivateGameOverScene) {

	g := game{
		fieldWidth:       fieldWidth,
		fieldHeight:      fieldHeight,
		scene:            scene,
		activateGameOver: activateGameOver,
		snake:            newSnake(fieldHeight, fieldWidth),
	}

	scene.Activate(fieldWidth, fieldHeight)
	g.outputAllSnake()
	g.generateFruit()
	g.scene.OutputScore(0)

	scene.DrawFrame()

	g.handleInput(ctx, in)
}

func (g *game) handleInput(ctx context.Context, in <-chan input.Input) {
	dirHandler := func(d direction) func() {
		return func() { g.snake.rotateHead(d) }
	}
	keyHandlers := map[input.Input]func(){
		input.UpKey:    dirHandler(directionUp),
		input.LeftKey:  dirHandler(directionLeft),
		input.DownKey:  dirHandler(directionDown),
		input.RightKey: dirHandler(directionRight),
	}
	timer := time.NewTicker(time.Millisecond * baseTurnDelayMs)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			ateFruit, ok := g.makeMove()
			if !ok {
				g.activateGameOver(g.score)
				return
			}
			if ateFruit {
				timer.Reset(time.Millisecond * time.Duration(baseTurnDelayMs-g.score*speedUpMs))
			}
		case key := <-in:
			if action, ok := keyHandlers[key]; ok {
				action()
			}
		}
	}
}

func (g *game) outputAllSnake() {
	for i, p := range g.snake.body {
		if i == len(g.snake.body)-1 {
			g.scene.OutputSnakeHead(p.X, p.Y)
		} else {
			g.scene.OutputSnakeBody(p.X, p.Y)
		}
	}
}

func (g *game) makeMove() (ateFruit bool, ok bool) {
	previousHeadPosition, currentHeadPosition, tail, ok := g.snake.move()
	if !ok {
		return false, false
	}

	g.scene.OutputField(tail.X, tail.Y)
	g.scene.OutputSnakeBody(previousHeadPosition.X, previousHeadPosition.Y)
	g.scene.OutputSnakeHead(currentHeadPosition.X, currentHeadPosition.Y)

	ateFruit = false
	if currentHeadPosition == g.fruit {
		ateFruit = true
		g.score++
		g.snake.increaseTail(tail)
		g.scene.OutputSnakeBody(tail.X, tail.Y)
		g.generateFruit()
		g.scene.OutputScore(g.score)
	}

	g.scene.DrawFrame()
	return ateFruit, true
}

func (g *game) generateFruit() {
	position := rand.Intn(g.snake.freeCells)
	for p, isFree := range g.snake.freeCellsMap {
		if !isFree {
			continue
		}
		if position == 0 {
			g.fruit = p
			g.scene.OutputFruit(p.X, p.Y)
			return
		}
		position--
	}
	panic("Fruit position not found")
}
