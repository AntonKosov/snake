package terminal

import (
	"fmt"
	"snake/input"
	"snake/view"
	"snake/view/terminal/colors"
	"snake/view/terminal/game"
	"snake/view/terminal/gameover"
	"snake/view/terminal/start"
	"sync"

	tcell "github.com/gdamore/tcell/v2"
)

const (
	minScreenWidth  = 58
	minScreenHeight = 20
	fieldWidth      = 29
	fieldHeight     = 16
	fieldStartY     = 4
)

type Factory struct {
	screen             tcell.Screen
	screenParams       view.ScreenParams
	input              chan input.Input
	terminateWaitGroup sync.WaitGroup
	terminateInputCh   chan struct{}
	errCh              chan<- error
}

func New(errCh chan<- error) (*Factory, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}

	defStyle := tcell.StyleDefault.
		Background(colors.Background).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	f := Factory{
		screen: s,
		screenParams: view.ScreenParams{
			MinScreenWidth:  minScreenWidth,
			MinScreenHeight: minScreenHeight,
			FieldWidth:      fieldWidth,
			FieldHeight:     fieldHeight,
			FieldStartY:     fieldStartY,
		},
		input:            make(chan input.Input),
		terminateInputCh: make(chan struct{}),
		errCh:            errCh,
	}

	f.runInputHandler()

	return &f, nil
}

var keyMapping map[tcell.Key]input.Input

func init() {
	keyMapping = map[tcell.Key]input.Input{
		tcell.KeyEscape: input.EscapeKey,
		tcell.KeyCtrlC:  input.QuitKey,
		tcell.KeyEnter:  input.EnterKey,
		tcell.KeyUp:     input.UpKey,
		tcell.KeyRight:  input.RightKey,
		tcell.KeyLeft:   input.LeftKey,
		tcell.KeyDown:   input.DownKey,
	}
}

func (f *Factory) CreateStartScene() view.Start {
	return start.New(f.screen, f.screenParams)
}

func (f *Factory) CreateGameScene() view.Game {
	return game.New(f.screen, f.screenParams)
}

func (f *Factory) CreateGameOverScene() view.GameOver {
	return gameover.New(f.screen, f.screenParams)
}

func (f *Factory) Close() {
	close(f.terminateInputCh)
	f.terminateWaitGroup.Wait()
	f.screen.Fini()
}

func (f *Factory) Input() <-chan input.Input {
	return f.input
}

func (f *Factory) runInputHandler() {
	eventCh := make(chan tcell.Event)
	f.terminateWaitGroup.Add(1)
	go func() {
		defer f.terminateWaitGroup.Done()
		failOnResize := false
		for {
			select {
			case <-f.terminateInputCh:
				return
			case e := <-eventCh:
				switch event := e.(type) {
				case *tcell.EventResize:
					if failOnResize {
						f.errCh <- fmt.Errorf("screen size has been changed")
						return
					}
					w, h := f.screen.Size()
					if w < f.screenParams.MinScreenWidth {
						f.errCh <- fmt.Errorf("screen width must be at least %d", f.screenParams.MinScreenWidth)
						return
					}
					if h < f.screenParams.MinScreenHeight {
						f.errCh <- fmt.Errorf("screen height must be at least %d", f.screenParams.MinScreenHeight)
						return
					}
					failOnResize = true
				case *tcell.EventKey:
					if key, ok := keyMapping[event.Key()]; ok {
						f.input <- key
					}
				}
			}
		}
	}()

	f.terminateWaitGroup.Add(1)
	go func() {
		defer f.terminateWaitGroup.Done()
		f.screen.ChannelEvents(eventCh, f.terminateInputCh)
	}()
}
